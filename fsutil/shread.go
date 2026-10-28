package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Shred overwrites a regular file with three random passes and an optional final
// zero-fill pass, then removes it. It does not guarantee physical erasure.
func Shred(path string, zeroFill bool) error {
	file, err := openWritableFile(path)

	if err != nil {
		return fmt.Errorf("shred %q: %w", path, err)
	}

	err = overwriteAndClose(file, zeroFill)

	if err != nil {
		return fmt.Errorf("shred %q: %w", path, err)
	}

	err = os.Remove(path)

	if err != nil {
		return fmt.Errorf("shred %q: remove file: %w", path, err)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// openWritableFile opens a regular file without truncation, adding owner-write
// permission if necessary. The caller must prevent concurrent path replacement.
func openWritableFile(path string) (*os.File, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return nil, fmt.Errorf("inspect file: %w", err)
	}

	if !info.Mode().IsRegular() {
		return nil, errors.New("not a regular file")
	}

	file, err := os.OpenFile(path, os.O_WRONLY, 0)

	if err == nil {
		return file, nil
	}

	if !errors.Is(err, os.ErrPermission) {
		return nil, fmt.Errorf("open file: %w", err)
	}

	err = os.Chmod(path, info.Mode().Perm()|0200)

	if err != nil {
		return nil, fmt.Errorf("add owner-write permission: %w", err)
	}

	file, err = os.OpenFile(path, os.O_WRONLY, 0)

	if err != nil {
		return nil, fmt.Errorf("open file after permission change: %w", err)
	}

	return file, nil
}

// overwriteAndClose performs the overwrite passes and closes the file,
// preserving both overwrite and close errors.
func overwriteAndClose(file *os.File, zeroFill bool) error {
	overwriteErr := overwriteFile(file, zeroFill)
	closeErr := file.Close()

	if closeErr != nil {
		closeErr = fmt.Errorf("close file: %w", closeErr)
	}

	return errors.Join(overwriteErr, closeErr)
}

// overwriteFile performs three random passes and an optional final zero-fill
// pass without closing the file.
func overwriteFile(file *os.File, zeroFill bool) error {
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("inspect opened file: %w", err)
	}

	if !info.Mode().IsRegular() {
		return errors.New("opened file is not a regular file")
	}

	buffer := make([]byte, 1024*1024)

	for pass := range 3 {
		err = overwritePass(file, info.Size(), buffer, false)

		if err != nil {
			return fmt.Errorf("random pass %d: %w", pass+1, err)
		}
	}

	if zeroFill {
		err = overwritePass(file, info.Size(), buffer, true)

		if err != nil {
			return fmt.Errorf("zero-fill pass: %w", err)
		}
	}

	return nil
}

// overwritePass writes size bytes of random or zero-filled data from offset zero
// and syncs the file. The buffer must not be empty.
func overwritePass(file *os.File, size int64, buffer []byte, zeroFill bool) error {
	if len(buffer) == 0 {
		return errors.New("empty overwrite buffer")
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("seek to beginning: %w", err)
	}

	if zeroFill {
		for i := range buffer {
			buffer[i] = 0
		}
	}

	for remaining := size; remaining > 0; {
		chunk := buffer

		if remaining < int64(len(chunk)) {
			chunk = chunk[:int(remaining)]
		}

		if !zeroFill {
			if _, err := rand.Read(chunk); err != nil {
				return fmt.Errorf("generate random data: %w", err)
			}
		}

		written, err := file.Write(chunk)

		if err != nil {
			return fmt.Errorf("write data: %w", err)
		}

		if written != len(chunk) {
			return fmt.Errorf("write data: %w", io.ErrShortWrite)
		}

		remaining -= int64(written)
	}

	if err := file.Sync(); err != nil {
		return fmt.Errorf("sync file: %w", err)
	}

	return nil
}
