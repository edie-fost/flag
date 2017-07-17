package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var sdf string
	fmt.Println("Syntax: cmd -sdf=??? mm-yyy")
	flag.StringVar(&sdf, "sdf", "", "sdf")
	flag.Parse()
	s := flag.Args()
	parsedates(s)
	log.Println("s:", s, "Len: ", len(s), "N: ", flag.NArg())
}

func parsedates(t []string) {
	log.Println("t:", t, "Len: ", len(t))
	log.Println("first ", t[0], t[1])
}
