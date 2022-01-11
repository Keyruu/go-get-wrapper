package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Started the NoCoDB Wrapper for Releases DB...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}

		req, err := http.NewRequest("GET", "http://keyruu.de:8001/nc/music_m7ws/api/v1/releases", nil)
		if err != nil {
			panic(err)
		}

		req.Header.Add("xc-token", os.Getenv("NOCODB_API_TOKEN"))

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		w.Write(bytes)
	})

	port := ":8787"

	fmt.Printf("Listening on port \"%s\"...\n", port)

	http.ListenAndServe(port, nil)
}