// IMyAidlInterface.aidl
package com.example.piasy.startup;

// Declare any non-default types here with import statements
//import com.example.piasy.startup.Permit;

interface IMyAidlInterface {
    /**
     * Demonstrates some basic types that you can use as parameters
     * and return values in AIDL.
     */
    void basicTypes(int anInt, long aLong, boolean aBoolean, float aFloat,
            double aDouble, String aString);
}
