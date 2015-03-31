package com.example.piasy.startup.location;

import android.content.Context;
import android.graphics.Canvas;
import android.graphics.Color;
import android.graphics.Paint;
import android.util.AttributeSet;
import android.view.View;

import java.util.ArrayList;

/**
 * Created by piasy on 15/3/30.
 */
public class PaintBoard extends View {
    public PaintBoard(Context context) {
        super(context);
        init();
    }

    public PaintBoard(Context context, AttributeSet attrs) {
        super(context, attrs);
        init();
    }

    public PaintBoard(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
        init();
    }

    static Paint pointPaint, circlePaint, linePaint;
    private void init() {
        pointPaint = new Paint(Paint.ANTI_ALIAS_FLAG);
        pointPaint.setColor(Color.BLUE);
        pointPaint.setStyle(Paint.Style.FILL);

        circlePaint = new Paint(Paint.ANTI_ALIAS_FLAG);
        circlePaint.setColor(Color.YELLOW);
        circlePaint.setAlpha(0x77);
        circlePaint.setStyle(Paint.Style.FILL);

        linePaint = new Paint(Paint.ANTI_ALIAS_FLAG);
        linePaint.setColor(Color.GREEN);
        linePaint.setStyle(Paint.Style.FILL);
    }

    private static float gps2length(float lon1, float lat1, float lon2, float lat2) {
        double R = 6378.137; // Radius of earth in KM
        double dLat = Math.abs(lat2 - lat1) * Math.PI / 180;
        double dLon = Math.abs(lon2 - lon1) * Math.PI / 180;
        double a = Math.sin(dLat/2) * Math.sin(dLat/2) +
                Math.cos(lat1 * Math.PI / 180) * Math.cos(lat2 * Math.PI / 180) *
                        Math.sin(dLon/2) * Math.sin(dLon/2);
        double c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));
        double d = R * c;
        return (float) (d * 1000); // meters
    }

    ArrayList<Shape> mAboveShapes = new ArrayList<Shape>();
    ArrayList<Shape> mUnderShapes = new ArrayList<Shape>();
    static float maxLon, maxLat, minLon, minLat;
    boolean inited = false;
    static float pixPerMeter = 1.7f;
    static float xOff = 80, yOff = 80;

    public void addAboveShape(Shape shape) {
        measure(shape);
        mAboveShapes.add(shape);
        invalidate();
    }

    public void addUnderShape(Shape shape) {
        measure(shape);
        mUnderShapes.add(shape);
        invalidate();
    }

    private void measure(Shape shape) {
        if (!inited && shape instanceof Pointer) {
            maxLon = ((Pointer) shape).getX();
            minLon = maxLon;
            maxLat = ((Pointer) shape).getY();
            minLat = maxLat;
            inited = true;
        }

        if (shape instanceof Pointer) {
            float lon = ((Pointer) shape).getX();
            float lat = ((Pointer) shape).getY();
            if (lon > maxLon) {
                maxLon = lon;
            }
            if (lon < minLon) {
                minLon = lon;
            }
            if (lat > maxLat) {
                maxLat = lat;
            }
            if (lat < minLat) {
                minLat = lat;
            }
        }

        float width = gps2length(minLon, minLat, maxLon, minLat);
        float height = gps2length(minLon, minLat, minLon, maxLat);
        if (width * height != 0) {
            pixPerMeter = Math.min(pixPerMeter, Math.min(getHeight() / height, getWidth() / width));
        }
    }

    @Override
    protected void onDraw(Canvas canvas) {
        for (Shape shape : mUnderShapes) {
            shape.draw(canvas);
        }

        for (int i = 0; i < mAboveShapes.size(); i++) {
            if (i == mAboveShapes.size() - 1) {
                pointPaint.setColor(Color.RED);
                mAboveShapes.get(i).draw(canvas);
                pointPaint.setColor(Color.BLUE);
            } else {
                mAboveShapes.get(i).draw(canvas);
            }
        }
    }

    public static abstract class Shape {
        public abstract void draw(Canvas canvas);
    }

    interface Pointer {
        float getX();
        float getY();
    }

    public static class Point extends Shape implements Pointer {
        private float x, y, radius;
        public Point(float x, float y, float radius) {
            this.x = x;
            this.y = y;
            this.radius = radius;
        }

        @Override
        public void draw(Canvas canvas) {
            float w = gps2length(minLon, maxLat, x, maxLat);
            float h = gps2length(minLon, maxLat, minLon, y);
            canvas.drawCircle(w * pixPerMeter + xOff, h * pixPerMeter + yOff, radius, pointPaint);
        }

        @Override
        public float getX() {
            return x;
        }

        @Override
        public float getY() {
            return y;
        }
    }

    public static class Circle extends Shape implements Pointer {
        private float x, y, radius;
        public Circle(float x, float y, float radius) {
            this.x = x;
            this.y = y;
            this.radius = radius;
        }

        @Override
        public void draw(Canvas canvas) {
            float w = gps2length(minLon, maxLat, x, maxLat);
            float h = gps2length(minLon, maxLat, minLon, y);
            canvas.drawCircle(w * pixPerMeter + xOff, h * pixPerMeter + yOff, radius * pixPerMeter, circlePaint);
        }

        @Override
        public float getX() {
            return x;
        }

        @Override
        public float getY() {
            return y;
        }
    }

}
