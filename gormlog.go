package gormlog

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	gorm "gorm.io/gorm/logger"
)

var logLevelMap = map[gorm.LogLevel]zerolog.Level{
	gorm.Silent: zerolog.Disabled,
	gorm.Error:  zerolog.ErrorLevel,
	gorm.Warn:   zerolog.WarnLevel,
	gorm.Info:   zerolog.InfoLevel,
}

var _ gorm.Interface = (*logger)(nil)

type logger struct {
	log zerolog.Logger
}

func NewLogger(log zerolog.Logger) gorm.Interface {
	return &logger{log: log}
}

func (l *logger) LogMode(level gorm.LogLevel) gorm.Interface {
	if newlevel, ok := logLevelMap[level]; ok {
		newlogger := *l
		newlogger.log = l.log.Level(newlevel)
		return &newlogger
	}
	return l // return the same logger again, because we can't return an error
}

func (l *logger) Info(context context.Context, msg string, data ...interface{}) {
	l.log.Info().Msgf(msg, data...)
}

func (l *logger) Warn(context context.Context, msg string, data ...interface{}) {
	l.log.Warn().Msgf(msg, data...)
}

func (l *logger) Error(context context.Context, msg string, data ...interface{}) {
	l.log.Error().Msgf(msg, data...)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	l.log.Trace().
		TimeDiff("time", time.Now(), begin).
		Str("sql", sql).
		Int64("rowsAffected", rowsAffected).
		Err(err).
		Send()
}
