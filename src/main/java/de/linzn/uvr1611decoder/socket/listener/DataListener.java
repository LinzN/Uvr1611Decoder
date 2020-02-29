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

package de.linzn.uvr1611decoder.socket.listener;

import de.linzn.uvr1611decoder.UVRDataRequest;
import de.linzn.zSocket.components.events.IListener;
import de.linzn.zSocket.components.events.ReceiveDataEvent;
import de.linzn.zSocket.components.events.handler.EventHandler;
import org.json.JSONObject;

import java.io.*;

public class DataListener implements IListener {

    @EventHandler(channel = "heater_data")
    public void onData(ReceiveDataEvent event) {
        DataInputStream in = new DataInputStream(new ByteArrayInputStream(event.getDataInBytes()));

        try {
            String command = in.readUTF();
            if (command.equalsIgnoreCase("request")) {
                JSONObject jsonObject = new UVRDataRequest().getDataObject();

                if (jsonObject != null) {

                    ByteArrayOutputStream byteArrayOutputStream = new ByteArrayOutputStream();
                    DataOutputStream dataOutputStream = new DataOutputStream(byteArrayOutputStream);
                    dataOutputStream.writeUTF(jsonObject.toString());

                    event.getClientConnection().writeOutput("heater_data", byteArrayOutputStream.toByteArray());
                } else {
                    System.out.println("No data to send!");
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }

    }
}
