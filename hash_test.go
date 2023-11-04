package pikpakhash

import "testing"

func BenchmarkHashBuf128K(b *testing.B) {
	ph := NewPikPakHash(DefaultCalcRoutines, 1<<17)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}

func BenchmarkHashBuf256K(b *testing.B) {
	ph := NewPikPakHash(DefaultCalcRoutines, 1<<18)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}
func BenchmarkHashBuf512K(b *testing.B) {
	ph := NewPikPakHash(DefaultCalcRoutines, 1<<19)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}

func BenchmarkHashBuf1M(b *testing.B) {
	ph := NewPikPakHash(DefaultCalcRoutines, 1<<20)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}

func BenchmarkHashRoutine8(b *testing.B) {
	ph := NewPikPakHash(8, DefaultBufferSize)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}

func BenchmarkHashRoutine12(b *testing.B) {
	ph := NewPikPakHash(12, DefaultBufferSize)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}
func BenchmarkHashRoutine16(b *testing.B) {
	ph := NewPikPakHash(16, DefaultBufferSize)
	for i := 0; i < b.N; i++ {
		ph.HashFromPath("./b")
	}
}
