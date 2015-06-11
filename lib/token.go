package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	
	"github.com/dgrijalva/jwt-go"
)

var algType string = "RS256"
var key string = "data/weakgame"

// Helper func:  Read input from specified file or stdin
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	
	return ioutil.ReadAll(rdr)
}

// Verify a token and output the claims.  This is a great example
// of how to verify and view a token.
func VerifyToken(tokData []byte) (*jwt.Token, error) {
	var token *jwt.Token
	
	// trim possible whitespace from token
	tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})
	
	// Parse the token.  Load the key from command line option
	token, err := jwt.Parse(string(tokData), func(t *jwt.Token) (interface{}, error) {
		return loadData(key)
	})

	// Print an error if we can't parse for some reason
	if err != nil {
		return token, fmt.Errorf("Couldn't parse token: %v", err)
	}

	// Is token invalid?
	if !token.Valid {
		return token, fmt.Errorf("Token is invalid")
	}

	return token, nil
}

// Create, sign, and output a token.  This is a great, simple example of
// how to use this library to create and sign a token.
func SignToken(tokData []byte) (string, error) {
	var token *jwt.Token
	
	// parse the JSON of the claims
	var claims map[string]interface{}
	if err := json.Unmarshal(tokData, &claims); err != nil {
		return "", fmt.Errorf("Couldn't parse claims JSON: %v", err)
	}

	// get the key
	keyData, err := loadData(key)
	if err != nil {
		return "", fmt.Errorf("Couldn't read key: %v", err)
	}

	// get the signing alg
	alg := jwt.GetSigningMethod(algType)
	if alg == nil {
		return "", fmt.Errorf("Couldn't find signing method: %v", algType)
	}

	// create a new token
	token = jwt.New(alg)
	token.Claims = claims

	if token, err := token.SignedString(keyData); err == nil {
		return token, nil
	} else {
		return token, fmt.Errorf("Error signing token: %v", err)
	}
}