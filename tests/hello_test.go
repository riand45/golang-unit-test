package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHello(t *testing.T) {
	result := HelloFirst()
	if result != "Hello First" {
		t.Fail()
	}
}

//Sub Test
func TestSubTest(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result := HelloWorld("John")
		require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
	})
	t.Run("John Doe", func(t *testing.T) {
		result := HelloWorld("John Doe")
		require.Equal(t, "Hello John Doe", result, "Result must be 'Hello John Doe'")
	})
	t.Run("Lorem Ipsum", func(t *testing.T) {
		result := HelloWorld("Lorem Ipsum")
		require.Equal(t, "Hello Lorem Ipsum", result, "Result must be 'Hello Lorem Ipsum'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("World")
	require.Equal(t, "Hello World", result, "Result must be 'Hello World'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("John")
	require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Doe")
	assert.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "World",
			request:  "World",
			expected: "Hello World",
		},
		{
			name:     "John",
			request:  "John",
			expected: "Hello John",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
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
			request: "John doe lorem ipsum dolor sit amet",
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