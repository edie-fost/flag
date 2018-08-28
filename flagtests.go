// comments here

package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var sdf, vdf string
	fmt.Println("Syntax: cmd -v=???  -p= ??? mm-yyy")
	flag.StringVar(&sdf, "p", "", "p")
	flag.StringVar(&vdf, "v", "", "v")
	flag.Parse()
	s := flag.Args()
	fmt.Println("num: ", flag.NArg(), "args: ", s)
	parsedates(s)
	log.Println("s:", s, "Len: ", len(s), "N: ", flag.NArg())
}

func parsedates(t []string) {
	log.Println("t:", t, "Len: ", len(t))
	log.Println("first ", t[0], t[1])
}
