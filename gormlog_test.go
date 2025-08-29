package gormlog_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/petoem/gormlog"
	"github.com/rs/zerolog"
	gorm "gorm.io/gorm/logger"
)

func TestLogLevels(t *testing.T) {
	t.Run("info", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := gormlog.NewLogger(zerolog.New(out))
		log.Info(context.TODO(), "test %v", "one")
		if got, want := out.String(), fmt.Sprintln(`{"level":"info","message":"test one"}`); got != want {
			t.Errorf("invalid log output\ngot: %vwant: %v", got, want)
		}
	})

	t.Run("warn", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := gormlog.NewLogger(zerolog.New(out))
		log.Warn(context.TODO(), "test %v", "two")
		if got, want := out.String(), fmt.Sprintln(`{"level":"warn","message":"test two"}`); got != want {
			t.Errorf("invalid log output\ngot: %vwant: %v", got, want)
		}
	})

	t.Run("error", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := gormlog.NewLogger(zerolog.New(out))
		log.Error(context.TODO(), "test %v", "three")
		if got, want := out.String(), fmt.Sprintln(`{"level":"error","message":"test three"}`); got != want {
			t.Errorf("invalid log output\ngot: %vwant: %v", got, want)
		}
	})

	t.Run("trace", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := gormlog.NewLogger(zerolog.New(out))
		log.Trace(context.TODO(), time.Now(), func() (string, int64) {
			return "SELECT * FROM TEST;", 10
		}, errors.New("testing"))
		match, err := regexp.MatchString("(?i)\\{\"level\":\"trace\",\"time\":[0-9]*\\.[0-9]+,\"sql\":\"SELECT \\* FROM TEST;\",\"rowsAffected\":10,\"error\":\"testing\"\\}", out.String())
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("invalid log output\ngot: %v", out.String())
		}
	})
}

func TestLogModes(t *testing.T) {
	t.Run("silent", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := gormlog.NewLogger(zerolog.New(out))
		log = log.LogMode(gorm.Silent)
		log.Info(context.TODO(), "test %v", "silent")
		if got, want := out.String(), ""; got != want {
			t.Errorf("invalid log output\ngot: %vwant: %v", got, want)
		}
	})
}
