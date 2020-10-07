package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/polynomialspace/pprofui"
)

func defaulthandler(w http.ResponseWriter, r *http.Request) {
	loremIpsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fakeLoad := make([]string, 0)
			for _, word := range strings.Split(loremIpsum, " ") {
				fakeLoad = append(fakeLoad, word)
			}
			_ = fakeLoad
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Fprintln(w, loremIpsum)
}

func main() {
	http.HandleFunc("/", defaulthandler)

	go func() {
		storage := pprofui.NewMemStorage(1, 0)
		debugsrv := pprofui.NewServer(storage, nil)
		debugmux := http.NewServeMux()
		debugmux.HandleFunc("/", debugsrv.ServeHTTP)
		log.Println("pprofUI on http://localhost:8081/")
		log.Fatalln(http.ListenAndServe("localhost:8081", debugmux))
	}()

	log.Println("Listening on http://localhost:8080/")
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
