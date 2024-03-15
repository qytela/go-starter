package helper

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"os"
)

type InitProjectRender struct {
	sharedPath  string
	gatewayPath string
	workDir     string
}

func NewInitProjectRender(sharedPath string, gatewayPath string) *InitProjectRender {
	workDir := GetWorkDir()

	return &InitProjectRender{
		sharedPath:  sharedPath,
		gatewayPath: gatewayPath,
		workDir:     workDir,
	}
}

func (t *InitProjectRender) Init() error {
	if err := t.EnvInit(); err != nil {
		return err
	}
	if err := t.MiddlewareInit(); err != nil {
		return err
	}
	if err := t.ResponsesInit(); err != nil {
		return err
	}
	if err := t.GatewayInit(); err != nil {
		return err
	}

	return nil
}

func (t *InitProjectRender) EnvInit() error {
	source, err := mustache.RenderFile("codeshared/templates/env.mustache")
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/.env", os.Getenv("PROJECT_DIR")),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitProjectRender) MiddlewareInit() error {
	source, err := mustache.RenderFile("codeshared/templates/api/middleware/auth.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/middleware/auth.go", t.sharedPath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitProjectRender) ResponsesInit() error {
	sourceErr, err := mustache.RenderFile("codeshared/templates/api/responses/error_response.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	sourceOk, err := mustache.RenderFile("codeshared/templates/api/responses/success_response.mustache")
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/responses/error_response.go", t.sharedPath),
		[]byte(sourceErr),
		0644,
	); err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/responses/success_response.go", t.sharedPath),
		[]byte(sourceOk),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitProjectRender) GatewayInit() error {
	source, err := mustache.RenderFile("codebase/templates/gateway_cmd.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/main.go", t.gatewayPath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}
