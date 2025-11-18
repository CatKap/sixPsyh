package handlers

import (
    "encoding/json"
    "net/http"

    "database/sql"
    loger "github.com/CatKap/sixPsyh/loger"
)

type Handler struct {
    db   *sql.DB
    log  *loger.Loger
}

type Error struct {
	message string `json:"error"`
}

func NewHandler(db *sql.DB, log *loger.Loger) *Handler {
    return &Handler{db: db, log: log}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"app": "myapp", "status": "ok"})
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

//func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
//    switch r.Method {
//    case http.MethodGet:
//        users, err := h.repo.ListUsers(r.Context())
//        if err != nil {
//            h.log.Error("list users:", err)
//            http.Error(w, "internal error", http.StatusInternalServerError)
//            return
//        }
//        w.Header().Set("Content-Type", "application/json")
//        json.NewEncoder(w).Encode(users)
//    default:
//        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
//    }
//}
