package tests

import "testing"

func TestHello(t *testing.T) {
	result := HelloWorld("Hello")
	if result != "Hello" {
		t.Fail()
	}
}

// Basic Benchmark
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Welcome John Doe")
	}
}

// Sub Benchmark
func BenchmarkSubHello(b *testing.B) {
	b.Run("Welcome John Doe", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Welcome John Doe")
		}
	})
	b.Run("Welcome John Doe Ipsum Dolor", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Welcome John Doe Ipsum Dolor")
		}
	})
}

// Benchmark Table
func BenchmarkTable(b *testing.B) {
	benchmark := []struct{
		name string
		request string
	}{
		{
			name: "John doe",
			request: "John doe",
		},
		{
			name: "John doe lorem ipsum",
			request: "John doe lorem ipsum",
		},
		{
			name: "John doe lorem ipsum dolor sit amet",
			request: "John doe lorem ipsum dolot sit amet",
		},
	}

	for _, benchmark := range benchmark {
		b.Run(benchmark.name, func (b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}