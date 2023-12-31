package controllers

import (
	"html/template"
	"log"
	"net/http"
	"storeapp/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func LoadEditProduct(w http.ResponseWriter, r *http.Request) {
	editId := r.URL.Query().Get("id")
	product := models.GetProductById(editId)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("ERROR  while converting price: ", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("ERROR while converting quantitiy: ", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedId, err := strconv.Atoi(id)

		if err != nil {
			log.Println("ERROR  while converting id: ", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("ERROR  while converting price: ", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("ERROR while converting quantitiy: ", err)
		}

		productToUpdate := models.ConvertDataToProduct(convertedId, name, description, convertedPrice, convertedQuantity)
		models.UpdateProduct(*productToUpdate)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DelectProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
