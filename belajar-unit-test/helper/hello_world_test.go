package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// UNTUK MENJALANKAN TEST DI CLI BISA PAKAI COMMAND INI
// go test -v -run=TestNamaFunction

// UNTUK MENJALANKAN SALAH SATU SUBTEST DI CLI BISA PAKAI COMMAND INI
// go test -v -run=TestNamaFunction/NamaSubTest

// UNTUK MENJALANKAN SEMUA SUBTEST
// go test -v -run=/NamaSubTest

// before and after unit test
func TestMain(m *testing.M) {
	println("Sebelum unit test")

	m.Run()

	println("Setelah unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syakir")
	// pakai assertion lebih mudah dan ringkas, kalau pakai if bakal banyak nanti test code nya
	assert.Equal(t, "Halo wak Syakir", result, "Result must be 'Halo wak Syakir'")
	fmt.Println("Function executed")

	// if result != "Halo wak Syakir" {
	// 	// result error pakai error langsung memanggil function t.Fail()
	// 	t.Error("Result must be 'Halo wak Syakir'")
	// }
}

func TestHelloDunia(t *testing.T) {
	result := HelloWorld("Dunia")
	// kalau pakai require, kalau test nya fail langsung panggil t.FailNow()
	require.Equal(t, "Halo wak Dunia", result, "Result must be 'Halo wak Dunia'")
	fmt.Println("Function executed with require done")

	// if result != "Halo wak Sempardak" {
	// 	// result error pakai fatal langsung memanggil function t.FailNow()
	// 	t.Fatal("Result must be 'Halo wak Dunia'")
	// }

}

func TestHelloError(t *testing.T) {
	result := HelloError()
	if result != "ini pasti error" {
		// result error pakai error langsung memanggil function t.Fail()
		t.Error("Result must be 'ini pasti error'")
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skip dulu jika operating system nya windows")
	}
	result := HelloWorld("Matahari")
	require.Equal(t, "Halo wak Matahari", result, "Result must be 'Halo wak Matahari'")
	fmt.Println("Function executed with require done")
}

// test sub test ini memungkinkan kita membuat function unit test di dalam unit test
func TestSubTest(t *testing.T) {
	t.Run("Bulan", func(t *testing.T) {
		result := HelloWorld("Bulan")
		assert.Equal(t, "Halo wak", result, "Result must be 'Halo wak Bulan'")
	})
	t.Run("Blontok", func(t *testing.T) {
		result := HelloWorld("Blontok")
		require.Equal(t, "Halo wak Blontok", result, "Result must be 'Halo wak Blontok'")
	})
}

// cobain konsep table test
// konsep ini memudahkan kita untuk mengetest beberapa functions dengan cara iterasi

func TestHelloWorldTable(t *testing.T) {
	testsTable := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Bulan)",
			request:  "Bulan",
			expected: "Halo wak Bulan",
		},
		{
			name:     "HelloWorld(Matahari)",
			request:  "Matahari",
			expected: "Halo wak Matahari",
		},
		{
			name:     "HelloWorld(Dunias)",
			request:  "Dunias",
			expected: "Halo wak Dunia",
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}
