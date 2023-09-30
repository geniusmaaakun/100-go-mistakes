package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	if err := listing1(); err != nil {
		panic(err)
	}
	if err := listing2(); err != nil {
		panic(err)
	}
	if err := listing3(); err != nil {
		panic(err)
	}
}

type Event1 struct {
	ID int
	time.Time
}

func listing1() error {
	event := Event1{
		ID:   1234,
		Time: time.Now(),
	}

	//埋め込みのMarshalJSONが呼ばれるため、timeしか出力されない
	//独自のMarshalJSONを定義する必要がある。これで回避できる
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

type Event2 struct {
	ID int
	// 名前をつけることで、埋め込みではなくなる
	Time time.Time
}

func listing2() error {
	event := Event2{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

//埋め込み型のままにしておきたい場合は、独自のMarshalJSONを定義する必要がある
type Event3 struct {
	ID int
	time.Time
}

func (e Event3) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			ID   int
			Time time.Time
		}{
			ID:   e.ID,
			Time: e.Time,
		},
	)
}

func listing3() error {
	event := Event3{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
