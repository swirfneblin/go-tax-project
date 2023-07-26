package main

import (
	// "encoding/json"
	"net/http"

	// "github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"
	"github.com/labstack/echo"
	"github.com/swirfneblin/go-tax-project/internal/entity"
)

func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Order)
	// http.ListenAndServe(":8888", r)
	e := echo.New()
	e.GET("/order", Order)
	e.Logger.Fatal(e.Start(":8888"))
}

// order endpoint for echo framework
func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "1232",
		Price: 1000,
		Tax:   10,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}

// func Order(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(w).Encode("Method not allowed")
// 		return
// 	}

// 	order, err := entity.NewOrder("123", 1000, 10)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	order.CalculateFinalPrice()
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(order)

// }
