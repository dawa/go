package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
  "strconv"
)

func responseTime(url string, c chan string) {
    start := time.Now()

    res, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    elapsed := time.Since(start).Seconds()

    c <- url + " took " + strconv.FormatFloat(elapsed, 'f', 6, 64) + " seconds \n"

}

func receiver(c chan string) {
  for msg := range c {
    fmt.Println(msg)
  }
}

func main() {
  urls := make([]string, 3)

  urls[0] = "https://www.usa.gov/"

  urls[1] = "https://www.gov.uk/"

  urls[2] = "http://www.gouvernement.fr/"

  c := make(chan string, 3)

  for _, u := range urls {

    go responseTime(u, c)

  }
  
  fmt.Println("Pushed 3 messages to channel")

  time.Sleep(time.Second * 2)

  close(c)

  receiver(c)
}

