package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//go get github.com/PuerkitoBio/goquery  설치

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJob struct {
	title    string
	location string
	summary  string
}

func main() {
	var jobs []extractedJob

	totalPages := getPages()

	// 각 페이지 별로 getPage 함수 호출
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		// 두 개의 배열을 합침 [ c c c]
		// ...를 안하면 배열 안에 배열이 추가됨 [[] [] []]
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	//writer 생성, writer에 데이터 입력, 모든 데이터를 파일에 저장
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Title", "Location", "Summary"}
	//w.Write()를 보면 error를 리턴함 따라서 에러 체크
	//w.Write()에는 []string이 입력돼야함
	wErr := w.Write(headers)
	checkErr(wErr)

	//반복문이 끝나면 defer가 실행되고 데이터가 파일에 입력됨
	for _, job := range jobs {
		jobSlice := []string{job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

//각 페이지에 있는 job 반환
func getPage(page int) []extractedJob {
	var jobs []extractedJob

	//Itoa : integer to ascii
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".job_seen_beacon")

	fmt.Println(searchCards)
	//s : 각각의 div를 의미
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}
func extractJob(card *goquery.Selection) extractedJob {
	title := cleanString(card.Find(".jobTitle>a").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{
		title:    title,
		location: location,
		summary:  summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

//총 페이지 숫자
func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	//res.Body : byte인데, 입력과 출력(I/O)임 - 함수가 끝나면 닫아줘야 함

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status : ", res.StatusCode)
	}
}
