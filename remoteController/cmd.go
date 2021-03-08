package main
// 命令行,只做中间层,不负责HTTP相关操作
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "os/exec"
    "strconv"
)
const (
    SETSHUTDOWNPLAN_CMD = 1
    CANCELSHUTDOWNPLAN_CMD = 2
    GETMASTERVOLUME_CMD = 3
    SETMASTERVOLUME_CMD  = 4
    GETSCREENCAPTURE_CMD = 5
)
var WindowControllerCmd = "WindowsRemoteController.exe"
func SetMasterVolume(volume float64) (float64, error) {
    args := []string{"2"}
    args =  append(args, fmt.Sprintf("%f", volume))
    cmd := exec.Command(WindowControllerCmd, args...)
    output, err := cmd.CombinedOutput()
    if nil != err {
        return 0, err
    }
    volume1, err := strconv.ParseFloat(string(output), 10)
    return volume1, err
}
func GetMasterVolume() (float64, error) {
    args := []string{"1"}
    cmd := exec.Command(WindowControllerCmd, args...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return 0, err
    }
    volume, err := strconv.ParseFloat(string(output), 10)
    return volume, err
}
func SetShutdownPlan(time int) (string, error) {
    args := []string{"-s", "-t"}
    args = append(args,fmt.Sprintf("%d", time))
    cmd := exec.Command("shutdown", args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}
func CancelShutdownPlan() (string, error) {
    args := []string{"-a"}
    cmd := exec.Command("shutdown", args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

func GetScreenCapture() ([]byte, error){
    args := []string{"3", "test.bmp"}
    cmd := exec.Command(WindowControllerCmd, args...)
    _, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }
    output, err := ioutil.ReadFile(args[1])
    return output, err
}
func shutdown2(c *gin.Context)  {
    args :=[]string{"-s","-t","30"}
    cmd := exec.Command("shutdown",args...)
    err := cmd.Run()
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message":"服务器内部错误"})
    }else{
        globalStatus.Status = "电脑将在30s后关机"
        c.JSON(http.StatusOK, gin.H{"message":globalStatus.Status})
    }
}
func shutdown(c *gin.Context){
    args :=[]string{"-s","-t","30"}
    t,err := template.ParseFiles("index.html")
    if err != nil {
        log.Println("err : ", err)
        return
    }
    cmd := exec.Command("shutdown",args...)
    err = cmd.Run()
    if err != nil {
        t.Execute(c.Writer, status{"设置关机任务失败"})
    }else{
        globalStatus.Status = "电脑将在30s后关机"
        t.Execute(c.Writer,globalStatus )
    }
}

func cancelShutdown(c *gin.Context){
    args := []string{"-a"}
    cmd := exec.Command("shutdown",args...)
    err := cmd.Run()
    if err != nil {
        c.JSON(http.StatusInternalServerError,gin.H{
            "message":fmt.Sprintf("取消关机失败,err : %v",err),
        })
    }else{
        c.JSON(http.StatusOK,gin.H{
            "message":"取消关机成功",
        })
    }
}