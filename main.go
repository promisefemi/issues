package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	body := request(http.MethodGet, "https://api.github.com/repos/gin-gonic/gin/issues")

	fmt.Printf("%s", body)
	_ = body
}

func request(method string, url string) []byte {
	fmt.Println("Is Running")
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header = http.Header{
		"Authorization": {"token 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"},
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
