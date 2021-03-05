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
	SlowThreshold                       time.Duration
	LogLevel                            glogger.LogLevel
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *GormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glogger.Info {
		ormlooger.Info(msg, zap.String("detail", fmt.Sprintf("%v",append([]interface{}{utils.FileWithLineNum()}, data...)...)))
	}
}

// Warn print warn messages
func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glogger.Warn {
		ormlooger.Warn(msg, zap.String("detail", fmt.Sprintf("%v",append([]interface{}{utils.FileWithLineNum()}, data...)...)))
	}
}

// Error print error messages
func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glogger.Error {
		ormlooger.Error(msg, zap.String("detail",fmt.Sprintf("%v",append([]interface{}{utils.FileWithLineNum()}, data...))))
	}
}

// Trace print sql message
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > glogger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= glogger.Error:
			sql, rows := fc()
			if rows == -1 {

				ormlooger.Info(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				ormlooger.Info(fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= glogger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				ormlooger.Info(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				ormlooger.Info(fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		default:
			sql, rows := fc()
			if rows == -1 {
				ormlooger.Info(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
			} else {
				ormlooger.Info(fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
			}
		}
	}
}
