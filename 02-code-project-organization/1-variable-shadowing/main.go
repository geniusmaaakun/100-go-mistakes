package main

import (
	"log"
	"net/http"
)

//　変数スコープのシャドーイング

func listing1() error {
	var client *http.Client
	if tracing {
		// この変数は、このブロック内でしか使われない
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		// この変数は、このブロック内でしか使われない
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	_ = client
	return nil
}

// 解決策
func listing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		// 上書きする
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}

//解決策2
//こっちが良い
func listing3() error {
	var client *http.Client
	var err error
	if tracing {
		// =で代入
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
	} else {
		// =で代入
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

func listing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	// ここでエラーをチェックする
	// まとめて確認できる
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
