package main

import (
	"fmt"
	"github.com/eden-framework/plugins"
	"path"
)

var Plugin GenerationPlugin

type GenerationPlugin struct {
}

func (g *GenerationPlugin) GenerateEntryPoint(opt plugins.Option, cwd string) string {
	globalPkgPath := path.Join(opt.PackageName, "internal/global")
	globalFilePath := path.Join(cwd, "internal/global")
	tpl := fmt.Sprintf(`,
		{{ .UseWithoutAlias "github.com/eden-framework/eden-framework/pkg/application" "" }}.WithConfig(&{{ .UseWithoutAlias "%s" "%s" }}.QueueConfig)`, globalPkgPath, globalFilePath)
	return tpl
}

func (g *GenerationPlugin) GenerateFilePoint(opt plugins.Option, cwd string) []*plugins.FileTemplate {
	file := plugins.NewFileTemplate("global", path.Join(cwd, "internal/global/queue.go"))
	file.WithBlock(`
var QueueConfig = struct {
	Producer *{{ .UseWithoutAlias "github.com/eden-framework/plugin-queue/queue" "" }}.Producer
	Consumer *{{ .UseWithoutAlias "github.com/eden-framework/plugin-queue/queue" "" }}.Consumer
}{
	Producer: &{{ .UseWithoutAlias "github.com/eden-framework/plugin-queue/queue" "" }}.Producer{
		Driver: {{ .UseWithoutAlias "github.com/eden-framework/plugin-queue/queue" "" }}.DRIVER__REDIS,
		Host:   "localhost",
		Port:   6379,
		Topic:  "default",
	},
	Consumer: &{{ .UseWithoutAlias "github.com/eden-framework/plugin-queue/queue" "" }}.Consumer{
		Driver:  {{ .UseWithoutAlias "github.com/eden-framework/plugin-queue/queue" "" }}.DRIVER__REDIS,
		Brokers: []string{"localhost:6379"},
		Topic:   "default",
	},
}
`)

	return []*plugins.FileTemplate{file}
}
