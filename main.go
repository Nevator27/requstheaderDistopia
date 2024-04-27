package main

import (
	"fmt"
	"net/http"
)

func main() {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//Faz não seguir redirecionamentos
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get("https://distopia.savi2w.workers.dev/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	fmt.Println("Resposta do Header Distopia: ", resp.Header.Get("Distopia"))
}
