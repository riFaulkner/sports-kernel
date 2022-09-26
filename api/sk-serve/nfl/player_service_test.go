package playernfl

import (
	"context"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

// Methods implemented below
type mockPlayerRepository struct {
	mock.Mock
}

func (m mockPlayerRepository) Create(ctx context.Context, player model.PlayerNfl) (*model.PlayerNfl, error) {
	args := m.Called(player)

	if args.Get(0) != nil {
		return nil, args.Error(0)
	}

	return &player, args.Error(0)
}

func (m mockPlayerRepository) GetAll(ctx context.Context) ([]*model.PlayerNfl, bool) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}

	return args.Get(0).([]*model.PlayerNfl), args.Bool(1)
}

func (m mockPlayerRepository) GetPlayersByPosition(ctx context.Context, position model.PlayerPosition) ([]*model.PlayerNfl, bool) {
	//TODO implement me
	panic("implement me")
}

func (m mockPlayerRepository) GetPlayersWithLimit(ctx context.Context, numberOfResults int) ([]*model.PlayerNfl, bool) {
	args := m.Called(numberOfResults)

	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}

	return args.Get(0).([]*model.PlayerNfl), args.Bool(1)
}

func (m mockPlayerRepository) GetPlayerById(ctx context.Context, playerId *string) (*model.PlayerNfl, bool) {
	//TODO implement me
	panic("implement me")
}

func TestPlayerService_CreatePlayer(t *testing.T) {
	type fields struct {
		MockSetup func() *mockPlayerRepository
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
		{"Test Basic defaults met",
			fields{func() *mockPlayerRepository {
				mockRepository := new(mockPlayerRepository)
				mockRepository.
					On("Create", mock.Anything).
					Return()
				return mockRepository
			}},
			args{
				context.Background(),
				model.NewPlayerNfl{
					PlayerName: "Test Player",
					Position:   "RB",
					Team:       model.NflTeamFa,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPlayerRepository := tt.fields.MockSetup()
			p := &PlayerService{
				playerRepository: mockPlayerRepository,
			}
			got, err := p.CreatePlayer(tt.args.ctx, tt.args.playerInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePlayer() got = %v, want %v", got, tt.want)
			}
			if !mockPlayerRepository.AssertExpectations(t) {
				t.Errorf("Failed to assert mock expectations")
			}
		})
	}
}

func TestPlayerService_GetAllPlayers(t *testing.T) {
	type fields struct {
		MockSetup func() *mockPlayerRepository
	}
	type args struct {
		ctx             context.Context
		numberOfResults func() *int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.PlayerNfl
		wantErr bool
	}{
		{"Successfully gets empty data",
			fields{func() *mockPlayerRepository {
				mockRepository := new(mockPlayerRepository)
				mockRepository.On("GetAll").Return(make([]*model.PlayerNfl, 0), true)
				return mockRepository
			}},
			args{context.Background(), func() *int {
				return nil
			},
			},
			make([]*model.PlayerNfl, 0),
			false,
		},
		{"Successfully gets empty data, with limited number of players requested",
			fields{func() *mockPlayerRepository {
				mockRepository := new(mockPlayerRepository)
				mockRepository.On("GetPlayersWithLimit", 1).Return(make([]*model.PlayerNfl, 0), true)
				return mockRepository
			}},
			args{context.Background(), func() *int {
				expectedResults := 1
				return &expectedResults
			},
			},
			make([]*model.PlayerNfl, 0),
			false,
		},
		{"Failed to get data",
			fields{func() *mockPlayerRepository {
				mockRepository := new(mockPlayerRepository)
				mockRepository.On("GetAll", mock.Anything).Return(nil, false)
				return mockRepository
			}},
			args{context.Background(), func() *int {
				return nil
			}},
			nil,
			true,
		},
		{"Failed to get data, from filter",
			fields{func() *mockPlayerRepository {
				mockRepository := new(mockPlayerRepository)
				mockRepository.On("GetPlayersWithLimit", 1).Return(nil, false)
				return mockRepository
			}},
			args{context.Background(), func() *int {
				filterResults := 1
				return &filterResults
			}},
			nil,
			true,
		},
		{"Invalid number of results requested, should return error",
			fields{func() *mockPlayerRepository {
				mockRepository := new(mockPlayerRepository)

				return mockRepository
			}},
			args{context.Background(), func() *int {
				expectedResults := 0
				return &expectedResults
			}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := tt.fields.MockSetup()
			p := &PlayerService{
				playerRepository: mockRepository,
			}
			got, err := p.GetAllPlayers(tt.args.ctx, tt.args.numberOfResults())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPlayers() got = %v, want %v", got, tt.want)
			}
			if !mockRepository.AssertExpectations(t) {
				t.Errorf("Failed to assert expectations!")
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
				playerRepository: tt.fields.PlayerRepository,
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
