package com.example.piasy.startup;

import android.app.Activity;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.util.Log;
import android.view.View;
import android.webkit.ConsoleMessage;
import android.webkit.JsResult;
import android.webkit.WebChromeClient;
import android.webkit.WebSettings;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.widget.Toast;

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
        mWebview.setWebChromeClient(new WebChromeClient() {
            @Override
            public boolean onConsoleMessage(@NonNull ConsoleMessage consoleMessage) {
                Log.d("WebViewActivity", consoleMessage.message() + " -- From line "
                        + consoleMessage.lineNumber() + " of "
                        + consoleMessage.sourceId());

                return true;
            }

            @Override
            public boolean onJsAlert(WebView view, String url, String message,
                                     JsResult result) {
                Log.d("WebViewActivity", message);
                Toast.makeText(getApplicationContext(), message, Toast.LENGTH_SHORT).show();
                result.confirm();
                return true;
            }
        });
        mWebview.loadUrl("http://101.5.236.43:8080/web/welcome.html");
    }
}
