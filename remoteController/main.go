package main

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "html/template"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

type status struct {
    Status string
}
var globalStatus = status{Status: "正常"}
const TIME_LAYOUT = "2006-01-02T15:04"
var exit_app_server = false

// 连接结构体,
type RCTLConnection struct {
    conn net.Conn
    readQueue []byte
    writeQueue []byte
}

/*
* app连接线程
*/
func appConnHandler(conn net.Conn) {
    if nil == conn {
        log.Println("错误APP连接")
        return
    }
    go func() {
        for !exit_app_server {
            readbuf := make([]byte, 4096)
            n,err := conn.Read(readbuf)
            if nil != err || n < 0 {
                log.Println("读取数据失败")
                break
            }
            fmt.Printf("App server read data size = %d\n", n)
            fmt.Println("App server read data :" + string(readbuf[:n + 1]))
            time.Sleep(100*time.Millisecond)
        }
    }()
    
    go func() {
        for !exit_app_server {
            writebuf := []byte("SERVER SAY HELLO")
            n, err := conn.Write(writebuf)
            if nil != err || n != len(writebuf) {
                fmt.Println("App Server write data failed.\n")
                break
            }
            time.Sleep(100*time.Millisecond)
        }
    }()
}
func server_notice() {
    upd, err := net.ResolveUDPAddr("udp4", ":1400")
    conn, err := net.ListenUDP("udp4", upd)
    defer conn.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    readBuf := make([]byte, 1024)
    writeBuf := []byte("321")
    for true {
        time.Sleep(50 * time.Millisecond)
        n, remoteAddr, err := conn.ReadFromUDP(readBuf)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("server notice read data size = %d\n", n)
        fmt.Println("客户端探测,ip:", remoteAddr.IP, "端口:", remoteAddr.Port, "数据:", string(readBuf[:n]))
        peerData := string(readBuf[:n])
        fmt.Println("peer data..", peerData)
        if strings.Compare("123", peerData) == 0{
            peerConn, err := net.DialUDP("udp4", nil, remoteAddr)
            if err != nil {
                fmt.Println(err)
                continue
            }
            fmt.Println("发送响应....")
            peerConn.Write(writeBuf)
            peerConn.Close()
        }else{
            fmt.Println("未知数据....")
        }
    }
}
func App_Server()  {
    server, err := net.Listen("udp", ":1399")
    if nil != err {
        fmt.Println(err)
    }
    for !exit_app_server {
        fmt.Println("等待连接....")
        conn, err := server.Accept()
        fmt.Println("接受连接....")
        if nil != err{
            fmt.Println(err)
            
        }else{
            go appConnHandler(conn)
        }
        time.Sleep(200 * time.Millisecond)
    }
}

func Http_Server()  {
    e:= gin.Default()
    e.POST("/cmd",HandleController)
    e.GET("/", func(context *gin.Context) {
        t,err := template.ParseFiles("index.html")
        if err != nil {
            log.Println("err : ", err)
            return
        }
        globalStatus.Status = "正常"
        log.Println(t.Execute(context.Writer, globalStatus))
    })
    
    e.POST("/SetShutdownTime", func(context *gin.Context) {
        t,err := template.ParseFiles("index.html")
        if err != nil {
            log.Println("err : ", err)
            return
        }
        r := context.Request
        err = r.ParseForm()
        if err != nil {
            t.Execute(context.Writer, status{Status: err.Error()})
            return
        }
        arrTime := r.Form["time"]
        if len(arrTime) < 1 {
            t.Execute(context.Writer, status{Status: "未设置关机时间"})
            return
        }
        shutdownTime := arrTime[0]
        if "" == shutdownTime {
            t.Execute(context.Writer, status{Status: "错误时间格式"})
            return
        }
        times , err := time.Parse(TIME_LAYOUT,shutdownTime)
        if nil != err {
            t.Execute(context.Writer, status{Status: "错误时间格式"})
            return
        }
        d := times.Unix() - time.Now().Unix()
        if d > 0 {
            args :=[]string{"-s","-t"}
            args = append(args, strconv.FormatInt(d, 10))
            cmd := exec.Command("shutdown",args...)
            err := cmd.Run()
            if err != nil {
                t.Execute(context.Writer, status{"设置关机任务失败"})
            }else{
                globalStatus.Status = "电脑将于:"+ shutdownTime + "关闭"
                t.Execute(context.Writer, globalStatus)
            }
            return
        }
        t.Execute(context.Writer, status{Status: "错误时间格式"})
        return
    })
    err := e.Run(":9999")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("service start at http://localhost:9999")
}

