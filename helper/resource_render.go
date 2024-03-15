package helper

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"os"
	"regexp"
)

type ResourceRender struct {
	servicePath       string
	resourceName      string
	snakeResourceName string
	workDir           string
}

func NewResourceRender(servicePath string, resourceName string) *ResourceRender {
	workDir := GetWorkDir()

	return &ResourceRender{
		servicePath:       servicePath,
		resourceName:      resourceName,
		snakeResourceName: regexp.MustCompile("-").ReplaceAllString(resourceName, "_"),
		workDir:           workDir,
	}
}

func (t *ResourceRender) AddResource() error {
	source, err := mustache.RenderFile("codebase/templates/resource.mustache", map[string]interface{}{
		"name": CapitalizeWords(t.resourceName),
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/resources/%s_resource.go", t.servicePath, t.snakeResourceName),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ResourceRender) AddCollection() error {
	source, err := mustache.RenderFile("codebase/templates/resource_collection.mustache", map[string]interface{}{
		"name": CapitalizeWords(t.resourceName),
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/resources/%s_collection.go", t.servicePath, t.snakeResourceName),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}
