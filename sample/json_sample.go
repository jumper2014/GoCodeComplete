package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main()  {
	b := []byte(`{"Name":"Bob","Body1":"Pickle"}`)
	var m Message
	err := json.Unmarshal(b, &m)
	if err == nil {
		fmt.Print(m)
	}

}


