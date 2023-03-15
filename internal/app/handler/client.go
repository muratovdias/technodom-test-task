package handler

import (
	"net/http"
)

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	statusCode, err := h.service.Client.Redirect(link)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusBadRequest)
		return
	}
	w.WriteHeader(statusCode)
	w.Write([]byte("success"))
}
