package main

import (
	"cmp"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/edr3x/chi-template/internal/server"
)

func init() {
	logger := createProductionLogger()
	zap.ReplaceGlobals(logger)
}

func main() {
	serv := server.NewServer()

	port := cmp.Or(os.Getenv("PORT"), "8080")
	zap.L().Info("Starting server", zap.String("port", port))
	http.ListenAndServe("0.0.0.0:"+port, serv.NewHandler())
}

func createProductionLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	return zap.Must(config.Build())
}
