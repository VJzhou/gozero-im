package orm

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

type customLog struct {
	LogLevel logger.LogLevel
}

// Error implements logger.Interface.
func (l *customLog) Error(ctx context.Context, format string, ops ...interface{}) {
	logx.WithContext(ctx).Errorf(format, ops...)
}

// Info implements logger.Interface.
func (l *customLog) Info(ctx context.Context, format string, ops ...interface{}) {
	logx.WithContext(ctx).Infof(format, ops...)
}

// Trace implements logger.Interface.
func (l *customLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).
		Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds()/1e6), rows, sql)

}

// Warn implements logger.Interface.
func (l *customLog) Warn(ctx context.Context, format string, ops ...interface{}) {
	// the logx package don't have warnf function
	logx.WithContext(ctx).Infof(format, ops...)
}

func (l *customLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}
