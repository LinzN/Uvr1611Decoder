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

package de.linzn.uvr1611decoder.socket;

import de.linzn.zSocket.components.IZMask;

import java.util.concurrent.Executors;

public class SocketMask implements IZMask {
    @Override
    public void runThread(Runnable runnable) {
        Executors.newSingleThreadExecutor().submit(runnable);
    }

    @Override
    public boolean isDebugging() {
        return true;
    }
}
