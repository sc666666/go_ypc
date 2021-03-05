package requests

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"time"
)

var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改 gin 框架中的 Validator 引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 在校验器注册自定义的校验方法
		if err := v.RegisterValidation("checkDate", customFunc); err != nil {
			return err
		}

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用（fallback）的语言环境，后面的参数是支持的语言环境
		uni := ut.New(enT, zhT, enT)

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}

		// 注意！因为这里会使用到trans实例
		// 所以这一步注册要放到trans初始化的后面
		if err := v.RegisterTranslation(
			"checkDate",
			trans,
			registerTranslator("checkDate", "{0}必须要晚于当前日期"),
			translate,
		); err != nil {
			return err
		}

		return
	}
	return
}

// ParseCustomErrors 解析自定义错误消息
func ParseCustomErrors(model interface{}, err error, messages map[string]map[string]string) []map[string]string {
	// 1、将 err 强制转换为 ValidationErrors 类型
	ve := err.(validator.ValidationErrors)
	// 2、验证失败的字段以及错误消息会存放进来
	InvalidFields := make([]map[string]string, 0)

	// 3、循环处理验证失败的数据
	for _, e := range ve {
		errors := map[string]string{}
		jT := jsonTag(model, e.Field())
		// 3.1、原始的错误消息
		errors[jT] = e.Translate(trans)
		// 3.2、先判断有没有自定义错误消息
		if f, exist := messages[e.Field()]; exist {
			// 3.3、再去获取自定义的错误消息，存在就覆盖原始得错误消息
			if m, exist := f[e.Tag()]; exist {
				errors[jT] = m
			}
		}
		// 3.4、追加到返回数据中
		InvalidFields = append(InvalidFields, errors)
	}

	return InvalidFields
}

// jsonTag 获取 json tag 名称
func jsonTag(model interface{}, name string) string {
	if field, ok := reflect.TypeOf(model).Elem().FieldByName(name); ok {
		return field.Tag.Get("json")
	}
	return ""
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

// customFunc 自定义字段级别校验方法
func customFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}
