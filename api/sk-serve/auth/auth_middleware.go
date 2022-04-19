package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"log"
	"net/http"
	"strings"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(client firestore.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if token == "" {
				http.Error(w, "User unauthenticated", http.StatusForbidden)
				return
				//next.ServeHTTP(w, r)
			}

			userId, err := validateAndGetUserID(token)

			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// get the user from the database
			user := getUserByID(client, userId)

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

func validateAndGetUserID(tokenString string) (string, error) {
	tokenHeaderStrings := strings.Split(tokenString, " ")
	tokenString = tokenHeaderStrings[len(tokenHeaderStrings)-1]

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return token.Header["kid"], nil
	})

	log.Printf("token %v", token)
	//token.Claims

	//jwt.MapClaims{}
	return "", nil
}

func getUserByID(client firestore.Client, userId string) *model.User {
	return &model.User{
		ID:        userId,
		OwnerName: "rick",
		Email:     "ktarfum@gmail.com",
	}
}
