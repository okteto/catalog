package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/justinas/alice"
	"github.com/okteto/divert"
)

func main() {
	chain := alice.New(divert.InjectDivertHeader())

	http.Handle("/", chain.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(http.MethodGet, "http://api:8080/data", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.Header.Set(divert.DivertHeaderName, divert.FromContext(r.Context()))

		client := http.Client{}
		resp, err := client.Do(req)
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
	}))

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
