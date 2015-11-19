package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"
)

type Hello struct{}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func(s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for i:= 0; i < 100*100; i++ {
		fmt.Fprint(w, s)
	}
}

func (s* Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w, s)
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello!\n")

	//time.Sleep(3000 * time.Millisecond)

	//fmt.Fprint(w, "There\n")
	for i := 0; i < 100*100; i++ {
		fmt.Fprint(w, i)
		fmt.Fprint(w, ",")
	}
}

func main() {
	http.Handle("/string", String("App, Lol\n"))
	http.Handle("/struct", &Struct{"Barbossa", ":", "Haul your wind and hold your water!"})
	http.Handle("/hello", Hello{})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
