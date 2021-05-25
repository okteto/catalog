package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://api:8080/data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("non-200 api response of %d", resp.StatusCode), http.StatusInternalServerError)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintln(w, string(body))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
