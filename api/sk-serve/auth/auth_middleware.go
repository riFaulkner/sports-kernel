package auth

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"log"
	"net/http"
	"strings"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user_roles"}
var userIdCtxKey = &contextKey{"user_id"}

type contextKey struct {
	name string
}

type UserRoles []*model.UserRoles

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func LoadUserRoles(client firestore.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject

		userRoles, err := getUserRolesByID(r.Context(), client, userId)
		if err != nil {
			// return and don't let it go any further, they don't have any roles. or something bad happened
		}

		// Add the user id to the context
		ctx := context.WithValue(r.Context(), userIdCtxKey, userId)

		// put roles in context
		ctx = context.WithValue(ctx, userCtxKey, userRoles)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}

func GetUserIdFromContext(ctx context.Context) string {
	raw, _ := ctx.Value(userIdCtxKey).(string)

	return raw
}
func GetUserRolesFromContext(ctx context.Context) UserRoles {
	raw, _ := ctx.Value(userCtxKey).([]*model.UserRoles)
	if raw == nil {
		return make([]*model.UserRoles, 0)
	}
	return raw
}

func (r UserRoles) ContainsRole(role model.Role, ctx context.Context) bool {
	graphVariables := graphql.GetOperationContext(ctx).Variables
	leagueID := ""
	if graphVariables["leagueId"] != nil {
		leagueID = fmt.Sprintf("%s", graphVariables["leagueId"])
	}
	teamID := ""
	if graphVariables["teamId"] != nil {
		teamID = fmt.Sprintf("%s", graphVariables["teamId"])
	}
	acceptableRoles := getAcceptableRoleStrings(role, leagueID, teamID)
	for _, item := range r {
		for _, role := range acceptableRoles {
			if *role == item.Role {
				return true
			}
		}
	}

	return false
}

func getUserRolesByID(cxt context.Context, client firestore.Client, userId string) ([]*model.UserRoles, error) {
	results, err := client.Collection(firestore.UsersCollection).
		Doc(userId).Collection(firestore.UserRolesCollection).Documents(cxt).GetAll()

	if err != nil {
		return nil, err
	}
	userRoles := make([]*model.UserRoles, 0)

	for _, result := range results {
		role := new(model.UserRoles)
		err = result.DataTo(&role)
		id := result.Ref.ID
		role.ID = id
		if err != nil {
			log.Printf("Error getting all user roles together: %v", err)
		}

		userRoles = append(userRoles, role)
	}
	return userRoles, nil
}

func getAcceptableRoleStrings(role model.Role, leagueID string, teamID string) []*string {
	acceptableRoles := make([]*string, 0, 2)
	// Admin is always accepted
	acceptableRoles = append(acceptableRoles, getAdminRole())

	// league managers can do any action that concerns their league
	if leagueID != "" {
		acceptableRoles = append(acceptableRoles, GetLeagueManagerRole(leagueID))
	}

	if leagueID != "" && strings.Contains(strings.ToLower(role.String()), "league") {
		acceptableRoles = append(acceptableRoles, GetLeagueMemberRole(leagueID))
	}
	if teamID != "" && strings.Contains(strings.ToLower(role.String()), "team") {
		acceptableRoles = append(acceptableRoles, GetTeamManagerRole(teamID))
	}

	return acceptableRoles
}

func getAdminRole() *string {
	role := fmt.Sprintf("skAdmin")
	return &role
}

func GetTeamManagerRole(teamID string) *string {
	role := fmt.Sprintf("%s:%s", model.RoleTeamOwner, teamID)
	return &role
}

func GetLeagueMemberRole(leagueID string) *string {
	role := fmt.Sprintf("%s:%s", model.RoleLeagueMember, leagueID)
	return &role
}

func GetLeagueManagerRole(leagueID string) *string {
	role := fmt.Sprintf("%s:%s", model.RoleLeagueManager, leagueID)
	return &role
}
