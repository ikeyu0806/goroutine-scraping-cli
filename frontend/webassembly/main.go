package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func fetchContent(this js.Value, p []js.Value) interface{} {
	url := "https://jsonplaceholder.typicode.com/photos"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	content := string(body)
	js.Global().Get("console").Call("log", "Fetched content:", content)
	return nil
}

func registerCallbacks() {
	js.Global().Set("fetchContent", js.FuncOf(fetchContent))
}

func main() {
	c := make(chan struct{}, 0)

	registerCallbacks()

	<-c
}
