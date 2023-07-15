package infra

import (
	humbleFilesystem "github.com/biosvos/markadr/infra/internal/filesystem"
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
	workspace := humbleFilesystem.NewWorkspace(humbleFilesystem.NewFilesystem(), assetPath)
	extension := humbleFilesystem.NewExtension(workspace, workspace, "json")
	server := web.NewWeb(8123, filesystem.NewRepository(extension, extension))
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
