package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErorr(w, 400, "something went wrong")
}
