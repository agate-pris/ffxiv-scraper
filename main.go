package main

import (
	"log"

	"github.com/agate-pris/ffxiv-scraper/login"
	"github.com/agate-pris/ffxiv-scraper/quest"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Language  string `short:"l" long:"lang" alias:"language" description:"Specify subdomain for lodestone" default:"jp"`
		Session   string `short:"s" long:"session" description:"Specify session (ldst_sess)" required:"true"`
		Character int    `short:"c" long:"character" description:"Specify character id" required:"true"`
	}
	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "-s <ldst_sess> -c <character-id>"
	_, e := parser.Parse()
	if e != nil {
		if fe, ok := e.(*flags.Error); ok && fe.Type == flags.ErrHelp {
			return
		}
		log.Println(e)
		return
	}

	client, e := login.Login(opts.Language, opts.Session)
	if e != nil {
		log.Println(e)
		return
	}
	e = quest.Scrape(client, opts.Language, opts.Character)
	if e != nil {
		log.Println(e)
		return
	}
}
