package lib

//import (
//    "fmt"
//    "encoding/json"
//)

type SingleJson struct {
    Hostname                    string   `json:"hostname"` 
    Bytes_in                    string    `json:"bytes_in"`
    Bytes_out                   string    `json:"bytes_out"`
    Conn_total                  string    `json:"conn_total"`
    Req_total                   string    `json:"req_total"`
    Http_2xx                    string    `json:"http_2xx"`
    Http_3xx                    string    `json:"http_3xx"`
    Http_4xx                    string    `json:"http_4xx"` 
    Http_5xx                    string    `json:"http_5xx"`
    Http_other_status           string    `json:"http_other_status"`
    Rt                          string    `json:"rt"`
    Ups_req                     string    `json:"ups_req"`
    Ups_rt                      string    `json:"ups_rt"`
    Ups_tries                   string    `json:"ups_tries"`
    Http_200                    string    `json:"http_200"` 
    Http_206                    string    `json:"http_206"`
    Http_302                    string    `json:"http_302"`
    Http_304                    string    `json:"http_304"`
    Http_403                    string    `json:"http_403"`
    Http_404                    string    `json:"http_404"`
    Http_416                    string    `json:"http_416"`
    Http_499                    string    `json:"http_499"`
    Http_500                    string    `json:"http_500"`
    Http_502                    string    `json:"http_502"`
    Http_503                    string    `json:"http_503"`
    Http_504                    string    `json:"http_504"`
    Http_508                    string    `json:"http_508"`
    Http_other_detail_status    string    `json:"http_other_detail_status"`
    Http_ups_4xx                string    `json:"http_ups_4xx"` 
    Http_ups_5xx                string    `json:"http_ups_5xx"`
}

type NoticeJson struct {
    Trigger_id  string  `json:"trigger_id"`
    Endpoint    string  `json:"endpoint"`
    Metric      string  `json:"metric"`
    Metric_id   string  `json:"metric_id"`
    Value       string  `json:"value"`
    Tags        string  `json:"tags"`
    Note        string  `json:"note"`
    Severity    string  `json:"severity"`
    Status      string  `json:"status"`
    Message     string  `json:"message"`
}

type SingelStdoutJson struct {
    Endpoint    string  `json:"endpoint"`
    Tags        string  `json:"tags"`
    Timestamp   int64  `json:"timestamp"`
    Metric      string  `json:"metric"`
    Value       float64  `json:"value"`
    CounterType string  `json:"counterType"`
    Step        int  `json:"step"`
}

type StdoutJson []SingelStdoutJson
