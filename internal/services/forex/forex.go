package forex


import (
	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/utils"
)

func Routes() *chi.Mux {
	log := utils.GetLogger()
	log.Info("In details.Routes(), registering new details routes")
	router := chi.NewRouter()
	router.Get("/{curOne}/{curTwo}/rate", nil)
	return router
}
