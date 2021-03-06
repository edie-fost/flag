// comments here would be a good idea

package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var sdf, vdf string
	fmt.Printf("Testing flags and args.\n")
	fmt.Println("Syntax: cmd v=???  p= ??? mm-yyy")
	flag.StringVar(&sdf, "p", "", "p")
	flag.StringVar(&vdf, "v", "", "v")
	flag.Parse()
	s := flag.Args()
	if len(s) > 0 {
		fmt.Println("num: ", flag.NArg(), "args: ", s)
		parsedates(s)
		log.Println("s:", s, "Len: ", len(s), "N: ", flag.NArg())
	}
}

// parse comments
func parsedates(t []string) {
	log.Println("t:", t, "Len: ", len(t))
	log.Println("first ", t[0], t[1])
}
