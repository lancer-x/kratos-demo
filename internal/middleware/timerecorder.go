package middleware

import (
    "fmt"
    bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
    "github.com/prometheus/common/log"
    "time"
)

type timeRecorder struct {
    Start time.Duration
    End time.Duration
}

func TimeRecorder() bm.HandlerFunc {
    return func (ctx *bm.Context)  {
        log.Info("start")
        log.Info("end")
        start := time.Now()
        defer func(start time.Time) {
            fmt.Println(time.Since(start))
        }(start)
        ctx.Next()
    }
}

