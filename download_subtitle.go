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


func check_url(url string){
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil{
		log.Fatal(err);
	}

	res, err := http.DefaultClient.Do(req);
	if err != nil{
		log.Fatal(err);
	}

	if res.StatusCode != 200 {
		log.Fatal("url is invalid")
		os.Exit(0)
	}
}


func crawl(url string){
	check_url(url);
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

func string_replace(s string, old string, new string) string {
	var new_string string;
	for i:=0; i<len(s); i++{
		letter := string(s[i])
		if (letter == old){
			new_string = strings.Replace(s, old, new, i)
		}
	}

	return new_string
}

func main() {

	// url := "https://subtitlestar.com/persian-subtitles-wednesday/"
	// var url string;
	url_addr := flag.String("url", "", "the link of page that you want crawl it")
	name_addr := flag.String("movie", "", "the movie name")
	flag.Parse()
	name := *name_addr;
	url := *url_addr;

	if (name == ""){
		fmt.Println("The movie name is fucking empty");
		os.Exit(0)
	}
	
	name = string_replace(name, "_", "-")
	fmt.Println(name)
	if (url == ""){
		url = fmt.Sprintf("https://subtitlestar.com/persian-subtitles-%s/", name)
	}
	crawl(url);
}
