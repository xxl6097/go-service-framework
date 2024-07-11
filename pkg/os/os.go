package os

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func IsMacOs() bool {
	if strings.Compare(runtime.GOOS, "darwin") == 0 {
		return true
	}
	return false
}

// GetOsInfo
// <table>
// <tr>
// <td>Data 1</td>
// <td>Data 2</td>
// </tr>
// <tr>
// <td>Data 4</td>
// <td>Data 5</td>
// </tr>
// </table>
func GetOsInfo() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	info := "<table>"
	info += "<tr>"
	info += "<td>系统名称</td>"
	info += "<td>" + runtime.GOOS + "</td>"
	info += "</tr>"
	info += "<td>系统架构</td>"
	info += "<td>" + runtime.GOARCH + "</td>"
	info += "</tr>"
	info += "<td>cup数量</td>"
	info += "<td>" + strconv.Itoa(runtime.NumCPU()) + "</td>"
	info += "</tr>"
	info += "<td>GOMIPS</td>"
	info += "<td>" + os.Getenv("GOMIPS") + "</td>"
	info += "</tr>"
	info += "<td>Goroutine数量</td>"
	info += "<td>" + strconv.Itoa(runtime.NumGoroutine()) + "</td>"
	info += "</tr>"
	info += "<td>内存分配</td>"
	info += "<td>" + fmt.Sprintf("%d", m.Alloc/1024/1024) + "G</td>"
	info += "</tr>"
	info += "<td>总内存分配</td>"
	info += "<td>" + fmt.Sprintf("%d", m.TotalAlloc/1024/1024) + "G</td>"
	info += "</tr>"
	info += "<td>GC次数</td>"
	info += "<td>" + fmt.Sprintf("%d", m.NumGC) + "</td>"
	info += "</tr>"
	info += "<td>Go Version</td>"
	info += "<td>" + runtime.Version() + "</td>"
	info += "</tr>"
	info += "</table>"
	//info += "系统名称：" + runtime.GOOS
	//info += "\r\n系统架构：" + runtime.GOARCH
	//info += "\r\ncup数量：" + strconv.Itoa(runtime.NumCPU())
	//info += "\r\nGOMIPS：" + os.Getenv("GOMIPS")
	//info += "\r\nGoroutine数量：" + strconv.Itoa(runtime.NumGoroutine())
	//info += fmt.Sprintf("内存分配:%dMB", m.Alloc/1024/1024)
	//info += fmt.Sprintf("总内存分配%dMB", m.TotalAlloc/1024/1024)
	//info += fmt.Sprintf("GC次数:%d", m.NumGC)
	//info += fmt.Sprintf("Go Version: %s", runtime.Version())
	return info
}

func SomeFunction() {
	var pc [1]uintptr
	n := runtime.Callers(1, pc[:])
	if n == 0 {
		return
	}
	fn := runtime.FuncForPC(pc[0])
	fmt.Printf("Function call: %s\n", fn.Name())
}
