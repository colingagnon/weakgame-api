//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package products

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/colingagnon/weakgame-api/db/mysql/products"
	dbi "github.com/colingagnon/weakgame-api/dbi/products"
	users "github.com/colingagnon/weakgame-api/db/mysql/users"
	"github.com/colingagnon/weakgame-api/lib"
)

func Purchase(res http.ResponseWriter, req *http.Request) {
	var authId int;
	
	res, req, authId = lib.GetAuthUser(res, req)
	
	userId := uint32(authId);
	if userId > 0 {
		user, _ := users.GetById(userId)
		
		// Now get product information
		var err error
		vars := mux.Vars(req)
	
		param := vars["id"]
	
		i, err := strconv.Atoi(param)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		id := uint32(i)
	
		product, get_err := products.GetById(id)
		if get_err != nil {
			fmt.Println(get_err)
			fmt.Fprintln(res, get_err)
			return
		}
		
		user.Unicorns += product.Tokens
		
		users.Save(user)
		
		enc := json.NewEncoder(res)
		enc.Encode(user)
	} else {
		res.WriteHeader(500)
		fmt.Fprint(res, "You must login before purchasing something")
	}
}

func List(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	enc := json.NewEncoder(res)

	shouldFilter := false
	filterObject := &dbi.Product{}
	
	/*
	req.ParseForm()
	if len(req.Form) > 0 {

		if _, ok := req.Form["id"]; ok {
			form_id := req.FormValue("id")
			shouldFilter = true
			i, _ := strconv.Atoi(form_id)
			parsedField := uint32(i)
			filterObject.Id = parsedField
		}

		if _, ok := req.Form["name"]; ok {
			form_name := req.FormValue("name")
			shouldFilter = true
			parsedField := form_name
			filterObject.Name = parsedField
		}

		if _, ok := req.Form["tokens"]; ok {
			form_tokens := req.FormValue("tokens")
			shouldFilter = true
			i, _ := strconv.Atoi(form_tokens)
			parsedField := int32(i)
			filterObject.Tokens = parsedField
		}

		if _, ok := req.Form["price"]; ok {
			form_price := req.FormValue("price")
			shouldFilter = true
			i, _ := strconv.ParseFloat(form_price, 32)
			parseField := float32(i)
			filterObject.Price = parsedField
		}

		if _, ok := req.Form["available"]; ok {
			form_available := req.FormValue("available")
			shouldFilter = true
			parsedField := []byte(form_available)
			filterObject.Available = parsedField
		}

	}
	*/
	
	if shouldFilter {
		rows, err := products.Find(filterObject)
		if err != nil {
			log.Println(err)
		}
		enc.Encode(rows)
		return
	}

	rows, err := products.All()
	if err != nil {
		log.Println(err)
	}
	enc.Encode(rows)
}

func Get(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	enc := json.NewEncoder(res)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	row, _ := products.GetById(id)
	enc.Encode(row)
}

func Post(res http.ResponseWriter, req *http.Request) {
	err := save(req)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, err)
	}
}

func Put(res http.ResponseWriter, req *http.Request) {
	var err error
	vars := mux.Vars(req)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	_, get_err := products.GetById(id)
	if get_err != nil {
		fmt.Println(get_err)
		fmt.Fprintln(res, get_err)
		return
	}
	err = save(req)
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, err)
	}
}

func save(req *http.Request) error {
	dec := json.NewDecoder(req.Body)
	row := &dbi.Product{}
	dec.Decode(&row)
	return products.Save(row)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	products.DeleteById(id)
}
