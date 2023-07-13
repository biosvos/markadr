package infra

import (
	"github.com/biosvos/markadr/infra/internal/repository/filesystem"
	"github.com/biosvos/markadr/infra/internal/web"
	"log"
	"os"
)

const (
	envAssetPath = "ASSET_PATH"
)

func Run() {
	assetPath := getEnv(envAssetPath)
	server := web.NewWeb(8123, filesystem.NewRepository(assetPath))
	err := server.Run()
	panicIfErr(err)
}

func getEnv(env string) string {
	ret, ok := os.LookupEnv(env)
	if !ok {
		log.Panicf("env %v not found", env)
	}
	return ret
}

func panicIfErr(err error) {
	if err != nil {
		log.Panicf("%+v", err)
	}
}
