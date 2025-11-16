package router

import (
    "net/http"
)

// RouteHandler — функция для обработки маршрутов
type RouteHandler func(w http.ResponseWriter, r *http.Request)

// Router хранит маршруты по методу и пути
type Router struct {
    routes map[string]map[string]RouteHandler // method -> path -> handler
}

// NewRouter создаёт новый роутер
func New() *Router {
    return &Router{
        routes: make(map[string]map[string]RouteHandler),
    }
}

// Handle регистрирует обработчик для метода и пути
func (rt *Router) Handle(method string, path string, handler RouteHandler) {
    method = methodUpper(method)
    if rt.routes[method] == nil {
        rt.routes[method] = make(map[string]RouteHandler)
    }
    rt.routes[method][path] = handler
}

// Алиасы для удобства
func (rt *Router) GET(path string, handler RouteHandler) {
    rt.Handle("GET", path, handler)
}

func (rt *Router) POST(path string, handler RouteHandler) {
    rt.Handle("POST", path, handler)
}

func (rt *Router) PUT(path string, handler RouteHandler) {
    rt.Handle("PUT", path, handler)
}

func (rt *Router) DELETE(path string, handler RouteHandler) {
    rt.Handle("DELETE", path, handler)
}

// ServeHTTP реализует http.Handler
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    method := methodUpper(r.Method)
    path := r.URL.Path

    if methodRoutes, ok := rt.routes[method]; ok {
        if handler, ok := methodRoutes[path]; ok {
            handler(w, r)
            return
        }
    }

    http.NotFound(w, r)
}

func methodUpper(method string) string {
    return http.CanonicalHeaderKey(method)
}

