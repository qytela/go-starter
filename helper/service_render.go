package helper

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"os"
	"regexp"
)

type ServiceRender struct {
	serviceName string
	servicePort string
	moduleName  string
	servicePath string
	workDir     string
}

func NewServiceRender(serviceName string, servicePort string, moduleName string, servicePath string) *ServiceRender {
	workDir := GetWorkDir()

	return &ServiceRender{
		serviceName: serviceName,
		servicePort: servicePort,
		moduleName:  moduleName,
		servicePath: servicePath,
		workDir:     workDir,
	}
}

func (t *ServiceRender) Render() error {
	if err := t.RoutesRender(); err != nil {
		return err
	}
	if err := t.ServerRender(); err != nil {
		return err
	}
	if err := t.ServerHandlerRender(); err != nil {
		return err
	}
	if err := t.ConfigRender(); err != nil {
		return err
	}
	if err := t.ServiceCmdRender(); err != nil {
		return err
	}

	return nil
}

func (t *ServiceRender) RoutesRender() error {
	source, err := mustache.RenderFile("codebase/templates/routes.mustache", map[string]interface{}{
		"name":        CapitalizeWords(t.serviceName),
		"service":     t.serviceName,
		"moduleName":  t.moduleName,
		"module":      CapitalizeWords(t.moduleName),
		"handlerName": CapitalizeWords(t.moduleName),
		"workDir":     t.workDir,
		"prefix":      regexp.MustCompile("-").ReplaceAllString(t.moduleName, "_"),
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/api/routes/routes.go", t.servicePath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ServiceRender) ServerRender() error {
	source, err := mustache.RenderFile("codebase/templates/server.mustache", map[string]interface{}{
		"name":        CapitalizeWords(t.serviceName),
		"service":     t.serviceName,
		"module":      CapitalizeWords(t.moduleName),
		"handlerName": CapitalizeWords(t.moduleName),
		"workDir":     t.workDir,
		"prefix":      regexp.MustCompile("-").ReplaceAllString(t.serviceName, "_"),
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/internal/server.go", t.servicePath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ServiceRender) ServerHandlerRender() error {
	source, err := mustache.RenderFile("codebase/templates/server_handler.mustache", map[string]interface{}{
		"name":       CapitalizeWords(t.moduleName),
		"service":    t.serviceName,
		"moduleName": t.moduleName,
		"workDir":    t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/internal/handler.go", t.servicePath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ServiceRender) ConfigRender() error {
	sourceDB, err := mustache.RenderFile("codebase/templates/pgsql_db.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/config/db.go", t.servicePath),
		[]byte(sourceDB),
		0644,
	); err != nil {
		return err
	}

	sourceRedisDB, err := mustache.RenderFile("codebase/templates/redis_db.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/config/redis.go", t.servicePath),
		[]byte(sourceRedisDB),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *ServiceRender) ServiceCmdRender() error {
	source, err := mustache.RenderFile("codebase/templates/service_cmd.mustache", map[string]interface{}{
		"service": t.serviceName,
		"workDir": t.workDir,
		"port":    t.servicePort,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/cmd/main.go", t.servicePath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}
