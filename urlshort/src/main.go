package main

import (
	"fmt"
	"net/http"

	"urlshort"
	"flag"
	"io/ioutil"
)

func readArguments() string{
	filename1 := flag.String("yaml-filename","redirect.yaml","input yaml file name")
	flag.Parse()
	return *filename1
}
func main() {
	var yamlfilename string;
	yamlfilename = readArguments()
	yamlcontent,err := ioutil.ReadFile(yamlfilename)
	if err != nil{
		fmt.Println("erro opeing yaml file")
	}


	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlcontent), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
