package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// 模拟内存分配

type mem struct {
    firstp int
    lastp int
    callocable bool
}
type Memory struct {
    firstP int
    lastP int
    callocTable []mem
    idleTable []mem
}

func allocMem(m Memory, n int) string {
    if n == 0 {
        return "error"
    }
    for i:=0; i<len(m.idleTable); i++{
        if m.idleTable[i].lastp - m.idleTable[i].firstp + 1 > n {
            m.callocTable = append(m.callocTable, mem{
                firstp:     m.idleTable[i].firstp,
                lastp:      m.idleTable[i].firstp + n - 1,
                callocable: false,
            })
            p := strconv.Itoa(m.idleTable[i].firstp)
            m.idleTable[i].firstp = m.idleTable[i].firstp + n
            return p
        }
    }
    return "error"
}
func freeMem(m Memory, n int) string {
    for i:=0; i < len(m.callocTable); i-- {
        if m.callocTable[i].firstp == n {
            for j:=0; j < len(m.idleTable); j++ {
                if m.idleTable[j].lastp + 1 == m.callocTable[i].firstp {
                    m.idleTable[j].lastp = m.callocTable[i].lastp
                }
                if m.idleTable[j].firstp == m.callocTable[i].lastp + 1 {
                    m.idleTable[j].firstp = m.callocTable[i].firstp
                }
            }
            m.idleTable = append(m.idleTable, m.callocTable[i])
            m.callocTable = append(m.callocTable[ :i], m.callocTable[i+1:]...)
            return ""
        }
    }
    return "error"
}
func main() {
    // 维护一个分配表和空闲表
    op := []string{}
    m := Memory{
        firstP:      0,
        lastP:       99,
        callocTable: []mem{},
        idleTable: []mem{{
            firstp:     0,
            lastp:      99,
            callocable: false,
        }},
    }
    in := bufio.NewReader(os.Stdin)
    for str,err := in.ReadString('\n');err == nil;str,err = in.ReadString('\n') {
        str = strings.Replace(str,"\r\n", "", -1)
        str = strings.Replace(str,"\n", "", -1)
        if "" == str {
            break
        }
        op = append(op, str)
    }
    for i := 1; i < len(op); i++{
        cmd :=strings.Split(op[i], "=")
        if "REQUEST" == cmd[0] {
            n,_ := strconv.Atoi(cmd[1])
            fmt.Println(allocMem(m, n))
        }else if "RELEASE" == cmd[0] {
            n, _ := strconv.Atoi(cmd[1])
            fmt.Println(freeMem(m, n))
        }
    }
}
