package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkEchoMethod1(b *testing.B) {
	// Simulate command-line arguments
	os.Args = []string{"cmd", "arg1", "arg2", "arg3", "arg4", "arg5"}

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
}

func BenchmarkEchoMethod2(b *testing.B) {
	// Simulate command-line arguments
	os.Args = []string{"cmd", "arg1", "arg2", "arg3", "arg4", "arg5"}

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		var s, sep string
		for _, arg := range os.Args {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkEchoMethod3(b *testing.B) {
	// Simulate command-line arguments
	os.Args = []string{"cmd", "arg1", "arg2", "arg3", "arg4", "arg5"}

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = (strings.Join(os.Args[1:], " "))
	}
}
