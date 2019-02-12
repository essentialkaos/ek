// Package directio provides methods for reading/writing files with direct io
package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io"
	"os"
	"sync"
	"unsafe"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var blockPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, BLOCK_SIZE+ALIGN_SIZE)
	},
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ReadFile read file with Direct IO without buffering data in page cache
// (currently only Unix and Mac OS X is supported)
func ReadFile(file string) ([]byte, error) {
	fd, err := openFile(file, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	info, err := fd.Stat()

	if err != nil {
		return nil, err
	}

	return readData(fd, info)
}

// WriteFile write file with Direct IO without buffering data in page cache
// (currently only Unix and Mac OS X is supported)
func WriteFile(file string, data []byte, perms os.FileMode) error {
	fd, err := openFile(file, os.O_CREATE|os.O_WRONLY, perms)

	if err != nil {
		return err
	}

	defer fd.Close()

	return writeData(fd, data)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readData(fd *os.File, info os.FileInfo) ([]byte, error) {
	var buf []byte

	block := allocateBlock()
	blockSize := len(block)
	chunks := (int(info.Size()) / blockSize) + 1

	for i := 0; i < chunks; i++ {
		n, err := fd.ReadAt(block, int64(i*blockSize))

		if err != nil && err != io.EOF {
			return nil, err
		}

		buf = append(buf, block[:n]...)
	}

	freeBlock(block)

	return buf, nil
}

func writeData(fd *os.File, data []byte) error {
	block := allocateBlock()
	blockSize := len(block)
	dataSize := len(data)
	pointer := 0

	for {
		if pointer+blockSize >= dataSize {
			copy(block, data[pointer:])
		} else {
			copy(block, data[pointer:pointer+blockSize])
		}

		_, err := fd.Write(block)

		if err != nil {
			return err
		}

		pointer += blockSize

		if pointer >= dataSize {
			break
		}
	}

	freeBlock(block)

	return fd.Truncate(int64(dataSize))
}

func allocateBlock() []byte {
	block := blockPool.Get().([]byte)

	if ALIGN_SIZE == 0 {
		return block
	}

	var offset int

	alg := alignment(block, ALIGN_SIZE)

	if alg != 0 {
		offset = ALIGN_SIZE - alg
	}

	return block[offset : offset+BLOCK_SIZE]
}

func freeBlock(block []byte) {
	blockPool.Put(block)
}

func alignment(block []byte, alignment int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(alignment-1))
}
