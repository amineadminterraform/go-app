package api

import (
	"github.com/amineadminterraform/go-app/utils"
	"github.com/go-playground/validator/v10"
)
	

var validTemplateType validator.Func = func(fl validator.FieldLevel) bool {
	if templatetype,ok := fl.Field().Interface().(string); ok {
		return utils.IsSupportedTemplateType(templatetype)
	}
	return  false
}

var validGithubURL validator.Func = func(fl validator.FieldLevel) bool {
	if githuburl,ok := fl.Field().Interface().(string); ok {
		return utils.IsGitHubURL(githuburl)
	}
	return  false
}

