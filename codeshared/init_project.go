package codeshared

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go-starter/helper"
	"os"
)

type InitProject struct {
	ctx *cli.Context
}

func NewInitProject(ctx *cli.Context) *InitProject {
	return &InitProject{ctx}
}

func (c *InitProject) Init() error {
	projectDir := os.Getenv("PROJECT_DIR")
	sharedPath := fmt.Sprintf("%s/api", projectDir)
	gatewayPath := fmt.Sprintf("%s/services/gateway/cmd", projectDir)

	if err := os.MkdirAll(sharedPath, 0777); err != nil {
		return err
	}
	if err := os.MkdirAll(gatewayPath, 0777); err != nil {
		return err
	}

	makeDirs := []string{
		"middleware",
		"responses",
	}
	for _, dir := range makeDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", sharedPath, dir), 0777); err != nil {
			return err
		}
	}

	if err := helper.NewInitProjectRender(sharedPath, gatewayPath).Init(); err != nil {
		return err
	}

	return nil
}
