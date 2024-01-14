package main

import "testing"

var args = []string{"hello", "world"}
var sep = " "

func BenchmarkEchoJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoJoin(args, sep)
	}
}

func BenchmarkEchoConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoConcat(args, sep)
	}
}
