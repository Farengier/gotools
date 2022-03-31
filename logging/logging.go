package logging

var logger Log

type Log interface {
	Fatal(args ...any)
	Panic(args ...any)
	Error(args ...any)
	Warn(args ...any)
	Info(args ...any)
	Debug(args ...any)
}

func SetLogger(l Log) {
	logger = l
}

func Fatal(args ...any) {
	logger.Fatal(args...)
}
func Panic(args ...any) {
	logger.Panic(args...)
}
func Error(args ...any) {
	logger.Error(args...)
}
func Warn(args ...any) {
	logger.Warn(args...)
}
func Info(args ...any) {
	logger.Info(args...)
}
func Debug(args ...any) {
	logger.Debug(args...)
}
