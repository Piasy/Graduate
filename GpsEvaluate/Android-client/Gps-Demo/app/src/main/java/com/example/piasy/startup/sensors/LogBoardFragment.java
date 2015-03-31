package com.example.piasy.startup.sensors;

import com.example.piasy.startup.Controller;
import com.example.piasy.startup.R;

import android.app.Activity;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.text.TextUtils;
import android.text.method.ScrollingMovementMethod;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.TextView;

import butterknife.ButterKnife;
import butterknife.InjectView;
import butterknife.OnClick;

/**
 * Created by piasy on 15/3/5.
 */
public class LogBoardFragment extends Fragment implements Logger {

    private final static String START = "Start";
    private final static String STOP = "Stop";

    @InjectView(R.id.bt_start)
    Button mBtStart;

    @InjectView(R.id.tv_board)
    TextView mTvBoard;

    private boolean started = false;
    private Controller mController;

    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.log_board_fragment, container, false);
        ButterKnife.inject(this, view);
        mTvBoard.setMovementMethod(new ScrollingMovementMethod());
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
    public void onResume() {
        super.onResume();
        if (mBtStart != null) {
            mBtStart.setText(START);
            started = false;
        }

        if (mTvBoard != null) {
            mTvBoard.setText("");
        }
    }

    //@XLog
    @Override
    public void log(String text) {
        if (started) {
            if (mTvBoard.getText().length() > 100 * 1000) {
                onClearClick();
            }
            mTvBoard.setText(mTvBoard.getText() + text + "\n");
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

    @OnClick(R.id.bt_clear)
    void onClearClick() {
        mTvBoard.setText("");
    }
}
