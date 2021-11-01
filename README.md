# Search Tel Num

## About
- use `golang.org/x/net/html`
- If you give searchtelnum your company name and zip code, it will return your phone number.
- Use Google to search for it
- It's faster than goquery for this purpose alone.

## Install

```bash
$ go get -u github.com/sohey-dr/searchtelnum
```

## Usage

### Example

```go
package main

import (
        "log"
        "github.com/sohey-dr/searchtelnum"
)

func main() {
	telNum, err := searchtelnum.Run("Example Inc.", "〒XXX-XXXX")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(telNum)
}
```

## License

Released under the [MIT License](https://github.com/sohey-dr/searchtelnum/blob/main/LICENSE).
