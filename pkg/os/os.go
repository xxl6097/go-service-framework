package os

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

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

func IsMacOs() bool {
	if strings.Compare(runtime.GOOS, "darwin") == 0 {
		return true
	}
	return false
}

func IsLinux() bool {
	if strings.Compare(runtime.GOOS, "linux") == 0 {
		return true
	}
	return false
}

func IsWindows() bool {
	if strings.Compare(runtime.GOOS, "windows") == 0 {
		return true
	}
	return false
}

func IsFreebsd() bool {
	if strings.Compare(runtime.GOOS, "freebsd") == 0 {
		return true
	}
	return false
}

func IsOpenbsd() bool {
	if strings.Compare(runtime.GOOS, "openbsd") == 0 {
		return true
	}
	return false
}

func IsNetbsd() bool {
	if strings.Compare(runtime.GOOS, "netbsd") == 0 {
		return true
	}
	return false
}

func IsDragonfly() bool {
	if strings.Compare(runtime.GOOS, "dragonfly") == 0 {
		return true
	}
	return false
}

func IsAndroid() bool {
	if strings.Compare(runtime.GOOS, "android") == 0 {
		return true
	}
	return false
}

func kill(pid int) error {
	res := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(pid))
	res.Stderr = os.Stderr
	res.Stdout = os.Stdout
	return res.Run()
}

func Kill(process *os.Process) error {
	if IsWindows() {
		return kill(process.Pid)
	}
	return process.Kill()
}

func SysKill(proc *os.Process) {
	// 终止 Java 进程及其子进程
	if err := syscall.Kill(-proc.Pid, syscall.SIGKILL); err != nil {
		fmt.Println("Error killing Java process:", err)
		return
	}
}

func Killpgid(proc *os.Process) {
	// 终止 Java 进程及其子进程
	pgid, err := syscall.Getpgid(proc.Pid)
	if err != nil {
		fmt.Println("Error getting process group ID:", err)
		return
	}

	if err := syscall.Kill(-pgid, syscall.SIGKILL); err != nil {
		fmt.Println("Error killing Java process group:", err)
		return
	}
}
