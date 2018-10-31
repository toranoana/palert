package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	flag.Parse()

	slackUrl := flag.Arg(0)
	message := flag.Arg(1)

	response, err := http.PostForm(slackUrl,
		url.Values{
			//"payload": {`{"text:"` + message + `"}`},
			"payload": { `{"text": "` + message + `"}` },
		})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", response.Status)
}