package com.init;

import io.lightflame.bootstrap.LightFlame;

public class App 
{
    public static void main( String[] args )
    {
        new LightFlame()
            .port(8080)
            .start(App.class);
    }
}
