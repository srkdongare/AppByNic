package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
	Tag	string `json:"-"`
}

type Products []*Product

func (ps *Products) ToJSON(wr io.Writer) error {
	encoder := json.NewEncoder(wr)
	return encoder.Encode(ps)
}

func (p *Product) FromJSON(rd io.Reader) error {
	decoder := json.NewDecoder(rd)
	return decoder.Decode(p)
}

//
func GetProducts() Products {
	return ProductList
}

//
func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func getNextID() int {
	prod := ProductList[len(ProductList)-1]
	return prod.ID + 1
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	ProductList[pos] = prod
	return nil
}

//Error if proct id not found
var ErrProductNotFound = fmt.Errorf("Product Not Found.")

func findProduct(id int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

//
var ProductList = []*Product{
	&Product{
		ID:          1000,
		Name:        "Tea",
		Description: "Masala Tea",
		Price:       20,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          10001,
		Name:        "Coffee",
		Description: "Espresso",
		Price:       50,
		SKU:         "abc456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
