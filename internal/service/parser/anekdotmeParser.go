package parser

import (
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"

	"golang.org/x/text/encoding/charmap"
)

// http://anekdotme.ru/
type ParserAnekdotme struct {
}

func (p *ParserAnekdotme) ParseAnekdots() ([]models.Anekdot, error) {
	anekdotList := make([]models.Anekdot, 0)

	var urls = []string{
		// "http://anekdotme.ru/random",
		// "http://anekdotme.ru/lenta",
		"http://anekdotme.ru/lenta/page_" + strconv.Itoa(rand.Intn(750)+1),
		// "http://anekdotme.ru/lenta/page_" + strconv.Itoa(rand.Intn(750)+1),
		// "http://anekdotme.ru/lenta/page_" + strconv.Itoa(rand.Intn(750)+1),
		// "http://anekdotme.ru/lenta/page_" + strconv.Itoa(rand.Intn(750)+1),
		// "http://anekdotme.ru/anekdoti_pro-chukchu",
		// "http://anekdotme.ru/anekdoti_pro-kino",
		// "http://anekdotme.ru/anekdoti_pro-studentov",
		// "http://anekdotme.ru/anekdoti_cherniy-yumor",
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: urls,
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.anekdot").Each(func(i int, s *goquery.Selection) {
				text := s.Find("div.anekdot_text").Text()
				id := s.Find("a.number").Text()

				decoder := charmap.Windows1251.NewDecoder()
				decodedText, err := decoder.String(text)
				if err != nil {
					log.Println(err)
					return
				}

				decodedText = strings.ReplaceAll(decodedText, "\t", "")
				decodedText = strings.Replace(decodedText, "\n", "", 1)
				decodedText = strings.Replace(decodedText, "  ", "", -1)
				decodedText = strings.Replace(decodedText, "вЂ”", "—", -1)

				var tmp = models.Anekdot{
					Sender: models.Sender{
						ID: 1,
					},
					Text:       decodedText,
					ExternalID: id,
					Status:     int(models.StatusAnekdotOK),
				}
				anekdotList = append(anekdotList, tmp)
			})
		},
		LogDisabled: false,
	}).Start()

	return anekdotList, nil
}

func NewParserAnekdotme() Parser {
	return &ParserAnekdotme{}
}
