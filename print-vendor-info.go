// Copyright 2019 William Hubbs
// Copyright 2019 Sony Interactive Entertainment
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
)

func findRepo(importPath string, required bool) string {
	repo := ""
	resp, err := http.Get("https://" + importPath + "?go-get=1")
	checkError(err)
	defer resp.Body.Close()

	// If we get a 404 then we check to see if we have a versioned uri
	// i.e. github.com/minio/gokrb5/v7
	// Then we try to check for the v[0-9] - if so then we rerun with
	// the version removed

	if resp.StatusCode == 404 {
		base := path.Base(importPath)
		if regexp.MustCompile(`^v[0-9]`).MatchString(base) {
			return findRepo(path.Dir(importPath), true)
		}
	}

	if strings.HasPrefix(importPath, "github.com/") {
		if required {
			return importPath
		} else {
			return repo
		}
	}

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
	repo := findRepo(importPath, false)
	if len(repo) == 0 {
		fmt.Printf("\t\"%s %s\"\n", importPath, version)
	} else {
		fmt.Printf("\t\"%s %s %s\"\n", importPath, version, repo)
	}
}
