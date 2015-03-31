package com.example.piasy.startup.location;

import com.example.piasy.startup.Controller;
import com.example.piasy.startup.R;
import com.promegu.xlog.base.XLog;

import org.json.JSONException;
import org.json.JSONObject;

import android.app.Activity;
import android.location.Location;
import android.os.Bundle;
import android.os.Environment;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.text.TextUtils;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.PrintWriter;

import butterknife.ButterKnife;
import butterknife.InjectView;
import butterknife.OnClick;

/**
 * Created by piasy on 15/3/5.
 */
public class PaintBoardFragment extends Fragment {

    private final static String START = "Start";
    private final static String STOP = "Stop";

    @InjectView(R.id.bt_start)
    Button mBtStart;
    @InjectView(R.id.paintBoard)
    PaintBoard mPaintBoard;

    private boolean started = false;
    private Controller mController;
    PrintWriter mWriter = null;

    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.paint_board_fragment, container, false);
        ButterKnife.inject(this, view);
        try {
            mWriter = new PrintWriter(new FileOutputStream(new File(
                    Environment.getExternalStorageDirectory() + "/location-capture-" +
                            System.currentTimeMillis() + ".txt")));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        return view;
    }

    @Override
    public void onAttach(Activity activity) {
        super.onAttach(activity);
        if (activity instanceof Controller) {
            mController = (Controller) activity;
        } else {
            throw new IllegalStateException("Activity must implement Controller");
        }
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        mWriter.close();
    }

    @Override
    public void onResume() {
        super.onResume();
        if (mBtStart != null) {
            mBtStart.setText(START);
            started = false;
        }
    }

    @OnClick(R.id.bt_start)
    void onStartClick() {
        if (TextUtils.equals(mBtStart.getText(), START)) {
            started = true;
            mBtStart.setText(STOP);
            mController.start();
        } else if (TextUtils.equals(mBtStart.getText(), STOP)) {
            started = false;
            mBtStart.setText(START);
            mController.stop();
        }
    }

    @XLog
    public void updateLocation(Location location) {
        if (started && mWriter != null) {
            JSONObject loc = new JSONObject();
            try {
                loc.put("time", location.getTime());
                loc.put("latitude", location.getLatitude());
                loc.put("longitude", location.getLongitude());
                if (location.hasAltitude()) {
                    loc.put("altitude", location.getAltitude());
                }
                if (location.hasSpeed()) {
                    loc.put("speed", location.getSpeed());
                }
                if (location.hasAccuracy()) {
                    loc.put("accuracy", location.getAccuracy());
                }
                if (location.hasBearing()) {
                    loc.put("bearing", location.getBearing());
                }

                mWriter.println(loc.toString());

                mPaintBoard.addAboveShape(new PaintBoard.Point((float) location.getLatitude(),
                        (float) location.getLongitude(), 2));
                if (location.hasAccuracy()) {
                    mPaintBoard.addUnderShape(new PaintBoard.Circle((float) location.getLatitude(),
                            (float) location.getLongitude(), location.getAccuracy()));
                }
            } catch (JSONException e) {
                e.printStackTrace();
            }
        }
    }
}
