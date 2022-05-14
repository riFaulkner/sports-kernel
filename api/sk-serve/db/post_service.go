package db

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"time"

	firestore "cloud.google.com/go/firestore"
	appFirestore "github.com/rifaulkner/sports-kernel/api/sk-serve/firestore"
	"github.com/rifaulkner/sports-kernel/api/sk-serve/graph/model"
)

type PostImpl struct {
	Client appFirestore.Client
}

func (p *PostImpl) GetAll(ctx context.Context, leagueId string, numberOfResults *int) ([]*model.LeaguePost, error) {

	postref := p.Client.Collection(leagueCollection).Doc(leagueId).Collection("posts")

	posts := make([]*model.LeaguePost, 0)

	var results []*firestore.DocumentSnapshot
	var err error

	if *numberOfResults < -1 {
		return nil, err
	}

	if *numberOfResults == -1 {
		//-1 to return all players, no query limit
		results, err = postref.OrderBy("PostDate", firestore.Asc).Documents(ctx).GetAll()
	} else {
		//else return the desired number of players
		results, err = postref.OrderBy("PostDate", firestore.Asc).Limit(*numberOfResults).Documents(ctx).GetAll()
	}

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		post := new(model.LeaguePost)
		err = result.DataTo(&post)
		post.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *PostImpl) GetPostById(ctx context.Context, leagueId *string, postId *string) (*model.LeaguePost, error) {
	result, err := p.Client.Collection(*leagueId).Doc(*postId).Get(ctx)
	if err != nil {
		return nil, err
	}
	post := new(model.LeaguePost)
	err = result.DataTo(&post)
	id := result.Ref.ID
	post.ID = id
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostImpl) Create(ctx context.Context, leagueId string, inputPost model.NewLeaguePost) (*model.LeaguePost, error) {
	posts := p.Client.Collection(leagueCollection).Doc(leagueId).Collection("posts")

	titlehash := hashTitle(inputPost.Title)

	newPost := model.LeaguePost{
		ID:       titlehash,
		Author:   inputPost.Author,
		PostDate: time.Now(),
		Content:  inputPost.Content,
		Title:    inputPost.Title,
		Comments: make([]*model.PostComment, 0),
	}

	_, err := posts.Doc(newPost.ID).Set(ctx, newPost)
	if err != nil {
		return nil, err
	}

	return &newPost, nil
}

func (p *PostImpl) AddComment(ctx context.Context, leagueId string, postId string, inputComment model.NewPostComment) (*model.PostComment, error) {
	post := p.Client.Collection(collectionLeague).Doc(leagueId).Collection("posts").Doc(postId)

	newComment := model.PostComment{
		Author:      inputComment.Author,
		Content:     inputComment.Content,
		CommentDate: time.Now(),
		ID:          hashTitle(inputComment.Content),
	}

	_, err := post.Collection("comments").Doc(newComment.ID).Set(ctx, newComment)

	if err != nil {
		return nil, err
	}

	return &newComment, nil
}

func (p *PostImpl) GetComments(ctx context.Context, leagueId string, postId string) ([]*model.PostComment, error) {
	post := p.Client.Collection(collectionLeague).Doc(leagueId).Collection("posts").Doc(postId)

	comments := make([]*model.PostComment, 0)

	results, err := post.Collection("comments").Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		comment := new(model.PostComment)
		err = result.DataTo(&comment)
		comment.ID = result.Ref.ID
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func hashTitle(name string) string {
	hashString := []byte(name)
	md5string := md5.Sum(hashString)
	b64String := base64.RawURLEncoding.EncodeToString(md5string[:])
	return b64String
}
