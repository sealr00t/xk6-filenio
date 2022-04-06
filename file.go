// Copyright 2022 sealr00t
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE.txt file.

package filenio

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/file", new(FILENIO))
}

// FILE is the k6 extension
type FILENIO struct{}

// Write string to file
func (*FILENIO) WriteString(path string, s string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(s); err != nil {
		return err
	}
	return nil
}

// Append string to file
func (*FILENIO) AppendString(path string, s string) error {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(s); err != nil {
		return err
	}
	return nil
}

func removeLines(fn string, start, n int) (err error) {
	if start < 1 {
		return errors.New("invalid request.  line numbers start at 1.")
	}
	if n < 0 {
		return errors.New("invalid request.  negative number to remove.")
	}
	var f *os.File
	if f, err = os.OpenFile(fn, os.O_RDWR, 0); err != nil {
		return
	}
	defer func() {
		if cErr := f.Close(); err == nil {
			err = cErr
		}
	}()
	var b []byte
	if b, err = ioutil.ReadAll(f); err != nil {
		return
	}
	cut, ok := skip(b, start-1)
	if !ok {
		return fmt.Errorf("less than %d lines", start)
	}
	if n == 0 {
		return nil
	}
	tail, ok := skip(cut, n)
	if !ok {
		return fmt.Errorf("less than %d lines after line %d", n, start)
	}
	t := int64(len(b) - len(cut))
	if err = f.Truncate(t); err != nil {
		return
	}
	if len(tail) > 0 {
		_, err = f.WriteAt(tail, t)
	}
	return
}

func skip(b []byte, n int) ([]byte, bool) {
	for ; n > 0; n-- {
		if len(b) == 0 {
			return nil, false
		}
		x := bytes.IndexByte(b, '\n')
		if x < 0 {
			x = len(b)
		} else {
			x++
		}
		b = b[x:]
	}
	return b, true
}
