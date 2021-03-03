package main

// #include <stddef.h>
// #include <stdint.h>
// #include <stdlib.h>
import "C"
import (
	"GraalVM_Go/go/tools"
	"sync"
	"time"

	"github.com/ClarkGuan/jni"
)

var (
	mtx            sync.Mutex
	callbackMethod uintptr
	clazzID        uintptr
	envID          jni.Env
	started        bool
)

func getGoByteArr(env uintptr, iarr uintptr) []byte {
	var buf []byte
	arrLen := jni.Env(env).GetArrayLength(iarr)
	for i := 0; i < arrLen; i++ {
		b := jni.Env(env).GetByteArrayElement(iarr, i)
		buf = append(buf, b)
	}
	return buf
}
func setEnv(env uintptr, clazz uintptr) {
	if envID == 0 {
		envID = jni.Env(env)
	}
	if clazzID == 0 {
		clazzID = clazz
	}
	if callbackMethod == 0 {
		callbackMethod = jni.Env(env).GetStaticMethodID(clazz, "callback", "(I)V")
		started = true
	}
}

//export Java_org_pkg_apinative_Native_getGasForData
func Java_org_pkg_apinative_Native_getGasForData(env uintptr, clazz uintptr, iarr uintptr) uint64 {
	mtx.Lock()
	defer mtx.Unlock()
	return getGasForData(getGoByteArr(env, iarr))
}

//export Java_org_pkg_apinative_Native_run
func Java_org_pkg_apinative_Native_run(env uintptr, clazz uintptr, iarr uintptr) uintptr {
	mtx.Lock()
	defer mtx.Unlock()

	setEnv(env, clazz)
	rarr := run(getGoByteArr(env, iarr))
	jarr := jni.Env(env).NewByteArray(len(rarr))
	jni.Env(env).SetByteArrayRegion(jarr, 0, rarr)
	return jarr
}
func main() {}

/*///////////////////////////////////////////////////////////////////////////////
WARNING: DON'T MODIFY UPPER PART. QA TESTER WILL GENERATE AN ERROR AFTER SUBMISSION
ONLY IMPORT SECTION CAN BE MODIFIED.
/////////////////////////////////////////////////////////////////////////////////*/

func init() {
	clazzID = 0
	callbackMethod = 0
	envID = 0
	started = false
	go sendCallback()
}

// getGasForData - Returns back gas required to execute the contract
func getGasForData(arr []byte) uint64 {
	// calculate gas here
	return uint64(5000000)
}

// run - Runs the contract, It receive data as parsed byte and returns back a parsed byte array
func run(arr []byte) []byte {
	return tools.NumToBytes(time.Now().Unix())
}

func sendCallback() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		if started {
			vm, _ := envID.GetJavaVM()
			newEnv, _ := vm.AttachCurrentThread()
			newEnv.CallStaticVoidMethodA(clazzID, callbackMethod, jni.IntValue(123456))
			vm.DetachCurrentThread()
		}
	}
}
