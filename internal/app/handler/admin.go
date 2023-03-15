package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/technodom-test-task/internal/app/models"
	"github.com/muratovdias/technodom-test-task/internal/app/service"
)

func (h *Handler) allLinks(w http.ResponseWriter, r *http.Request) {
	queryPage := r.URL.Query().Get("page")
	if len(queryPage) == 0 {
		queryPage = "1"
	}
	// fmt.Println(queryPage)
	page, err := strconv.Atoi(queryPage)
	if err != nil || page < 1 || strings.TrimSpace(queryPage) == "" {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	links, err := h.service.GetLinks(page)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if len(*links) == 0 {
		w.Write([]byte("Nothing Found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(links); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getLink(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusBadRequest)
		return
	}
	link, err := h.service.GetLinkByID(id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			w.Write([]byte(err.Error()))
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(link); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateLink(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusBadRequest)
		return
	}
	var update models.Link
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateLink(id, update.ActiveLink); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Active link updated successfully"))
}

func (h *Handler) deleteLink(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(paramID)
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteLink(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Deleted successfully"))

}

func (h *Handler) createLink(w http.ResponseWriter, r *http.Request) {
	var link models.Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateLink(link.ActiveLink); err != nil {
		if errors.Is(err, service.ErrAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Created successfully"))
}
