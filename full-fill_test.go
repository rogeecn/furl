package furl

import (
	"testing"
)

func TestFUrl_Get(t *testing.T) {
	pageUrl := "http://baidu.com/aaa/bbb/ccc?query=123&ran"

	pathUrl := "//a/b?123"
	fUrl := New(pageUrl, pathUrl)
	resultUrl, err := fUrl.Get()
	t.Log(pathUrl, ":", resultUrl, err)

	pathUrl = "/a/b?query=123#123123"
	fUrl = New(pageUrl, pathUrl)
	resultUrl, err = fUrl.Get()
	t.Log(pathUrl, ":", resultUrl, err)

	pathUrl = "a/b?query=123#123123"
	fUrl = New(pageUrl, pathUrl)
	resultUrl, err = fUrl.Get()
	t.Log(pathUrl, ":", resultUrl, err)
}
