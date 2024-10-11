package routeutils

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

type resultWrapper struct {
	Error     bool        `json:"error"`
	ErrorSlug string      `json:"error_slug,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

const ErrorMessageContextKey = "error-message"
const ErrorStackContextKey = "error-stack"
const ErrorNameContextKey = "error-name"
const ErrorPathContextKey = "error-path"
const SuccessMessageResponse = "success"

/////////////////////////////////////////

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	ErrorType  string `json:"error_type,omitempty"` // Optional field for more detailed errors
}

// Custom HTTP Error Handler for Echo
// Custom HTTP Error Handler for Echo
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message string

	// Log the error with the original cause for debugging (prints the detailed error to the terminal)
	log.Printf("Error: %+v", err)

	// Default to 500 Internal Server Error for unknown errors
	code = http.StatusInternalServerError
	message = http.StatusText(http.StatusInternalServerError)

	// Check if the error is an HTTPError from Echo
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string) // Use Echo's message for HTTP errors
	}

	// If it's a validation or user error, return 400 Bad Request
	cause := errors.Cause(err)
	if cause != nil {
		message = cause.Error()

		// Custom logic for specific errors (e.g., validation errors)
		if code == http.StatusInternalServerError { // Only change the code if it's still 500
			code = http.StatusBadRequest // Change to 400 if it's a user error
		}
	}

	// Return the JSON response with the appropriate status code and message
	if !c.Response().Committed {
		c.JSON(code, ErrorResponse{
			StatusCode: code,
			Message:    message,
			ErrorType:  "validation_error", // Or other types based on error
		})
	}
}

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// Success Handlers

func ResponseAPIOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       data,
	})
}

func ResponseCreated(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, SuccessResponse{
		StatusCode: http.StatusCreated,
		Message:    "Resource created successfully",
		Data:       data,
	})
}

func ResponseNoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// Error Handlers

func ResponseBadRequest(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusBadRequest, ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    errMsg,
	})
}

func ResponseUnauthorized(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusUnauthorized, ErrorResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    errMsg,
	})
}

func ResponseForbidden(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusForbidden, ErrorResponse{
		StatusCode: http.StatusForbidden,
		Message:    errMsg,
	})
}

func ResponseNotFound(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{
		StatusCode: http.StatusNotFound,
		Message:    errMsg,
	})
}

func ResponseConflict(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusConflict, ErrorResponse{
		StatusCode: http.StatusConflict,
		Message:    errMsg,
	})
}

func ResponseInternalServerError(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusInternalServerError, ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    errMsg,
	})
}

func ResponseServiceUnavailable(c echo.Context, errMsg string) error {
	return c.JSON(http.StatusServiceUnavailable, ErrorResponse{
		StatusCode: http.StatusServiceUnavailable,
		Message:    errMsg,
	})
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////

// func ResponseAPIOK(c echo.Context, data interface{}) error {
// 	returnValue := resultWrapper{
// 		Error:   false,
// 		Message: SuccessMessageResponse,
// 		Data:    data,
// 	}

// 	return responseJSON(c, http.StatusOK, returnValue)
// }

// func responseJSON(c echo.Context, code int, i interface{}) (err error) {
// 	rw, ok := i.(resultWrapper)
// 	if ok && rw.Data == nil {
// 		rw.Data = struct{}{}
// 	}

// 	b, err := json.Marshal(rw)
// 	if err != nil {
// 		return
// 	}
// 	return responseJSONBlob(c, code, b)
// }

// func responseJSONBlob(c echo.Context, code int, b []byte) (err error) {
// 	return responseBlob(c, code, echo.MIMEApplicationJSONCharsetUTF8, b)
// }

// func responseBlob(c echo.Context, code int, contentType string, b []byte) (err error) {
// 	c.Response().Header().Set(echo.HeaderContentType, contentType)
// 	c.Response().WriteHeader(code)

// 	_, err = c.Response().Write(b)
// 	return
// }
