package os

import (
	"fmt"
	"github.com/xxl6097/go-service-framework/pkg/version"
	"io"
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
	info += "<tr>"
	info += "<td>系统架构</td>"
	info += "<td>" + runtime.GOARCH + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>cup数量</td>"
	info += "<td>" + strconv.Itoa(runtime.NumCPU()) + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>GOMIPS</td>"
	info += "<td>" + os.Getenv("GOMIPS") + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>Goroutine数量</td>"
	info += "<td>" + strconv.Itoa(runtime.NumGoroutine()) + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>内存分配</td>"
	info += "<td>" + fmt.Sprintf("%d", m.Alloc/1024/1024) + "G</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>总内存分配</td>"
	info += "<td>" + fmt.Sprintf("%d", m.TotalAlloc/1024/1024) + "G</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>GC次数</td>"
	info += "<td>" + fmt.Sprintf("%d", m.NumGC) + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>Go Version</td>"
	info += "<td>" + runtime.Version() + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>AppName</td>"
	info += "<td>" + version.AppName + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>DisplayName</td>"
	info += "<td>" + version.DisplayName + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>Description</td>"
	info += "<td>" + version.Description + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>AppVersion</td>"
	info += "<td>" + version.AppVersion + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>BuildVersion</td>"
	info += "<td>" + version.BuildVersion + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>BuildTime</td>"
	info += "<td>" + version.BuildTime + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>GitRevision</td>"
	info += "<td>" + version.GitRevision + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>GitBranch</td>"
	info += "<td>" + version.GitBranch + "</td>"
	info += "</tr>"
	info += "<tr>"
	info += "<td>GitRevision</td>"
	info += "<td>" + version.GitRevision + "</td>"
	info += "</tr>"
	info += "</table>"
	return info
}

var debug bool = false

func IsDebug() bool {
	return debug
}
func SetDebug(b bool) {
	debug = b
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

//func SysKill(proc *os.Process) {
//	// 终止 Java 进程及其子进程
//	if err := syscall.Kill(-proc.Pid, syscall.SIGKILL); err != nil {
//		fmt.Println("Error killing Java process:", err)
//		return
//	}
//}
//
//func Killpgid(proc *os.Process) {
//	// 终止 Java 进程及其子进程
//	pgid, err := syscall.Getpgid(proc.Pid)
//	if err != nil {
//		fmt.Println("Error getting process group ID:", err)
//		return
//	}
//
//	if err := syscall.Kill(-pgid, syscall.SIGKILL); err != nil {
//		fmt.Println("Error killing Java process group:", err)
//		return
//	}
//}

func IsExecutable(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	file.Stat()
	defer file.Close()

	// 读取文件头信息
	header := make([]byte, 4)
	if _, err := io.ReadFull(file, header); err != nil {
		return false, err
	}

	// 检查 ELF 文件头（Linux 和 macOS）
	if header[0] == 0x7f && header[1] == 'E' && header[2] == 'L' && header[3] == 'F' {
		return true, nil
	}

	// 检查 PE 文件头（Windows）
	if header[0] == 'M' && header[1] == 'Z' {
		return true, nil
	}

	// 使用 os.Stat 获取文件信息
	fileInfo, err := file.Stat() //os.Stat(filePath)
	if err != nil {
		return false, err
	}
	// 检查用户执行权限
	if fileInfo.Mode().Perm()&0111 == 0111 {
		return true, nil
	}

	// 如果当前用户是文件的所有者，检查所有者执行权限
	if fileInfo.Sys().(*syscall.Stat_t).Uid == uint32(os.Getuid()) {
		return (fileInfo.Mode().Perm() & 0100) == 0100, nil
	}

	// 检查组执行权限
	if fileInfo.Sys().(*syscall.Stat_t).Gid == uint32(os.Getgid()) {
		return (fileInfo.Mode().Perm() & 0010) == 0010, nil
	}

	// 检查其他用户的执行权限
	return (fileInfo.Mode().Perm() & 0001) == 0001, nil

	return false, nil
}
