package Products

import (
	"encoding/json"
	"fmt"
	 "math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Declare a Product object the struct togather data related to the Products:
type Product struct {
	ID    string `json:"id"`
	Name string `json:"itle"`
	Maker *Maker `json:"maker"`
}

// Declare a Maker struct:
type Maker struct {
	Brand string `json:"brand"`
	Owner string `json:"owner"`
}

// declare a type to represent all the Products:
type Store struct {
	Products []Product
}

// Instqntiqte q new Maker:
func NewMaker(brand, owner string) *Maker {
	Maker := new(Maker)
	Maker.Brand = brand
	Maker.Owner = owner
	return Maker
}

// Instantiate a new Product:
func NewProduct(id, name, brand, owner string) *Product {
	Product := new(Product)
	Product.ID = id
	Product.Name = name
	Product.Maker = NewMaker(brand, owner)
	return Product
}

// Instantiate a new sequence of Products:
func InitProducts() *Store {
	return new(Store)
}

// Create a global instance of the store
var AllProducts = InitProducts()


// Get all Products handler:
func GetAll(wr http.ResponseWriter, rq *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(AllProducts)
}

// Get a specific Product handler:
func GetOne(wr http.ResponseWriter, rq *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rq)
	for _, product := range AllProducts.Products {
		if product.ID == params["id"] {
			json.NewEncoder(wr).Encode(product)
			break  
		}
	}
}

// Create a new Product:
func CreateProduct(wr http.ResponseWriter, rq *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(rq.Body).Decode(&product)
	randomId := strconv.Itoa(rand.Intn(100000000))
	fmt.Println(randomId)
	AllProducts.AppendProduct(randomId, "name "+ randomId, "brand "+ randomId, "owner "+ randomId)
	json.NewEncoder(wr).Encode(AllProducts)
}

// Update an existing product:
func UpdateProduct(wr http.ResponseWriter, rq *http.Request) {
	// Set the json content type:
	wr.Header().Set("Content-Type", "application/json")
	// Get params:
	params := mux.Vars(rq)
	// loop over the products
	for index, product := range AllProducts.Products {
		if product.ID == params["id"] {
			fmt.Println("Found")
			AllProducts.Products[index].ID = "Updated"
			AllProducts.Products[index].Name= "Updated"
			AllProducts.Products[index].Maker.Brand = "Updated"
			AllProducts.Products[index].Maker.Owner= "Updated"
			json.NewEncoder(wr).Encode(AllProducts.Products[index])
			break
		}
	}
	
}

// Delete an existing product:
func DeleteProduct(wr http.ResponseWriter, rq *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rq)
	for index, product := range AllProducts.Products {
		if product.ID == params["id"] {
			AllProducts.Products = append(AllProducts.Products[:index], AllProducts.Products[index+1:]...)
			fmt.Println(AllProducts)
			json.NewEncoder(wr).Encode(AllProducts)
			break
		}
	}
}

// Append a new product to the market:
func (store *Store) AppendProduct(id, name, brand, owner string) {
	instance := NewProduct(id, name, brand, owner)
	store.Products = append(store.Products, *instance)
}

// Display the content of my Store:
func (store *Store) Display() {
	for _, item := range store.Products {
		fmt.Printf("The product has the id %v and name %v and brand is %v the maker is %v\n", item.ID, item.Name, item.Maker.Brand, item.Maker.Owner)
	}
}

// Routing function:
func Routing() {
	// Create a new instance of the mux router
	router := mux.NewRouter()
	router.HandleFunc("/products", GetAll).Methods("GET")
	router.HandleFunc("/products/{id}", GetOne).Methods("GET")
	router.HandleFunc("/create", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("Delete")
	fmt.Println("Starting a server at port: 8080")
	http.ListenAndServe(":8080", router)
}
