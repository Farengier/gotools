package logging

import (
	"fmt"
	"github.com/rs/zerolog"
)

//goland:noinspection GoUnusedExportedFunction
func SetLoggerZeroLog(l zerolog.Logger) {
	SetLogger(zeroAdapter{internal: l})
}

type zeroAdapter struct {
	internal zerolog.Logger
}

func (za zeroAdapter) Fatal(args ...any) {
	za.internal.Fatal().Msg(fmt.Sprint(args...))
}
func (za zeroAdapter) Panic(args ...any) {
	za.internal.Panic().Msg(fmt.Sprint(args...))
}
func (za zeroAdapter) Error(args ...any) {
	za.internal.Error().Msg(fmt.Sprint(args...))
}
func (za zeroAdapter) Warn(args ...any) {
	za.internal.Warn().Msg(fmt.Sprint(args...))
}
func (za zeroAdapter) Info(args ...any) {
	za.internal.Info().Msg(fmt.Sprint(args...))
}
func (za zeroAdapter) Debug(args ...any) {
	za.internal.Debug().Msg(fmt.Sprint(args...))
}
