package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/cespare/xxhash/v2"
)

func main() {
	ls := []string{
		"pageviews",
		"mobile",
		"tablet",
		"laptop",
		"desktop",
	}
	seen := map[string]struct{}{
		"mobile":  {},
		"tablet":  {},
		"laptop":  {},
		"desktop": {},
	}
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".json" {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var o []string
		err = json.Unmarshal(b, &o)
		if err != nil {
			return err
		}
		for _, v := range o {
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			ls = append(ls, v)
		}
		return os.Remove(path)
	})
	if err != nil {
		log.Fatal(err)
	}
	sort.Strings(ls)
	var b bytes.Buffer
	b.WriteString(`
	// DO NOT EDIT Code generated by ua/index/main.go"
	package ua 

	var CommonIndex =map[uint64]string{
	`)
	for _, k := range ls {
		h := xxhash.Sum64String(k)
		fmt.Fprintf(&b, "%d:%q,\n", h, k)
	}
	b.WriteString("}\n")
	x, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("index.go", x, 0600)
}
