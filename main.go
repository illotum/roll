package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/illotum/roll/table"
)

type config struct {
	Text string              `toml:"text"`
	Data map[string][]string `toml:"data"`
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	if len(os.Args) < 2 {
		log.Fatal("Nothing to roll")
	}

	fname := os.Args[1]
	var c config
	md, err := toml.DecodeFile(fname, &c)
	if err != nil {
		log.Fatal(err)
	}

	if !md.IsDefined("text") {
		log.Fatal("'text' string is required in a table")
	}

	templ, err := parseTemplate(c.Text)
	if err != nil {
		log.Fatal(err)
	}
	data, err := parseTables(c.Data)
	if err != nil {
		log.Fatal(err)
	}

	err = templ.Execute(os.Stdout, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func parseTables(ms map[string][]string) (map[string]table.Table, error) {
	ts := make(map[string]table.Table)
	for k, ss := range ms {
		t := table.New(ss)
		ts[k] = t
	}
	return ts, nil
}
