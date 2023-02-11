package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

const (
	defaultSlowThreshold = 1
)

var (
	gormToZerologLevel = map[logger.LogLevel]zerolog.Level{
		logger.Silent: zerolog.Disabled,
		logger.Error:  zerolog.ErrorLevel,
		logger.Warn:   zerolog.WarnLevel,
		logger.Info:   zerolog.InfoLevel,
	}
)

type Option func(*ZerologGormOption)

type ZerologGormOption struct {
	level         zerolog.Level
	slowThreshold int
}

type ZerologGorm struct {
	log    *zerolog.Logger
	option ZerologGormOption
}

func SetSlowlogThreshold(th int) Option {
	return func(o *ZerologGormOption) {
		o.slowThreshold = th
	}
}

func NewZerologToGormLogger(ctx context.Context, options ...Option) ZerologGorm {
	op := ZerologGormOption{
		level:         zerolog.GlobalLevel(),
		slowThreshold: defaultSlowThreshold,
	}

	for _, option := range options {
		option(&op)
	}

	return ZerologGorm{log: log.Ctx(ctx), option: op}
}

func (z ZerologGorm) LogMode(level logger.LogLevel) logger.Interface {
	z.option.level = gormToZerologLevel[level]
	return z
}

func (z ZerologGorm) Info(ctx context.Context, msg string, data ...interface{}) {
	logger := z.log.Info()

	for _, d := range data {
		logger = logger.Str(utils.FileWithLineNum(), d.(string))
	}
	logger.Msg(msg)
}

func (z ZerologGorm) Warn(ctx context.Context, msg string, data ...interface{}) {
	logger := z.log.Warn()

	for _, d := range data {
		logger = logger.Str(utils.FileWithLineNum(), d.(string))
	}
	logger.Msg(msg)

}
func (z ZerologGorm) Error(ctx context.Context, msg string, data ...interface{}) {
	logger := z.log.Error()

	for _, d := range data {
		logger = logger.Str(utils.FileWithLineNum(), d.(string))
	}
	logger.Msg(msg)

}
func (z ZerologGorm) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	lgr := z.log.Info()

	switch {
	case err != nil && z.option.level >= zerolog.ErrorLevel:
		sql, rows := fc()
		if rows == -1 {
			lgr = lgr.
				Err(err).
				Str("file", utils.FileWithLineNum()).
				Str("sql", sql).
				Float64("duration", float64(elapsed.Nanoseconds())/1e6)
		} else {
			lgr = lgr.
				Err(err).
				Str("file", utils.FileWithLineNum()).
				Int64("rows", rows).
				Str("sql", sql).
				Float64("duration", float64(elapsed.Nanoseconds())/1e6)
		}
	case z.option.level >= zerolog.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", z.option.slowThreshold)
		if rows == -1 {
			lgr = lgr.Err(err).
				Str("slow log", slowLog).
				Str("file", utils.FileWithLineNum()).
				Str("sql", sql).
				Float64("duration", float64(elapsed.Nanoseconds())/1e6)
		} else {
			lgr = lgr.Err(err).
				Str("slow log", slowLog).
				Str("file", utils.FileWithLineNum()).
				Str("sql", sql).
				Int64("rows", rows).
				Float64("duration", float64(elapsed.Nanoseconds())/1e6)
		}
	case z.option.level <= zerolog.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			lgr = lgr.Err(err).
				Str(utils.FileWithLineNum(), sql).
				Str("file", utils.FileWithLineNum()).
				Str("sql", sql).
				Float64("duration", float64(elapsed.Nanoseconds())/1e6)
		} else {
			lgr = lgr.Err(err).
				Str(utils.FileWithLineNum(), sql).
				Str("file", utils.FileWithLineNum()).
				Int64("rows", rows).
				Str("sql", sql).
				Float64("duration", float64(elapsed.Nanoseconds())/1e6)
		}
	}

	lgr.Send()
}
