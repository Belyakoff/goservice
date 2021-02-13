package handlers

import (
	"net/http"
	"github.com/Belyakoff/goservice/data"
)


func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod.Name)
	data.AddProduct(prod)
}