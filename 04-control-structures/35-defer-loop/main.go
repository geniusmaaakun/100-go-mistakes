package main

import "os"

// よくない例
func readFiles1(ch <-chan string) error {
	// readfilesが終了するまでcloseされないので、無限にループする場合はリークする
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		// Do something with file
	}
	return nil
}

// 改善例
func readFiles2(ch <-chan string) error {
	for path := range ch {
		if err := readFile(path); err != nil {
			return err
		}
	}
	return nil
}

// ファイルを読み込む関数を切り出す
// 結果的には都度closeするのとは変わらない
func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	// Do something with file
	return nil
}

// クロージャにする例
func readFiles3(ch <-chan string) error {
	for path := range ch {
		// この関数を実行後にcloseする
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()

			// Do something with file
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
