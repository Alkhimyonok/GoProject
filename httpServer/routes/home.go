package routes

import (
	"fmt"
	"htttpServer/models"
	"htttpServer/utils"
	"net/http"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	products, err := models.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	utils.ExecuteTemplate(w, "home.html", struct {
		Categories []models.Category
		Products   []models.Product
	}{
		Categories: categories,
		Products:   products,
	})
}
func homePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")
	r.ParseForm()
	search := r.PostForm.Get("search")
	products, err := models.SearchProducts(search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	count := len(products)
	var html string = ""
	if count > 0 {
		html += "<table class='table table-bordered'>"
		html += fmt.Sprintf("<th> Id </th> <th> Category </th> <th> Name </th> <th> Price </th> <th> Quantity </th> <th> Amount </th>")
		for _, product := range products {
			html += "<tr>"
			html += fmt.Sprintf("<td> %d </td> <td> %s </td> <td> %s </td> <td> %.2f $ </td> <td> %d </td> <td> %.2f </td>",
				product.Id, product.Category.Description, product.Name, product.Price, product.Quantity, product.Amount)
			html += "</tr>"
		}
		html += "</table>"
	} else {
		html += fmt.Sprintf("<p class='alert alert-info'> Nothing found with <code>'<strong>%s</strong>'</code></p>", search)
	}
	w.Write([]byte(html))
}
