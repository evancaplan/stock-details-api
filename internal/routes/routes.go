package routes

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/stock-details-api/internal/endpoints/details"
	"github.com/stock-details-api/internal/utils"
)

func Routes() *chi.Mux {

	log := utils.GetLogger()
	log.Info("in routes/Routes(), initializing new chi router")
	router := chi.NewRouter()

	log.Info("Setting cors values...")
	_cors := cors.New(cors.Options{
		AllowOriginFunc:  allowedOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	log.Info("Setting middlewares in router")
	router.Use( // Define middlewares for the router to use
		render.SetContentType(render.ContentTypeJSON), // Set content-type headers as application/json
		middleware.Logger,          // Logging
		middleware.DefaultCompress, // Compress results, mostly assets + JSON
		middleware.Recoverer,       // Recover from panics without crashing server
		_cors.Handler,              //Cross Origin Resource Sharing
	)
	router.Use(_cors.Handler)
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/details", details.Routes())
	})

	return router

}

func allowedOriginFunc(r *http.Request, origin string) bool {
	originsWhitelist := []string{"bwssolutions.com", "localhost"}
	for _, whitelistedOrigin := range originsWhitelist {
		if strings.Contains(origin, whitelistedOrigin) {
			return true
		}
	}
	return false
}
