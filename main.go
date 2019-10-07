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
	fileNamePtr := flag.String("f", "vendor/modules.txt", "modules.txt file")
	flag.Parse()

	file, err := os.Open(*fileNamePtr)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("EGO_VENDOR=(")
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.HasPrefix(s, "# ") {
			continue
		}
		printVendorInfo(s)
	}
	checkError(scanner.Err())
	fmt.Println(")")
}
