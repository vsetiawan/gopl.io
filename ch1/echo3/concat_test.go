package main

import (
	"strings"
	"testing"
)

var testString = []string{"testing", "testing2", "testing3", "testing4", "testing5", "testing6", "testing7", "testing8", "testing9", "testing10"}

func concat(args []string) {
	var s string
	for _, arg := range args {
		s += arg
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(testString)
	}
}

func stringJoin(args []string) {
	strings.Join(args, " ")
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoin(testString)
	}
}
