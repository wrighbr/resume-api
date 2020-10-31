package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
	"github.com/shaj13/go-guardian/store"
	"github.com/wrighbr/resume-api/client"
	"github.com/wrighbr/resume-api/models"
)

var authenticator auth.Authenticator
var cache store.Cache

// func createToken(w http.ResponseWriter, r *http.Request) {
// 	token := uuid.New().String()
// 	body := fmt.Sprintf("token: %s \n", token)
// 	w.Write([]byte(body))
// }

func createToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	user := auth.NewDefaultUser("medium", "1", nil, nil)
	tokenStrategy := authenticator.Strategy(bearer.CachedStrategyKey)
	auth.Append(tokenStrategy, token, user, r)
	body := fmt.Sprintf("token: %s \n", token)
	w.Write([]byte(body))
}

func setupGoGuardian() {
	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*10)

	basicStrategy := basic.New(validateUser, cache)
	tokenStrategy := bearer.New(bearer.NoOpAuthenticate, cache)

	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)
}

func validateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	// here connect to db or any other service to fetch user and validate it.
	// if userName == "medium" && password == "medium" {
	// 	return auth.NewDefaultUser("medium", "1", nil, nil), nil
	// }
	encodedpass := base64.StdEncoding.EncodeToString([]byte(password))
	info := client.GetUser(userName)

	var myData models.Auth
	mapstructure.Decode(info, &myData)
	if userName == myData.Username && encodedpass == myData.Password {
		return auth.NewDefaultUser(userName, "1", nil, nil), nil
	}

	return nil, fmt.Errorf("Invalid credentials")
}

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		user, err := authenticator.Authenticate(r)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Printf("User %s Authenticated\n", user.UserName())
		next.ServeHTTP(w, r)
	})
}
