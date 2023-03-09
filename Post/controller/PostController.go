package controller

import (
	"blog/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostController struct {
	PostUsecase model.PostUsecase
}

func PostRouter(PostUsecase model.PostUsecase) {
	r := mux.NewRouter()
	http.Handle("/", r)
	handler := &PostController{
		PostUsecase: PostUsecase,
	}
	r.HandleFunc("/", handler.GetAll)
	r.HandleFunc("/detail/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/insert", handler.Insert).Methods("POST")
	r.HandleFunc("/update/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/delete/{id}", handler.Delete).Methods("DELETE")
}

func (pc *PostController) GetAll(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var res []model.Post
	res, err := pc.PostUsecase.GetAll()
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
	response.Message = "Get Data Success"
	response.Data = res
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)
}
func (pc *PostController) GetByID(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var res model.Post
	vars := mux.Vars(r)

	getid := vars["id"]
	id, err := strconv.Atoi(getid)

	res, err = pc.PostUsecase.GetByID(id)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
	response.Message = "Get Data Success"
	response.Data = res
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
func (pc *PostController) Insert(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var p *model.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.Message = err.Error()
	}
	_, err = pc.PostUsecase.Insert(p)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
	fmt.Println("Insert Succes")

	response.Message = "Insert Succes"
	response.Data = p
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (pc *PostController) Update(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var p *model.Post
	vars := mux.Vars(r)
	getId := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.Message = err.Error()
	}
	p.Id, err = strconv.Atoi(getId)
	_, err = pc.PostUsecase.Update(p)
	if err != nil {
		response.Message = err.Error()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println("Update Succes")

	response.Message = "Update Succes"
	response.Data = p
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
func (pc *PostController) Delete(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	vars := mux.Vars(r)
	getId := vars["id"]
	id, err := strconv.Atoi(getId)
	var res model.Post
	err = pc.PostUsecase.Delete(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Println("Delete Succes")

	response.Message = "Delete Succes"
	response.Data = res
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
