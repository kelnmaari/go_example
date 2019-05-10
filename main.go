package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://mc-t.ru/"

func mct_api(method string, kwargs map[string] string) string{

	buffer := "{"
	for key, value := range kwargs {
		buffer += "\"" + key + "\":\"" + value + "\","
	}
	buffer = buffer[:len(buffer) - 1]
	buffer += "}"

	var jsonStr = []byte(`{"method":"`+ method +`","kwargs":`+ buffer +`}`)
	req, err := http.NewRequest("POST", url + "api/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	var arr = make(map[string]string)

	arr["msg"] = "hello"

	fmt.Println(mct_api("getEcho", arr))

}
