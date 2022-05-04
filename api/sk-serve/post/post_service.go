package post

import (
	"context"

	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type LeaguePost interface {
	GetAll(ctx context.Context, leagueId string, numberOfResults *int) ([]*model.LeaguePost, error)
	GetPostById(ctx context.Context, leagueId *string, postId *string) (*model.LeaguePost, error)
	Create(ctx context.Context, leagueId string, post model.NewLeaguePost) (*model.LeaguePost, error)
	AddComment(ctx context.Context, leagueId string, postId string, comment model.NewPostComment) (*model.PostComment, error)
}
