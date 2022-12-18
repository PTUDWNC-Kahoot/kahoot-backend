package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var logger zerolog.Logger //nolint:gochecknoglobals // allow here pls

func Initialize() {
	var l zerolog.Level

	switch strings.ToLower(viper.GetString("LOG_LEVEL")) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger = zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
}

func Debug(message interface{}, args ...interface{}) {
	msg("debug", message, args...)
}

func Info(message string, args ...interface{}) {
	log(message, args...)
}

func Warn(message string, args ...interface{}) {
	log(message, args...)
}

func Error(message interface{}, args ...interface{}) {
	if logger.GetLevel() == zerolog.DebugLevel {
		Debug(message, args...)
	}

	msg("error", message, args...)
}

func Fatal(message interface{}, args ...interface{}) {
	msg("fatal", message, args...)

	os.Exit(1)
}

func log(message string, args ...interface{}) {
	if len(args) == 0 {
		logger.Info().Msg(message)
	} else {
		logger.Info().Msgf(message, args...)
	}
}

func msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		log(msg.Error(), args...)
	case string:
		log(msg, args...)
	default:
		log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
