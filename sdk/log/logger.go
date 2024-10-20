package log

import "go.uber.org/zap"

var (
	sugar *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()
}
func Info(args ...interface{}) {

	sugar.Info(args)
}
func Debug(args ...interface{}) {
	sugar.Debug(args)
}
func Trace(args ...interface{}) {
	//sugar.Debug(args)
}
func Error(args ...interface{}) {
	sugar.Error(args)
}
func Warning(args ...interface{}) {
	sugar.Warn(args)
}
