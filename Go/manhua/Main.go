package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "time"
)

/**
全局变量区
 */
var day string = string(time.Now().Format("2006_01_02"))
var wd,_ = os.Getwd()
var log_file string = wd +"\\" + day + "_log.txt"
var rpc_url = "http://localhost:6800/jsonrpc"

var root_url string = ""
var index_url string
var list_prefix string
var totalPageNumber int
var images_pattern string
var image_pattern string
var script_tmp_path,_ = os.Getwd()
var script_tmp_name = "tmp.html"
var base_save_dir string

func log(s string,flag bool) {
    var now_data string = time.Now().Format("2006/1/2 15:04:05")
    log_message := fmt.Sprintf("时间:%s : log : %s \n",now_data,s)
    fl,err := os.OpenFile(log_file,os.O_APPEND|os.O_CREATE,0644)
    if err != nil{
        fmt.Println("打开文件失败")
    }
    defer fl.Close()
    n,err := fl.WriteString(log_message)
    if err != nil || n < len(log_message){
        fmt.Println("保存日志失败")
    }
    if flag {
        fmt.Println(log_message)
    }
}
func initConfig() {
   var f interface{}
   data,err := ioutil.ReadFile("conf.json")
   err = json.Unmarshal(data,&f)
   if err == nil{
       conf := f.(map[string]interface{})
       root_url = conf["root_url"].(string)
       index_url = conf["index_url"].(string)
       list_prefix = conf["list_prefix"].(string)
       totalPageNumber =int(conf["total_page_number"].(float64))
       image_pattern = conf["image_pattern"].(string)
       images_pattern = conf["images_pattern"].(string)
   }
    println("================conf===================")
    println(fmt.Sprintf("root_url:%s",root_url))
    println(fmt.Sprintf("index_url:%s",index_url))
    println(fmt.Sprintf("list_prefix:%s",list_prefix))
    println(fmt.Sprintf("totalPageNumber:%d",totalPageNumber))
    println(fmt.Sprintf("image_pattern:%s",image_pattern))
    println(fmt.Sprintf("images_pattern:%s",images_pattern))
    println("=======================================")
}

func main() {
    initConfig()
    
    log("Hello world",true)
    var download_list []string
    data,err := ioutil.ReadFile("list.json")
    if err == nil{
        var f []interface{}
        json.Unmarshal(data,&f)
        for _,data := range f{
            download_list = append(download_list, data.(string))
        }
    }
    for i := 1; i < totalPageNumber; i++ {
        index := fmt.Sprintf("%d",i)
        new_list_url := index_url + strings.Replace(list_prefix,"{}",index,1)
        url_list := get_index_info(new_list_url)
        for _,i := range url_list{
            println(i)
        }
    }
}
func addDownloadTask(url string, dir string, out string) string {
    type rpcParams struct {
        urls []string
        options map[string]string
    }
    type requestBody struct {
        jsonrpc string
        id string
        method string
        params []string
    }
    body := requestBody{
        jsonrpc: "2.0",
        id:      "QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA==",
        method:  "aria2.addUri"
    }
    bodyString,_ := json.Marshal(body)
    res,err := http.Post(rpc_url,"application/json",strings.NewReader(string(bodyString)))
    if err == nil{
        if res.StatusCode == 200{
            var re []byte
            res.Body.Read(re)
            var f interface{}
            json.Unmarshal(re,&f)
            return f.(map[string]string)["result"]
        }else{
            log("无法调用aria2c",true)
        }
    }
    return ""
}
func download_status(gid string) string {
    postdatax
    return ""
}
func download(url string,dir string,out string) int {
    log(fmt.Sprintf("开始下载:%s",url),true)
    gid := addDownloadTask(url,dir,out)
    status := download_status(gid)
    for true {
        if status == "active" {
            time.Sleep(time.Duration(2)*time.Second)
            
        }
    }
    return -1
}

func htmlContent(url string) (string, error) {
    status := download(url,script_tmp_path,script_tmp_name)
    if status == 0{
        data,err := ioutil.ReadFile(script_tmp_name)
        if err == nil {
             return string(data),err
        }
    }
    return _,Error("无法读取文件")
}

func get_index_info(url string) []string {
    var result []string
    response_data,err = htmlContent(url)
    if err == nil{
    
    }
    return result
}


