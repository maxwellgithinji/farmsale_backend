package userscontroller

import (
	"context"
	"encoding/json"
	"farmsale_backend/config/mdb"
	"farmsale_backend/models/usersmodel"
	"farmsale_backend/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	user := &usersmodel.User{}

	var users []*usersmodel.User

	params := mux.Vars(req)

	var email = params["email"]

	filter := bson.D{{"email", email}}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	//find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"email": email})
	if err != nil {
		err := ErrorResponse{
			Err: "Email is invalid",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	if err = filterCursor.All(ctx, &users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	if len(users) == 0 {
		err := ErrorResponse{
			Err: `User with email (` + email + `) not found`,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	//TODO: replace search user by ID so they can delete by ID
	deleteUser, err := mdb.Users.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: `Deletion Failed`,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	msg := utils.MessageResponse{
		Msg: "Deletion Successful successful",
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteUser.DeletedCount)
	json.NewEncoder(w).Encode(msg)
}