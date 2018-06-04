package util

import (
"testing"
)

func TestGetAccessToken(t *testing.T) {
	x:=RemoveTag("<p>test</p>")
	if x!="test"{
		t.Error("error:token is null")
	} else {
		t.Log("pass")
	}
}
