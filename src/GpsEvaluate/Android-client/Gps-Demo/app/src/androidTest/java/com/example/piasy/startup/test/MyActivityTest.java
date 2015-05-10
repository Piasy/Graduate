package com.example.piasy.startup.test;

import com.example.piasy.startup.MyActivity;

import android.test.ActivityInstrumentationTestCase2;

/**
 * Created by piasy on 14/10/26.
 */
public class MyActivityTest extends ActivityInstrumentationTestCase2<MyActivity> {
    @SuppressWarnings("deprecation")
    public MyActivityTest() {
        super("com.example.piasy.startup", MyActivity.class);
    }

    @Override
    protected void setUp() throws Exception {
        super.setUp();
        getActivity();
    }

    public void testButton() throws Exception {
//        onView(withId(R.id.button)).check(matches(allOf(isDisplayed(), withText("New Button Test"))));
//        onView(withId(R.id.textView1)).check(matches(withText("TEXT1")));
//        onView(withId(R.id.textView2)).check(matches(withText("TEXT2")));
//
//        onView(withId(R.id.button)).perform(click());

        Thread.sleep(5000);
    }
}
