// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type LeaguePost struct {
	ID       string         `json:"id"`
	Author   string         `json:"author"`
	Title    string         `json:"title"`
	PostDate time.Time      `json:"postDate"`
	Content  string         `json:"content"`
	Comments []*PostComment `json:"comments"`
}

type LeagueTeamFiltering struct {
	TeamID  *string `json:"teamId"`
	OwnerID *string `json:"ownerId"`
}

type NewLeaguePost struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewPlayerNfl struct {
	PlayerName   string  `json:"playerName"`
	Position     string  `json:"position"`
	PositionRank *int    `json:"positionRank"`
	Team         NflTeam `json:"team"`
	Birthday     *string `json:"birthday"`
	Avatar       *string `json:"avatar"`
	OverallRank  *int    `json:"overallRank"`
}

type NewPostComment struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type NewUser struct {
	OwnerName string `json:"ownerName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type NewUserRole struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

type PlayerNfl struct {
	ID           string  `json:"id"`
	OverallRank  int     `json:"overallRank"`
	PlayerName   string  `json:"playerName"`
	Position     string  `json:"position"`
	PositionRank int     `json:"positionRank"`
	Team         NflTeam `json:"team"`
	Birthday     string  `json:"birthday"`
	Age          int     `json:"age"`
	Avatar       string  `json:"avatar"`
}

type PostComment struct {
	ID          string    `json:"id"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	CommentDate time.Time `json:"commentDate"`
}

type Transaction struct {
	TransactionType TransactionType `json:"transactionType"`
	OccurrenceDate  int             `json:"occurrenceDate"`
	TransactionData string          `json:"transactionData"`
}

type TransactionInput struct {
	TransactionType TransactionType `json:"transactionType"`
	TransactionData string          `json:"transactionData"`
}

