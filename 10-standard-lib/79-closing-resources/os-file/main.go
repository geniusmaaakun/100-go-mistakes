package main

import (
	"log"
	"os"
)

func listing1(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v\n", err)
		}
	}()

	return nil
}

func writeToFile1(filename string, content []byte) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	// 書き込みに成功した場合にはcloseのエラーを返す。名前付きエラーを使う事でできる
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = f.Write(content)
	return
}

// ディスク上に書き込まれる事が保証できないので、バッファに残ったままフラッシュされない可能性
// があるので、sync()を呼び出す
func writeToFile2(filename string, content []byte) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	// バッファに残ったままフラッシュされない可能性があるので、sync()を呼び出す
	// コミットする
	// 内容がディスクに書き込まれてからreturn する。性能は落ちる
	return f.Sync()
}
