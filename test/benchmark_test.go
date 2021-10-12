package test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	count := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", count)
	}
}

func BenchmarkFormat(b *testing.B) {
	count := int64(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(count, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	count := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(count)
	}
}
