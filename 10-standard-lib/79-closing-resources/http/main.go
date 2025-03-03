package main

import (
	"io"
	"log"
	"net/http"
)

// bodyをクローズしないとリソースリークになる
func (h handler) getBody1() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (h handler) getBody2() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// ここでクローズしないとリソースリークになる
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	return string(body), nil
}

// 読み込みに関係なくクローズする
// KeepAliveが有効な場合は、必要がなくて読み込むべき
func (h handler) getStatusCode1(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	return resp.StatusCode, nil
}

func (h handler) getStatusCode2(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	// Copyを使って、読み込みを行う。コピーせずに破棄するio.Readallよりも効率的
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.StatusCode, nil
}

type handler struct {
	client http.Client
	url    string
}
