package middleware

import (
	"net/http"
)

// Basic middleware type definition
type Middleware func(http.Handler) http.Handler
