package hscan

import (
	"testing"
)

func BenchmarkMd5(b *testing.B) {
	CreateHashFiles("dump.txt", 1)
}

func TestMd5(t *testing.T) {
	CreateHashFiles("dump.txt", 4)
}
