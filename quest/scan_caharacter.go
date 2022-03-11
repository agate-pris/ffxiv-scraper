package quest

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/agate-pris/ffxiv-scraper/login"
)

const characterUrl = "https://%s.finalfantasyxiv.com/lodestone/character/%d/quest/?page=%d"

func getCharacterPagesLen(client *http.Client, subDomain string, characterId int) (int, error) {
	resp, doc, e := login.Get(client, fmt.Sprintf(characterUrl, subDomain, characterId, 1))
	if e != nil {
		log.Println(e)
		return 0, e
	}
	defer func() {
		e := resp.Body.Close()
		if e != nil {
			log.Println(e)
		}
	}()
	a := doc.Find(".btn__pager__next--all")
	if a == nil {
		return 1, nil
	}
	href, exists := a.Attr("href")
	if !exists {
		return 1, nil
	}
	url, e := url.Parse(href)
	if e != nil {
		log.Println(e)
		return 0, e
	}
	page, e := strconv.Atoi(url.Query().Get("page"))
	if e != nil {
		log.Println(e)
		return 0, e
	}
	return page, nil
}

func scrapeCharacterPage(client *http.Client, subDomain string, characterId int, page int) (map[string]time.Time, error) {
	resp, doc, e := login.Get(client, fmt.Sprintf(characterUrl, subDomain, characterId, page))
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer func() {
		e := resp.Body.Close()
		if e != nil {
			log.Println(e)
		}
	}()
	re := regexp.MustCompile(`ldst_strftime\s*\(\s*(\d+),`)
	times := make(map[string]time.Time)
	doc.Find(".entry__quest").Each(func(i int, li *goquery.Selection) {
		s := li.Find("div").First()
		if s != nil {
			path, exists := s.Attr("href")
			if !exists {
				log.Println("path not found.")
				return
			}
			script := s.Find("script").Text()
			submatched := re.FindStringSubmatch(script)
			if len(submatched) < 2 {
				log.Println("time not found.")
				return
			}
			t, e := strconv.ParseInt(submatched[1], 10, 64)
			if e != nil {
				log.Println(e)
				return
			}
			times[path] = time.Unix(t, 0)
		}
	})
	return times, nil
}
