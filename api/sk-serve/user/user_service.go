package user

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/auth"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/user/crossfunctional"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	UserRepository UserRepository
}

func (u UserService) AddTeamToUser(ctx context.Context, decodedAccessCode crossfunctional.DecodedAccessCode, userID string) (*UserPreferences, error) {
	//	Get the User's data, specifically user preferences
	ok := false
	//	If they have a user preferences record add the new league to their list of leagues, you'll need the new league's ID and name
	userPreferences, err := u.GetUserPreferences(ctx, userID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			//	If they don't have a user preferences record, create a new default one with the new league in the list of leagues.
			userPreferences, ok = u.createDefaultUserPreferences(ctx, userID)
			if !ok {
				return nil, gqlerror.Errorf("Error creating default user preferences")
			}
		} else {
			log.Errorf(ctx, "Error getting user preferences")
			return nil, err
		}
	}

	ok = u.UserRepository.AddLeagueToUserPreferences(ctx, userID, UserPreferencesLeagueSnippet{
		Id:           decodedAccessCode.LeagueID,
		LeagueName:   decodedAccessCode.LeagueName,
		RoleInLeague: decodedAccessCode.Role,
	})

	if ok {
		//	Add the user roles for their new league
		//	Included in the payload is the level of permissions to give the user, use that to give the user permissions
		ok = u.createRolesForNewUserLeague(ctx, userID, decodedAccessCode)
		if ok {
			return userPreferences, nil
		} else {
			err = gqlerror.Errorf("Error updating new roles for user")
		}
	} else {
		err = gqlerror.Errorf("Error adding team to user preferences")
	}
	return nil, err
}

func (u UserService) GetAll(ctx context.Context) ([]*model.User, error) {
	return u.UserRepository.GetAll(ctx)
}

func (u UserService) Create(ctx context.Context, user UserPreferences) error {
	return u.UserRepository.Create(ctx, user)
}

func (u UserService) GetUserPreferences(ctx context.Context, userId string) (*UserPreferences, error) {
	return u.UserRepository.GetUserPreferences(ctx, userId)
}

func (u UserService) CreateUserRole(ctx context.Context, newRole *model.NewUserRole) (*model.UserRoles, error) {
	return u.UserRepository.CreateUserRole(ctx, newRole)
}

func (u UserService) GetUserRoles(ctx context.Context, userID *string) ([]*model.UserRoles, error) {
	return u.UserRepository.GetUserRoles(ctx, userID)
}

func (u UserService) createDefaultUserPreferences(ctx context.Context, userID string) (*UserPreferences, bool) {
	isAdmin := false
	userPreferences := UserPreferences{
		ID:        userID,
		OwnerName: "",
		Leagues:   []*league.League{},
		IsAdmin:   &isAdmin,
	}
	err := u.Create(ctx, userPreferences)
	if err != nil {
		return nil, false
	}
	return &userPreferences, true
}
func (u UserService) createRolesForNewUserLeague(ctx context.Context, userID string, decodedAccessCode crossfunctional.DecodedAccessCode) bool {
	leagueMemberRole := model.NewUserRole{
		UserID: userID,
		Role:   *auth.GetLeagueMemberRole(decodedAccessCode.LeagueID),
	}
	_, err := u.UserRepository.CreateUserRole(ctx, &leagueMemberRole)
	if err != nil {
		log.Errorf(ctx, "Error creating user role League member")
		return false
	}
	if decodedAccessCode.Role == model.RoleLeagueMember.String() {
		return true
	}

	teamOwnerRole := model.NewUserRole{
		UserID: userID,
		Role:   *auth.GetTeamManagerRole(decodedAccessCode.TeamID),
	}

	_, err = u.UserRepository.CreateUserRole(ctx, &teamOwnerRole)
	if err != nil {
		log.Errorf(ctx, "Error creating user role Team owner")
		return false
	}

	if decodedAccessCode.Role == model.RoleTeamOwner.String() {
		return true
	}

	if model.RoleLeagueManager.String() == decodedAccessCode.Role {
		leagueManagerRole := model.NewUserRole{
			UserID: userID,
			Role:   *auth.GetLeagueManagerRole(decodedAccessCode.LeagueID),
		}
		_, err := u.UserRepository.CreateUserRole(ctx, &leagueManagerRole)
		if err != nil {
			return false
		}
		return true
	}

	return false
}
