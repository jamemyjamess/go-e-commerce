package response

import "github.com/gofiber/fiber/v2"

type IResponse interface {
	// Print()
	Success(code int, data any) IResponse
	Error(code int, traceId string, msg string) IResponse
	Res() error
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"msg"`
}

type Response struct {
	Context    *fiber.Ctx
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	IsError    bool
}

func NewResponse(ctx *fiber.Ctx) IResponse {
	return &Response{
		Context: ctx,
	}
}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	return r
}

func (r *Response) Error(code int, traceId string, msg string) IResponse {
	r.StatusCode = code
	r.ErrorRes = &ErrorResponse{
		TraceId: traceId,
		Msg:     msg,
	}
	r.IsError = true
	return r
}

func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() interface{} {
		if r.IsError {
			return r.ErrorRes
		}
		return r.Data
	}())
}
