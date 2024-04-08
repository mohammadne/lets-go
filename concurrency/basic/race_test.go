package basic

import "testing"

func TestIncementWithRace(t *testing.T) {
	resultShouldBe := 50000
	result := incementWithRace(resultShouldBe, 50)

	if result != uint64(resultShouldBe) {
		t.Fatal("invalid result")
	}
}

func BenchmarkIncementWithRace(b *testing.B) {
	b.Log(b.N)
	incementWithRace(b.N, b.N/1000+1)
}

func TestIncementWithoutRace1(t *testing.T) {
	resultShouldBe := 50000
	result := incementWithoutRace1(resultShouldBe, 50)

	if result != uint64(resultShouldBe) {
		t.Fatal("invalid result")
	}
}

func BenchmarkIncementWithoutRace1(b *testing.B) {
	b.Log(b.N)
	incementWithoutRace1(b.N, b.N/1000+1)
}

func TestIncementWithoutRace2(t *testing.T) {
	resultShouldBe := 50000
	result := incementWithoutRace2(resultShouldBe, 50)

	if result != uint64(resultShouldBe) {
		t.Fatal("invalid result")
	}
}

func BenchmarkIncementWithoutRace2(b *testing.B) {
	b.Log(b.N)
	incementWithoutRace2(b.N, b.N/1000+1)
}
