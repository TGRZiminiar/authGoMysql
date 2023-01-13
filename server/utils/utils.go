package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)


var Validate *validator.Validate

func ValidatorConnect(){

	Validate = validator.New();
}

const SecretKey = "2e9joafhxcv9j234kndfaouhwfknweug"

func GenerateToken(userId int64) (string, error) {
	
	strUserId := strconv.FormatInt(userId, 10)

	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strUserId,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	

	token, err := claims.SignedString([]byte(SecretKey))
	
	return token,err;

}

func GetUserIdFromToken(tokenString string) (int64, error) {
	// Validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		
		// Return the secret key
		return []byte(SecretKey), nil
	})

	if err != nil {
		// Handle the error
		return 0, err
	}

	// Check if the token is valid
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// The token is valid, so you can access the claims like this:
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims type")
	}

	// Get the issuer (user ID) from the claims
	userId := claims["iss"].(string)
	
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)

	return userIdInt, nil
}

func GenerateuuidInt() int64 {
	uniqueNumber := time.Now().UnixNano()/(1<<22)/999999
	return uniqueNumber;
}

func GenerateuuidBigInt()int64{
	
	uuid := uuid.NewV4().String()
    var i big.Int
	// i, _ := strconv.ParseInt(uuid, 10, 64)
    // fmt.Println(i.String())
	
	i.SetString(strings.Replace(uuid, "-", "", 1), 16)
    //or if your uuid is [16]byte
    //i.SetBytes(uuid[:])
	
	return i.Int64();

}

