/*
Package api defines the API handlers for the web application.

It is responsible for receiving requests, validating input data,
and delegating the processing to the corresponding service layer.
*/

package api

import (
	"errors"
	"fmt"
	"goweb/global"
	"goweb/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Base provides a base structure for API handlers.
// It includes common functionalities like error handling, request data binding, and response formatting.
type Base struct {
	Ctx    *gin.Context       // The Gin context for the current request.
	Errors error              // Stores errors encountered during every request processing.
	Logger *zap.SugaredLogger // A logger for logging messages and errors.
}

// NewBase creates a new instance of the Base struct.
func NewBase() Base {
	return Base{
		Logger: global.Logger,
	}
}

// BuildRequestOption defines options for building a request object.
// It includes the Gin context, data transfer object (DTO), and a flag to indicate whether to bind parameters from the URL.
type BuildRequestOption struct {
	Ctx               *gin.Context // The Gin context for the current request.
	DTO               any          // The data transfer object (DTO) for the request.
	BindParamsFromUrl bool         // Indicates whether to bind parameters from the URL.
}

// AddError appends a new error to the existing errors.
func (me *Base) AddError(errNew error) {
	me.Errors = utils.AppendError(me.Errors, errNew)
}

// GetError retrieves the accumulated errors.
func (me *Base) GetError() error {
	return me.Errors
}

// BuildRequest builds a request object.
// It binds the DTO to the request data, handling validation errors.
func (me *Base) BuildRequest(opt BuildRequestOption) error {
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
			// Parse validation errors and add them to the error list.
			errRes = me.ParseValidDataErrors(errRes, opt.DTO)
			me.AddError(errRes)
			// Send an error response to the client.
			Fail(me.Ctx, ResponseJson{
				Msg: me.GetError().Error(),
			})
		}
	}
	return errRes
}

// ParseValidDataErrors parses validation errors and extracts user-friendly error messages.
func (me *Base) ParseValidDataErrors(_errs error, target any) (errRes error) {
	if errs, ok := _errs.(validator.ValidationErrors); ok {
		fields := reflect.TypeOf(target).Elem()
		for _, fieldErr := range errs {
			field, _ := fields.FieldByName(fieldErr.Field())
			// Retrieve the error message from the tag.
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

// Fail sends an error response to the client.
func (me *Base) Fail(resp ResponseJson) {
	Fail(me.Ctx, resp)
}

// OK sends a success response to the client.
func (me *Base) OK(resp ResponseJson) {
	OK(me.Ctx, resp)
}

// ServerFail sends a server error response to the client.
func (me *Base) ServerFail(resp ResponseJson) {
	ServerFail(me.Ctx, resp)
}
