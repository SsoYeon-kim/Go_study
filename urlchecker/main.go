package main

/*
import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//빈 map 만들기
	//var results map[sring]sring 은 초기화하지 않았기 때문에 값을 추가할 수 없음, nil임
	//var results = map[string]string{}
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
*/

/*
func coolCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is cool", i)
		time.Sleep(time.Second)
	}
}
*/

/*
import (
	"fmt"
	"time"
)

func main() {
	//Gorouines는 프로그램이 동작하는 동안만 유효 (메인함수 실행 동안만)
	//main하무는 goroutines를 기다리지 않음, 메인 함수가 끝나면 goroutine도 소멸됨
	//go coolCount("so"), coolCount("yeon")이 동시에 실행됨
	// go coolCount("so")
	// coolCount("yeon")

	//위는 main이 coolCount("yeon")을 가리키고 있었기 때문에 동작
	//time.Sleep() 추가
	// go coolCount("so")
	// go coolCount("yeon")

	//chan 보낼정보의 type
	//▶ c := make(chan string)
	c := make(chan bool)
	people := [2]string{"so", "yeon"}

	for _, person := range people {
		go isCool(person, c)
	}
	//true를 채널로 보냈고, 채널의 메세지를 result로 받음
	//time.sleep을 지워도 동작함
	//채널로부터 받을 때 기다림
	//메세지를 받는 것은 blocking operation, 이 작업이 끝날 때까지 멈춤

	//goroutines 개수만큼 메세지를 받을 수 있음

	//resultOne := <-c
	//resultTwo := <-c
	//fmt.Println("Waiting for messages")
	//fmt.Println("Received this message:", <-c)
	//메세지를 얻어야 다음라인으로 넘어옴
	//fmt.Println(<-c)

	//동시성때문에 실행할때마다 결과는 다르게 나옴
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i)
		fmt.Println(<-c)
	}
}

//c chan bool : c의 type은 chan, 이 채널을 통해 보낼 데이터의 type은 bool이다
//채널을 두개의 함수에 보내줬고 5초뒤 true 2개의 메세지를 채널로 전송함
func isCool(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	//▶ c <- person + "is cool"
	c <- true
}
*/

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

//이 함수는 데이터를 받을 수 없고 보내기만 가능
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
		c <- requestResult{url: url, status: status}
	}
	c <- requestResult{url: url, status: status}

}
