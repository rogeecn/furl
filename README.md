# FURL (full fill url)
------------

fill page relative link to absolute url

usage:

```go
	pageUrl := "http://google.com/aaa/bbb/ccc?query=123&ran"
	pathUrl := "//a/b?123"
	fUrl := New(pageUrl, pathUrl)
	resultUrl, err := fUrl.Get()
	fmt.Println(resultUrl,err)
	// output
	// http://baidu.com/a/b?123 <nil>
```

more usage see `full-fill_test.go`