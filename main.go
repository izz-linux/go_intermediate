package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// create a new channel for comms
	c := make(chan string)

	for _,link := range links {
		//var j int
		// fork child routines with go keyword: (need channels added)
		go checkLink(link, c)
	}

	// function literal insert:
	for l := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkLink(url, c)
			/*j := 0
			for i:=j; i < 4; i++ {
				checkLink(url, c)
				fmt.Printf("\nThe total number completed is %v",)
			}*/
		}(l) // these parentheses are here to denote "CALL IT!" with link l outside the scoped function literal ugh...
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
// this is ran in serial
func checkLink(link string, c chan string) {
	// not here because this is stupid
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		currentTime := time.Now()
		d1 := []byte(currentTime.Format("2006-01-02 15:04:05"))
		d2 := []byte(" -- [ALERT] -- this  is down\n")
		var d3 []byte = append(d1, d2...)
		// NEED TO APPEND, NOT OVERWRITE
		err := ioutil.WriteFile("/Users/izz/develop/custom/engineering/golang/Learning/intmon_go/dat1", d3, 0644)
		check(err)
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	currentTime2 := time.Now()
	d7 := []byte(currentTime2.Format("2006-01-02 15:04:05"))
	d8 := []byte(" -- [INFO] -- this  is up\n")
	var d9 []byte = append(d7, d8...)
	// NEED TO APPEND, NOT OVERWRITE
	err = ioutil.WriteFile("/Users/izz/develop/custom/engineering/golang/Learning/intmon_go/dat1", d9, 0644)
	check(err)
	fmt.Println(link, "is up!")
	c <- link
}
