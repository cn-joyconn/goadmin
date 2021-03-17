package middleware

import (
	"context"
	"fmt"
	"time"

	golog "github.com/cn-joyconn/gologs"
	zap "go.uber.org/zap"
	glogger "gorm.io/gorm/logger"

	"gorm.io/gorm/utils"
)

var ormlooger = golog.GetLogger("orm").With(zap.String("origin", "orm"))




type GormLogger struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      glogger.LogLevel
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (g *GormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	newLogger := *g
	newLogger.LogLevel = level
	return &newLogger
}

func (g *GormLogger) Info(ctx context.Context, message string, data ...interface{}) {
	if g.LogLevel >= glogger.Info {
		g.Printf(g.infoStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GormLogger) Warn(ctx context.Context, message string, data ...interface{}) {
	if g.LogLevel >= glogger.Warn {
		g.Printf(g.warnStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GormLogger) Error(ctx context.Context, message string, data ...interface{}) {
	if g.LogLevel >= glogger.Error {
		g.Printf(g.errStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if g.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && g.LogLevel >= glogger.Error:
			sql, rows := fc()
			if rows == -1 {
				g.Printf(g.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(g.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= glogger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", g.SlowThreshold)
			if rows == -1 {
				g.Printf(g.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(g.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case g.LogLevel >= glogger.Info:
			sql, rows := fc()
			if rows == -1 {
				g.Printf(g.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				g.Printf(g.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func (g *GormLogger) Printf(message string, data ...interface{}) {
	 ormlooger.Info(message, zap.String("detail", fmt.Sprintf("%v",append([]interface{}{utils.FileWithLineNum()}, data...)...)))
	// ormlooger.Info(
	// 	"gorm",
	// 	zap.String("type", "sql"),
	// 	zap.Any("src", data[0]),
	// 	zap.Any("duration", data[1]),
	// 	zap.Any("rows", data[2]),
	// 	zap.Any("sql", data[3]),
	// )
}