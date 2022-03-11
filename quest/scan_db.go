package quest

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/agate-pris/ffxiv-scraper/login"
)

const dbUrl = "https://%s.finalfantasyxiv.com/lodestone/playguide/db/quest/?page=%d"

func getDbPagesLen(client *http.Client, subDomain string) (int, error) {
	resp, doc, e := login.Get(client, fmt.Sprintf(dbUrl, subDomain, 1))
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
	a := doc.Find(".next_all > a")
	if a == nil {
		return 1, nil
	}
	href, exists := a.First().Attr("href")
	if !exists {
		return 1, nil
	}
	url, e := url.Parse(href)
	if e != nil {
		log.Println(e)
		return 0, e
	}
	page := url.Query().Get("page")
	if page == "" {
		return 1, nil
	} else {
		page, e := strconv.Atoi(page)
		if e != nil {
			log.Println(e)
			return 0, e
		}
		return page, nil
	}
}

func scrapeDbPage(client *http.Client, subDomain string, page int) ([]Quest, error) {
	resp, doc, e := login.Get(client, fmt.Sprintf(dbUrl, subDomain, page))
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
	quests := make([]Quest, 0, 50)
	tbody := doc.Find(".db-table > tbody")
	tbody.Find("tr").Each(func(_ int, tr *goquery.Selection) {
		var quest = Quest{}
		tr.Find("td").Each(func(i int, td *goquery.Selection) {
			switch i {
			case 0:
				anchors := td.Find("a")
				switch anchors.Length() {
				case 3:
					anchors.Each(func(i int, a *goquery.Selection) {
						switch i {
						case 0:
							quest.MainCategory = a.Text()
						case 1:
							quest.SubCategory = a.Text()
						case 2:
							quest.Title = a.Text()
							href, e := a.First().Attr("href")
							if !e {
								log.Println("path not found.")
							} else {
								quest.Path = href
							}
						}
					})
				case 4:
					anchors.Each(func(i int, a *goquery.Selection) {
						switch i {
						case 0:
							quest.MainCategory = a.Text()
						case 1:
							quest.SubCategory = a.Text()
						case 2:
							quest.SubSubCategory = a.Text()
						case 3:
							quest.Title = a.Text()
							href, e := a.First().Attr("href")
							if !e {
								log.Println("path not found.")
							} else {
								quest.Path = href
							}
						}
					})
				}
				quest.LatestPatch = 0 < td.Find(".latest_patch__major__icon").Length()
			case 1:
				quest.Area = td.Text()
			case 2:
				level, e := strconv.Atoi(td.Text())
				if e != nil {
					log.Println(e)
				} else {
					quest.Level = level
				}
			}
		})
		quests = append(quests, quest)
	})
	return quests, nil
}
