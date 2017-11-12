package lib

import (
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Write(jsonData []byte) {
    pwd, _ := os.Getwd()
    if jsonData != nil {
        err := ioutil.WriteFile(pwd + "/jsondata", jsonData, 0644)
        check(err)
    }
}

func Read() []byte {
    pwd, _ := os.Getwd()
    file := pwd + "/jsondata"
    var data []byte
    if ExistCheck() {
        dat, err := ioutil.ReadFile(file)
        check(err)
        data = dat
    }
    return data
}

func ExistCheck() bool {
    pwd, _ := os.Getwd()
    file := pwd + "/jsondata"
    _, err := os.Stat(file)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}
