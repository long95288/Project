package main

import (
    "context"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "sync"
    "syscall"
    "time"
)
const RUNNING = 1
const STOP    = 2

var serviceStatus int = 1
var server *http.Server
var lock sync.Mutex
var exitCh chan int
func StartService() {
    log.Println("StartService")
    lock.Lock()
    serviceStatus = RUNNING
    lock.Unlock()
    
    // 开启服务
    exitCh = make(chan int)
    for serviceStatus == RUNNING {
        log.Println("Service is running")
        // code ....
        
        
        
        
        time.Sleep(3 * time.Second)
    }
    close(exitCh)
}
func StopService() {
    log.Println("StopService")
    // 关闭
    lock.Lock()
    serviceStatus = STOP
    lock.Unlock()
    serviceExit := false
    timeout := 2 * time.Second
    for !serviceExit {
        select {
        case <- exitCh:
            serviceExit = true
        case <- time.After(timeout):
            serviceExit = true
        }
    }
    log.Println("Shutdown Http Service")
    err := server.Shutdown(context.Background())
    if err != nil {
        log.Printf("Service Shutdown , err %s\n", err)
        os.Exit(-1)
    }else{
        log.Println("Service Shutdown")
        os.Exit(0)
    }
}
func ReloadService() {
    log.Println("ReloadService")
    // 关闭start服务,重新启动
    lock.Lock()
    serviceStatus = STOP
    lock.Unlock()
    serviceExit := false
    timeout := 2 * time.Second
    for !serviceExit {
        select {
        case <- exitCh:
            serviceExit = true
        case <- time.After(timeout):
            serviceExit = true
        }
    }
    StartService()
}

func RequestHandler(writer http.ResponseWriter, request *http.Request) {
    if err := request.ParseForm(); err != nil {
        log.Printf("RequestHandler ParseForm %s\n", err)
        return
    }
    for key, value := range request.Form {
        // 取出值
        log.Printf("Request key:%s, value:%s", key, value)
        if key == "cmd" && len(value) > 0 {
           
           if "status" == value[0] {
               writer.Write([]byte(fmt.Sprintf("%d", serviceStatus)))
           }else if "reload" == value[0] {
               writer.Write([]byte(value[0]))
               ReloadService()
           }else if "stop" == value[0] {
               writer.Write([]byte(value[0]))
               StopService()
           }else if "start" == value[0] {
               writer.Write([]byte(value[0]))
               ReloadService()
           } else {
               log.Printf("UNKNOWN CMD %s\n", value[0])
               writer.Write([]byte(value[0]))
           }
           
           return
        }
    }
}

func getServiceStatus() (int, error) {
    out, err := sendCmd("status")
    if err != nil {
        return STOP, nil
    }
    status, err := strconv.Atoi(string(out))
    if err != nil {
        return 0, err
    }
    if status != RUNNING && status != STOP {
        return 0, errors.New("UnKnown Status")
    }
    return status, nil
}

func sendCmd(cmd string) (out []byte, err error) {
    url := fmt.Sprintf("http://localhost:56001/?cmd=%s", cmd)
    request, err := http.NewRequest("GET", url,nil)
    if err != nil {
        log.Printf("create request %s faile, %s\n",url, err)
        return nil, err
    }
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    out, err = ioutil.ReadAll(response.Body)
    return
}
func main() {
    args := os.Args
    cmd := ""
    if len(args) <= 1 {
        cmd = "start"
    }else {
        cmd = args[1]
    }
    status, err := getServiceStatus()
    if err != nil {
        os.Exit(-1)
    }
    if status == RUNNING {
        log.Println("Service is Active")
        if cmd == "start" || cmd == "reload" || cmd == "stop" {
            _, err := sendCmd(cmd)
            if err != nil {
                log.Printf("send cmd err , %v", err)
            }
            return
        }else{
            log.Printf("UNKNOW CMD %s\n", cmd)
            return
        }
    }
    
    // 不是处于启动状态
    srvMux := http.NewServeMux()
    srvMux.HandleFunc("/", RequestHandler)
    server = &http.Server{
        Addr: ":56001",
        Handler: srvMux,
    }
    go func() {
        err := server.ListenAndServe()
        if err != nil {
            log.Printf("Start service Failed, err %s", err)
        }
    }()
    
    // 处理退出
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    select {
    case sig := <-sigs:
       log.Printf("Notify Signal %v\n", sig)
       serviceStatus = STOP
       err := server.Shutdown(context.Background())
       if err != nil {
           log.Printf("Service shutdown , err %s\n", err)
           os.Exit(-1)
       }else{
           log.Println("Service Shutdown")
           os.Exit(0)
       }
    }
}
