package goredditscraper

import (
	"github.com/gocolly/colly"
)
type Data struct {
	link [] string
}
// const URL = "https://www.reddit.com/r/IndonesianGirlsOnly/top/?t=day"
func (d *Data) AddLink(link string) {
	d.link=append(d.link,link)
}
func GetRedditLinkImage(URL string,scraper *colly.Collector)(Data,error) {
	data:=Data{}
	scraper.OnHTML("img[alt='Post image']", func(e *colly.HTMLElement) {
		data.AddLink(e.Attr("src"))
	})
	err:=scraper.Visit(URL)
	if err!=nil{
		return data,nil
	}
	return data,nil
}

