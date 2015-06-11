//WARNING.
//THIS HAS BEEN GENERATED AUTOMATICALLY BY AUTOAPI.
//IF THERE WAS A WARRANTY, MODIFYING THIS WOULD VOID IT.

package monsters

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/colingagnon/weakgame-api/db/mysql/monsters"
	dbi "github.com/colingagnon/weakgame-api/dbi/monsters"
)

func List(res http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(res)

	shouldFilter := false
	filterObject := &dbi.Monster{}
	
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

		if _, ok := req.Form["url"]; ok {
			form_url := req.FormValue("url")
			shouldFilter = true
			parsedField := form_url
			filterObject.Url = parsedField
		}

		if _, ok := req.Form["monsterlevel"]; ok {
			form_monsterlevel := req.FormValue("monsterlevel")
			shouldFilter = true
			i, _ := strconv.Atoi(form_monsterlevel)
			parsedField := int32(i)
			filterObject.Monsterlevel = parsedField
		}

		if _, ok := req.Form["hp"]; ok {
			form_hp := req.FormValue("hp")
			shouldFilter = true
			i, _ := strconv.Atoi(form_hp)
			parsedField := int32(i)
			filterObject.Hp = parsedField
		}

		if _, ok := req.Form["accuracy"]; ok {
			form_accuracy := req.FormValue("accuracy")
			shouldFilter = true
			i, _ := strconv.ParseFloat(form_accuracy, 32)
			parseField := float32(i)
			filterObject.Accuracy = parsedField
		}

		if _, ok := req.Form["dodge"]; ok {
			form_dodge := req.FormValue("dodge")
			shouldFilter = true
			i, _ := strconv.ParseFloat(form_dodge, 32)
			parseField := float32(i)
			filterObject.Dodge = parsedField
		}

		if _, ok := req.Form["damagelow"]; ok {
			form_damagelow := req.FormValue("damagelow")
			shouldFilter = true
			i, _ := strconv.Atoi(form_damagelow)
			parsedField := int32(i)
			filterObject.Damagelow = parsedField
		}

		if _, ok := req.Form["damagehigh"]; ok {
			form_damagehigh := req.FormValue("damagehigh")
			shouldFilter = true
			i, _ := strconv.Atoi(form_damagehigh)
			parsedField := int32(i)
			filterObject.Damagehigh = parsedField
		}

		if _, ok := req.Form["experiencepoints"]; ok {
			form_experiencepoints := req.FormValue("experiencepoints")
			shouldFilter = true
			i, _ := strconv.Atoi(form_experiencepoints)
			parsedField := int32(i)
			filterObject.Experiencepoints = parsedField
		}

	}
	*/
	
	if shouldFilter {
		rows, err := monsters.Find(filterObject)
		if err != nil {
			log.Println(err)
		}
		enc.Encode(rows)
		return
	}

	rows, err := monsters.All()
	if err != nil {
		log.Println(err)
	}
	enc.Encode(rows)
}

func Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	vars := mux.Vars(req)

	enc := json.NewEncoder(res)

	param := vars["id"]

	i, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := uint32(i)

	row, _ := monsters.GetById(id)
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

	_, get_err := monsters.GetById(id)
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
	row := &dbi.Monster{}
	dec.Decode(&row)
	return monsters.Save(row)
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

	monsters.DeleteById(id)
}
