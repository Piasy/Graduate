package com.example.piasy.startup;

import com.example.piasy.startup.location.PaintBoardFragment;
import com.example.piasy.startup.sensors.LogBoardFragment;

import android.content.Context;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;
import android.location.Location;
import android.location.LocationListener;
import android.location.LocationManager;
import android.os.Bundle;
import android.support.v4.app.FragmentActivity;
import android.util.Log;

import java.math.BigDecimal;
import java.math.RoundingMode;


public class MyActivity extends FragmentActivity implements SensorEventListener, Controller {

    private static final int MODE_NONE = 0;
    private static final int MODE_SENSOR = 1;
    private static final int MODE_LOCATION = 2;

    private int mode = MODE_NONE;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        mode = MODE_LOCATION;
        switch (mode) {
            case MODE_SENSOR:
                measureSensors();
                break;
            case MODE_LOCATION:
                measureLocation();
                break;
            default:
                break;
        }
        mStarted = true;
    }

    PaintBoardFragment mPainter;
    LocationManager mLocationManager;
    LocationListener mLocationListener = new LocationListener() {
        @Override
        public void onLocationChanged(Location location) {
            mPainter.updateLocation(location);
        }

        @Override
        public void onStatusChanged(String provider, int status, Bundle extras) {

        }

        @Override
        public void onProviderEnabled(String provider) {

        }

        @Override
        public void onProviderDisabled(String provider) {

        }
    };
    private void measureLocation() {
        mPainter = new PaintBoardFragment();
        getSupportFragmentManager().beginTransaction()
                .add(android.R.id.content, mPainter, mPainter.getClass().getName()).commit();
        mLocationManager = (LocationManager) this.getSystemService(Context.LOCATION_SERVICE);
        mLocationManager.requestLocationUpdates(LocationManager.GPS_PROVIDER, 0, 0, mLocationListener);
    }

    LogBoardFragment mLogger;
    SensorManager mSensorManager;
    Sensor mAccelerometer, mGravity, mGyroscope, mLinearAcc;
    private void measureSensors() {
        mLogger = new LogBoardFragment();
        getSupportFragmentManager().beginTransaction()
                .add(android.R.id.content, mLogger, mLogger.getClass().getName()).commit();
        mSensorManager = (SensorManager) getSystemService(Context.SENSOR_SERVICE);
        mAccelerometer = mSensorManager.getDefaultSensor(Sensor.TYPE_ACCELEROMETER);
        mGravity = mSensorManager.getDefaultSensor(Sensor.TYPE_GRAVITY);
        mGyroscope = mSensorManager.getDefaultSensor(Sensor.TYPE_GYROSCOPE);
        mLinearAcc = mSensorManager.getDefaultSensor(Sensor.TYPE_LINEAR_ACCELERATION);
        mSensorManager.registerListener(this, mAccelerometer, SensorManager.SENSOR_DELAY_UI);
        mSensorManager.registerListener(this, mGravity, SensorManager.SENSOR_DELAY_UI);
        mSensorManager.registerListener(this, mLinearAcc, SensorManager.SENSOR_DELAY_UI);
    }

    @Override
    protected void onStart() {
        super.onStart();
        Log.d("xjltest", "MyActivity onStart");
    }

    @Override
    protected void onRestart() {
        super.onRestart();
        Log.d("xjltest", "MyActivity onRestart");
    }

    @Override
    protected void onResume() {
        super.onResume();
        Log.d("xjltest", "MyActivity onResume");
    }

    @Override
    protected void onPause() {
        super.onPause();
        Log.d("xjltest", "MyActivity onPause");
    }

    @Override
    protected void onStop() {
        super.onStop();
        Log.d("xjltest", "MyActivity onStop");
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        Log.d("xjltest", "MyActivity onDestroy");
        if (mSensorManager != null) {
            mSensorManager.unregisterListener(this);
        }
        if (mLocationManager != null) {
            mLocationManager.removeUpdates(mLocationListener);
        }
    }

    @Override
    protected void onUserLeaveHint() {
        super.onUserLeaveHint();
    }

    float [] mLastGravity = null;
    @Override
    public void onSensorChanged(SensorEvent event) {
        if (event.sensor.equals(mAccelerometer)) {
            mLogger.log("ACCELEROMETER: x = " + cutBit(event.values[0], 2) + " m/s^2, "
                + "y = " + cutBit(event.values[1], 2) + " m/s^2, "
                + "z = " + cutBit(event.values[2], 2) + " m/s^2.");
            if (mLastGravity != null) {
                mLogger.log("REAL: x = " + cutBit(event.values[0] - mLastGravity[0], 2) + " m/s^2, "
                        + "y = " + cutBit(event.values[1] - mLastGravity[1], 2) + " m/s^2, "
                        + "z = " + cutBit(event.values[2] - mLastGravity[2], 2) + " m/s^2.");
            }
        } else if (event.sensor.equals(mGravity)) {
            mLogger.log("GRAVITY: x = " + cutBit(event.values[0], 2) + " m/s^2, "
                    + "y = " + cutBit(event.values[1], 2) + " m/s^2, "
                    + "z = " + cutBit(event.values[2], 2) + " m/s^2.");
            mLastGravity = event.values;
        } else if (event.sensor.equals(mLinearAcc)) {
            mLogger.log("LINEAR_ACC: x = " + cutBit(event.values[0], 2) + " m/s^2, "
                    + "y = " + cutBit(event.values[1], 2) + " m/s^2, "
                    + "z = " + cutBit(event.values[2], 2) + " m/s^2.");
        }
    }

    @Override
    public void onAccuracyChanged(Sensor sensor, int accuracy) {

    }

    private float cutBit(float num, int bit) {
        return new BigDecimal(num).setScale(bit, RoundingMode.UP).floatValue();
    }

    boolean mStarted = false;
    @Override
    public void start() {
        if (mStarted) {
            return;
        }
        switch (mode) {
            case MODE_SENSOR:
                mSensorManager.registerListener(this, mAccelerometer, SensorManager.SENSOR_DELAY_UI);
                mSensorManager.registerListener(this, mGravity, SensorManager.SENSOR_DELAY_UI);
                mSensorManager.registerListener(this, mLinearAcc, SensorManager.SENSOR_DELAY_UI);
                break;
            case MODE_LOCATION:
                mLocationManager.requestLocationUpdates(LocationManager.GPS_PROVIDER, 0, 0, mLocationListener);
                break;
            default:
                break;
        }
        mStarted = true;
    }

    @Override
    public void stop() {
        if (!mStarted) {
            return;
        }
        switch (mode) {
            case MODE_SENSOR:
                mSensorManager.unregisterListener(this);
                break;
            case MODE_LOCATION:
                mLocationManager.removeUpdates(mLocationListener);
                break;
            default:
                break;
        }
        mStarted = false;
    }
}
