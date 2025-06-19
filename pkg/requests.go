package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Requests struct {
	Headers map[string]string
}

func (r *Requests) Get(url string) (int, map[string]interface{}, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 500, nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 500, nil, err
	}

	return resp.StatusCode, data, nil
}

func (r *Requests) Post(url string, data map[string]interface{}) (int, map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return 500, nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return 500, nil, err
	}

	//setting headers
	req.Header.Set("Content-Type", "application/json")
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 500, nil, err
	}

	return resp.StatusCode, res, nil
}

func (r *Requests) PostForm(url string, data *url.Values) (int, map[string]interface{}, error) {
	resp, err := http.PostForm(url, *data)
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 500, nil, err
	}

	return resp.StatusCode, res, nil
}
