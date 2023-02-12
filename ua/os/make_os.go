package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/gernest/vince/ua"
)

func main() {
	var b bytes.Buffer

	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by ua/os/make_os.go")
	fmt.Fprintln(&b, " package ua")

	err := generate(&b, "oss.yml")
	if err != nil {
		log.Fatal(err)
	}
	r, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("ua/ua_os.go", r, 0600)
}

type OsReg struct {
	Regex   string `yaml:"regex" json:"regex"`
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

func generate(b *bytes.Buffer, path string) error {
	var items []*OsReg
	err := ua.Read(path, &items)
	if err != nil {
		return err
	}
	var s strings.Builder
	for i, d := range items {
		if i != 0 {
			s.WriteByte('|')
		}
		s.WriteString(d.Regex)
	}
	if ua.IsStdRe(s.String()) {
		fmt.Fprintf(b, " var osAllRe= MatchRe(`%s`)\n", ua.Clean(s.String()))
	} else {
		fmt.Fprintf(b, " var osAllRe= MatchRe2(`%s`)\n", ua.Clean(s.String()))
	}
	fmt.Fprintf(b, "var osAll=[]*clientRe{\n")
	var buf bytes.Buffer
	for _, d := range items {
		buf.Reset()
		r := ua.Clean(d.Regex)
		if ua.IsStdRe(d.Regex) {
			fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
		} else {
			fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
		}
		fmt.Fprintf(b, "{%s,name:%q", &buf, d.Name)
		if d.Version != "" {
			fmt.Fprintf(b, ",version:%q", d.Version)
		}
		fmt.Fprintf(b, "},\n")
	}
	fmt.Fprintln(b, "}")
	return nil
}
