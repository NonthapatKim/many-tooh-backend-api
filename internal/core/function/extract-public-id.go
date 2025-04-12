package function

import "strings"

func ExtractPublicID(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return ""
	}
	last := parts[len(parts)-1]
	return strings.Split(last, ".")[0]
}
