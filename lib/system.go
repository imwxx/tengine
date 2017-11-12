package lib

import (
    "os"
)

func GetHostname() (string, int){
    host, err := os.Hostname()
    var Host string
    var Err int = 0
    if err != nil {
        Host = "null"
        Err = 0
    } else {
        Host = host
        Err = 1
    }
    return Host, Err
}
