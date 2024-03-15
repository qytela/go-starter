package codebase

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"go-starter/helper"
	"os"
	"regexp"
)

type AddResource struct {
	ctx *cli.Context
}

func NewAddResource(ctx *cli.Context) *AddResource {
	return &AddResource{ctx}
}

func (c *AddResource) Add() error {
	projectDir := os.Getenv("PROJECT_DIR")

	serviceName := c.ctx.String("service")
	resourceName := c.ctx.String("resource")
	servicePath := fmt.Sprintf("%s/services/%s", projectDir, serviceName)

	if ok := helper.IsPathExists(servicePath); !ok {
		return errors.New("service name doesn't exists")
	}

	isCollection := c.ctx.Bool("collection")
	snakeResourceName := regexp.MustCompile("-").ReplaceAllString(resourceName, "_")

	if isCollection {
		if ok := helper.IsFileExists(
			fmt.Sprintf("%s/api/resources/%s_collection.go", servicePath, snakeResourceName),
		); ok {
			return errors.New("collection already exists")
		}
	} else {
		if ok := helper.IsFileExists(
			fmt.Sprintf("%s/api/resources/%s_resource.go", servicePath, snakeResourceName),
		); ok {
			return errors.New("resource already exists")
		}
	}

	makeDirs := []string{
		"api/resources",
	}
	for _, dir := range makeDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", servicePath, dir), 0777); err != nil {
			return err
		}
	}

	serviceRender := helper.NewResourceRender(servicePath, resourceName)

	if isCollection {
		if err := serviceRender.AddCollection(); err != nil {
			return err
		}
	} else {
		if err := serviceRender.AddResource(); err != nil {
			return err
		}
	}

	return nil
}
