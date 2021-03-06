package productscontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
	"github.com/maxwellgithinji/farmsale_backend/models/productsmodel"

	"github.com/globalsign/mgo/bson"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

// GetProducts godoc
// @Summary Get details of all products
// @Description Get details of all products
// @Tags auth User
// @Produce  json
// @Success 200 {object} []productsmodel.Product{}
// @Router /auth/products [get]
// @Security ApiKeyAuth
func Index(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	var prods = []productsmodel.Product{}
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	ctx := context.Background()
	cur, err := mdb.Products.Find(ctx, &bson.M{})
	if err != nil {
		fmt.Printf("%+v\n", err)
		err := ErrorResponse{
			Err: "Error finding products",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = cur.All(ctx, &prods)
	if err != nil {
		fmt.Printf("%+v\n", err)
		err := ErrorResponse{
			Err: "Error finding all products",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(prods)
}
