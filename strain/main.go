package strain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetStrains(BaseURL string) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", BaseURL+"strains", nil)
	req.Header.Add("APIUser", os.Getenv("APIUser"))
	req.Header.Add("APIKey", os.Getenv("APIKey"))
	resp, err := client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error Getting Strains!"
	}
	fmt.Println("data is ", string(data))
	return string(data)
}

func GetStrain() string {
	return ""
}

func AddStrain() string {
	return ""
}

func UpdateStrain() string {
	return ""
}

func RemoveStrain() string {
	return ""
}
