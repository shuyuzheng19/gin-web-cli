package configs

import (
	"gin-web/helper"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig 日志配置
type LoggerConfig struct {
	Dev         bool   `yaml:"dev" json:"dev"`
	Encoding    string `yaml:"encoding" json:"encoding"`
	OutputPaths string `yaml:"outputPaths" json:"outputPaths"`
	ErrorPaths  string `yaml:"errorPaths" json:"errorPaths"`
	Level       string `yaml:"level" json:"level"`
	LoggerDir   string `yaml:"loggerDir" json:"loggerDir"`
}

var maps = map[string]zapcore.Level{
	"info":  zapcore.InfoLevel,
	"debug": zapcore.DebugLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
}

var LOGGER *zap.Logger

// InitLogger 函数用于加载日志配置
func InitLogger(loggerConfig LoggerConfig) {
	// 检查日志文件目录是否为空
	if loggerConfig.LoggerDir == "" {
		helper.PanicErrorAndMessage(nil, "日志文件目录不能为空")
	}

	// 创建一个默认的生产配置
	config := zap.NewProductionConfig()

	// 设置日志编码格式
	config.Encoding = loggerConfig.Encoding

	// 设置是否为开发模式
	config.Development = loggerConfig.Dev

	// 获取日志级别，若未指定则默认为 Info 级别
	var level zapcore.Level

	if l, found := maps[loggerConfig.Level]; found {
		level = l
	} else {
		level = zap.InfoLevel
	}

	// 设置时间格式
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 设置日志级别
	config.Level = zap.NewAtomicLevelAt(level)

	if loggerConfig.OutputPaths != "" {

		var outPaths = make([]string, 0)

		var paths = strings.Split(loggerConfig.OutputPaths, ",")

		for _, path := range paths {
			if path == "stdout" {
				outPaths = append(outPaths, path)
			} else {
				outPaths = append(outPaths, filepath.Join(loggerConfig.LoggerDir, path))
			}
		}

		config.OutputPaths = outPaths
	}

	// 设置错误输出路径
	if loggerConfig.ErrorPaths != "" {
		var errorPaths = make([]string, 0)

		var paths = strings.Split(loggerConfig.ErrorPaths, ",")

		for _, path := range paths {
			errorPaths = append(errorPaths, filepath.Join(loggerConfig.LoggerDir, path))
		}

		config.ErrorOutputPaths = errorPaths
	}

	// 构建日志对象
	logger, err := config.Build()
	// 若构建失败则触发异常
	helper.PanicErrorAndMessage(err, "加载日志配置失败")

	// 设置全局日志对象
	LOGGER = logger

	// 在函数结束时同步日志对象
	defer LOGGER.Sync()
}
