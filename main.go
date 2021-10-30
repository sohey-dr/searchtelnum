package main

import (
    "bytes"
    "fmt"
    "log"
    "strings"
    "regexp"

    "golang.org/x/net/html"
)

func FindData(node *html.Node) string {
    for _, v := range node.Attr {
        if v.Key == "class" && v.Val == "BNeawe s3v9rd AP7Wnd" {
            // fmt.Println(v.Val)
        }
    }

    var buff bytes.Buffer
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.TextNode {
            buff.WriteString(c.Data)

            re := regexp.MustCompile(`(\d{2,4})-(\d{2,4})-(\d{3,4})`)
            if re.MatchString(buff.String()) {
                return buff.String()
            }
        }
    }

    return ""
}

func SearchTelNum(node *html.Node, telNum *string) {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.ElementNode && c.Data == "span" {
            num := FindData(c)
            if num != "" {
                *telNum = num
            }
        }
        SearchTelNum(c, telNum)
    }
}

func main() {

    r := strings.NewReader(`
<html>
<head></head>
<body>
<div class="vbShOe kCrYT"><div class="AVsepf"><div class="BNeawe s3v9rd AP7Wnd"><span><span class="BNeawe s3v9rd AP7Wnd">住所</span></span>： <span><span class="BNeawe tAd8D AP7Wnd">〒905-0401 沖縄県国頭郡今帰仁村仲宗根９９−３</span></span></div></div><div class="AVsepf"><div class="BNeawe s3v9rd AP7Wnd"><span><span class="BNeawe s3v9rd AP7Wnd">時間</span></span>： <span><span class="BNeawe tAd8D AP7Wnd">営業時間外 ⋅ 営業開始: 9:00 月</span></span></div></div><div class="AVsepf u2x1Od"><div class="BNeawe s3v9rd AP7Wnd"><span><span class="BNeawe s3v9rd AP7Wnd">電話番号</span></span>： <span><span class="BNeawe tAd8D AP7Wnd">0120-954-062</span></span></div></div></div>
</body>
</html>
`)

    node, err := html.Parse(r)
    if err != nil {
        log.Fatal(err)
    }

    var telNum string
    SearchTelNum(node, &telNum)
    fmt.Println(telNum)
}