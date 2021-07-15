package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

var UndefinedError = errors.New("UndefinedError: %s")


type RecoverStats struct {
	errorStatusCodeList map[error]int
}

type Response struct {
	Message string `json:"message"`
}

func NewRecover(errorStatusCodeList map[error]int) *RecoverStats {
	return &RecoverStats{
		errorStatusCodeList,
	}

}

func (s *RecoverStats) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				fmt.Println(err)
				if !ok {
					err = fmt.Errorf("%v", r)

				} else {
					jsonWithAbort(c, s.getStatusCode(err), err.Error())
				}

			}
		}()
		return next(c)
	}
}

func (s *RecoverStats) getStatusCode(e error) (statusCode int) {
	if val, ok := s.errorStatusCodeList[e]; ok {
		return val
	}
	return 500
}

func jsonWithAbort(c echo.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{message})
}

func AddRecover(e *echo.Echo, errorStatusCodeList map[error]int) {
	s := NewRecover(errorStatusCodeList)
	e.Use(s.Handle)
}
