package searchtelnum

import (
    "net/http"
    "bytes"
    "fmt"
    "log"
    "strings"
    "regexp"

    "golang.org/x/net/html"
)

var (
    telNum string
    existsAddress bool
)


func findData(node *html.Node, postalCode string) string {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.TextNode {
            var buff bytes.Buffer
            buff.WriteString(c.Data)

            if !existsAddress {
                existsAddress = strings.HasPrefix(buff.String(), postalCode)
                break
            }

            re := regexp.MustCompile(`(\d{2,4})-(\d{2,4})-(\d{3,4})`)
            if existsAddress && re.MatchString(buff.String()) {
                return buff.String()
            }
        }
    }

    return ""
}

func searchTelNum(node *html.Node, postalCode string) {
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.ElementNode && c.Data == "span" {
            num := findData(c, postalCode)
            if num != "" {
                telNum = num
            }
        }

        searchTelNum(c, postalCode)
    }
}

func getHtml(companyName string) (*html.Node, error) {
    resp, err := http.Get("https://www.google.com/search?q=" + companyName)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    node, err := html.Parse(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    return node, nil
}

func Run(companyName string, postalCode string) (string, error) {
    if companyName == "" {
        return "", fmt.Errorf("company name is empty")
    }

    if postalCode == "" {
        return "", fmt.Errorf("postalCode is empty")
    }

    node, err := getHtml(companyName)
    if err != nil {
        return "", fmt.Errorf("getHtml error: %v", err)
    }

    searchTelNum(node, postalCode)

    return telNum, nil
}