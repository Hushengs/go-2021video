package split

import (
	"reflect"
	"testing"
)

func TestSplitV1(t *testing.T) {
	got := SplitV1("babc", "b")
	want := []string{"", "a", "c"}
	if !reflect.DeepEqual(got, want) {
		//测试失败
		t.Errorf("want:%v but got:%v", want, got)
	}
}

func TestSplitV2(t *testing.T) {
	got := SplitV1("babc", "b")
	want := []string{"", "a", "c"}
	if !reflect.DeepEqual(got, want) {
		//测试失败
		t.Errorf("want:%v but got:%v", want, got)
	}
}

//基准测试SplitV1
func BenchmarkSplitV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitV1("babc", "b")
	}
}

//基准测试SplitV2
// go test -bench=SplitV2
func BenchmarkSplitV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitV2("babc", "b")
	}
}

//性能比较测试
//go test -bench=Fib
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}
