package testutils

// This file adds some nice helpers for performing controller tests

import (
	log "github.com/Sirupsen/logrus"
	"github.com/derekdowling/bursa/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Formats a post form submission url
func buildFormUrl(path string, form url.Values) string {
	url := baseUrl(path)
	url.RawQuery = form.Encode()
	return url.String()
}

func GetRequest(path string) *http.Request {
	url := buildUrl(path)
	return buildRequest("GET", url, nil)
}

func buildUrl(path string) string {
	url := baseUrl(path)
	return url.String()
}

func baseUrl(path string) *url.URL {
	return &url.URL{
		Host: getTestHost(),
		Path: path,
	}
}

// Checks our config to get our local test path
func getTestHost() string {
	location := "localhost"
	port := config.App.GetString("ports.http")
	return strings.Join([]string{location, port}, ":")
}

func buildRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return req
}

// Builds a nice post form submissions request you can use in testing
func FormPostRequest(path string, form url.Values) *http.Request {
	url := buildFormUrl(path, form)
	return buildRequest("POST", url, nil)
}
