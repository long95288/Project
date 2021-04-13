package main

import (
    "io/ioutil"
    "log"
)

func checkStartCode3(buf []byte) bool{
    if len(buf) < 3 {
        return false
    }
    if buf[0] == 0 && buf[1] == 0 && buf[2] == 1 {
        return true
    }
    return false
}
func checkStartCode4(buf []byte) bool{
    if len(buf) < 4 {
        return false
    }
    if buf[0] == 0 && buf[1] == 0 && buf[2] == 0 && buf[3] == 1  {
        return true
    }
    return false
}

func simplest_h264_parser(path string) {
    // 一直往前读,读到0x00 00 01之后开始分割
    data, err := ioutil.ReadFile(path)
    log.Println(data)
    log.Println(err)
    //file, err :=  os.Open(path)
    //if err != nil {
    //    log.Println(err)
    //    return
    //}
    //buf := make([]byte, 1)
    //index := 0
    //buffer := make([]byte, 0)
    //for _, err := file.Read(buf);err == nil; {
    //    index += 1
    //    buffer = append(buffer, buf...)
    //    if index <= 4 {
    //        continue
    //    }
    //
    //    if checkStartCode3(buffer) {
    //        buffer = buffer[3:]
    //        log.Println("startCode3 Index ", index)
    //    }else if checkStartCode4(buffer) {
    //        log.Println("startCode4 Index ", index)
    //        buffer = buffer[4:]
    //    }else{
    //        buffer = buffer[1:]
    //    }
    //}
}

func main() {
    simplest_h264_parser("aa.h264")
}
