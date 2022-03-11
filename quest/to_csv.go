package quest

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func write_to_csv(quests *[]Quest) error {
	f, e := os.Create("quest.csv")
	if e != nil {
		log.Println(e)
		return e
	}
	defer func() {
		e := f.Close()
		if e != nil {
			log.Println(e)
		}
	}()
	w := csv.NewWriter(f)
	for i := 0; i < len(*quests); i++ {
		var done string
		var latestPatch string
		var time string
		if (*quests)[i].time != nil {
			done = "x"
			time = (*quests)[i].time.String()
		}
		if (*quests)[i].LatestPatch {
			latestPatch = "New"
		}
		record := []string{
			done,
			latestPatch,
			(*quests)[i].Title,
			(*quests)[i].MainCategory,
			(*quests)[i].SubCategory,
			(*quests)[i].SubSubCategory,
			(*quests)[i].Area,
			strconv.Itoa((*quests)[i].Level),
			time,
			(*quests)[i].Path,
		}
		if e := w.Write(record); e != nil {
			log.Println(e)
			continue
		}
	}
	w.Flush()
	if e := w.Error(); e != nil {
		log.Println(e)
		return e
	}
	return nil
}
