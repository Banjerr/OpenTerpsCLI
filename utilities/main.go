package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetData(BaseURL, dataType string) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", BaseURL+dataType, nil)
	req.Header.Add("APIUser", os.Getenv("APIUser"))
	req.Header.Add("APIKey", os.Getenv("APIKey"))
	resp, err := client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error Getting " + dataType + "!"
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(data), "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return "false"
	}
	fmt.Println(dataType+":", prettyJSON.String())
	return prettyJSON.String()
}
