package products

import (
	"log"
	"net/http"

	"github.com/jabrogena5100/MarketGoLang.git/internal/json"
)

type handler struct { 
	service Service 
}

func NewHandler(service Service) *handler { 
	return &handler{
		service: service, 
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil { 
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

// delete this since we already are using products
// 	products := struct { 
// 	Products []string `json:"products"`
// }{}

json.Write(w, http.StatusOK, products)
}