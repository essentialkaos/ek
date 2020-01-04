package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleReader_Read() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	reader := NewReader(fd)
	reader.Comma = ','

	for {
		data, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", data)
	}
}

func ExampleReader_ReadTo() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	reader := NewReader(fd)
	reader.Comma = ','

	data := make([]string, 10)

	for {
		err := reader.ReadTo(data)

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", data)
	}
}
