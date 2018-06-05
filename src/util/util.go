package util

import (
	"regexp"
)

func RemoveTagClosure() func(string)string{
	var tag_br = regexp.MustCompile(`(<br>)|(<br/>)|(<BR>)|(<BR/>)`)
	//var tag_p = regexp.MustCompile(`<[pP]>([^<]*)</[pP]>`)
	var tags = regexp.MustCompile(`<\s*/?\s*[a-zA-Z0-9]+.*?>`)

	replace :=func (htmlstring string)string{
		htmlstring = tag_br.ReplaceAllString(htmlstring, "\n")
		//htmlstring = tag_p.ReplaceAllString(htmlstring, "$1\n")
		htmlstring = tags.ReplaceAllString(htmlstring, "")
		return htmlstring
	}
	return replace
}

var RemoveTag = RemoveTagClosure()
