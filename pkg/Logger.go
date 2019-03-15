package pkg

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

var xlog = logrus.New()

//Init is Initial log config .
func init() {
	// load config
	cfg := GetCfg()
	var (
		writerD io.Writer
		writerI io.Writer
		writerW io.Writer
		writerE io.Writer
		writerF io.Writer
		writerP io.Writer
	)
	if cfg == nil {
		cfg = new(Config)
		cfg.Parse("./config/app.properties")
		SetCfg(cfg)
	}

	for k, v := range cfg.Log {
		switch k {
		case "default":
			var (
				maxAgeD       int
				rotationTimeD int
			)
			maxAgeD, _ = strconv.Atoi(v["maxAg"])
			rotationTimeD, _ = strconv.Atoi(v["rotationTime"])
			if v["switch"] == "true" {
				writerI, _ = rotatelogs.New(
					"./info-"+"%Y-%m-%d.log",
					//baseLogPaht+".%Y%m%d",
					rotatelogs.WithMaxAge(time.Duration(maxAgeD)*24*time.Hour),             // 文件最大保存时间
					rotatelogs.WithRotationTime(time.Duration(rotationTimeD)*24*time.Hour), // 日志切割时间间隔
				)

			} else {
				writerI = nil
			}
			break

		case "error":
			var (
				maxAgeE       int
				rotationTimeE int
			)
			maxAgeE, _ = strconv.Atoi(v["maxAg"])
			rotationTimeE, _ = strconv.Atoi(v["rotationTime"])
			if v["switch"] == "true" {
				writerE, _ = rotatelogs.New(
					"./error-"+"%Y-%m-%d.log",
					//baseLogPaht+".%Y%m%d",
					rotatelogs.WithMaxAge(time.Duration(maxAgeE)*24*time.Hour),             // 文件最大保存时间
					rotatelogs.WithRotationTime(time.Duration(rotationTimeE)*24*time.Hour), // 日志切割时间间隔
				)

			} else {
				writerE = nil
			}
			break

		}
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writerD, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writerI,
		logrus.WarnLevel:  writerW,
		logrus.ErrorLevel: writerE,
		logrus.FatalLevel: writerF,
		logrus.PanicLevel: writerP,
	}, xlog.Formatter)
	xlog.AddHook(lfHook)

}

// Info send log info to logstash
func Info(format string, args ...interface{}) {
	xlog.SetLevel(logrus.InfoLevel)
	xlog.Info(format, args)
}

// Error send log error to logstash
func Error(format string, args ...interface{}) {
	xlog.SetLevel(logrus.ErrorLevel)
	xlog.Error(format, args)
}

// Warn send log warn to logstash
func Warn(format string, args ...interface{}) {
	xlog.SetLevel(logrus.WarnLevel)
	xlog.Warn(format, args)
}

// Debug send log debug to logstash
func Debug(format string, args ...interface{}) {
	xlog.SetLevel(logrus.DebugLevel)
	xlog.Debug(format, args)
}

// Fatal send log fatal to logstash
func Fatal(format string, args ...interface{}) {
	xlog.SetLevel(logrus.FatalLevel)
	xlog.Fatal(format, args)
}
