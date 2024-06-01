package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	CsbProfileBaseUrl = "https://www.cloudskillsboost.google/public_profiles/%s"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	store := persistence.NewInMemoryStore(time.Second)
	v1 := r.Group("/v1")
	{
		v1.GET("/csb_profile", cache.CachePage(store, time.Minute, GetCsbProfile))
	}
	if err := r.Run(); err != nil {
		log.Println(err)
	}
}

func GetCsbProfile(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	res, err := http.Get(fmt.Sprintf(CsbProfileBaseUrl, id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var hs []gin.H
	document.Find(".profile-badge").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Find(".badge-image").First().Attr("href")
		image, _ := selection.Find("img").First().Attr("src")
		title := selection.Find(".ql-title-medium").First()
		earned := selection.Find(".ql-body-medium").First()

		// Regular expression to match the desired date format, then find all matches within the text
		re := regexp.MustCompile(`[A-Z][a-z]{2}\s*\d{1,2}, \d{4}`)
		matches := re.FindAllString(strings.TrimSpace(earned.Text()), -1)

		hs = append(hs, gin.H{
			"link":   link,
			"image":  image,
			"detail": strings.TrimSpace(title.Text()),
			"earned": matches[0],
		})
	})

	c.JSON(http.StatusOK, hs)
}
