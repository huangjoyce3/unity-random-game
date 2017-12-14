package handlers

import (
	"net/http"
)

/* TODO: implement a CORS middleware handler, as described
in https://drstearns.github.io/tutorials/cors/ that responds
with the following headers to all requests:

  Access-Control-Allow-Origin: *
  Access-Control-Allow-Methods: GET, PUT, POST, PATCH, DELETE
  Access-Control-Allow-Headers: Content-Type, Authorization
  Access-Control-Expose-Headers: Authorization
  Access-Control-Max-Age: 600
*/

type CORSHandler struct {
	Handler http.Handler
}

func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Max-Age", "600")

	// preflight
	if r.Method != "OPTIONS" {
		ch.Handler.ServeHTTP(w, r)
	}
}

func NewCORSHandler(handlerToWrap http.Handler) *CORSHandler {
	return &CORSHandler{handlerToWrap}
}
