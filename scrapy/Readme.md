### 使用Golang实现网页爬虫

大多数的爬虫都是`Python`实现的，也可以使用`Golang`语言实现爬虫，使用的第三方库是[goquery](https://github.com/PuerkitoBio/goquery),这个库是类似`Jquery`一样的`Golang`实现，可以像`Jquery`一样操作`html`,爬取的页面是[Github Trending](https://github.com/trending)

#### 安装

```
go get github.com/PuerkitoBio/goquery
```

#### 获取html页面,输出到控制台

```
import (
	"io"
	"log"
	"net/http"
	"os"
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
	io.Copy(os.Stdout, resp.Body)
}

func main() {
	ScrapeHtml()
}
```

#### 使用 `goquery` 获取页面标题

```
package main

import (
	"fmt"
	"log"
	"net/http"

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
}

func main() {
	ScrapeHtml()
}

```

*输出结果*

```
Trending  repositories on GitHub today · GitHub
```

#### 使用 `goquery`，通过`tag`获取单个或多个页面元素

```
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

func main() {
	ScrapeHtml()
}

```

*输出结果*

```
Trending  repositories on GitHub today · GitHub
apachecn / awesome-algorithm
cfenollosa / os-tutorial
Tencent / MMKV
Zulko / eagle.js
TheAlgorithms / Python
Eloston / ungoogled-chromium
mjavascript / mastering-modular-javascript
imhuay / Algorithm_Interview_Notes-Chinese
slothking-online / graphql-editor
pubkey / rxdb
hoya012 / deep_learning_object_detection
curl / curl
alibaba / arthas
cgarciae / pypeln
Jam3 / math-as-code
you-dont-need / You-Dont-Need-Momentjs
brannondorsey / chattervox
dntzhang / westore
ihucos / plash
Yorko / mlcourse.ai
huanghaibin-dev / CalendarView
ocean1 / awesome-thesis
BNMetrics / logme
alibaba / easyexcel
Algram / ytdl-webserver
```

#### 通过元素的属性(Attr)获取页面元素

```
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
```

*输出结果*

```
https://github.com/apachecn/awesome-algorithm
https://github.com/cfenollosa/os-tutorial
https://github.com/Tencent/MMKV
https://github.com/Zulko/eagle.js
https://github.com/TheAlgorithms/Python
https://github.com/Eloston/ungoogled-chromium
https://github.com/mjavascript/mastering-modular-javascript
https://github.com/imhuay/Algorithm_Interview_Notes-Chinese
https://github.com/slothking-online/graphql-editor
https://github.com/pubkey/rxdb
https://github.com/hoya012/deep_learning_object_detection
https://github.com/curl/curl
https://github.com/alibaba/arthas
https://github.com/cgarciae/pypeln
https://github.com/Jam3/math-as-code
https://github.com/you-dont-need/You-Dont-Need-Momentjs
https://github.com/brannondorsey/chattervox
https://github.com/dntzhang/westore
https://github.com/ihucos/plash
https://github.com/Yorko/mlcourse.ai
https://github.com/huanghaibin-dev/CalendarView
https://github.com/ocean1/awesome-thesis
https://github.com/BNMetrics/logme
https://github.com/alibaba/easyexcel
https://github.com/Algram/ytdl-webserver
```

#### 通过 `class` 或者其他的属性获取页面元素

```
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
```

*输出结果*

```
979 stars today
745 stars today
732 stars today
663 stars today
605 stars today
485 stars today
418 stars today
359 stars today
348 stars today
268 stars today
209 stars today
240 stars today
184 stars today
185 stars today
181 stars today
170 stars today
157 stars today
150 stars today
148 stars today
146 stars today
139 stars today
135 stars today
132 stars today
121 stars today
122 stars today
```

#### 获取元素的子节点

```
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

```

*输出结果*

```
apachecn / awesome-algorithm
cfenollosa / os-tutorial
Tencent / MMKV
Zulko / eagle.js
TheAlgorithms / Python
Eloston / ungoogled-chromium
mjavascript / mastering-modular-javascript
imhuay / Algorithm_Interview_Notes-Chinese
slothking-online / graphql-editor
pubkey / rxdb
hoya012 / deep_learning_object_detection
curl / curl
alibaba / arthas
cgarciae / pypeln
Jam3 / math-as-code
you-dont-need / You-Dont-Need-Momentjs
brannondorsey / chattervox
dntzhang / westore
ihucos / plash
Yorko / mlcourse.ai
huanghaibin-dev / CalendarView
ocean1 / awesome-thesis
BNMetrics / logme
alibaba / easyexcel
Algram / ytdl-webserver
```

#### 获取前后的兄弟节点

```
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
```

*输出结果*

```
TheAlgorithms / Python
Zulko / eagle.js
Eloston / ungoogled-chromium
```

#### 组合在一起(爬取 Github Trending)

```
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
```

*输出结果*

```
apachecn / awesome-algorithm     979 stars today         https://github.com/apachecn/awesome-algorithm
cfenollosa / os-tutorial         745 stars today         https://github.com/cfenollosa/os-tutorial
Tencent / MMKV   732 stars today         https://github.com/Tencent/MMKV
Zulko / eagle.js         663 stars today         https://github.com/Zulko/eagle.js
TheAlgorithms / Python   605 stars today         https://github.com/TheAlgorithms/Python
Eloston / ungoogled-chromium     485 stars today         https://github.com/Eloston/ungoogled-chromium
mjavascript / mastering-modular-javascript       418 stars today         https://github.com/mjavascript/mastering-modular-javascript
imhuay / Algorithm_Interview_Notes-Chinese       359 stars today         https://github.com/imhuay/Algorithm_Interview_Notes-Chinese
slothking-online / graphql-editor        348 stars today         https://github.com/slothking-online/graphql-editor
pubkey / rxdb    268 stars today         https://github.com/pubkey/rxdb
hoya012 / deep_learning_object_detection         209 stars today         https://github.com/hoya012/deep_learning_object_detection
curl / curl      240 stars today         https://github.com/curl/curl
alibaba / arthas         184 stars today         https://github.com/alibaba/arthas
cgarciae / pypeln        185 stars today         https://github.com/cgarciae/pypeln
Jam3 / math-as-code      181 stars today         https://github.com/Jam3/math-as-code
you-dont-need / You-Dont-Need-Momentjs   170 stars today         https://github.com/you-dont-need/You-Dont-Need-Momentjs
brannondorsey / chattervox       157 stars today         https://github.com/brannondorsey/chattervox
dntzhang / westore       150 stars today         https://github.com/dntzhang/westore
ihucos / plash   148 stars today         https://github.com/ihucos/plash
Yorko / mlcourse.ai      146 stars today         https://github.com/Yorko/mlcourse.ai
huanghaibin-dev / CalendarView   139 stars today         https://github.com/huanghaibin-dev/CalendarView
ocean1 / awesome-thesis          135 stars today         https://github.com/ocean1/awesome-thesis
BNMetrics / logme        132 stars today         https://github.com/BNMetrics/logme
alibaba / easyexcel      121 stars today         https://github.com/alibaba/easyexcel
Algram / ytdl-webserver          122 stars today         https://github.com/Algram/ytdl-webserver
```

[更多精彩内容](http://coderminer.com)
[源码](https://github.com/coderminer/go-learn/tree/master/scrapy)
[译自](https://www.thetaranights.com/web-scraping-using-golang/?utm_campaign=Master%20the%20World%20of%20Golang&utm_medium=email&utm_source=Revue%20newsletter)