package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ///////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNew() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	r := NewReader(fd, ';')

	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", row)
	}
}

func ExampleReader_Read() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	r := NewReader(fd, ';')

	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", row)
	}
}

func ExampleReader_ReadTo() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	r := NewReader(fd, ';')
	row := make(Row, 10)

	for {
		err := r.ReadTo(row)

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", row)
	}
}

func ExampleReader_Seq() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	r := NewReader(fd, ';')

	for line, row := range r.Seq {
		fmt.Printf("%d: %#v\n", line, row)
	}
}

func ExampleReader_WithHeader() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	r := NewReader(fd, ';').WithHeader(true)

	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", row)
	}
}

func ExampleReader_Line() {
	fd, err := os.Open("file.csv")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer fd.Close()

	r := NewReader(fd, ';')

	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}

		if !strings.HasPrefix(row.Get(0), "id-") {
			fmt.Printf("Invalid value in row %d: value in cell 0 must have \"id-\" prefix", r.Line())
			return
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRow_Size() {
	r := Row{"1", "John", "Doe", "0.34"}

	fmt.Printf("Size: %d\n", r.Size())
	// Output:
	// Size: 4
}

func ExampleRow_Cells() {
	r := Row{"1", "", "", "0.34"}

	fmt.Printf("Size: %d\n", r.Size())
	fmt.Printf("Cells: %d\n", r.Cells())
	// Output:
	// Size: 4
	// Cells: 2
}

func ExampleRow_IsEmpty() {
	r1 := Row{"1", "John", "Doe", "0.34"}
	r2 := Row{"", "", "", ""}

	fmt.Printf("r1 is empty: %t\n", r1.IsEmpty())
	fmt.Printf("r2 is empty: %t\n", r2.IsEmpty())
	// Output:
	// r1 is empty: false
	// r2 is empty: true
}

func ExampleRow_Has() {
	r := Row{"1", "John", "", "0.34"}

	fmt.Printf("Has cell 1: %t\n", r.Has(1))
	fmt.Printf("Has cell 2: %t\n", r.Has(2))
	fmt.Printf("Has cell 100: %t\n", r.Has(100))
	// Output:
	// Has cell 1: true
	// Has cell 2: false
	// Has cell 100: false
}

func ExampleRow_Get() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetI() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetI8() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI8(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetI16() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI16(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetI32() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI32(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetI64() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI64(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetF() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetF32() {
	r := Row{"1", "John", "Doe", "0.34"}

	id, err := r.GetI(0)

	if err != nil {
		panic(err.Error())
	}

	balance, err := r.GetF32(3)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("First name: %s\n", r.Get(1))
	fmt.Printf("Last name: %s\n", r.Get(2))
	fmt.Printf("Balance: %g\n", balance)
	// Output:
	// ID: 1
	// First name: John
	// Last name: Doe
	// Balance: 0.34
}

func ExampleRow_GetU() {
	r := Row{"1846915341", "user@domain.com", "Yes"}

	id, err := r.GetU(0)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Email: %s\n", r.Get(1))
	fmt.Printf("Is active: %t\n", r.GetB(2))
	// Output:
	// ID: 1846915341
	// Email: user@domain.com
	// Is active: true
}

func ExampleRow_GetU8() {
	r := Row{"184", "user@domain.com", "Yes"}

	id, err := r.GetU8(0)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Email: %s\n", r.Get(1))
	fmt.Printf("Is active: %t\n", r.GetB(2))
	// Output:
	// ID: 184
	// Email: user@domain.com
	// Is active: true
}

func ExampleRow_GetU16() {
	r := Row{"18469", "user@domain.com", "Yes"}

	id, err := r.GetU16(0)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Email: %s\n", r.Get(1))
	fmt.Printf("Is active: %t\n", r.GetB(2))
	// Output:
	// ID: 18469
	// Email: user@domain.com
	// Is active: true
}

func ExampleRow_GetU32() {
	r := Row{"1846915341", "user@domain.com", "Yes"}

	id, err := r.GetU32(0)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Email: %s\n", r.Get(1))
	fmt.Printf("Is active: %t\n", r.GetB(2))
	// Output:
	// ID: 1846915341
	// Email: user@domain.com
	// Is active: true
}

func ExampleRow_GetU64() {
	r := Row{"1846915341", "user@domain.com", "Yes"}

	id, err := r.GetU64(0)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Email: %s\n", r.Get(1))
	fmt.Printf("Is active: %t\n", r.GetB(2))
	// Output:
	// ID: 1846915341
	// Email: user@domain.com
	// Is active: true
}

func ExampleRow_GetB() {
	r := Row{"1846915341", "user@domain.com", "Yes"}

	id, err := r.GetU(0)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Email: %s\n", r.Get(1))
	fmt.Printf("Is active: %t\n", r.GetB(2))
	// Output:
	// ID: 1846915341
	// Email: user@domain.com
	// Is active: true
}

func ExampleRow_ForEach() {
	r := Row{"John", "Do"}

	err := r.ForEach(func(index int, value string) error {
		if len(value) < 3 {
			return fmt.Errorf("Cell %d contains invalid value %q", index, value)
		}

		return nil
	})

	fmt.Println(err)
	// Output:
	// Cell 1 contains invalid value "Do"
}

func ExampleRow_ToString() {
	r := Row{"1", "John", "Doe", "0.34"}

	fmt.Println(r.ToString(';'))
	// Output:
	// 1;John;Doe;0.34
}

func ExampleRow_ToBytes() {
	r := Row{"1", "John", "Doe", "0.34"}

	fmt.Println(string(r.ToBytes(';')))
	// Output:
	// 1;John;Doe;0.34
}
