package main

import (
    "bytes"
    "fmt"
    "log"
    "strings"
    "regexp"

    "golang.org/x/net/html"
)

func findData(node *html.Node, existsAddress *bool) string {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.TextNode {
            var buff bytes.Buffer
            buff.WriteString(c.Data)

            if !*existsAddress {
                *existsAddress = strings.HasPrefix(buff.String(), "〒905-0401")
                break
            }

            re := regexp.MustCompile(`(\d{2,4})-(\d{2,4})-(\d{3,4})`)
            if *existsAddress && re.MatchString(buff.String()) {
                return buff.String()
            }
        }
    }

    return ""
}

func SearchTelNum(node *html.Node, existsAddress *bool, telNum *string) {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.ElementNode && c.Data == "span" {
            num := findData(c, existsAddress)
            if num != "" {
                *telNum = num
            }
        }

        SearchTelNum(c, existsAddress, telNum)
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

    var existsAddress bool
    var telNum string
    SearchTelNum(node, &existsAddress, &telNum)
    fmt.Println(telNum)
}