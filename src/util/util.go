package util

import (
	"regexp"
)



func RemoveTag(htmlstring string)string{
	r := regexp.MustCompile("<\\s*/?\\s*[a-zA-Z0-9]+.*?>")
	return r.ReplaceAllString(htmlstring, "")
}
