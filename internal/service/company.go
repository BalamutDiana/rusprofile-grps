package company

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ContentsData struct {
	UL []Info `json:"ul"`
}

type Info struct {
	Name    string `json:"name"`
	Link    string `json:"link"`
	INN     string `json:"inn"`
	CeoName string `json:"ceo_name"`
}

func GetMainInfo(inn string) {
	rand.Seed(time.Now().Unix())

	rowUrl := "https://www.rusprofile.ru/ajax.php?query=%s&action=search&cacheKey=%.12f"
	url := fmt.Sprintf(rowUrl, inn, rand.Float64())

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respOne ContentsData

	err = json.Unmarshal(body, &respOne)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(respOne.UL[0].Name)
	fmt.Println(respOne.UL[0].INN)
	fmt.Println(GetKPP(inn))
	fmt.Println(respOne.UL[0].CeoName)
}

func GetKPP(inn string) string {
	rowUrl := "https://www.rusprofile.ru/search?query=%s&type=ul"
	url := fmt.Sprintf(rowUrl, inn)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var spans []string
	doc.Find(".dashboard-page .company-col dd").Each(func(i int, s *goquery.Selection) {
		item := s.Find("span").Text()
		spans = append(spans, item)
	})
	return spans[3]
}
