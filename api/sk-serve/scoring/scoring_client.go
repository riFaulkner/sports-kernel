package scoring

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"io"
	"log"
	"net/http"
	"os"
)

func makePostRequest(r io.Reader, w io.Writer, targetURL string, audience string) error {
	var resp *http.Response
	var err error
	if os.Getenv("ENV") != "PROD" {
		request, err := http.NewRequest(http.MethodPost, targetURL, r)
		tokenSource, err := iDTokenTokenSource(context.Background(), audience)
		if err != nil {
			return fmt.Errorf("Error getting token %v", err)
		}
		token, err := tokenSource.Token()

		token.SetAuthHeader(request)
		request.Header.Set("Content-Type", "application/json")

		resp, err = http.DefaultClient.Do(request)
	} else {
		log.Printf("Using production client version")
		ctx := context.Background()
		var client *http.Client
		client, err = idtoken.NewClient(ctx, audience)
		if err != nil {
			return fmt.Errorf("idtoken.NewClient: %v", err)
		}

		resp, err = client.Post(targetURL, "application/json", r)
		log.Printf("Response: %v", resp)
		log.Printf("Response error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("scoring.service - Return status %d", resp.StatusCode)
		return fmt.Errorf("scoring.service - Return status %v", resp.StatusCode)
	}

	if _, err = io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	return nil
}

func iDTokenTokenSource(ctx context.Context, audience string) (oauth2.TokenSource, error) {
	// First we try the idtoken package, which only works for service accounts
	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		if err.Error() != `idtoken: credential must be service_account, found "authorized_user"` {
			return nil, err
		}
		// If that fails, we use our Application Default Credentials to fetch an id_token on the fly
		gts, err := google.DefaultTokenSource(ctx)
		if err != nil {
			return nil, err
		}
		ts = oauth2.ReuseTokenSource(nil, &idTokenSource{TokenSource: gts})
	}
	return ts, nil
}

// idTokenSource is an oauth2.TokenSource that wraps another
// It takes the id_token from TokenSource and passes that on as a bearer token
type idTokenSource struct {
	TokenSource oauth2.TokenSource
}

func (s *idTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.TokenSource.Token()
	if err != nil {
		return nil, err
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("token did not contain an id_token")
	}

	return &oauth2.Token{
		AccessToken: idToken,
		TokenType:   "Bearer",
		Expiry:      token.Expiry,
	}, nil
}
