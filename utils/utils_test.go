package utils

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/panshiqu/analyst/define"
)

func TestSymbol2Address(t *testing.T) {
	fmt.Println(Symbol2Address("btc"))
	fmt.Println(Symbol2Address("eth"))
	fmt.Println(Symbol2Address("matic"))
	fmt.Println(Symbol2Address("uni"))
	fmt.Println(Symbol2Address("0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6"))
	fmt.Println(Symbol2Address("0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6,8"))
	fmt.Println(Symbol2Address(""))
	fmt.Println(Symbol2Address("0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6,A"))
}

func TestAddress2Symbol(t *testing.T) {
	fmt.Println(Address2Symbol("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"))
	fmt.Println(Address2Symbol("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"))
	fmt.Println(Address2Symbol("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"))
	fmt.Println(Address2Symbol("0xb33eaad8d922b1083446dc23f610c2567fb5180f"))
}

func TestIsDisableNotification(t *testing.T) {
	fmt.Println(IsDisableNotification(""))
	fmt.Println(IsDisableNotification("-"))
	fmt.Println(IsDisableNotification("8-"))
	fmt.Println(IsDisableNotification("8-a"))
	fmt.Println(IsDisableNotification("8-13"))
	fmt.Println(IsDisableNotification("8-13,15-22"))
}

func TestName2Address(t *testing.T) {
	fmt.Println(Name2Address("Maer"))
	fmt.Println(Name2Address("Panshi"))
	fmt.Println(Name2Address("Zhuge"))
	fmt.Println(Name2Address("0xa67153e17bb2f4b51b127c3dd3869b7bc3e256C1"))
}

func TestTransferString(t *testing.T) {
	t1 := &define.Transfer{
		TokenSymbol:  "USDC",
		TokenDecimal: 6,
		Value:        "123456",
	}
	fmt.Println(t1)

	t2 := &define.Transfer{
		TokenSymbol:  "USDC",
		TokenDecimal: 6,
		Value:        "12345",
	}
	fmt.Println(t2)

	t3 := &define.Transfer{
		TokenSymbol:  "USDC",
		TokenDecimal: 6,
		Value:        "1234",
	}
	fmt.Println(t3)

	t4 := &define.Transfer{
		TokenSymbol:  "USDC",
		TokenDecimal: 6,
		Value:        "1234567",
	}
	fmt.Println(t4)

	t5 := &define.Transfer{
		TokenSymbol:  "USDC",
		TokenDecimal: 6,
		Value:        "12345678",
	}
	fmt.Println(t5)

	t6 := &define.Transfer{
		TokenSymbol:  "USDC",
		TokenDecimal: 6,
		Value:        "12340000",
	}
	fmt.Println(t6)
}

func TestMapAdd(t *testing.T) {
	m := make(map[string]*big.Int)
	MapAdd(m, "a", big.NewInt(1))
	fmt.Println(m)
	MapAdd(m, "a", big.NewInt(1))
	fmt.Println(m)
	MapAdd(m, "b", big.NewInt(1))
	fmt.Println(m)
}
