package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	// "math/rand"
	// "time"
	"net/http"
	"flag"
	"github.com/gocolly/colly"
)


func crawl(url string){
	collector := colly.NewCollector(
		colly.AllowedDomains("www.subtitlestar.com", "subtitlestar.com"),
	)

	collector.OnHTML("#link-download", func(element *colly.HTMLElement){
		download_link := element.Attr("href");
		download_file(download_link);
	})

	collector.Visit(url)

}

func download_file(link string){
	// fetch file name
	splited_link := strings.Split(link, "/");
	file_name := splited_link[len(splited_link)-1]
	
	// create blank file
	file, err := os.Create(fmt.Sprintf("./%s", file_name))
	if err != nil{
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil{
		log.Fatal(err)
	}
	
	response, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	defer response.Body.Close()
	io.Copy(file, response.Body)
	defer file.Close()

	fmt.Println("download complete")
}

func main() {
	
	// url := "https://subtitlestar.com/persian-subtitles-wednesday/"
	// var url string;
	url := flag.String("url", "", "the link of page that you want crawl it")
	flag.Parse()
	if *url == ""{
		log.Fatal("invalid url")
	}
	crawl(*url);
}
