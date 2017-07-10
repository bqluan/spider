package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type pages []string

func (p *pages) String() string {
	return fmt.Sprintf("%d pages", len(*p))
}

func (p *pages) Set(v string) error {
	*p = append(*p, v)
	return nil
}

var pgs pages

func main() {
	rand.Seed(time.Now().Unix())

	flag.Var(&pgs, "page", "Web page url.")
	flag.Parse()

	if len(pgs) == 0 {
		fmt.Fprintln(os.Stderr, "error: missing required flag -page")
		flag.Usage()
		os.Exit(1)
	}

	for {
		p := pgs[rand.Intn(len(pgs))]

		resp, err := http.Get(p)
		if err != nil {
			log.Printf("%s %s", p, err)
			continue
		}
		resp.Body.Close()
		log.Printf("%s %d", p, resp.StatusCode)

		time.Sleep(time.Duration(3+rand.Intn(7)) * time.Second)
	}
}
