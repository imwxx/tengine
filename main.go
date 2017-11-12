package main

import (
    "fmt"
    "encoding/json"
    "os"
    "strconv"
    lib "tengine/lib"
//    "reflect"
)

func main() {
    response, timestamp := lib.ResHandle()
    //var jsonbody map[int]map[string]string
    var jsonbody map[string]lib.SingleJson
    err := json.Unmarshal(response, &jsonbody)
    if err != nil {
        panic(err)
    }
    if lib.ExistCheck() == false {
        lib.Write(response)
        os.Exit(1)
    }

    oldres := lib.Read()
    var oldjsonbody map[string]lib.SingleJson
    olderr := json.Unmarshal(oldres, &oldjsonbody)
    if olderr != nil {
        panic(olderr)
    }
    endpoint, Err := lib.GetHostname()
    if Err != 1 {
        fmt.Println("get hostname failed")
        os.Exit(1)
    }

//dataInfo := url.Values{}
//dataInfo.Set("email", Username[0])
//dataInfo.Add("content", content)
//DATA := dataInfo.Encode()
    var stdoutJson lib.StdoutJson
    for _, v := range jsonbody {
        hostname := v.Hostname
        reqTotalNum, _ := strconv.ParseFloat(v.Req_total, 32)
        oldreqTotalNum, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Req_total, 32)
        totalNum := reqTotalNum - oldreqTotalNum
        if totalNum > 0 {
            var reqTotal lib.SingelStdoutJson
            reqTotal.Endpoint = endpoint
            reqTotaltags := "flag=req_total,"+"domain="+hostname
            reqTotal.Tags = reqTotaltags
            reqTotal.Timestamp = timestamp
            reqTotal.Metric = "tengine.req.httpStatus.percent"
            reqTotal.Value = totalNum
            reqTotal.CounterType = "GAUGE"
            reqTotal.Step = 60
            stdoutJson = append(stdoutJson, reqTotal)
        }
        http200Num, _ := strconv.ParseFloat(v.Http_200, 32)
        oldhttp200Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_200, 32)
        h200Num := http200Num - oldhttp200Num
        if h200Num > 0 {
            h200Percent := h200Num / totalNum
            var h200Stru lib.SingelStdoutJson
            h200Stru.Endpoint = endpoint
            h200tags := "flag=http_200,"+"domain="+hostname
            h200Stru.Tags = h200tags
            h200Stru.Timestamp = timestamp
            h200Stru.Metric = "tengine.req.httpStatus.percent"
            h200Stru.Value = h200Percent
            h200Stru.CounterType = "GAUGE"
            h200Stru.Step = 60
            stdoutJson = append(stdoutJson, h200Stru)
        }

        http5xxNum, _ := strconv.ParseFloat(v.Http_5xx, 32)
        oldhttp5xxNum, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_5xx, 32)
        h5xxNum := http5xxNum - oldhttp5xxNum
        if h5xxNum > 0 {
            h5xxPercent := h5xxNum / totalNum
            var h5xxStru lib.SingelStdoutJson
            h5xxStru.Endpoint = endpoint
            h5xxtags := "flag=http_5xx,"+"domain="+hostname
            h5xxStru.Tags = h5xxtags
            h5xxStru.Timestamp = timestamp
            h5xxStru.Metric = "tengine.req.httpStatus.percent"
            h5xxStru.Value = h5xxPercent
            h5xxStru.CounterType = "GAUGE"
            h5xxStru.Step = 60
            stdoutJson = append(stdoutJson, h5xxStru)
        }

        http502Num, _ := strconv.ParseFloat(v.Http_502, 32)
        oldhttp502Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_502, 32)
        h502Num := http502Num - oldhttp502Num
        if h502Num > 0 {
            var h502Stru lib.SingelStdoutJson
            h502Stru.Endpoint = endpoint
            h502tags := "flag=http_502,"+"domain="+hostname
            h502Stru.Tags = h502tags
            h502Stru.Timestamp = timestamp
            h502Stru.Metric = "tengine.req.httpStatus.percent"
            h502Stru.Value = h502Num
            h502Stru.CounterType = "GAUGE"
            h502Stru.Step = 60
            stdoutJson = append(stdoutJson, h502Stru)
        }

        http500Num, _ := strconv.ParseFloat(v.Http_500, 32)
        oldhttp500Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_500, 32)
        h500Num := http500Num - oldhttp500Num
        if h500Num > 0 {
            var h500Stru lib.SingelStdoutJson
            h500Stru.Endpoint = endpoint
            h500tags := "flag=http_500,"+"domain="+hostname
            h500Stru.Tags = h500tags
            h500Stru.Timestamp = timestamp
            h500Stru.Metric = "tengine.req.httpStatus.percent"
            h500Stru.Value = h500Num
            h500Stru.CounterType = "GAUGE"
            h500Stru.Step = 60
            stdoutJson = append(stdoutJson, h500Stru)
        }

        //http4xxNum, _ := strconv.ParseFloat(v.Http_4xx, 32)
        //oldhttp4xxNum, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_4xx, 32)
        //h4xxNum := http4xxNum - oldhttp4xxNum
        //if h4xxNum > 0 {
        //    h4xxPercent := h4xxNum / totalNum
        //    var h4xxStru lib.SingelStdoutJson
        //    h4xxStru.Endpoint = endpoint
        //    h4xxtags := "flag=http_4xx,"+"domain="+hostname
        //    h4xxStru.Tags = h4xxtags
        //    h4xxStru.Timestamp = timestamp
        //    h4xxStru.Metric = "tengine.req.httpStatus.percent"
        //    h4xxStru.Value = h4xxPercent
        //    h4xxStru.CounterType = "GAUGE"
        //    h4xxStru.Step = 60
        //    stdoutJson = append(stdoutJson, h4xxStru)
        //}

        http499Num, _ := strconv.ParseFloat(v.Http_499, 32)
        oldhttp499Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_499, 32)
        h499Num := http499Num - oldhttp499Num
        if h499Num > 0 && totalNum > 180 {
            h499Percent := h499Num / totalNum
            var h499Stru lib.SingelStdoutJson
            h499Stru.Endpoint = endpoint
            h499tags := "flag=http_499,"+"domain="+hostname
            h499Stru.Tags = h499tags
            h499Stru.Timestamp = timestamp
            h499Stru.Metric = "tengine.req.httpStatus.percent"
            h499Stru.Value = h499Percent
            h499Stru.CounterType = "GAUGE"
            h499Stru.Step = 60
            stdoutJson = append(stdoutJson, h499Stru)
        }

        http404Num, _ := strconv.ParseFloat(v.Http_404, 32)
        oldhttp404Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_404, 32)
        h404Num := http404Num - oldhttp404Num
        if h404Num > 0 && totalNum > 180 {
            h404Percent := h404Num / totalNum
            var h404Stru lib.SingelStdoutJson
            h404Stru.Endpoint = endpoint
            h404tags := "flag=http_404,"+"domain="+hostname
            h404Stru.Tags = h404tags
            h404Stru.Timestamp = timestamp
            h404Stru.Metric = "tengine.req.httpStatus.percent"
            h404Stru.Value = h404Percent
            h404Stru.CounterType = "GAUGE"
            h404Stru.Step = 60
            stdoutJson = append(stdoutJson, h404Stru)
        }
    }
        Re, _ := json.Marshal(stdoutJson)
        fmt.Println(string(Re))

    lib.Write(response)
}
