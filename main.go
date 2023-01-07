package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"image/png"

	"golang.org/x/image/webp"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("starting %s\n", arg)
		if filepath.Ext(arg) == ".webp" {
			f, err := os.Open(arg)
			if err != nil {
				f.Close()
				fmt.Printf("%s couldn't be opened, skipping (%s)", arg, err.Error())
				continue
			}

			r, err := os.Create(arg + ".png")
			if err != nil {
				f.Close()
				r.Close()
				fmt.Printf("%s.png couldn't be created, skipping (%s)", arg, err.Error())
				continue
			}

			webpIn, err := webp.Decode(f)
			if err != nil {
				f.Close()
				r.Close()
				log.Fatalf("%s couldn't be decoded (%s)", arg, err.Error())
				continue
			}

			err = png.Encode(r, webpIn)
			if err != nil {
				f.Close()
				r.Close()
				log.Fatalf("%s.png couldn't be encoded (%s)", arg, err.Error())
				continue
			}

			f.Close()
			r.Close()
		}
	}
	fmt.Println("job completed")
}
