package tools

import (
	"encoding/binary"
	"log"
)

// Check checks for the error, logs it and sends back the bool value.
func Check(err error, info ...string) bool {
	if err != nil {
		log.Println(info, err.Error())
		return false
	}
	return true
}

// NumToBytes - converts integer to bytes, supports int32 and int64
func NumToBytes(num interface{}) []byte {
	var len int
	switch num.(type) {
	case int64:
		len = 8
		b := make([]byte, len)
		binary.BigEndian.PutUint64(b, uint64(num.(int64)))
		return b
	case int:
		len = 4
		b := make([]byte, len)
		binary.BigEndian.PutUint32(b, uint32(num.(int)))
		return b
	case int32:
		len = 4
		b := make([]byte, len)
		binary.BigEndian.PutUint32(b, uint32(num.(int32)))
		return b
	}
	return []byte{}
}

// BytesToInt - converts byte array to int
func BytesToInt(arr []byte) int {
	return int(binary.BigEndian.Uint32(arr))
}

// BytesToInt64 - converts byte array to int
func BytesToInt64(arr []byte) int64 {
	return int64(binary.BigEndian.Uint64(arr))
}
