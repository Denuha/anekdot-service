package parser

import (
	"fmt"
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

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"http://anekdotme.ru/random"},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.anekdot").Each(func(i int, s *goquery.Selection) {
				text := s.Find("div.anekdot_text").Text()
				id := s.Find("a.number").Text()

				decoder := charmap.Windows1251.NewDecoder()
				decodedText, err := decoder.String(text)
				if err != nil {
					fmt.Println(err)

				}

				decodedText = strings.ReplaceAll(decodedText, "\t", "")
				decodedText = strings.Replace(decodedText, "\n", "", 1)
				decodedText = strings.Replace(decodedText, "  ", "", -1)
				decodedText = strings.Replace(decodedText, "вЂ”", "—", -1)

				var tmp = models.Anekdot{
					SenderID:   1,
					Text:       decodedText,
					ExternalID: id,
					Status:     int(models.StatusAnekdotOK),
				}
				anekdotList = append(anekdotList, tmp)
			})
		},
		LogDisabled: true,
	}).Start()

	return anekdotList, nil
}

func NewParserAnekdotme() Parser {
	return &ParserAnekdotme{}
}
