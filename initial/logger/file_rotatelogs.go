package logger

import (
	"backend/global"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	logFile := path.Join(global.AppConfig.Zap.Path, global.AppConfig.App.Name+"-%Y-%m-%d"+".log")
	if global.AppConfig.Zap.FileShunt {
		logFile = path.Join(global.AppConfig.Zap.Path, global.AppConfig.App.Name+"-%Y-%m-%d-"+level+".log")
	}

	fileWriter, err := rotatelogs.New(
		logFile,
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.AppConfig.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.AppConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
