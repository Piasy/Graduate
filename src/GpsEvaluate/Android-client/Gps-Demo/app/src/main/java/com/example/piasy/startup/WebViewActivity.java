package com.example.piasy.startup;

import android.app.Activity;
import android.os.Bundle;
import android.view.View;
import android.webkit.WebSettings;
import android.webkit.WebView;
import android.webkit.WebViewClient;

/**
 * Created by piasy on 15/6/9.
 */
public class WebViewActivity extends Activity {

    private WebView mWebview ;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.web_view_layout);

        mWebview = (WebView) findViewById(R.id.web_view);

        mWebview.setScrollBarStyle(View.SCROLLBARS_OUTSIDE_OVERLAY);
        WebSettings webSettings = mWebview.getSettings();
        webSettings.setJavaScriptEnabled(true);
        webSettings.setDomStorageEnabled(true);
        mWebview.setWebViewClient(new WebViewClient()
        {
            public boolean shouldOverrideUrlLoading(WebView view, String url)
            {
                view.loadUrl(url);
                return true;
            }
        });
        mWebview.loadUrl("http://101.5.236.43:8080/web/welcome.html");
    }
}
