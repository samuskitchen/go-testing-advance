package request_http

import (
	"io/ioutil"
	"net/http"
)

func fetchContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}


