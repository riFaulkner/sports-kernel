package playernfl

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"reflect"
	"testing"
)

func TestPlayerService_CreatePlayer(t *testing.T) {
	type fields struct {
		PlayerRepository PlayerRepository
	}
	type args struct {
		ctx         context.Context
		playerInput model.NewPlayerNfl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PlayerNfl
		wantErr bool
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlayerService{
				PlayerRepository: tt.fields.PlayerRepository,
			}
			got, err := p.CreatePlayer(tt.args.ctx, tt.args.playerInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePlayer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerService_GetAllPlayers(t *testing.T) {
	type fields struct {
		PlayerRepository PlayerRepository
	}
	type args struct {
		ctx             context.Context
		numberOfResults *int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.PlayerNfl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlayerService{
				PlayerRepository: tt.fields.PlayerRepository,
			}
			got, err := p.GetAllPlayers(tt.args.ctx, tt.args.numberOfResults)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPlayers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerService_GetPlayerById(t *testing.T) {
	type fields struct {
		PlayerRepository PlayerRepository
	}
	type args struct {
		ctx      context.Context
		playerId *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PlayerNfl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlayerService{
				PlayerRepository: tt.fields.PlayerRepository,
			}
			got, err := p.GetPlayerById(tt.args.ctx, tt.args.playerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertNewPlayerInputToPlayer(t *testing.T) {
	type args struct {
		newPlayerInput model.NewPlayerNfl
	}
	tests := []struct {
		name string
		args args
		want model.PlayerNfl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertNewPlayerInputToPlayer(tt.args.newPlayerInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertNewPlayerInputToPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePlayerId(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePlayerId(tt.args.name); got != tt.want {
				t.Errorf("generatePlayerId() = %v, want %v", got, tt.want)
			}
		})
	}
}
