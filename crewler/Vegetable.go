package crewler

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"farming_backend/models"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
)

const (
	// StartTime 开始时间
	// StartTime = "20_04_01"
	// EndTime 结束时间
	// EndTime     = "20_04_30"
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//GetURL 获取url
func getURL(startTime string, endTime string, page int) string {

	return "http://www.jnmarket.net/import/list-1/date-" + startTime + "-" + endTime + "/" + strconv.Itoa(page) + ".html"
}

// RandomString 随机字符串
func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// VegetableCrewler 蔬菜菜价格爬虫
func VegetableCrewler(startTime string, endTime string) {
	c := colly.NewCollector(
		// colly.UserAgent("xy"),
		colly.Debugger(&debug.LogDebugger{}),
	)
	extensions.RandomUserAgent(c)
	var pageNums string

	c.OnHTML("body > div.page-container > div > div.content > div.main.clearfix > div.flickr", func(el *colly.HTMLElement) {
		pageNums = el.ChildText("a:nth-child(8)")

	})
	var vegetablePrizeList []models.VegetablePrize

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			// dateTime 　:= el.ChildText("td:nth-child(5)")
			// a := strings.Split(dateTime, "-")

			// strings.Replace()
			dateTime := "20" + el.ChildText("td:nth-child(5)") + " 00:00:00"
			st, _ := time.Parse("2006-01-02 15:04:05", dateTime)
			vegetablePrize := models.VegetablePrize{
				Vegetable: el.ChildText("td:nth-child(1)"),
				Origin:    el.ChildText("td:nth-child(2)"),
				AvgPrice:  el.ChildText("td:nth-child(3)"),
				DateTime:  st,
			}
			vegetablePrizeList = append(vegetablePrizeList, vegetablePrize)
		})
		models.AddMulti(vegetablePrizeList)
		vegetablePrizeList = []models.VegetablePrize{}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	url := getURL(startTime, endTime, 1)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})
	c.Visit(url)

	groutineThinkTimeThrottle := time.Tick(120 * time.Second)

	n, _ := strconv.Atoi(pageNums)

	for i := 2; i < n; i++ {

		go func() {
			url := getURL(startTime, endTime, i)
			c.Visit(url)
		}()
		<-groutineThinkTimeThrottle
	}

	// fmt.Println(pageNums)

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", " ")

	// Dump json to the standard output

	// enc.Encode(vegetablePrizeList)

}
