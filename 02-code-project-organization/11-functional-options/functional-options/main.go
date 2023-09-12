package main

import (
	"errors"
	"net/http"
)

//関数オプションパターン

const defaultHTTPPort = 8080

type options struct {
	port *int
}

type Option func(options *options) error

//慣例的にWithをつける
func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

//空のオプションを渡す必要はない
func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	// オプションを順番に実行し、設定する
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	// At this stage, the options struct is built and contains the config
	// Therefore, we can implement our logic related to port configuration
	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	_ = port
	return nil, nil
}

func client() {
	// _, _ = NewServer("localhost", WithPort(8080))

	// 複数のオプションを渡すこともできる
	_, _ = NewServer("localhost", WithPort(8080), WithPort(8080))

}

func randomPort() int {
	return 4 // Chosen by fair dice roll, guaranteed to be random.
}

func main() {
	client()
}