type CMDCommonRequestDTO struct {
    CMD int `json:"CMD"`
}

func HandleController(c *gin.Context)  {
    body,err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        ResponseError(c, 200001, "获取body失败")
        return
    }
    cmdDTO := CMDCommonRequestDTO{}
    err = json.Unmarshal(body, &cmdDTO)
    if err != nil {
        ResponseError(c, 200002, "解析body失败")
        return
    }
    switch cmdDTO.CMD {
    case SETSHUTDOWNPLAN_CMD:
        setShutdownPlanHandler(c, body)
    case CANCELSHUTDOWNPLAN_CMD:
        cancelShutdownPlanHandler(c, body)
    case GETMASTERVOLUME_CMD:
        getMasterVolumeHandler(c, body)
    case SETMASTERVOLUME_CMD:
        setMasterVolumeHandler(c, body)
    default:
        c.JSON(
            http.StatusBadRequest,
            gin.H{
                "message":"bad Request",
            })
    }
}
func ResponseError(c *gin.Context, status int, message string){
    c.JSON( http.StatusOK, gin.H{ "OptionStatus": status, "Message": message})
    return
}
func ResponseSuccess(c *gin.Context, message string, data interface{}){
    c.JSON(http.StatusOK, gin.H{"OptionStatus":2000000, "Message":message, "Data":data})
}


type SetShutdownPlanRequestDTO struct {
    ShutdownTime int `json:"ShutdownTime"`
}
type SetShutdownPlanResponseDTO struct {
    ShutdownTime int `json:"ShutdownTime"`
}
func setShutdownPlanHandler(c *gin.Context, body []byte) {
    var dto SetShutdownPlanRequestDTO
    err := json.Unmarshal(body, &dto)
    if err != nil {
        ResponseError(c, 210001, "解析body失败")
        return
    }
    retMsg, err := SetShutdownPlan(dto.ShutdownTime)
    if err != nil {
        ResponseError(c, 210002, fmt.Sprintf("执行计划失败,%s", err))
        return
    }
    ResponseSuccess(c, retMsg, SetShutdownPlanResponseDTO{ShutdownTime: dto.ShutdownTime})
}
func cancelShutdownPlanHandler(c *gin.Context, body []byte) {
    _,err := CancelShutdownPlan()
    if err != nil {
        ResponseError(c, 220001, fmt.Sprintf("执行失败,err:%v", err))
        return
    }
    ResponseSuccess(c, "success", nil)
    return
}

type GetMasterVolumeRequestDTO struct {
}
type GetMasterVolumeResponseDTO struct {
    Volume int `json:"Volume"`
}
func getMasterVolumeHandler(c *gin.Context, body []byte) {
    volume, err := GetMasterVolume()
    if err != nil {
        ResponseError(c, 230001, fmt.Sprintf("查询主音量失败,%s", err))
        return
    }
    ResponseSuccess(c, "success", GetMasterVolumeResponseDTO{Volume: int(volume)})
}

type SetMasterVolumeRequestDTO struct {
    Volume int `json:"Volume"`
}
type SetMasterVolumeResponseDTO struct {
    Volume int `json:"Volume"`
}
func setMasterVolumeHandler(c *gin.Context,body []byte) {
    requestDTO := SetMasterVolumeRequestDTO{}
    err := json.Unmarshal(body, &requestDTO)
    if err != nil {
        ResponseError(c, 240001, "解析body参数失败")
        return
    }
    volume, err := SetMasterVolume(float64(requestDTO.Volume))
    if err != nil {
        ResponseError(c, 240002, fmt.Sprintf("执行命令失败,%v", err))
        return
    }
    ResponseSuccess(c, "success", SetMasterVolumeResponseDTO{Volume: int(volume)})
}



func main() {
    // go App_Server()
    go server_notice()
    Http_Server()
}

