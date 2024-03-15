package codebase

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"go-starter/helper"
	"os"
)

type AddModule struct {
	ctx *cli.Context
}

func NewAddModule(ctx *cli.Context) *AddModule {
	return &AddModule{ctx}
}

func (c *AddModule) AddModule() error {
	projectDir := os.Getenv("PROJECT_DIR")

	serviceName := c.ctx.String("service")
	if serviceName == "" {
		serviceName = c.ctx.String("name")
	}

	moduleName := c.ctx.String("module")
	servicePath := fmt.Sprintf("%s/services/%s", projectDir, serviceName)

	if ok := helper.IsPathExists(servicePath); !ok {
		return errors.New("service name doesn't exists")
	}

	if ok := helper.IsPathExists(fmt.Sprintf("%s/api/modules/%s", servicePath, moduleName)); ok {
		return errors.New("module name already exists")
	}

	makeDirs := []string{
		fmt.Sprintf("api/modules/%s/repository", moduleName),
		fmt.Sprintf("api/modules/%s/services", moduleName),
		fmt.Sprintf("api/modules/%s/handlers", moduleName),
	}
	for _, dir := range makeDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", servicePath, dir), 0777); err != nil {
			return err
		}
	}

	serviceRender := helper.NewModuleRender(serviceName, moduleName, servicePath)
	if err := serviceRender.Render(); err != nil {
		return err
	}

	return nil
}
