package userscontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
	"github.com/maxwellgithinji/farmsale_backend/models/usersmodel"
	"github.com/maxwellgithinji/farmsale_backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ActivateDeactivateAccount godoc
// @Summary Admin activates or deactivates the account
// @Description ActivateDeactivateAccount is a safer option than deleting accounts which have interracted with the application
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} usersmodel.SuccessMessage
// @Router /admin/profile/activate/{id} [put]
// @Security ApiKeyAuth
func ActivateDeactivateAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	var users []*usersmodel.User

	params := mux.Vars(req)

	//id from params
	strID := params["id"]

	//Convert the id to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	// filter by the id
	filter := bson.D{{"_id", id}}

	// find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"_id": id})
	if err != nil {
		err := ErrorResponse{
			Err: "ID is invalid",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}

	if err = filterCursor.All(ctx, &users); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if len(users) == 0 {
		err := ErrorResponse{
			Err: `User with id (` + strID + `) not found`,
		}
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}

	var update bson.D
	var msg = map[string]interface{}{}

	if users[0].Isactive {
		update = bson.D{{"$set",
			bson.D{
				{"isactive", false},
			}}}

		msg["Message"] = "Account deactivation successful"

		updateUser, err := mdb.Users.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: `Update Failed`,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		if updateUser.MatchedCount != 0 {
			fmt.Println("matched and replaced an existing document dd") //TODO: delete in prod
		}
		if updateUser.UpsertedCount != 0 {
			fmt.Printf("inserted a new document with ID dd %v\n", updateUser.UpsertedID) //TODO: delete in prod
		}

	} else {

		update = bson.D{{"$set",
			bson.D{
				{"isactive", true},
			}}}

		msg["Message"] = "Account activation successful"

		updateUser, err := mdb.Users.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: `Update Failed`,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		if updateUser.MatchedCount != 0 {
			fmt.Println("matched and replaced an existing document") //TODO: delete in prod
		}
		if updateUser.UpsertedCount != 0 {
			fmt.Printf("inserted a new document with ID %v\n", updateUser.UpsertedID) //TODO: delete in prod
		}
	}
	json.NewEncoder(w).Encode(msg)
}

//TODO: Remember to log out user in the front end after deactivation
// DeactivateAccount godoc
// @Summary Owner of the account activates or deactivates the account
// @Description DeactivateAccount is a safer option than deleting accounts which have interracted with the application
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} usersmodel.SuccessMessage
// @Router /profile/deactivate/{id} [put]
// @Security ApiKeyAuth
func DeactivateAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	var users []*usersmodel.User

	params := mux.Vars(req)

	//id from params
	strID := params["id"]

	//Convert the id to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Println(err)
	}

	//filter by the id
	filter := bson.D{{"_id", id}}

	// find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"_id": id})
	if err != nil {
		err := ErrorResponse{
			Err: "ID is invalid",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}

	if err = filterCursor.All(ctx, &users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}

	if len(users) == 0 {
		err := ErrorResponse{
			Err: `User with id (` + strID + `) not found`,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}

	update := bson.D{{"$set",
		bson.D{
			{"isactive", false},
		}}}

	msg := utils.MessageResponse{
		Msg: "Account deactivation successful",
	}

	updateUser, err := mdb.Users.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: `Update Failed`,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if updateUser.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document") //TODO: delete in prod
		json.NewEncoder(w).Encode(msg)
		return
	}
	if updateUser.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", updateUser.UpsertedID) //TODO: delete in prod
	}
	json.NewEncoder(w).Encode(msg)
}
