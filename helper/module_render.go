package helper

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"os"
	"regexp"
)

type ModuleRender struct {
	serviceName       string
	moduleName        string
	moduleNameReplace string
	servicePath       string
	workDir           string
}

func NewModuleRender(serviceName string, moduleName string, servicePath string) *ModuleRender {
	workDir := GetWorkDir()

	moduleNameReplace := regexp.MustCompile("-").ReplaceAllString(moduleName, "_")

	return &ModuleRender{
		serviceName:       serviceName,
		moduleName:        moduleName,
		moduleNameReplace: moduleNameReplace,
		servicePath:       servicePath,
		workDir:           workDir,
	}
}

func (t *ModuleRender) Render() error {
	if err := t.HandlerRender(); err != nil {
		return err
	}
	if err := t.ServiceRender(); err != nil {
		return err
	}
	if err := t.RepositoryRender(); err != nil {
		return err
	}

	return nil
}

func (t *ModuleRender) HandlerRender() error {
	upServiceName := CapitalizeWords(t.moduleName)

	source, err := mustache.RenderFile("codebase/templates/handler.mustache", map[string]interface{}{
		"name": upServiceName,
	})
	if err != nil {
		return err
	}

	sourceImpl, err := mustache.RenderFile("codebase/templates/handler_impl.mustache", map[string]interface{}{
		"name":    upServiceName,
		"service": t.serviceName,
		"module":  t.moduleName,
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/modules/%s/handlers/%s.go", t.servicePath, t.moduleName, fmt.Sprintf("%s_handler", t.moduleNameReplace)),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/modules/%s/handlers/%s.go", t.servicePath, t.moduleName, fmt.Sprintf("%s_handler_impl", t.moduleNameReplace)),
		[]byte(sourceImpl),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ModuleRender) ServiceRender() error {
	upServiceName := CapitalizeWords(t.moduleName)

	source, err := mustache.RenderFile("codebase/templates/service.mustache", map[string]interface{}{
		"name": upServiceName,
	})
	if err != nil {
		return err
	}

	sourceImpl, err := mustache.RenderFile("codebase/templates/service_impl.mustache", map[string]interface{}{
		"name":    upServiceName,
		"service": t.serviceName,
		"module":  t.moduleName,
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/modules/%s/services/%s.go", t.servicePath, t.moduleName, fmt.Sprintf("%s_service", t.moduleNameReplace)),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/modules/%s/services/%s.go", t.servicePath, t.moduleName, fmt.Sprintf("%s_service_impl", t.moduleNameReplace)),
		[]byte(sourceImpl),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ModuleRender) RepositoryRender() error {
	upServiceName := CapitalizeWords(t.moduleName)

	source, err := mustache.RenderFile("codebase/templates/repository.mustache", map[string]interface{}{
		"name": upServiceName,
	})
	if err != nil {
		return err
	}

	sourceImpl, err := mustache.RenderFile("codebase/templates/repository_impl.mustache", map[string]interface{}{
		"name":    upServiceName,
		"service": t.serviceName,
		"module":  t.moduleName,
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/modules/%s/repository/%s.go", t.servicePath, t.moduleName, fmt.Sprintf("%s_repository", t.moduleNameReplace)),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/modules/%s/repository/%s.go", t.servicePath, t.moduleName, fmt.Sprintf("%s_repository_impl", t.moduleNameReplace)),
		[]byte(sourceImpl),
		0644,
	); err != nil {
		return err
	}

	return nil
}
