package main

import (
	"os"
	"strings"
)

func init() {
	os.Setenv("CANDIDATES", "HELLO,WORLD,QUITE,FANCY,FRESH,PANIC,CRAZY,BUGGY,SCARE")
}

func main() {
	candidate_env := os.Getenv("CANDIDATES")
	CANDIDATES := strings.Split(candidate_env, ",")

	println("CANDIDATES: ", strings.Join(CANDIDATES, ";"))
}
