#!/bin/bash
export JAVA_HOME=/graal/Pipeline_softwares/graalvm-ee-java11-21.0.0.2/
export PATH=$JAVA_HOME/bin:$PATH
$JAVA_HOME/bin/java -cp ./target/graalvm-native-1.0-SNAPSHOT.jar org.pkg.apinative.Native /root/GraalVM_JavaNative/nativeImpl/target/libnativeimpl.so