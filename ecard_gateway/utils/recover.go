package utils

import (
    "fmt"
    "ecard_gateway/log"
    "os"
    "runtime"
)

func RecoverAndLog() {
    defer func() {
        if err := recover(); err != nil {
            log.Warn("panic:%s", err)
            buf := make([]byte, 2048)
            n := runtime.Stack(buf, false)
            stackInfo := fmt.Sprintf("%s", buf[:n])
            log.Warn("panic stack info %s", stackInfo)
            // 项目初期可直接让程序退出，方便定位错误，以修正代码
            os.Exit(1)
        }
    }()
}
