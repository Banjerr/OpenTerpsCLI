package terpene

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetTerpenes(BaseURL string) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", BaseURL+"terpenes", nil)
	req.Header.Add("APIUser", os.Getenv("APIUser"))
	req.Header.Add("APIKey", os.Getenv("APIKey"))
	resp, err := client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error Getting Terpenes!"
	}
	fmt.Println("data is ", string(data))
	return string(data)
}

func GetTerpene() string {
	return ""
}

func AddTerpene() string {
	return ""
}

func UpdateTerpene() string {
	return ""
}

func RemoveTerpene() string {
	return ""
}