type User struct {
	ID        string `json:"id"`
	OwnerName string `json:"ownerName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UserRoles struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

type ContractRestructureStatus string

const (
	ContractRestructureStatusEligible               ContractRestructureStatus = "ELIGIBLE"
	ContractRestructureStatusIneligibleFinalYear    ContractRestructureStatus = "INELIGIBLE_FINAL_YEAR"
	ContractRestructureStatusPreviouslyRestructured ContractRestructureStatus = "PREVIOUSLY_RESTRUCTURED"
)

var AllContractRestructureStatus = []ContractRestructureStatus{
	ContractRestructureStatusEligible,
	ContractRestructureStatusIneligibleFinalYear,
	ContractRestructureStatusPreviouslyRestructured,
}

func (e ContractRestructureStatus) IsValid() bool {
	switch e {
	case ContractRestructureStatusEligible, ContractRestructureStatusIneligibleFinalYear, ContractRestructureStatusPreviouslyRestructured:
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

type ContractStatus string

const (
	ContractStatusActive          ContractStatus = "ACTIVE"
	ContractStatusInactiveExpired ContractStatus = "INACTIVE_EXPIRED"
	ContractStatusInactiveDropped ContractStatus = "INACTIVE_DROPPED"
)

var AllContractStatus = []ContractStatus{
	ContractStatusActive,
	ContractStatusInactiveExpired,
	ContractStatusInactiveDropped,
}

func (e ContractStatus) IsValid() bool {
	switch e {
	case ContractStatusActive, ContractStatusInactiveExpired, ContractStatusInactiveDropped:
		return true
	}
	return false
}

func (e ContractStatus) String() string {
	return string(e)
}

func (e *ContractStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContractStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContractStatus", str)
	}
	return nil
}

func (e ContractStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlayerPosition string

const (
	PlayerPositionQb PlayerPosition = "QB"
	PlayerPositionRb PlayerPosition = "RB"
	PlayerPositionWr PlayerPosition = "WR"
	PlayerPositionTe PlayerPosition = "TE"
)

var AllPlayerPosition = []PlayerPosition{
	PlayerPositionQb,
	PlayerPositionRb,
	PlayerPositionWr,
	PlayerPositionTe,
}

func (e PlayerPosition) IsValid() bool {
	switch e {
	case PlayerPositionQb, PlayerPositionRb, PlayerPositionWr, PlayerPositionTe:
		return true
	}
	return false
}

func (e PlayerPosition) String() string {
	return string(e)
}

func (e *PlayerPosition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlayerPosition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlayerPosition", str)
	}
	return nil
}

func (e PlayerPosition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin         Role = "ADMIN"
	RoleLeagueManager Role = "LEAGUE_MANAGER"
	RoleLeagueMember  Role = "LEAGUE_MEMBER"
	RoleTeamOwner     Role = "TEAM_OWNER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleLeagueManager,
	RoleLeagueMember,
	RoleTeamOwner,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleLeagueManager, RoleLeagueMember, RoleTeamOwner:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionType string

const (
	TransactionTypeContractRestructure TransactionType = "CONTRACT_RESTRUCTURE"
	TransactionTypeDropPlayer          TransactionType = "DROP_PLAYER"
)

var AllTransactionType = []TransactionType{
	TransactionTypeContractRestructure,
	TransactionTypeDropPlayer,
}

func (e TransactionType) IsValid() bool {
	switch e {
	case TransactionTypeContractRestructure, TransactionTypeDropPlayer:
		return true
	}
	return false
}

func (e TransactionType) String() string {
	return string(e)
}

func (e *TransactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionType", str)
	}
	return nil
}

func (e TransactionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NflTeam string

const (
	NflTeamAri NflTeam = "ARI"
	NflTeamAtl NflTeam = "ATL"
	NflTeamBal NflTeam = "BAL"
	NflTeamBuf NflTeam = "BUF"
	NflTeamCar NflTeam = "CAR"
	NflTeamChi NflTeam = "CHI"
	NflTeamCin NflTeam = "CIN"
	NflTeamCle NflTeam = "CLE"
	NflTeamDal NflTeam = "DAL"
	NflTeamDen NflTeam = "DEN"
	NflTeamDet NflTeam = "DET"
	NflTeamFa  NflTeam = "FA"
	NflTeamGb  NflTeam = "GB"
	NflTeamKc  NflTeam = "KC"
	NflTeamHou NflTeam = "HOU"
	NflTeamInd NflTeam = "IND"
	NflTeamJac NflTeam = "JAC"
	NflTeamLac NflTeam = "LAC"
	NflTeamLar NflTeam = "LAR"
	NflTeamLv  NflTeam = "LV"
	NflTeamMin NflTeam = "MIN"
	NflTeamMia NflTeam = "MIA"
	NflTeamNe  NflTeam = "NE"
	NflTeamNo  NflTeam = "NO"
	NflTeamNyg NflTeam = "NYG"
	NflTeamNyj NflTeam = "NYJ"
	NflTeamPit NflTeam = "PIT"
	NflTeamPhi NflTeam = "PHI"
	NflTeamSea NflTeam = "SEA"
	NflTeamSf  NflTeam = "SF"
	NflTeamTb  NflTeam = "TB"
	NflTeamTen NflTeam = "TEN"
	NflTeamWas NflTeam = "WAS"
)

var AllNflTeam = []NflTeam{
	NflTeamAri,
	NflTeamAtl,
	NflTeamBal,
	NflTeamBuf,
	NflTeamCar,
	NflTeamChi,
	NflTeamCin,
	NflTeamCle,
	NflTeamDal,
	NflTeamDen,
	NflTeamDet,
	NflTeamFa,
	NflTeamGb,
	NflTeamKc,
	NflTeamHou,
	NflTeamInd,
	NflTeamJac,
	NflTeamLac,
	NflTeamLar,
	NflTeamLv,
	NflTeamMin,
	NflTeamMia,
	NflTeamNe,
	NflTeamNo,
	NflTeamNyg,
	NflTeamNyj,
	NflTeamPit,
	NflTeamPhi,
	NflTeamSea,
	NflTeamSf,
	NflTeamTb,
	NflTeamTen,
	NflTeamWas,
}

func (e NflTeam) IsValid() bool {
	switch e {
	case NflTeamAri, NflTeamAtl, NflTeamBal, NflTeamBuf, NflTeamCar, NflTeamChi, NflTeamCin, NflTeamCle, NflTeamDal, NflTeamDen, NflTeamDet, NflTeamFa, NflTeamGb, NflTeamKc, NflTeamHou, NflTeamInd, NflTeamJac, NflTeamLac, NflTeamLar, NflTeamLv, NflTeamMin, NflTeamMia, NflTeamNe, NflTeamNo, NflTeamNyg, NflTeamNyj, NflTeamPit, NflTeamPhi, NflTeamSea, NflTeamSf, NflTeamTb, NflTeamTen, NflTeamWas:
		return true
	}
	return false
}

func (e NflTeam) String() string {
	return string(e)
}

func (e *NflTeam) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NflTeam(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid nflTeam", str)
	}
	return nil
}

func (e NflTeam) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
