package web

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func (c *IServWebClient) GetExercises() ([]types.IServExercise, error) {
	// get html page of exercise overview
	res, err := c.httpClient.Get(c.iServUrl + "/exercise")
	if err != nil {
		return []types.IServExercise{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return []types.IServExercise{}, err
	}

	// for each task get the url
	var urls []string
	doc.Find("#crud-table tbody tr").Each(func(i int, s *goquery.Selection) {
		if !s.HasClass("info") {
			url, _ := s.Children().Eq(1).Children().First().Attr("href")
			urls = append(urls, url)
		}
	})

	return []types.IServExercise{}, nil
}
