package org.pkg.apinative;

import java.io.IOException;
import java.util.*;
import java.nio.ByteBuffer;

public class Native {

  static native byte[] run(byte[] arr);

  static native long getGasForData(byte[] arr);

  public static long toLong(byte[] b) {
    ByteBuffer buffer = ByteBuffer.wrap(b);
    return buffer.getLong();
  }

  static void callback(int secret) {
    getTime();
  }

  static void getTime() {
    byte[] arr = "".getBytes();
    byte[] rarr = run(arr);
    long time = toLong(Arrays.copyOfRange(rarr, 0, 8));
    System.out.printf("Time from Go: %d\n", time);
  }

  public static void main(String[] args) {
    String fullPath = System.getProperty("user.dir") + "/go/gotimer.so";
    System.out.println("Loading shared library...");
    System.out.println(fullPath);

    System.load(fullPath);
    System.out.printf("Gas required: %d\n", getGasForData("Test".getBytes()));
    getTime();
    try {
      System.in.read();
    } catch (IOException e) {
      e.printStackTrace();
    }
  }


}
