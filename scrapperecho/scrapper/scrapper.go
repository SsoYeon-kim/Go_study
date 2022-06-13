package scrapper

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

type extractedJob struct {
	title    string
	location string
	summary  string
}

//Scrape indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"

	var jobs []extractedJob

	//[]extractedJob을 받을거니까 channel의 type이여야 함
	c := make(chan []extractedJob)

	totalPages := getPages(baseURL)

	// 각 페이지 별로 getPage 함수 호출
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	//위에서 goroutine이 totalPages만큼 생김 (그만큼의 메세지를 채널을 통해 받아야함)
	for i := 0; i < totalPages; i++ {
		//메세지를 기다리다가 채널에 전송되면 extractedJob에 저장
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

//각 페이지에 있는 job 반환
//goroutine을 생성해 일자리를 전달받고, main함수의 채널로 전송함
func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob

	c := make(chan extractedJob)

	//Itoa : integer to ascii
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".job_seen_beacon")

	//fmt.Println(searchCards)

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		//메세지가 전달되기를 기다렸다가, 받으면 jobs 배열에 추가
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

//채널을 통해 메세지 전송
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	title := CleanString(card.Find(".jobTitle>a").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	summary := CleanString(card.Find(".job-snippet").Text())
	c <- extractedJob{
		title:    title,
		location: location,
		summary:  summary}
}

//CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

//총 페이지 숫자
func getPages(url string) int {
	pages := 0

	res, err := http.Get(url)
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
