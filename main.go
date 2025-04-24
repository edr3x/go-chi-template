package main

import (
	"cmp"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/edr3x/chi-template/internal/server"
)

func init() {
	zconf := zap.NewProductionConfig()
	zconf.DisableStacktrace = true

	logger := zap.Must(zconf.Build())
	zap.ReplaceGlobals(logger)
}

func main() {
	serv := server.NewServer()

	port := cmp.Or(os.Getenv("PORT"), "8080")
	zap.L().Info("Starting server", zap.String("port", port))
	http.ListenAndServe("0.0.0.0:"+port, serv.NewHandler())
}
