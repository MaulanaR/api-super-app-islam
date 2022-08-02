package app

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

type LogInfo struct {
	Error      string            `json:"error,omitempty"`
	Version    string            `json:"version,omitempty"`
	Env        string            `json:"env,omitempty"`
	Method     string            `json:"method,omitempty"`
	Path       string            `json:"path,omitempty"`
	IP         []string          `json:"ip,omitempty"`
	Slug       string            `json:"slug,omitempty"`
	Email      string            `json:"email,omitempty"`
	ClientName string            `json:"client_name,omitempty"`
	Referer    string            `json:"referer,omitempty"`
	Trace      map[string]string `json:"trace,omitempty"`
}

func NewLogInfo() LogInfo {
	return LogInfo{
		Version: APP_VERSION,
		Env:     APP_ENV,
	}
}

func GetTrace(skip int) map[string]string {
	temp := map[int]string{}
	trace := map[string]string{}
	for i := 0; i <= 15; i++ {
		fun, fileName, lineNo, _ := runtime.Caller(i + skip)
		funcName := runtime.FuncForPC(fun).Name()
		if fileName != "" {
			wd, _ := os.Getwd()
			if wd != "" {
				projectFile := strings.Split(fileName, wd+"/")
				if len(projectFile) > 1 {
					fileName = projectFile[1]
				}
			}
			modFile := strings.Split(fileName, "/pkg/mod/")
			if len(modFile) > 1 {
				fileName = modFile[1]
			}
			projectFunc := strings.Split(funcName, "/")
			funcName = projectFunc[len(projectFunc)-1]
			projectFunc = strings.Split(funcName, ".")
			if len(projectFunc) > 2 {
				funcName = projectFunc[len(projectFunc)-2] + "." + projectFunc[len(projectFunc)-1]
			}
			temp[i] = fmt.Sprintf("%s🔹 %s:%d", funcName, fileName, lineNo)
		}
	}
	index := make([]int, 0)
	for i := range temp {
		index = append(index, i)
	}
	sort.Ints(index)
	for _, i := range index {
		trace[fmt.Sprintf("#%02d", i)] = temp[i]
	}
	return trace
}
