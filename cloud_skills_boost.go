package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	CsbProfileBaseUrl = "https://www.cloudskillsboost.google/public_profiles/%s"
)

type CloudSkillsBoostProfile struct {
	gorm.Model
	ProfileId string
}

func GetCsbProfile(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, &Response{
			Error: "id is required",
		})
		return
	}

	res, err := http.Get(fmt.Sprintf(CsbProfileBaseUrl, id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &Response{
			Error: err.Error(),
		})
		return
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &Response{
			Error: err.Error(),
		})
		return
	}

	name := strings.TrimSpace(document.Find(".ql-display-small").First().Text())
	avatar, _ := document.Find(".profile-avatar").First().Attr("src")

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

	c.JSON(http.StatusOK, &Response{
		Data: gin.H{
			"avatar": avatar,
			"name":   name,
			"badges": hs,
		},
	})
}
