package controllers

import (
	"clean-architecture-golang-example/entities"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ProductController struct {
	versionApi string
	logger     *log.Logger
	http       *http.ServeMux
	usecase    *entities.ProductUsecase
}

func NewProductController(versionApi string, logger *log.Logger, http *http.ServeMux, usecase *entities.ProductUsecase) *ProductController {
	controller := &ProductController{versionApi, logger, http, usecase}
	controller.Route()

	return controller
}

func (Controller *ProductController) Gets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		controller := *Controller.usecase

		if id == 0 {
			result, err := controller.Gets()
			for _, each := range result.Data {
				Controller.logger.Println("[SUCCESS] ID : ", each.ID)
			}
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Controller.logger.Println("[ERROR]:" + err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				Controller.logger.Println("[SUCCESS]: Gets product is success")
			}

			json.NewEncoder(w).Encode(result)
		} else {
			result, err := controller.GetOne(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Controller.logger.Println("[ERROR]:" + err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				Controller.logger.Println("[SUCCESS]: Get product by id is success")
			}

			json.NewEncoder(w).Encode(result)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		var data entities.CreateProductRequest

		controller := *Controller.usecase
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		}

		result, err := controller.Create(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		} else {
			w.WriteHeader(http.StatusCreated)
			Controller.logger.Println("[SUCCESS]: Create product is success")
		}
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPut {
		var data entities.UpdateProductRequest

		controller := *Controller.usecase
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		}

		result, err := controller.Update(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			Controller.logger.Println("[SUCCESS] Update product is success")

		}
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodDelete {
		var data entities.DeleteProductRequest

		controller := *Controller.usecase
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		}

		result, err := controller.DeleteByID(data.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			Controller.logger.Println("[SUCCESS] Delete product is success")

		}
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Route() {
	Controller.http.HandleFunc(Controller.versionApi+"/product/gets", Controller.Gets)
	Controller.http.HandleFunc(Controller.versionApi+"/product/create", Controller.Create)
	Controller.http.HandleFunc(Controller.versionApi+"/product/update", Controller.Update)
	Controller.http.HandleFunc(Controller.versionApi+"/product/delete", Controller.Delete)
}
