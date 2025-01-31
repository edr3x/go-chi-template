package interfaces

import (
	"net/http"
)

type Responders interface {
	JSON(w http.ResponseWriter, i any, statusCode ...int)
	Error(w http.ResponseWriter, r *http.Request, err error)
}
