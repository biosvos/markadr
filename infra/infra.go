package infra

import (
	"github.com/biosvos/markadr/flow/service"
	"github.com/biosvos/markadr/infra/internal/broker/mqtt"
	humbleFilesystem "github.com/biosvos/markadr/infra/internal/filesystem"
	"github.com/biosvos/markadr/infra/internal/repository/filesystem"
	"github.com/biosvos/markadr/infra/internal/watcher/file"
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
	repo := filesystem.NewRepository(extension, extension)
	server := web.NewWeb(8123, repo)
	watcher, err := file.NewWatcher(assetPath)
	panicIfErr(err)
	broker, err := mqtt.NewBroker()
	panicIfErr(err)
	svc := service.NewService(watcher, broker, repo)
	err = svc.Start()
	panicIfErr(err)
	err = server.Run()
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
