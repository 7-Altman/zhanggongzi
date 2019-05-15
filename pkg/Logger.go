package pkg

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var xlog = logrus.New()
var cfg = GetCfg()

func SetLogConf(confPath string) {
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
		cfg.Parse(confPath)
		SetCfg(cfg)
	}

	var (
		maxAge       int
		rotationTime int
		path         string
	)
	for k, v := range cfg.Log {
		maxAge, _ = strconv.Atoi(v["maxAg"])
		rotationTime, _ = strconv.Atoi(v["rotationTime"])
		path = v["path"]
		pathExist := ChkPathCreateNotExist(path)
		if !pathExist {
			println("path error")
			continue
		}
		writerTmp, _ := rotatelogs.New(
			strings.TrimRight(path, "/")+"/mf-gateway-"+k+".%Y-%m-%d.log",                    //baseLogPaht+".%Y%m%d"
			rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),             // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*24*time.Hour), // 日志切割时间间隔
		)
		switch k {
		case "info":
		case "default":
			if v["switch"] == "true" {
				writerI = writerTmp
			} else {
				writerI = nil
			}
			break
		case "error":
			if v["switch"] == "true" {
				writerE = writerTmp
			} else {
				writerE = nil
			}
			break
		case "warn":
			if v["switch"] == "true" {
				writerW = writerTmp
			} else {
				writerW = nil
			}
			break
		case "debug":
			if v["switch"] == "true" {
				writerD = writerTmp
			} else {
				writerD = nil
			}
			break
		case "fatal":
			if v["switch"] == "true" {
				writerF = writerTmp
			} else {
				writerF = nil
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
	if (cfg.Log["info"]["switch"] != "" && strings.ToLower(cfg.Log["info"]["switch"]) != "true") ||
		(strings.ToLower(cfg.Log["default"]["switch"]) != "" && strings.ToLower(cfg.Log["default"]["switch"]) != "true") {
		fmt.Println("loginfoswitch:false")
		return
	}
	xlog.SetLevel(logrus.InfoLevel)
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.Info(format, args)
}
func InfoWithFields(msg string, fields map[string]interface{}) {
	xlog.WithFields(fields).Info(msg)
}

// Error send log error to logstash
func Error(format string, args ...interface{}) {
	if strings.ToLower(cfg.Log["error"]["switch"]) != "true" {
		return
	}
	xlog.SetLevel(logrus.ErrorLevel)
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.Error(format, args)
}

// Warn send log warn to logstash
func Warn(format string, args ...interface{}) {
	if strings.ToLower(cfg.Log["warn"]["switch"]) != "true" {
		return
	}
	xlog.SetLevel(logrus.WarnLevel)
	xlog.SetFormatter(&logrus.TextFormatter{})
	xlog.Warn(format, args)
}

// Debug send log debug to logstash
func Debug(format string, args ...interface{}) {
	if strings.ToLower(cfg.Log["debug"]["switch"]) != "true" {
		return
	}
	xlog.SetLevel(logrus.DebugLevel)
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.Debug(format, args)
}

// Fatal send log fatal to logstash
func Fatal(format string, args ...interface{}) {
	if strings.ToLower(cfg.Log["fatal"]["switch"]) != "true" {
		return
	}
	xlog.SetLevel(logrus.FatalLevel)
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.Fatal(format, args)
}

// panic send log fatal to logstash
func Panic(format string, args ...interface{}) {
	if strings.ToLower(cfg.Log["panic"]["switch"]) != "true" {
		return
	}
	xlog.SetLevel(logrus.PanicLevel)
	xlog.SetFormatter(&logrus.JSONFormatter{})
	xlog.Fatal(format, args)
}

//log path check
func ChkPathCreateNotExist(path string) (bool) {
	exist := true
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("path " + path + " dir error,maybe is not exist, maybe not")
		if os.IsNotExist(err) {
			fmt.Println(path + " dir is not exist")
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Println("mkdir "+path+" failed!", err)
				exist = false
			} else {
				fmt.Println("mkdir " + path)
			}
		}
	}
	return exist
}

