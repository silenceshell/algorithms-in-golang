package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
	"os"
)

func main() {
	x := utils.GenerateSliceInt32(13)
	fmt.Println(x)

	radixSort(x)
	fmt.Println(x)
}

const digit int = 4
const maxbit = -1 << 31

func radixSort(x []int32) {
	buf := bytes.NewBuffer(nil)
	asBytesArray := make([][]byte, len(x))
	var (
		err error
		n   int
	)
	for i, e := range x {
		err = binary.Write(buf, binary.LittleEndian, e^maxbit)
		if err != nil {
			fmt.Printf("write error, %v", err)
			os.Exit(-1)
		}

		asBytes := make([]byte, digit)
		_, err = buf.Read(asBytes)
		if err != nil {
			fmt.Printf("read error, err %v, n %v", err, n)
			os.Exit(-1)
		}
		asBytesArray[i] = asBytes
	}

	var length int
	buckets := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, v := range asBytesArray {
			buckets[v[i]] = append(buckets[v[i]], v)
		}
		length = 0
		for k, bucket := range buckets {
			copy(asBytesArray[length:], bucket)
			length += len(bucket)
			buckets[k] = bucket[:0]
		}
	}
	var w int32
	for i, asBytes := range asBytesArray {
		_, err = buf.Write(asBytes)
		if err != nil {
			fmt.Printf("write error, err %v", err)
			os.Exit(-1)
		}
		err = binary.Read(buf, binary.LittleEndian, &w)
		if err != nil {
			fmt.Printf("read error, err %v", err)
			os.Exit(-1)
		}
		x[i] = w ^ maxbit
	}
}
