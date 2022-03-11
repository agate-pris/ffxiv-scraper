package login

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func Login(subDomain string, ldstSess string) (*http.Client, error) {
	url, e := url.Parse(fmt.Sprintf("https://%s.finalfantasyxiv.com", subDomain))
	if e != nil {
		return nil, e
	}
	cookie := &http.Cookie{
		Name:   "ldst_sess",
		Value:  ldstSess,
		Domain: ".finalfantasyxiv.com",
	}
	cookies := []*http.Cookie{cookie}
	jar, e := cookiejar.New(nil)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	jar.SetCookies(url, cookies)
	client := &http.Client{Jar: jar}
	return client, nil
}

func Get(client *http.Client, url string) (*http.Response, *goquery.Document, error) {
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		log.Println(e)
		return nil, nil, e
	}
	resp, e := client.Do(req)
	if e != nil {
		log.Println(e)
		return nil, nil, e
	}
	if resp.StatusCode != 200 {
		log.Println(e)
		defer func() {
			e := resp.Body.Close()
			if e != nil {
				log.Println(e)
			}
		}()
		return nil, nil, fmt.Errorf("unexpected response. status: %s", resp.Status)
	}
	doc, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		log.Println(e)
		defer func() {
			e := resp.Body.Close()
			if e != nil {
				log.Println(e)
			}
		}()
		return nil, nil, e
	}
	return resp, doc, nil
}
