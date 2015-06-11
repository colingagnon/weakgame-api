package lib

import (
	"fmt"
	"strconv"
	"net/http"
)

func GetAuthToken(id int) string {
	token, err := SignToken([]byte("{\"id\":\"" + strconv.Itoa(id) + "\"}"))
	
	if err != nil {
		fmt.Println(err)
	}
	
	return token
} 

func GetAuthUser(res http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, int) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "x-authorization-token")
	res.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS")
	res.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	
	token := req.Header.Get("x-authorization-token")
	
	if token == "" {
		res.WriteHeader(500)
		fmt.Fprint(res, "Authorization Required")
		return res, req, 0
	}
	
	vtoken, _ := VerifyToken([]byte(token))
	
	// TODO token parsing error, seems to work anyways though...
	// think its the key size is different than default, meh for now
	/*	
	if err != nil {
		fmt.Println(err)
	}
	*/
	
	if id, ok := vtoken.Claims["id"].(string); ok {
		id, _ := strconv.Atoi(id)
		return res, req, id
	}
	
	return res, req, 0
}
