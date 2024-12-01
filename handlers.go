package weather

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/bartalcorn/gothat/pkg/controller"
	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OWM_API_KEY")

type handlers struct{}

var handle *handlers

type Handler interface {
	base(w http.ResponseWriter, r *http.Request)
}

func (h *handlers) base(w http.ResponseWriter, r *http.Request) {
	wthr, err := owm.NewCurrent("F", "EN", apiKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Println("OpenWeatherMap error:", err)
	}

	wthr.CurrentByZipcode("30047", "US")
	ctx := context.WithValue(r.Context(), "module", "Weather")
	ctx = controller.AddContent(ctx, base(wthr))
	controller.Render(w, r.WithContext(ctx))
}
