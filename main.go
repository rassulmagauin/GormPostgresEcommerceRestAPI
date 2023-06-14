package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rassulmagauin/jsonstore/models"
)

type DBclient struct {
	db *gorm.DB
}

type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}
type OrderResponse struct {
	Order models.Order `json:"order"`
	Data  interface{}  `json:"data"`
}

func (driver *DBclient) GetUsersByFirstName(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	name := r.FormValue("first_name")

	var query = "select * from \"user\" where data->>'first_name'=?"
	driver.db.Raw(query, name).Scan(&users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(users)
	w.Write(respJSON)
}

func (driver *DBclient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)
	driver.db.First(&user, vars["id"])
	var userData interface{}
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(response)
	w.Write(respJson)
}

func (driver *DBclient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	postBody, _ := ioutil.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func (driver *DBclient) GetOrder(w http.ResponseWriter, r *http.Request) {
	var order = models.Order{}
	vars := mux.Vars(r)
	driver.db.First(&order, vars["id"])
	var orderData interface{}
	json.Unmarshal([]byte(order.Data), &orderData)
	var response = OrderResponse{Order: order, Data: orderData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(response)
	w.Write(respJson)
}

func (driver *DBclient) PostOrder(w http.ResponseWriter, r *http.Request) {
	var order = models.Order{}
	postBody, _ := ioutil.ReadAll(r.Body)
	order.Data = string(postBody)

	driver.db.Save(&order)
	responseMap := map[string]interface{}{"id": order.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBclient{db: db}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/v1/user/{id:[a-zA-Z0-9]*}", dbclient.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", dbclient.PostUser).Methods("POST")
	r.HandleFunc("/v1/user", dbclient.GetUsersByFirstName).Methods("GET")
	r.HandleFunc("/v1/order", dbclient.PostOrder).Methods("POST")
	r.HandleFunc("/v1/order/{id:[a-zA-Z0-9]*}", dbclient.GetOrder).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	log.Fatal(srv.ListenAndServe())

}
