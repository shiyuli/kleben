package core

import (
	"io/ioutil"
	"net/http"
)

func downloadFile(fileUrl string) ([]byte, error) {
	response, err := http.Get(fileUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
