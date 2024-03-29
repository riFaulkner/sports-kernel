package team

import (
	"time"
)

type TeamMutations struct {
}

type Team struct {
	ID                       string               `json:"id"`
	FoundedDate              time.Time            `json:"foundedDate"`
	TeamName                 string               `json:"teamName"`
	Division                 *string              `json:"division"`
	CurrentContractsMetadata *ContractsMetadata   `json:"currentContractsMetadata"`
	ContractsMetadata        []*ContractsMetadata `json:"contractsMetadata"`
	TeamAssets               *TeamAssets          `json:"teamAssets"`
	TeamLiabilities          *TeamLiabilities     `json:"teamLiabilities"`
	TeamOwners               []string             `json:"teamOwners"`
	AccessCodes              []*string            `json:"accessCodes"`
	TeamScoring              []TeamScoring        `json:"teamScoring"`
}

type ContractsMetadata struct {
	Year               int                    `json:"year"`
	TotalUtilizedCap   int                    `json:"totalUtilizedCap"`
	TotalAvailableCap  int                    `json:"totalAvailableCap"`
	QbUtilizedCap      *CapUtilizationSummary `json:"qbUtilizedCap"`
	RbUtilizedCap      *CapUtilizationSummary `json:"rbUtilizedCap"`
	WrUtilizedCap      *CapUtilizationSummary `json:"wrUtilizedCap"`
	TeUtilizedCap      *CapUtilizationSummary `json:"teUtilizedCap"`
	DeadCapUtilizedCap *CapUtilizationSummary `json:"deadCapUtilizedCap"`
}

type TeamAssets struct {
	DraftPicks []*DraftYear `json:"draftPicks"`
}

type TeamLiabilities struct {
	DeadCap []*DeadCap `json:"deadCap"`
}

type DraftPick struct {
	Round           int     `json:"round"`
	Value           *int    `json:"value"`
	OriginalOwnerID *string `json:"originalOwnerId"`
}

type DraftYear struct {
	Year  int          `json:"year"`
	Picks []*DraftPick `json:"picks"`
}

type DeadCap struct {
	ID                   string        `json:"id"`
	AssociatedContractID *string       `json:"associatedContractId"`
	DeadCapYears         []DeadCapYear `json:"deadCapYears"`
	DeadCapNote          *string       `json:"deadCapNote"`
}

type DeadCapInput struct {
	AssociatedContractID *string            `json:"associatedContractId"`
	DeadCapYears         []DeadCapYearInput `json:"deadCapYears"`
	DeadCapNote          string             `json:"deadCapNote"`
}

type DeadCapYearInput struct {
	Year   int `json:"year"`
	Amount int `json:"amount"`
}

type DeadCapYear struct {
	Year   int `json:"year"`
	Amount int `json:"amount"`
}

type CapUtilizationSummary struct {
	CapUtilization int `json:"capUtilization"`
	NumContracts   int `json:"numContracts"`
}

type TeamScoring struct {
	Year    int                      `json:"year"`
	Summary TeamScoringSeasonSummary `json:"summary"`
	Weeks   []TeamScoringWeek        `json:"weeks"`
}

type TeamScoringSeasonSummary struct {
	Wins               int     `json:"wins"`
	Losses             int     `json:"losses"`
	Ties               int     `json:"ties"`
	CurrentStreak      int     `json:"currentStreak"`
	TotalPointsFor     float64 `json:"totalPointsFor"`
	TotalPointsAgainst float64 `json:"totalPointsAgainst"`
}

type TeamScoringWeek struct {
	Week          int     `json:"week"`
	PointsFor     float64 `json:"pointsFor"`
	PointsAgainst float64 `json:"pointsAgainst"`
}

type NewTeam struct {
	ID          string     `json:"id"`
	TeamName    string     `json:"teamName"`
	Division    *string    `json:"division"`
	FoundedDate *time.Time `json:"foundedDate"`
}
