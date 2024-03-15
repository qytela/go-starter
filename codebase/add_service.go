package codebase

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"go-starter/helper"
	"os"
)

type AddService struct {
	ctx *cli.Context
}

func NewAddService(ctx *cli.Context) *AddService {
	return &AddService{ctx}
}

func (c *AddService) AddService() error {
	projectDir := os.Getenv("PROJECT_DIR")

	serviceName := c.ctx.String("name")
	servicePort := c.ctx.String("port")
	moduleName := c.ctx.String("module")
	servicePath := fmt.Sprintf("%s/services/%s", projectDir, serviceName)

	if ok := helper.IsPathExists(servicePath); ok {
		return errors.New("service name already exists")
	}

	if err := os.MkdirAll(servicePath, 0777); err != nil {
		return err
	}

	makeDirs := []string{
		"api/routes",
		"config",
		"internal",
		"cmd",
	}
	for _, dir := range makeDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", servicePath, dir), 0777); err != nil {
			return err
		}
	}

	serviceRender := helper.NewServiceRender(serviceName, servicePort, moduleName, servicePath)
	if err := serviceRender.Render(); err != nil {
		return err
	}

	return nil
}
