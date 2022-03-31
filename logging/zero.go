package logging

import "github.com/rs/zerolog"

//goland:noinspection GoUnusedExportedFunction
func SetLoggerZeroLog(l zerolog.Logger) {
	SetLogger(zeroAdapter{internal: l})
}

type zeroAdapter struct {
	internal zerolog.Logger
}

func (za zeroAdapter) Fatal(args ...any) {
	za.internal.Fatal().Fields(args).Send()
}
func (za zeroAdapter) Panic(args ...any) {
	za.internal.Panic().Fields(args).Send()
}
func (za zeroAdapter) Error(args ...any) {
	za.internal.Error().Fields(args).Send()
}
func (za zeroAdapter) Warn(args ...any) {
	za.internal.Warn().Fields(args).Send()
}
func (za zeroAdapter) Info(args ...any) {
	za.internal.Info().Fields(args).Send()
}
func (za zeroAdapter) Debug(args ...any) {
	za.internal.Debug().Fields(args).Send()
}
