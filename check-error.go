// Copyright 2019 William Hubbs
// Copyright 2019 Sony Interactive Entertainment
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
