package main

import (
	"flag"
	"log"
)

var (
	dir = flag.String("dir", ".", "Root directory of image repository.")
)

func main() {
	flag.Parse()
	urls := flag.Args()
	log.Printf("Initializing an image repository with root: %s\n", *dir)
	r := NewRepository(*dir)
	errors := r.AddImages(urls)
	for e := range errors {
		log.Println(e)
	}
}
