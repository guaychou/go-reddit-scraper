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

func NewScraper() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent="Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:28.0) Gecko/20100101 Firefox/28.0"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Accept-Encoding","gzip, deflate")
		r.Headers.Set("Accept-Language","en-US,en;q=0.9")
	})
	return c
}

func getLinkImage(scraper *colly.Collector)(Data,error) {
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