package fetcher

import "strings"

var accepted = [...]string{"text/html", "text/plain"}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

func isValid(header string) bool {
	for _, accept := range accepted {
		if strings.Contains(header, accept) {
			return true
		}
	}
	return false
}
