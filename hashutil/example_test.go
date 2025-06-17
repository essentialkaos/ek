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

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleHash_String() {
	h := String("Test1234!", sha256.New())

	fmt.Printf("Hash: %s\n", h.String())

	// Output:
	// Hash: 0fadf52a4580cfebb99e61162139af3d3a6403c1d36b83e4962b721d1c8cbd0b
}

func ExampleHash_Bytes() {
	h := String("Test1234!", sha256.New())

	fmt.Printf("Hash: %v\n", h.Bytes())

	// Output:
	// Hash: [15 173 245 42 69 128 207 235 185 158 97 22 33 57 175 61 58 100 3 193 211 107 131 228 150 43 114 29 28 140 189 11]
}

func ExampleHash_Equal() {
	h1 := String("Test1234!", sha256.New())
	h2 := String("Test1234!", sha256.New())
	h3 := String("1234!Test", sha256.New())

	fmt.Printf("h1 == h2: %t\n", h1.Equal(h2))
	fmt.Printf("h2 == h3: %t\n", h2.Equal(h3))

	// Output:
	// h1 == h2: true
	// h2 == h3: false
}

func ExampleHash_EqualString() {
	h1 := String("Test1234!", sha256.New())
	h2 := "0fadf52a4580cfebb99e61162139af3d3a6403c1d36b83e4962b721d1c8cbd0b"

	fmt.Printf("h1 == h2: %t\n", h1.EqualString(h2))

	// Output:
	// h1 == h2: true
}

func ExampleHash_IsEmpty() {
	h := String("Test1234!", sha256.New())

	fmt.Printf("Is empty: %t\n", h.IsEmpty())

	// Output:
	// Is empty: false
}
