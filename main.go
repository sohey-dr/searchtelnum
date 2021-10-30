package main

import (
    "bytes"
    "fmt"
    "log"
    "strings"

    "golang.org/x/net/html"
)

type Anchor struct {
    Text string
    Href string
}

func NewAnchor(node *html.Node) *Anchor {
    // href属性の値を取得
    href := ""
    for _, v := range node.Attr {
        if v.Key == "href" {
            href = v.Val
            break
        }

        if v.Key == "class" {
            if v.Val == "testt" {
                // fmt.Println(node.Data)
            }
        }
    }

    var buff bytes.Buffer
    // A要素のテキストを取得
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.TextNode {
            buff.WriteString(c.Data)
        }
    }

    return &Anchor{Text: buff.String(), Href: href}
}

func FindAnchors(node *html.Node, collection *[]*Anchor) {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.ElementNode && c.Data == "span" {
            fmt.Println(c.Data)
            break
        }
        FindAnchors(c, collection)
    }
}

func main() {

    r := strings.NewReader(`
<html>
<head></head>
<body>
  <ul>
      <li><a class="testt" href="https://example.com/foo">foo</a></li>
      <span>bar</span>
      <li><a href="https://example.com/bar">bar</a></li>
      <li><a href="https://example.com/baz">baz</a></li>
  </ul>
</body>
</html>
`)

    node, err := html.Parse(r)
    if err != nil {
        log.Fatal(err)
    }

    var collection []*Anchor
    FindAnchors(node, &collection)

    for _, a := range collection {
        fmt.Println(a.Text, ":", a.Href)
    }
}