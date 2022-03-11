package quest

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Quest struct {
	MainCategory   string
	SubCategory    string
	SubSubCategory string
	Title          string
	Path           string
	Area           string
	Level          int
	LatestPatch    bool
	time           *time.Time
}

func Scrape(client *http.Client, subDomain string, characterId int) error {
	quests := make([]Quest, 0)
	{
		page, e := getDbPagesLen(client, subDomain)
		if e != nil {
			log.Println(e)
			return e
		}
		for i := 1; i <= page; i++ {
			fmt.Println("db page", i, "...")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			s, e := scrapeDbPage(client, subDomain, i)
			if e != nil {
				log.Println(e)
				continue
			}
			quests = append(quests, s...)
		}
	}

	completeList := make(map[string]time.Time)
	{
		page, e := getCharacterPagesLen(client, subDomain, characterId)
		if e != nil {
			log.Println(e)
			return e
		}
		for i := 1; i <= page; i++ {
			fmt.Println("character page", i, "...")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			m, e := scrapeCharacterPage(client, subDomain, characterId, i)
			if e != nil {
				log.Println(e)
				continue
			}
			for k, v := range m {
				completeList[k] = v
			}
		}
	}

	for i := 0; i < len(quests); i++ {
		if val, ok := completeList[quests[i].Path]; ok {
			quests[i].time = &val
		}
	}

	return write_to_csv(&quests)
}
