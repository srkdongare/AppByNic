package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/srkdongare/AppByNic/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.postProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI1", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI2", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI3", http.StatusBadRequest)
			return
		}
		p.putProducts(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Internal Server Error in json.Marshal", http.StatusInternalServerError)
	}
}

func (p *Products) postProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Internal Server Error in json.unMarshal", http.StatusBadRequest)
	}
	data.AddProduct(prod)

}

func (p *Products) putProducts(id int, rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle PUT Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Internal Server Error in json.unMarshal", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "ErrProductNotFound", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(rw, "Internal Server Error in updation", http.StatusInternalServerError)
		return
	}

}
