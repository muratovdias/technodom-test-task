package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	fmt.Println(link)

}
