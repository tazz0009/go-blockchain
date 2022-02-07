package main_test

import (
	"fmt"
	"math/big"
	"testing"
)

type testCase struct {
	difficulty int
	value      string
}

func TestMain_target(t *testing.T) {
	tc := testCase{16, "1000000000000000000000000000000000000000000000000000000000000"}
	target := big.NewInt(1)
	Difficulty := tc.difficulty
	target.Lsh(target, uint(256-Difficulty))

	fmt.Printf("target : %x\n", target)

	if fmt.Sprintf("%x", target) != tc.value {
		t.Error("Error!!!")
	}
}

// func TestMain_DeriveHash(t *testing.T) {
// 	tc := testCase{"Genesis", "81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5", "First Block after Genesis"}

// 	info := bytes.Join([][]byte{[]byte(tc.inputData2), []byte(tc.hashInputData1)}, []byte{})
// 	hash := sha256.Sum256(info)
// 	 = hash[:]
// }
