package kf

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func connect(config *KFConfig, path string, method string, bodyJson string) []byte {
	url := config.Url + path
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, url, strings.NewReader(bodyJson))
		req.Header.Add("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	req.Header.Add("Authorization", "Bearer "+config.Token)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	return body
}

func login(config *KFConfig) string {
	ps := url.Values{}
	ps.Add("userName", config.Username)
	ps.Add("password", config.Password)

	url := config.Url + "auth/local/"
	res, err := http.PostForm(url, ps)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	type TokenObj struct {
		Token string
	}
	var tokenobj TokenObj
	if err := json.Unmarshal(body, &tokenobj); err != nil {
		fmt.Println(err)
		return ""
	}
	config.Token = tokenobj.Token
	return tokenobj.Token
}
