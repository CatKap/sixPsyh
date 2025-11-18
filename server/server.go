package server

import (
    "context"
    "database/sql"
    "net/http"
    "time"

    // "github.com/lib/pq"
    "github.com/CatKap/sixPsyh/config"
    "github.com/CatKap/sixPsyh/handlers"
    "github.com/CatKap/sixPsyh/loger"
    "github.com/CatKap/sixPsyh/router"

		_ "github.com/mattn/go-sqlite3"
)

type Server struct {
    httpServer *http.Server
    db         *sql.DB
    loger     *loger.Loger
		router 		*router.Router
}

func New(cfg *config.Config, log *loger.Loger) (*Server, error) {
    db, err := sql.Open("sqlite3", cfg.DBFile)
    if err != nil {
        return nil, err
    }

    //simple ping to ensure connectivity
    if err := db.Ping(); err != nil {
        return nil, err
    }

    h := handlers.NewHandler(db, log)
		r := router.New()
		r.GET("/health/", h.Health)	

		r.GET("/events/", h.GetEvents)

		r.POST("/events/new/", h.AddEvent)

		r.GET("/cathegorys/", h.Cathegorys)

		r.POST("/cathegorys/new/", h.NewCathegory)
		
		r.DELETE("/entity/", h.Delete)

    //mux := http.NewServeMux()
    //mux.HandleFunc("/health", h.Health)
    //mux.HandleFunc("/users", h.Users) // example route
    //mux.HandleFunc("/", h.Index)

    srv := &http.Server{
        Addr:         cfg.Address,
        Handler:      logingMiddleware(r, log),
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,

    }

    return &Server{
        httpServer: srv,
        db:         db,
        loger:     log,
				router: r,
    }, nil
}

func (s *Server) Start() error {
    return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
    // close http server
    if err := s.httpServer.Shutdown(ctx); err != nil {
        return err
    }
    // close db
    return s.db.Close()
}

func logingMiddleware(next *router.Router, log *loger.Loger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Infof("%s %s %s", r.Method, r.URL.Path, time.Since(start))
    })
}
