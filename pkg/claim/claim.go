package claim

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
    ID int `json:"id"`
    jwt.StandardClaims
}

func (c *Claim) NewToken(signingString string) (string, error) {

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    
    return token.SignedString([]byte(signingString))
}

func GetFromToken(tokenString, signingString string) (*Claim, error) {

    token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
        return []byte(signingString), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("Invalid token")
    }

    claim, ok := token.Claims.(jwt.MapClaims)

    if !ok {
        return nil, errors.New("Invalid claim")
    }

    iID, ok := claim["id"]

    if !ok {
        return nil, errors.New("User id not found")
    }
    
    d, ok := iID.(float64)

    if !ok {
        return nil, errors.New("Invalid user id")
    }

    return &Claim{ID: int(d)}, nil
}