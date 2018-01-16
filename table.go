package main

import (
	"fmt"
	"strings"
	"text/template"
)

type config struct {
	Text string              `toml:"text"`
	Data map[string][]string `toml:"tables"`
}

type table struct {
	ls    []string
	ws    []int
	total int
}

func parseTemplate(text string) (*template.Template, error) {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"random": rollTable,
	})
	return t.Parse(text)
}

func parseTable(ss []string) table {
	var (
		ls       []string
		l        string
		ws       []int
		w, total int
	)
	for _, s := range ss {
		_, err := fmt.Sscanf(s, "%d:%s", &w, &l)
		if err != nil {
			l = s
			w = 1
		}
		ls = append(ls, strings.TrimSpace(l))
		ws = append(ws, w)
		total += w
	}

	return table{
		ls:    ls,
		ws:    ws,
		total: total,
	}
}

func parseTables(ms map[string][]string) (map[string]table, error) {
	ts := make(map[string]table)
	for k, ss := range ms {
		t := parseTable(ss)
		ts[k] = t
	}
	return ts, nil
}
