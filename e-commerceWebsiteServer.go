package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []Product

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	file, err := ioutil.ReadFile("products.json")
	handleError(err)
	json.Unmarshal(file, &products)

	http.HandleFunc("/ProductRegistration", handlePostRequest)
	http.HandleFunc("/productData", handleGetRequest)
	http.HandleFunc("/updateProduct", handlePutRequest)
	http.HandleFunc("/deleteProduct", handleDeleteRequest)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handleGetRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		json.NewEncoder(writer).Encode(products)
		fmt.Println("product data successfully send to client")
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(writer, "Invalid method")
	}
}

func handlePostRequest(writer http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodPost {
		var product Product
		err := json.NewDecoder(request.Body).Decode(&product)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "Invalid request: %v", err)
			return
		}
		products = append(products, product)
		writer.WriteHeader(http.StatusAccepted)
		fmt.Println("Item successfully added to the slice")

	} else {

		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("method not found")
	}
}

func handlePutRequest(writer http.ResponseWriter, request *http.Request) {

	var newProduct Product
	json.NewDecoder(request.Body).Decode(&newProduct)

	var productExist = false
	for i, product := range products {
		if product.ID == newProduct.ID {
			products[i] = newProduct
			productExist = true
			break
		}
	}

	if !productExist {
		products = append(products, newProduct)
	}

	json.NewEncoder(writer).Encode(products)
	fmt.Println("product successfully updated to the slice")
}

func handleDeleteRequest(writer http.ResponseWriter, request *http.Request) {
	var delProduct int
	json.NewDecoder(request.Body).Decode(&delProduct)

	for i, product := range products {
		if product.ID == delProduct {
			products = append(products[:i], products[i+1:]...)
			break
		}
	}
}
