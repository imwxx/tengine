package main

import (
    "fmt"
    "encoding/json"
    "os"
    "strconv"
    lib "tengine/lib"
)

func main() {
    response, timestamp := lib.ResHandle()
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

    var stdoutJson lib.StdoutJson
    for _, v := range jsonbody {
        hostname := v.Hostname
        reqTotalNum, _ := strconv.ParseFloat(v.Req_total, 32)
        oldreqTotalNum, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Req_total, 32)
        totalNum := reqTotalNum - oldreqTotalNum
        var totalNumTmp float64
        if totalNum > 0 {
            totalNumTmp = totalNum
        } else {
            totalNumTmp = 0
        }
        var reqTotal lib.SingelStdoutJson
        reqTotal.Endpoint = endpoint
        reqTotaltags := "flag=req_total," + "domain=" + hostname
        reqTotal.Tags = reqTotaltags
        reqTotal.Timestamp = timestamp
        reqTotal.Metric = "tengine.req.httpStatus.percent"
        reqTotal.Value = totalNumTmp
        reqTotal.CounterType = "GAUGE"
        reqTotal.Step = 60
        stdoutJson = append(stdoutJson, reqTotal)

        http200Num, _ := strconv.ParseFloat(v.Http_200, 32)
        oldhttp200Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_200, 32)
        h200Num := http200Num - oldhttp200Num
        var h200PercentTmp float64
        if h200Num > 0 {
            h200PercentTmp = h200Num / totalNum
        } else {
            h200PercentTmp = 0
        }
        var h200Stru lib.SingelStdoutJson
        h200Stru.Endpoint = endpoint
        h200tags := "flag=http_200," + "domain=" + hostname
        h200Stru.Tags = h200tags
        h200Stru.Timestamp = timestamp
        h200Stru.Metric = "tengine.req.httpStatus.percent"
        h200Stru.Value = h200PercentTmp
        h200Stru.CounterType = "GAUGE"
        h200Stru.Step = 60
        stdoutJson = append(stdoutJson, h200Stru)

        http5xxNum, _ := strconv.ParseFloat(v.Http_5xx, 32)
        oldhttp5xxNum, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_5xx, 32)
        h5xxNum := http5xxNum - oldhttp5xxNum
        var h5xxPercentTmp float64
        if h5xxNum > 0 {
            h5xxPercentTmp = h5xxNum / totalNum
        } else {
            h5xxPercentTmp = 0
        }
        var h5xxStru lib.SingelStdoutJson
        h5xxStru.Endpoint = endpoint
        h5xxtags := "flag=http_5xx," + "domain=" + hostname
        h5xxStru.Tags = h5xxtags
        h5xxStru.Timestamp = timestamp
        h5xxStru.Metric = "tengine.req.httpStatus.percent"
        h5xxStru.Value = h5xxPercentTmp
        h5xxStru.CounterType = "GAUGE"
        h5xxStru.Step = 60
        stdoutJson = append(stdoutJson, h5xxStru)

        http502Num, _ := strconv.ParseFloat(v.Http_502, 32)
        oldhttp502Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_502, 32)
        h502Num := http502Num - oldhttp502Num
        var h502NumTmp float64
        if h502Num > 0 {
            h502NumTmp = h502Num
        } else {
            h502NumTmp = 0
        }
        var h502Stru lib.SingelStdoutJson
        h502Stru.Endpoint = endpoint
        h502tags := "flag=http_502," + "domain=" + hostname
        h502Stru.Tags = h502tags
        h502Stru.Timestamp = timestamp
        h502Stru.Metric = "tengine.req.httpStatus.percent"
        h502Stru.Value = h502NumTmp
        h502Stru.CounterType = "GAUGE"
        h502Stru.Step = 60
        stdoutJson = append(stdoutJson, h502Stru)

        http500Num, _ := strconv.ParseFloat(v.Http_500, 32)
        oldhttp500Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_500, 32)
        h500Num := http500Num - oldhttp500Num
        var h500NumTmp float64
        if h500Num > 0 {
            h500NumTmp = h500Num
        } else {
            h500NumTmp = 0
        }
        var h500Stru lib.SingelStdoutJson
        h500Stru.Endpoint = endpoint
        h500tags := "flag=http_500," + "domain=" + hostname
        h500Stru.Tags = h500tags
        h500Stru.Timestamp = timestamp
        h500Stru.Metric = "tengine.req.httpStatus.percent"
        h500Stru.Value = h500NumTmp
        h500Stru.CounterType = "GAUGE"
        h500Stru.Step = 60
        stdoutJson = append(stdoutJson, h500Stru)

        http499Num, _ := strconv.ParseFloat(v.Http_499, 32)
        oldhttp499Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_499, 32)
        h499Num := http499Num - oldhttp499Num
        var h499PercentTmp float64
        if h499Num > 0 && totalNum > 180 {
            h499PercentTmp = h499Num / totalNum
        } else {
            h499PercentTmp = 0
        }

        var h499Stru lib.SingelStdoutJson
        h499Stru.Endpoint = endpoint
        h499tags := "flag=http_499," + "domain=" + hostname
        h499Stru.Tags = h499tags
        h499Stru.Timestamp = timestamp
        h499Stru.Metric = "tengine.req.httpStatus.percent"
        h499Stru.Value = h499PercentTmp
        h499Stru.CounterType = "GAUGE"
        h499Stru.Step = 60
        stdoutJson = append(stdoutJson, h499Stru)

        http404Num, _ := strconv.ParseFloat(v.Http_404, 32)
        oldhttp404Num, _ := strconv.ParseFloat(oldjsonbody[v.Hostname].Http_404, 32)
        h404Num := http404Num - oldhttp404Num
        var h404PercentTmp float64
        if h404Num > 0 && totalNum > 180 {
            h404PercentTmp = h404Num / totalNum
        } else {
            h404PercentTmp = 0
        }
        var h404Stru lib.SingelStdoutJson
        h404Stru.Endpoint = endpoint
        h404tags := "flag=http_404," + "domain=" + hostname
        h404Stru.Tags = h404tags
        h404Stru.Timestamp = timestamp
        h404Stru.Metric = "tengine.req.httpStatus.percent"
        h404Stru.Value = h404PercentTmp
        h404Stru.CounterType = "GAUGE"
        h404Stru.Step = 60
        stdoutJson = append(stdoutJson, h404Stru)
    }
    Re, _ := json.Marshal(stdoutJson)
    fmt.Println(string(Re))

    lib.Write(response)
}
