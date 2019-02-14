package types

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	ResponseWriter http.ResponseWriter
	Route          string
	Method         string
	Values         []byte
}

func (r *Request) SendRequest(merchantId, merchantKey string) ([]byte, error) {

	var values io.Reader
	if r.Values == nil {
		values = nil
	} else {
		values = bytes.NewBuffer(r.Values)
	}

	req, err := http.NewRequest(r.Method, r.Route, values)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]interface{})

	headers["MerchantId"] = merchantId
	headers["MerchantKey"] = merchantKey
	headers["Content-Type"] = "application/json"

	for key, value := range headers {
		req.Header.Add(key, fmt.Sprintf("%v", value))
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
