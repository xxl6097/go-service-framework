package version

import "fmt"

var (
	AppName      string // 应用名称
	DisplayName  string
	Description  string
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

// Version 版本信息
func Version() string {
	version := fmt.Sprintf("App Name:\t%s\n", AppName)
	version += fmt.Sprintf("App DisplayName:\t%s\n", DisplayName)
	version += fmt.Sprintf("App Description:\t%s\n", Description)
	version += fmt.Sprintf("App Version:\t%s\n", AppVersion)
	version += fmt.Sprintf("Build version:\t%s\n", BuildVersion)
	version += fmt.Sprintf("Build time:\t%s\n", BuildTime)
	version += fmt.Sprintf("Git revision:\t%s\n", GitRevision)
	version += fmt.Sprintf("Git branch:\t%s\n", GitBranch)
	version += fmt.Sprintf("Golang Version: %s\n", GoVersion)
	fmt.Println(version)
	return version
}

func VersionJson() map[string]interface{} {
	return map[string]interface{}{
		"appName":      AppName,
		"description":  Description,
		"appVersion":   AppVersion,
		"buildVersion": BuildVersion,
		"buildTime":    BuildTime,
		"gitRevision":  GitRevision,
		"gitBranch":    GitBranch,
		"goVersion":    GoVersion,
	}
}
