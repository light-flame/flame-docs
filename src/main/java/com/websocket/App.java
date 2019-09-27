package com.websocket;

import io.lightflame.bootstrap.LightFlame;

public class App {
    public static void main(String[] args) {
        new LightFlame()
                .addConfiguration(new Config().setDefautHandlers(), null)
                .addBasicLog4jConfig()
                .addHttpAndWsListener(8080)
                .start(App.class);
    }
}