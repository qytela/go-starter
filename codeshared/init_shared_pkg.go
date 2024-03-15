package codeshared

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"go-starter/helper"
	"os"
)

type InitSharedPkg struct {
	ctx *cli.Context
}

func NewInitSharedPkg(ctx *cli.Context) *InitSharedPkg {
	return &InitSharedPkg{ctx}
}

func (c *InitSharedPkg) Init() error {
	projectDir := os.Getenv("PROJECT_DIR")
	sharedPath := fmt.Sprintf("%s/pkg", projectDir)

	if ok := helper.IsPathExists(sharedPath); ok {
		return errors.New("pkg path already exists")
	}

	if err := os.MkdirAll(sharedPath, 0777); err != nil {
		return err
	}

	makeDirs := []string{
		"auth",
		"exception",
		"logger",
		"utils",
		"validation",
	}
	for _, dir := range makeDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", sharedPath, dir), 0777); err != nil {
			return err
		}
	}

	initSharedPkg := helper.NewInitSharedPkgRender(sharedPath)
	if err := initSharedPkg.Init(); err != nil {
		return err
	}

	return nil
}
