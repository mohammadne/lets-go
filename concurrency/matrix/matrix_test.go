package matrix

import "testing"

const (
	randomElementMin = 1
	randomElementMax = 100
)

func TestMultiplicationSequential(t *testing.T) {
	matrix1 := GenerateRandomMatrixSequential(1000, 1000, randomElementMin, randomElementMax)
	matrix2 := GenerateRandomMatrixSequential(1000, 1000, randomElementMin, randomElementMax)

	_, err := MultiplicationSequential(matrix1, matrix2)
	if err != nil {
		panic("!!!!!!!!!!!!!")
	}
}

func BenchmarkGenerateRandomMatrixSequential(b *testing.B) {
	b.Log(b.N)
	matrix := GenerateRandomMatrixSequential(b.N, b.N, randomElementMin, randomElementMax)
	if matrix == nil {
		panic("HI")
	}
}

func BenchmarkGenerateRandomMatrixConcurrent1(b *testing.B) {
	b.Log(b.N)
	matrix := GenerateRandomMatrixConcurrent1(b.N, b.N, randomElementMin, randomElementMax)
	if matrix == nil {
		panic("HI")
	}
}

func BenchmarkGenerateRandomMatrixConcurrent2(b *testing.B) {
	b.Log(b.N)
	matrix := GenerateRandomMatrixConcurrent2(b.N, b.N, randomElementMin, randomElementMax)
	if matrix == nil {
		panic("HI")
	}
}

func BenchmarkGenerateRandomMatrixConcurrent3(b *testing.B) {
	b.Log(b.N)
	matrix := GenerateRandomMatrixConcurrent3(b.N, b.N, randomElementMin, randomElementMax)
	if matrix == nil {
		panic("HI")
	}
}

func BenchmarkGenerateRandomMatrixConcurrent4(b *testing.B) {
	b.Log(b.N)
	matrix := GenerateRandomMatrixConcurrent4(b.N, b.N, randomElementMin, randomElementMax)
	if matrix == nil {
		panic("HI")
	}
}
