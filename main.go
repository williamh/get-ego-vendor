// Copyright 2019 William Hubbs
// Copyright 2019 Sony Interactive Entertainment
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileNamePtr := flag.String("f", "go.sum", "go.sum file")
	flag.Parse()

	file, err := os.Open(*fileNamePtr)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("EGO_SUM=(")
	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Fields(s)
		mod := line[0]
		version := line [1]
		fmt.Printf("\t\"%s %s\"\n", mod, version)
	}
	fmt.Println(")")
	fmt.Println("go-module_set_globals")
	checkError(scanner.Err())
}
