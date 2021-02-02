package main

import (
	"flag"
	"log"
	"os"

	"github.com/mdnsService/go-airplay"
)

var opts struct {
	position     float64
	showHelpFlag bool
}

func init() {
	flag.Float64Var(&opts.position, "p", 0.0, "Number of seconds to move (second)")
	flag.BoolVar(&opts.showHelpFlag, "h", false, "Show this message")
	flag.Parse()

	if opts.showHelpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if opts.position < 0 {
		log.Fatal("options: position should not negative")
	}
}

func main() {
	client, _ := airplay.NewClient(&airplay.ClientParam{
		Addr:     "",
		Port:     0,
		Password: "",
	})
	client.Scrub(opts.position)
}
