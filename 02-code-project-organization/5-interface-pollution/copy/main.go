package main

import "io"

// 　テストが簡単
func copySourceToDest(source io.Reader, dest io.Writer) error {
	b, err := io.ReadAll(source)
	if err != nil {
		return err
	}
	_, err = dest.Write(b)
	return err
}
