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

import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;

public class UVRDataRequest {
    private JSONObject dataObject;
    private int tries = 0;

    public UVRDataRequest() {
        while (dataObject == null && tries < 3) {
            readGoApplication();
        }
    }

    private void readGoApplication() {

        ProcessBuilder processBuilder = new ProcessBuilder();
        File uvrjson = new File("uvrjson_armv6");
        String[] cmd = {"/bin/sh", "-c", "ifdown can0 && ifup can0 && " + uvrjson.getAbsolutePath()};

        processBuilder.command(cmd);

        try {
            Process process = processBuilder.start();

            BufferedReader reader =
                    new BufferedReader(new InputStreamReader(process.getInputStream()));

            String line;
            StringBuilder stringBuilder = new StringBuilder();

            while ((line = reader.readLine()) != null) {
                stringBuilder.append(line);
            }

            int exitCode = process.waitFor();
            if (exitCode != 0) {
                System.out.println("Error while reading can-bus. Trying again! Try:" + tries);
                tries++;
            } else {
                dataObject = new JSONObject(stringBuilder.toString());
            }

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }

    }

    public JSONObject getDataObject() {
        return dataObject;
    }
}
