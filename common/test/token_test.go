package main

import "testing"

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//genXid() // 23.07 ns/op
		//genKsuid() //259.6 ns/op
		//genBetterGUID() //66.33ns/op
		//genUlid() //10215 ns/op
		//genSonyflake() //5294597 ns/op
		//genSid()//247.8 ns/op
		//genUUIDv4() // 196.2 ns/op
	}
}
