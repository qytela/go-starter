package helper

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"os"
)

type InitSharedPkgRender struct {
	sharedPath string
	workDir    string
}

func NewInitSharedPkgRender(sharedPath string) *InitSharedPkgRender {
	workDir := GetWorkDir()

	return &InitSharedPkgRender{
		sharedPath: sharedPath,
		workDir:    workDir,
	}
}

func (t *InitSharedPkgRender) Init() error {
	if err := t.AuthInit(); err != nil {
		return err
	}
	if err := t.ExceptionInit(); err != nil {
		return err
	}
	if err := t.LoggerInit(); err != nil {
		return err
	}
	if err := t.UtilsInit(); err != nil {
		return err
	}
	if err := t.ValidationInit(); err != nil {
		return err
	}

	return nil
}

func (t *InitSharedPkgRender) AuthInit() error {
	source, err := mustache.RenderFile("codeshared/templates/auth.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/auth/auth.go", t.sharedPath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitSharedPkgRender) ExceptionInit() error {
	source, err := mustache.RenderFile("codeshared/templates/error_response.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/exception/error_response.go", t.sharedPath),
		[]byte(source),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitSharedPkgRender) LoggerInit() error {
	sourceLog, err := mustache.RenderFile("codeshared/templates/logger.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/logger/logger.go", t.sharedPath),
		[]byte(sourceLog),
		0644,
	); err != nil {
		return err
	}

	sourceInLog, err := mustache.RenderFile("codeshared/templates/incoming_logger.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/logger/incoming_logger.go", t.sharedPath),
		[]byte(sourceInLog),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitSharedPkgRender) UtilsInit() error {
	sourceBcrypt, err := mustache.RenderFile("codeshared/templates/bcrypt.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/utils/bcrypt.go", t.sharedPath),
		[]byte(sourceBcrypt),
		0644,
	); err != nil {
		return err
	}

	sourceHelper, err := mustache.RenderFile("codeshared/templates/helper.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/utils/helper.go", t.sharedPath),
		[]byte(sourceHelper),
		0644,
	); err != nil {
		return err
	}

	sourceMailer, err := mustache.RenderFile("codeshared/templates/mailer.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/utils/mailer.go", t.sharedPath),
		[]byte(sourceMailer),
		0644,
	); err != nil {
		return err
	}

	sourcePaginate, err := mustache.RenderFile("codeshared/templates/paginate.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/utils/paginate.go", t.sharedPath),
		[]byte(sourcePaginate),
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (t *InitSharedPkgRender) ValidationInit() error {
	sourceLog, err := mustache.RenderFile("codeshared/templates/validation.mustache", map[string]interface{}{
		"workDir": t.workDir,
	})
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("%s/validation/validation.go", t.sharedPath),
		[]byte(sourceLog),
		0644,
	); err != nil {
		return err
	}

	return nil
}
