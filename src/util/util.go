package util

import (
	"regexp"
)



func RemoveTag(htmlstring string)string{
	//var r = regexp.MustCompile("<\\s*/?\\s*[a-zA-Z0-9]+.*?>")
	var r = regexp.MustCompile(`<\s*/?\s*[a-zA-Z0-9]+.*?>`)
	return r.ReplaceAllString(htmlstring, "")
}
