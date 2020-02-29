/*
 * Copyright (C) 2020. Niklas Linz - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the LGPLv3 license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the LGPLv3 license with
 * this file. If not, please write to: niklas.linz@enigmar.de
 *
 */

package de.linzn.uvr1611decoder;

import de.linzn.uvr1611decoder.appConfiguration.AppConfigurationModule;
import de.linzn.uvr1611decoder.socket.SocketManager;

public class UVR1611DecoderApp {
    public static UVR1611DecoderApp UVR1611DecoderApp;
    public static AppConfigurationModule appConfigurationModule;
    private SocketManager socketManager;

    public UVR1611DecoderApp() {
        this.socketManager = new SocketManager();
        this.socketManager.startConnection();
    }

    public static void main(String[] args) {
        appConfigurationModule = new AppConfigurationModule();
        UVR1611DecoderApp = new UVR1611DecoderApp();
    }
}
