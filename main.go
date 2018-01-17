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

var version = "nightly"

type config struct {
	Text string              `toml:"text"`
	Data map[string][]string `toml:"data"`
}

func main() {
	var v bool
	var seed int64
	flag.BoolVar(&v, "v", false, "print version and exit")
	flag.Int64Var(&seed, "seed", 0, "seed PRNG with a given number")
	flag.Parse()
	log.SetFlags(0)

	if v {
		log.Printf("roll %s\n", version)
		os.Exit(0)
	}
	if seed == 0 {
		seed = time.Now().Unix()
	}
	rand.Seed(seed)

	if len(os.Args) < 2 {
		log.Fatal("Nothing to roll")
	}
	var c config
	md, err := toml.DecodeFile(os.Args[1], &c)
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
