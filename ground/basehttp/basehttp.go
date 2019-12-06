package basehttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{}

func init() {
	client.Timeout = 3 * time.Second
}

func BaseRequest(method, url string, dataByte []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respByte, err
}
