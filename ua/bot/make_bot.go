package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"

	"github.com/gernest/vince/ua"
)

func main() {
	err := genBot()
	if err != nil {
		log.Fatal(err)
	}
}

type Producer struct {
	Name string `yaml:"name" json:"name"`
	Url  string `yaml:"url" json:"url"`
}

type Bot struct {
	Regex    string   `yaml:"regex" json:"regex"`
	Name     string   `yaml:"name" json:"name"`
	Category string   `yaml:"category" json:"category"`
	Url      string   `yaml:"url" json:"url"`
	Producer Producer `yaml:"producer" json:"producer"`
}

func genBot() error {
	var r []*Bot
	err := ua.Read("bots.yml", &r)
	if err != nil {
		return err
	}

	var s bytes.Buffer

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "// DO NOT EDIT Code generated by ua/bot/make_bot.go")
	fmt.Fprintln(&buf, " package vince")
	for i, x := range r {
		if i != 0 {
			s.WriteByte('|')
		}
		s.WriteString(x.Regex)
	}

	if ua.IsStdRe(s.String()) {
		fmt.Fprintf(&buf, " var allBotsReStandardMatch= MustCompile(`%s`)\n", ua.Clean(s.String()))
	} else {
		fmt.Fprintf(&buf, " var allBotsReStandardMatch= MustCompile2(`%s`)\n", ua.Clean(s.String()))
	}
	fmt.Fprintln(&buf, `
	type botRe struct{
		re *ReMatch
		name     string   
		category string   
		url      string 
		producerName string  
		producerURL string  
	}
	type botResult struct {
		name     string   
		category string   
		url      string 
		producerName string  
		producerURL string  
	}
	`)
	fmt.Fprintln(&buf, "var botsReList=[]*botRe{")
	for _, m := range r {
		s.Reset()
		if ua.IsStdRe(m.Regex) {
			fmt.Fprintf(&s, "re:MatchRe(`%s`)", ua.Clean(m.Regex))
		} else {
			fmt.Fprintf(&s, "re:MatchRe2(`%s`)", ua.Clean(m.Regex))
		}
		fmt.Fprintf(&buf, "{%s, name:%q,category:%q,url:%q,producerName:%q,producerURL:%q, },\n",
			&s, m.Name, m.Category, m.Url, m.Producer.Name, m.Producer.Url,
		)
	}
	fmt.Fprintln(&buf, "}")
	f, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return os.WriteFile("ua_bots.go", f, 0600)
}
