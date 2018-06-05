package util

import (
"testing"
	"regexp"
	"fmt"
)

func reRemoveCompare(t *testing.T,origin string, expect string)  {
	x:=RemoveTag(origin)
	if x!=expect{
		t.Error(x+"!="+expect)
	} else {
		t.Log("pass")
	}
}

func TestRemoveTag(t *testing.T) {
	text := `Hello 世界！123 Go.`
	reg := regexp.MustCompile(`(Hello)(.*)(Go)`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
	reRemoveCompare(t,"<br>","\n")
	reRemoveCompare(t,"<br/>","\n")
	reRemoveCompare(t,"<BR>","\n")
	reRemoveCompare(t,"<BR/>","\n")
	//reRemoveCompare(t,"<p>test</p>","test\n")
	//reRemoveCompare(t,"<P>test</P>","test\n")
	//reRemoveCompare(t,"<p>test1<br/>test2</p>","test1\ntest2\n")
	//reRemoveCompare(t,"<p>test1</p><p>test2</p>","test1\ntest2\n")


	reRemoveCompare(t,"<abc/>","")
	reRemoveCompare(t,"<abc>","")
}

