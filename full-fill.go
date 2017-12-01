package furl

import (
	"strings"
	"net/url"
	"fmt"
)

type FUrl struct {
	currentPage    string
	currentPageURL *url.URL
	urlPath        string
}

func New(currentPage, urlPath string) *FUrl {
	return &FUrl{
		currentPage: currentPage,
		urlPath:     urlPath,
	}
}

func (instance *FUrl) Get() (string, error) {
	var err error
	instance.currentPageURL, err = url.Parse(instance.currentPage)
	if err != nil {
		return "", err
	}
	return instance.process()
}

func (instance *FUrl) process() (string, error) {
	if strings.HasPrefix(instance.urlPath, "http") {
		return instance.urlPath, nil
	}

	if strings.HasPrefix(instance.urlPath, "//") {
		retUrl := fmt.Sprintf("%s://%s%s",
			instance.currentPageURL.Scheme,
			instance.currentPageURL.Host,
			instance.urlPath[1:len(instance.urlPath)])

		return retUrl, nil
	}

	if strings.HasPrefix(instance.urlPath, "/") {
		retUrl := fmt.Sprintf("%s://%s%s",
			instance.currentPageURL.Scheme,
			instance.currentPageURL.Host,
			instance.urlPath)

		return retUrl, nil
	}

	tmpUrl, err := url.Parse(fmt.Sprintf("http://tmp.url/%s", instance.urlPath))
	if err != nil {
		return "", err
	}

	tmpUrl.Scheme = instance.currentPageURL.Scheme
	tmpUrl.Host = instance.currentPageURL.Host

	pageUrlPath := strings.Split(instance.currentPageURL.Path, "/")
	pageUrlPath[len(pageUrlPath)-1] = strings.TrimLeft(tmpUrl.Path, "/")
	tmpUrl.Path = strings.Join(pageUrlPath, "/")

	return tmpUrl.String(), nil
}
