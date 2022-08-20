package user

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/league"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserOnboardingService struct {
	LeagueService league.LeagueRepository
	UserService   UserService
	TeamService   team.TeamService
}

func (s UserOnboardingService) GenerateAccessCode(ctx context.Context, leagueID string, teamID string, role model.Role) (string, error) {
	return s.TeamService.GenerateAccessCode(ctx, leagueID, teamID, role)
}

func (s UserOnboardingService) OnboardWithAccessCode(ctx context.Context, accessCode string, ownerID string) (*UserPreferences, error) {
	decodedAccessCode, ok := s.TeamService.ValidateAccessToken(ctx, accessCode)
	if ok {
		isUserAlreadyTeamOwnerInLeague := s.isUserPreexistingTeamOwnerInLeague(ctx, decodedAccessCode.LeagueID, ownerID)
		if isUserAlreadyTeamOwnerInLeague {
			return nil, gqlerror.Errorf("User is already team owner in league")
		}

		decodedAccessCode.LeagueName = s.getLeagueNameForOnBoarding(ctx, decodedAccessCode.LeagueID)

		ok = s.TeamService.AddUserToTeamAndConsumeAccessCode(ctx, decodedAccessCode, ownerID)
		if ok {
			return s.UserService.AddTeamToUser(ctx, decodedAccessCode, ownerID)
		}
	}
	return nil, gqlerror.Errorf("Unable to onboard user")
}

func (s UserOnboardingService) isUserPreexistingTeamOwnerInLeague(ctx context.Context, leagueID string, ownerID string) bool {
	_, err := s.TeamService.GetTeamByOwnerID(ctx, leagueID, ownerID)
	if err != nil {
		if status.Code(err) != codes.NotFound {
			gqlerror.Errorf("Error validating user is not already a member of the league")
			return true
		}
	} else {
		gqlerror.Errorf("Error adding user to team, user is already owner of a team in the league")
		return true
	}
	return false
}
func (s UserOnboardingService) getLeagueNameForOnBoarding(ctx context.Context, leagueID string) string {
	league, err := s.LeagueService.GetByLeagueId(ctx, leagueID)
	if err != nil {
		log.Printf("Error getting league to add user, please try again")
		return ""
	}
	return league.LeagueName
}
