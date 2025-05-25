package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCopy() {
	src := bytes.NewBufferString("Test Data")
	dst := &bytes.Buffer{}

	n, hash, err := Copy(dst, src, sha256.New())

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Dest: %s (%d)\n", dst.String(), n)
	fmt.Printf("Hash: %s\n", hash)

	// Output:
	// Dest: Test Data (9)
	// Hash: bcfe67172a6f4079d69fe2f27a9960f9d62edae2fcd4bb5a606c2ebb74b3ba65
}

func ExampleFile() {
	file := "/home/bob/data.txt"
	fmt.Printf("File %s hash is %q", file, File(file, sha256.New()))
}

func ExampleBytes() {
	data := []byte("Test Data\n")

	fmt.Printf("Data hash is %q\n", Bytes(data, sha256.New()))

	// Output:
	// Data hash is "15c94dc2cf2ff71f83d0e5b9f7e7c3eb5efa8ba215dca81108cabf91f1958834"
}

func ExampleString() {
	data := "Test Data\n"

	fmt.Printf("String hash is %q\n", String(data, sha256.New()))

	// Output:
	// String hash is "15c94dc2cf2ff71f83d0e5b9f7e7c3eb5efa8ba215dca81108cabf91f1958834"
}

func ExampleSum() {
	hasher := sha256.New()
	hasher.Write([]byte("Test Data\n"))

	fmt.Println(Sum(hasher))

	// Output:
	// 15c94dc2cf2ff71f83d0e5b9f7e7c3eb5efa8ba215dca81108cabf91f1958834
}

func ExampleNewReader() {
	buf := bytes.NewBufferString("Test Data")
	r, _ := NewReader(buf, sha256.New())
	data, _ := io.ReadAll(r)

	fmt.Printf("Data: %s\n", string(data))
	fmt.Printf("Hash: %s\n", r.Sum())

	// Output:
	// Data: Test Data
	// Hash: bcfe67172a6f4079d69fe2f27a9960f9d62edae2fcd4bb5a606c2ebb74b3ba65
}

func ExampleNewWriter() {
	buf := &bytes.Buffer{}
	w, _ := NewWriter(buf, sha256.New())

	w.Write([]byte("Test Data"))

	fmt.Printf("Data: %s\n", buf.String())
	fmt.Printf("Hash: %s\n", w.Sum())

	// Output:
	// Data: Test Data
	// Hash: bcfe67172a6f4079d69fe2f27a9960f9d62edae2fcd4bb5a606c2ebb74b3ba65
}
