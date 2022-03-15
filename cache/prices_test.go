package cache

import (
	"net/http"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	PrintPrices()
	AppendPrice("btc", 1)
	PrintPrices()
	AppendPrice("btc", 2)
	PrintPrices()
	AppendPrice("btc", 3)
	PrintPrices()
	AppendPrice("btc", 4)
	PrintPrices()
	AppendPrice("btc", 5)
	PrintPrices()
	AppendPrice("btc", 6)
	PrintPrices()
	AppendPrice("btc", 7)
	PrintPrices()

	// s := make([]int, 0, 3)

	// fmt.Println(s, len(s), cap(s))
	// fmt.Printf("%p\n", s)

	// s = append(s, 1)

	// fmt.Println(s, len(s), cap(s))
	// fmt.Printf("%p\n", s)

	// s = append(s, 2)

	// fmt.Println(s, len(s), cap(s))
	// fmt.Printf("%p\n", s)

	// s = append(s[1:], 3)

	// fmt.Println(s, len(s), cap(s))
	// fmt.Printf("%p\n", s)

	// s = append(s[1:], 4)

	// fmt.Println(s, len(s), cap(s))
	// fmt.Printf("%p\n", s)

	// s = append(s[1:], 5)

	// fmt.Println(s, len(s), cap(s))
	// fmt.Printf("%p\n", s)
}

func TestServe(t *testing.T) {
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 1)
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 2)
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 3)
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 4)
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 5)
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 6)
	// AppendPrice("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", 6)

	r, err := http.NewRequest("GET", "/prices?s=btc&n=2", nil)
	if err != nil {
		t.Error(err)
	}

	ServeHTTP(os.Stdout, r)
}
