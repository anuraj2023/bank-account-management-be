package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			c.Logger().Error(err)

			he, ok := err.(*echo.HTTPError)
			if ok {
				if he.Internal != nil {
					if herr, ok := he.Internal.(*echo.HTTPError); ok {
						he = herr
					}
				}
			} else {
				he = &echo.HTTPError{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
				}
			}

			// Send custom error response
			code := he.Code
			message := map[string]interface{}{
				"error": he.Message,
			}

			if c.Request().Method == http.MethodHead {
				return c.NoContent(code)
			}
			return c.JSON(code, message)
		}

		return nil
	}
}