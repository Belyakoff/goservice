package handlers 

import (
	"context"
	"net/http"
	"fmt"
	"github.com/Belyakoff/goservice/data"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){


		prod := &data.Product{}              

         
		err := data.FromJSON(prod, r.Body) 
		if err != nil {
			// we should never be here but log the error just incase
			p.l.Println("[ERROR] deserializing product", err)


			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		v := data.NewValidation()

	  

		errs := v.Validate(prod)

		fmt.Printf("Middle v%",prod)
		
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		} 


		ctx := context.WithValue(r.Context(), KeyProduct{}, *prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

	})
}