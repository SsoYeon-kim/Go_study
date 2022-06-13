package main

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/ssoyeon/learngo/scrapperecho/scrapper"
)

//go get github.com/labstack/echo (echo 패키지 설치)

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	//Attachment(파일 리턴)
	return c.Attachment("jobs.csv", "jobs.csv")
}

//1323 포트로 서버 접속 가능 (http://localhost:1323/)
func main() {
	//scrapper.Scrape("term")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
