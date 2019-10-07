// Copyright 2019 William Hubbs
// Copyright 2019 Sony Interactive Entertainment
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func findRepo(importPath string) string {
	repo := ""
	if strings.HasPrefix(importPath, "github.com/") {
		return repo
	}
	resp, err := http.Get("https://" + importPath + "?go-get=1")
	checkError(err)
	defer resp.Body.Close()
	imports, err := parseMetaGoImports(resp.Body, IgnoreMod)
	checkError(err)
	for _, i := range imports {
		if i.Prefix == importPath {
			repo = i.RepoRoot
			break
		}
	}
	repoUrl, err := url.Parse(repo)
	checkError(err)
	switch repoUrl.Host {
	case "go.googlesource.com":
		repo = "github.com/golang" + repoUrl.Path
	case "gopkg.in":
		host := "github.com"
		path := repoUrl.Path[:strings.LastIndex(repoUrl.Path, ".")]
		if strings.Count(path, "/") == 1 {
			pkg := strings.TrimPrefix(path, "/")
			path = "/go-" + pkg + path
		}
		repo = host + path
	default:
		repo = repoUrl.Host + repoUrl.Path
	}
	return repo
}

func printVendorInfo(s string) {
	data := strings.Fields(s)
	importPath := data[1]
	version := convertVersion(data[2])
	repo := findRepo(importPath)
	if len(repo) == 0 {
		fmt.Printf("\t\"%s %s\"\n", importPath, version)
	} else {
		fmt.Printf("\t\"%s %s %s\"\n", importPath, version, repo)
	}
}
