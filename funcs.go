package main

import "math/rand"

func rollTable(t table) string {
	r := rand.Intn(t.total)
	for i := range t.ws {
		r -= t.ws[i]
		if r < 0 {
			return t.ls[i]
		}
	}
	return "ERROR"
}
