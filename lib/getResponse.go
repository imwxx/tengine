package lib

import (
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
    "encoding/json"
)

func HttpGet(URL string) (string, int64) {
    timestamp := time.Now().Unix()
    client := &http.Client{}
    reqest, err := http.NewRequest("GET", URL, nil)
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(0)
    }
    reqest.Header.Add("Content-Type", "application/json; charset=UTF-8")
    response, _ := client.Do(reqest)
    var result string
    if response.StatusCode == 200 {
        body, _ := ioutil.ReadAll(response.Body)
        result = string(body)
    }
    return result, timestamp
}

func ResHandle() ([]byte, int64) {
    reqStatusArr := [...] string {"hostname","bytes_in","bytes_out","conn_total","req_total","http_2xx","http_3xx","http_4xx","http_5xx","http_other_status","rt","ups_req","ups_rt","ups_tries","http_200","http_206","http_302","http_304","http_403","http_404","http_416","http_499","http_500","http_502","http_503","http_504","http_508","http_other_detail_status","http_ups_4xx","http_ups_5xx"}
    URL := "http://127.0.0.1/us"
    response, timestamp := HttpGet(URL)
    res := strings.Split(response, "\n")
    hostResMap := make(map[string]map[string]string)
    for _, value := range res {
        valueArray := strings.Split(value, ",")
        if len(valueArray) != 1 {
            formatResMap := make(map[string]string)
            key := ""
            for index, value := range reqStatusArr {
                //fmt.Println(reflect.TypeOf(index))
                key = valueArray[0]
                formatResMap[value] = valueArray[index]
            }
            hostResMap[key] = formatResMap
        }
    }
    result, _ := json.Marshal(hostResMap)
    return result, timestamp
}
