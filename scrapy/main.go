package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeHtml() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status coder error: %d %s", resp.StatusCode, resp.Status)
	}
	//io.Copy(os.Stdout, resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Find("title").Text())

	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		fmt.Println(strings.TrimSpace(s.Find("h3").Text()))
	})
}

func ScrapeAttr() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status coder error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		href, has_attr := s.Find("a").First().Attr("href")
		if has_attr {
			fmt.Println("https://github.com" + href)
		}
	})
}

func ScrapeViaClass() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		fmt.Println(strings.TrimSpace(s.Find(".float-sm-right").Text()))
	})
}

func NaviChildren() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Find("title").Text())
	olSelection := doc.Find("ol")
	olSelection.Children().Each(func(i int, s *goquery.Selection) {
		fmt.Println(strings.TrimSpace(s.Find("h3").Text()))
	})
}

func NaviSlibing() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	liSelection := doc.Find("ol li")
	fifthEle := liSelection.Eq(4)
	fmt.Println(strings.TrimSpace(fifthEle.Find("h3").Text()))

	fourthEle := fifthEle.Prev()
	fmt.Println(strings.TrimSpace(fourthEle.Find("h3").Text()))

	sixthEle := fifthEle.Next()
	fmt.Println(strings.TrimSpace(sixthEle.Find("h3").Text()))
}

func ScrapeGithubTrending() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		repositoryName := strings.TrimSpace(s.Find("h3").Text())
		totalStarsToday := strings.TrimSpace(s.Find(".float-sm-right").Text())
		href, has_attr := s.Find("a").Attr("href")
		if !has_attr {
			href = "No valid url found"
		}
		fmt.Println(repositoryName, "\t", totalStarsToday, "\t", "https://github.com"+href)
	})
}

func main() {
	//ScrapeHtml()
	//ScrapeAttr()
	//ScrapeViaClass()
	//NaviChildren()
	//NaviSlibing()
	ScrapeGithubTrending()
}
