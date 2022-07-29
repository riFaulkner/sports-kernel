package db

import (
	"context"
	"encoding/json"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/contract"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/team"

	"reflect"
	"testing"
)

func TestTeamImpl_AddDeadCapToTeam(t *testing.T) {
	type fields struct {
		Client firestore.Client
	}
	type args struct {
		ctx      context.Context
		leagueID string
		teamID   string
		deadCap  []*team.DeadCap
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &TeamRepositoryImpl{
				Client: tt.fields.Client,
			}
			got := u.AddDeadCapToTeam(tt.args.ctx, tt.args.leagueID, tt.args.teamID, tt.args.deadCap)

			if got != tt.want {
				t.Errorf("AddDeadCapToTeam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeamImpl_Create(t *testing.T) {
	type fields struct {
		Client firestore.Client
	}
	type args struct {
		ctx       context.Context
		leagueId  string
		teamInput team.NewTeam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *team.Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &TeamRepositoryImpl{
				Client: tt.fields.Client,
			}
			got, err := u.Create(tt.args.ctx, tt.args.leagueId, tt.args.teamInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeamImpl_GetAllLeagueTeams(t *testing.T) {
	type fields struct {
		Client firestore.Client
	}
	type args struct {
		ctx      context.Context
		leagueId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*team.Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &TeamRepositoryImpl{
				Client: tt.fields.Client,
			}
			got, err := u.GetAllLeagueTeams(tt.args.ctx, tt.args.leagueId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllLeagueTeams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllLeagueTeams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeamImpl_GetTeamById(t *testing.T) {
	type fields struct {
		Client firestore.Client
	}
	type args struct {
		ctx      context.Context
		leagueId string
		teamId   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *team.Team
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &TeamRepositoryImpl{
				Client: tt.fields.Client,
			}
			got, err := u.GetTeamById(tt.args.ctx, tt.args.leagueId, tt.args.teamId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTeamById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTeamById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeamImpl_UpdateTeamContractMetaData(t *testing.T) {
	type fields struct {
		Client firestore.Client
	}
	type args struct {
		ctx           context.Context
		leagueID      string
		teamContracts []*contract.Contract
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test",
			args: args{
				ctx:           context.Background(),
				leagueID:      "TEST_LEAGUE_ID",
				teamContracts: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &TeamRepositoryImpl{
				Client: tt.fields.Client,
			}
			if err := u.UpdateTeamContractMetaData(tt.args.ctx, tt.args.leagueID, tt.args.teamContracts); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTeamContractMetaData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generateDefaultTeamContractsMetadata(t *testing.T) {
	tests := []struct {
		name string
		want *team.ContractsMetadata
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateDefaultTeamContractsMetadata(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateDefaultTeamContractsMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateTeamAssets(t *testing.T) {
	type args struct {
		teamID string
	}
	goodTeamID := "abc"

	tests := []struct {
		name string
		args args
		want *team.TeamAssets
	}{
		{
			name: "team ID gets added to result",
			args: args{teamID: goodTeamID},
			want: &team.TeamAssets{
				DraftPicks: []*team.DraftYear{
					{
						Year: 2022,
						Picks: []*team.DraftPick{
							{Round: 1, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 2, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 3, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 4, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 5, Value: nil, OriginalOwnerID: &goodTeamID},
						},
					},
					{
						Year: 2023,
						Picks: []*team.DraftPick{
							{Round: 1, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 2, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 3, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 4, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 5, Value: nil, OriginalOwnerID: &goodTeamID},
						},
					},
					{
						Year: 2024,
						Picks: []*team.DraftPick{
							{Round: 1, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 2, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 3, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 4, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 5, Value: nil, OriginalOwnerID: &goodTeamID},
						},
					},
					{
						Year: 2025,
						Picks: []*team.DraftPick{
							{Round: 1, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 2, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 3, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 4, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 5, Value: nil, OriginalOwnerID: &goodTeamID},
						},
					},
					{
						Year: 2026,
						Picks: []*team.DraftPick{
							{Round: 1, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 2, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 3, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 4, Value: nil, OriginalOwnerID: &goodTeamID},
							{Round: 5, Value: nil, OriginalOwnerID: &goodTeamID},
						},
					},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonWant, _ := json.Marshal(tt.want)
			if jsonGot, _ := json.Marshal(generateTeamAssets(tt.args.teamID)); !reflect.DeepEqual(string(jsonGot), string(jsonWant)) {
				t.Errorf("generateTeamAssets() =\n %v, want \n %v", string(jsonGot), string(jsonWant))
			}
		})
	}
}
