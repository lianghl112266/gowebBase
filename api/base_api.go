package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"goweb/global"
	"goweb/utils"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUrl bool
}

func (me *BaseApi) AddError(errNew error) {
	me.Errors = utils.AppendError(me.Errors, errNew)
}

func (me *BaseApi) GetError() error {
	return me.Errors
}

func (me *BaseApi) BuildRequest(opt BuildRequestOption) *BaseApi {
	var errRes error
	me.Ctx = opt.Ctx
	if opt.DTO != nil {
		if opt.BindParamsFromUrl {
			errRes = me.Ctx.ShouldBindUri(opt.DTO)
		} else {
			errRes = me.Ctx.ShouldBind(opt.DTO)
			fmt.Println(errRes)
		}
		if errRes != nil {
			errRes = me.ParseValidDataErrors(errRes, opt.DTO)
			me.AddError(errRes)
			Fail(me.Ctx, ResponseJson{
				Msg: me.GetError().Error(),
			})
		}
	}
	return me
}

func (me *BaseApi) ParseValidDataErrors(_errs error, target any) (errRes error) {

	if errs, ok := _errs.(validator.ValidationErrors); ok {

		//通过反射获得指针指向元素类型对象
		fields := reflect.TypeOf(target).Elem()

		for _, fieldErr := range errs {
			field, _ := fields.FieldByName(fieldErr.Field())
			errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
			errMessage := field.Tag.Get(errMessageTag)
			if errMessage == "" {
				errMessage = field.Tag.Get("message")
				if errMessage == "" {
					errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
				}
			}

			errRes = utils.AppendError(errRes, errors.New(errMessage))
		}
	}

	return
}

func (me *BaseApi) Fail(resp ResponseJson) {
	Fail(me.Ctx, resp)
}

func (me *BaseApi) OK(resp ResponseJson) {
	OK(me.Ctx, resp)
}

func (me *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(me.Ctx, resp)
}
