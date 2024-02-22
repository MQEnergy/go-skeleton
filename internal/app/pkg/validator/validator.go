package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	cnTrans "github.com/go-playground/validator/v10/translations/zh"
)

type XValidator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
}

type ErrorResponse struct {
	Error       bool
	Message     string
	FailedField string
	Tag         string
	Value       interface{}
}

// New ...
func New(local string) (*XValidator, error) {
	validate := validator.New()
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(zhT, zhT, enT) // 第一个参数是默认的语言环境,后面的参数是支持的语言环境

	translator, found := uni.GetTranslator(local)
	if !found {
		return nil, fmt.Errorf("uni.GetTranslator(%s) failed", local)
	}
	switch local {
	case "zh":
		cnTrans.RegisterDefaultTranslations(validate, translator) // 注册默认的英文翻译器
	default:
		enTrans.RegisterDefaultTranslations(validate, translator)
	}

	return &XValidator{
		validate: validate,
		uni:      uni,
		trans:    translator,
	}, nil
}

// Validate
// @Description: 验证
// @receiver v
// @param data
// @return []ErrorResponse
// @author cx
func (v *XValidator) Validate(data interface{}) []ErrorResponse {
	var validateErrors []ErrorResponse
	if errs := v.validate.Struct(data); errs != nil {
		for _, e := range errs.(validator.ValidationErrors) {
			elem := ErrorResponse{
				Error:       true,
				Message:     e.Translate(v.trans),
				FailedField: e.Field(),
				Tag:         e.Tag(),
				Value:       e.Value(),
			}
			validateErrors = append(validateErrors, elem)
		}
	}
	return validateErrors
}
