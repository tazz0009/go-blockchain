package main_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	main "github.com/tazz0009/go-blockchain"
)

type testCase struct {
	inputData1     string
	hashInputData1 string
	inputData2     string
}

func TestMain_CreateBlock(t *testing.T) {

	tc := testCase{"Genesis", "81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5", "First Block after Genesis"}
	block := main.CreateBlock(tc.inputData1, []byte{})

	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	fmt.Printf("Data in Block: %x\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)

	if tc.hashInputData1 != fmt.Sprintf("%x", block.Hash) {
		t.Error("Wrong result")
	}

	decodedByteArray, err := hex.DecodeString(tc.hashInputData1)

	if err != nil {
		fmt.Println("Unable to convert hex to byte. ", err)
	}

	fmt.Println(block.Hash)
	fmt.Println(decodedByteArray)

	encodedString := hex.EncodeToString(block.Hash)

	fmt.Println(tc.hashInputData1)
	fmt.Println(encodedString)
}

// func TestMain_DeriveHash(t *testing.T) {
// 	tc := testCase{"Genesis", "81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5", "First Block after Genesis"}

// 	info := bytes.Join([][]byte{[]byte(tc.inputData2), []byte(tc.hashInputData1)}, []byte{})
// 	hash := sha256.Sum256(info)
// 	 = hash[:]
// }
