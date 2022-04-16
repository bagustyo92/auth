package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestParams struct {
	Url     string
	Method  string
	Header  *http.Header
	Payload interface{}
}

func Request(params *RequestParams, out interface{}) error {
	var payload io.Reader = nil
	if params.Payload != nil {
		payloadByte, _ := json.Marshal(&params.Payload)
		payload = bytes.NewBuffer([]byte(payloadByte))
	}

	req, err := http.NewRequest(params.Method, params.Url, payload)
	if err != nil {
		return err
	}

	if params.Header != nil {
		req.Header = *params.Header
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errBody map[string]interface{}
		if err := json.Unmarshal(body, &errBody); err != nil {
			return err
		}

		return errors.New(fmt.Sprintf("%v", errBody))
	}

	if out != nil {
		json.Unmarshal(body, &out)
	}
	return nil
}
