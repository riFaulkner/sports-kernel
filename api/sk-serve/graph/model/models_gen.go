// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CapUtilizationSummary struct {
	CapUtilization int `json:"capUtilization"`
	NumContracts   int `json:"numContracts"`
}

type Contract struct {
	ID                  string                    `json:"id"`
	PlayerID            string                    `json:"playerId"`
	Player              *PlayerNfl                `json:"player"`
	TeamID              string                    `json:"teamId"`
	CurrentYear         int                       `json:"currentYear"`
	RestructureStatus   ContractRestructureStatus `json:"restructureStatus"`
	TotalContractValue  float64                   `json:"totalContractValue"`
	TotalRemainingValue float64                   `json:"totalRemainingValue"`
	ContractLength      int                       `json:"contractLength"`
	PlayerPosition      *string                   `json:"playerPosition"`
	ContractDetails     []*ContractYear           `json:"contractDetails"`
}

type ContractDetail struct {
	RestructuredContract      bool    `json:"restructuredContract"`
	TotalRemainingValue       float64 `json:"totalRemainingValue"`
	CurrentYearRemainingValue float64 `json:"currentYearRemainingValue"`
	Year1value                float64 `json:"year1value"`
	Year2value                float64 `json:"year2value"`
	Year3Value                float64 `json:"year3Value"`
	Year4value                float64 `json:"year4value"`
}

type ContractInput struct {
	PlayerID            string                    `json:"playerId"`
	TeamID              string                    `json:"teamId"`
	CurrentYear         int                       `json:"currentYear"`
	RestructureStatus   ContractRestructureStatus `json:"restructureStatus"`
	TotalContractValue  *float64                  `json:"totalContractValue"`
	TotalRemainingValue *float64                  `json:"totalRemainingValue"`
	ContractLength      *int                      `json:"contractLength"`
	PlayerPosition      string                    `json:"playerPosition"`
	ContractDetails     []*ContractYearInput      `json:"contractDetails"`
}

type ContractYear struct {
	Year             int     `json:"year"`
	TotalAmount      int     `json:"totalAmount"`
	PaidAmount       int     `json:"paidAmount"`
	GuaranteedAmount float64 `json:"guaranteedAmount"`
}

type ContractYearInput struct {
	Year             int     `json:"year"`
	TotalAmount      float64 `json:"totalAmount"`
	PaidAmount       float64 `json:"paidAmount"`
	GuaranteedAmount float64 `json:"guaranteedAmount"`
}

type ContractsMetadata struct {
	TotalUtilizedCap  int                    `json:"totalUtilizedCap"`
	TotalAvailableCap int                    `json:"totalAvailableCap"`
	QbUtilizedCap     *CapUtilizationSummary `json:"qbUtilizedCap"`
	RbUtilizedCap     *CapUtilizationSummary `json:"rbUtilizedCap"`
	WrUtilizedCap     *CapUtilizationSummary `json:"wrUtilizedCap"`
	TeUtilizedCap     *CapUtilizationSummary `json:"teUtilizedCap"`
}

type Division struct {
	DivisionName string `json:"divisionName"`
	LeadingWins  *int   `json:"leadingWins"`
}

type DraftPick struct {
	Round int  `json:"round"`
	Value *int `json:"value"`
}

type DraftYear struct {
	Year  int          `json:"year"`
	Picks []*DraftPick `json:"picks"`
}

type League struct {
	ID         string      `json:"id"`
	LeagueName string      `json:"leagueName"`
	LogoURL    string      `json:"logoUrl"`
	StartDate  time.Time   `json:"startDate"`
	Teams      []*Team     `json:"teams"`
	Divisions  []*Division `json:"divisions"`
}

type NewPlayerNfl struct {
	PlayerName   string  `json:"playerName"`
	Position     string  `json:"position"`
	PositionRank *int    `json:"positionRank"`
	TeamNfl      *string `json:"teamNFL"`
	Birthday     *string `json:"birthday"`
	Avatar       *string `json:"avatar"`
	OverallRank  *int    `json:"overallRank"`
}

type NewTeam struct {
	ID          string     `json:"id"`
	TeamName    string     `json:"teamName"`
	Division    *string    `json:"division"`
	FoundedDate *time.Time `json:"foundedDate"`
}

type NewUser struct {
	OwnerName string `json:"ownerName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type PlayerNfl struct {
	ID           string `json:"id"`
	OverallRank  int    `json:"overallRank"`
	PlayerName   string `json:"playerName"`
	Position     string `json:"position"`
	PositionRank int    `json:"positionRank"`
	TeamNfl      string `json:"teamNFL"`
	Birthday     string `json:"birthday"`
	Avatar       string `json:"avatar"`
}

type Team struct {
	ID                       string             `json:"id"`
	FoundedDate              time.Time          `json:"foundedDate"`
	TeamName                 string             `json:"teamName"`
	OwnerID                  string             `json:"ownerID"`
	Division                 *string            `json:"division"`
	CurrentContractsMetadata *ContractsMetadata `json:"currentContractsMetadata"`
	TeamAssets               *TeamAssets        `json:"teamAssets"`
}

type TeamAssets struct {
	DraftPicks []*DraftYear `json:"draftPicks"`
}

type User struct {
	ID        string `json:"id"`
	OwnerName string `json:"ownerName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UserPreferences struct {
	ID                string    `json:"id"`
	OwnerName         string    `json:"ownerName"`
	PreferredLeagueID *string   `json:"preferredLeagueId"`
	Leagues           []*League `json:"leagues"`
}

type ContractRestructureStatus string

const (
	ContractRestructureStatusEligible               ContractRestructureStatus = "ELIGIBLE"
	ContractRestructureStatusIneligible             ContractRestructureStatus = "INELIGIBLE"
	ContractRestructureStatusPreviouslyRestructured ContractRestructureStatus = "PREVIOUSLY_RESTRUCTURED"
)

var AllContractRestructureStatus = []ContractRestructureStatus{
	ContractRestructureStatusEligible,
	ContractRestructureStatusIneligible,
	ContractRestructureStatusPreviouslyRestructured,
}

func (e ContractRestructureStatus) IsValid() bool {
	switch e {
	case ContractRestructureStatusEligible, ContractRestructureStatusIneligible, ContractRestructureStatusPreviouslyRestructured:
		return true
	}
	return false
}

func (e ContractRestructureStatus) String() string {
	return string(e)
}

func (e *ContractRestructureStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContractRestructureStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContractRestructureStatus", str)
	}
	return nil
}

func (e ContractRestructureStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
