package gin

import (
	"html/template"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/weisd/com"
)

type Input struct {
	ctx *Context
}

func (c *Context) Input() *Input {
	return &Input{ctx: c}
}

func (input *Input) Query(name string) string {
	if input.ctx.Request.Form == nil {
		input.ctx.Request.ParseForm()
	}
	return input.ctx.Request.Form.Get(name)
}

func (input *Input) QueryTrim(name string) string {
	return strings.TrimSpace(input.Query(name))
}

// QueryStrings returns a list of results by given query name.
func (input *Input) QueryStrings(name string) []string {
	if input.ctx.Request.Form == nil {
		input.ctx.Request.ParseForm()
	}

	vals, ok := input.ctx.Request.Form[name]
	if !ok {
		return []string{}
	}
	return vals
}

// QueryEscape returns escapred query result.
func (input *Input) QueryEscape(name string) string {
	return template.HTMLEscapeString(input.Query(name))
}

// QueryInt returns query result in int type.
func (input *Input) QueryInt(name string) int {
	return com.StrTo(input.Query(name)).MustInt()
}

// QueryInt64 returns query result in int64 type.
func (input *Input) QueryInt64(name string) int64 {
	return com.StrTo(input.Query(name)).MustInt64()
}

// QueryFloat64 returns query result in float64 type.
func (input *Input) QueryFloat64(name string) float64 {
	v, _ := strconv.ParseFloat(input.Query(name), 64)
	return v
}

// Params returns value of given param name.
func (input *Input) Params(name string) string {
	return input.ctx.Params.ByName(name)
}

// // SetParams sets value of param with given name.
// func (input *Input) SetParams(name, val string) {
// 	if !strings.HasPrefix(name, ":") {
// 		name = ":" + name
// 	}
// 	ctx.params[name] = val
// }

// ParamsEscape returns escapred params result.
func (input *Input) ParamsEscape(name string) string {
	return template.HTMLEscapeString(input.Params(name))
}

// ParamsInt returns params result in int type.
// e.g. ctx.ParamsInt(":uid")
func (input *Input) ParamsInt(name string) int {
	return com.StrTo(input.Params(name)).MustInt()
}

// ParamsInt64 returns params result in int64 type.
// e.g. ctx.ParamsInt64(":uid")
func (input *Input) ParamsInt64(name string) int64 {
	return com.StrTo(input.Params(name)).MustInt64()
}

// ParamsFloat64 returns params result in int64 type.
// e.g. ctx.ParamsFloat64(":uid")
func (input *Input) ParamsFloat64(name string) float64 {
	v, _ := strconv.ParseFloat(input.Params(name), 64)
	return v
}

// GetFile returns information about user upload file by given form field name.
func (input *Input) GetFile(name string) (multipart.File, *multipart.FileHeader, error) {
	return input.ctx.Request.FormFile(name)
}
