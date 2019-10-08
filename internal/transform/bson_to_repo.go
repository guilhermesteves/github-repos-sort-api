package transform

import (
	"github.com/guilhermesteves/github-repos-sort-api/internal/model"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BsonToRepos(src primitive.A) []*model.Repo {
	repos := []*model.Repo{}

	if src != nil {
		for _, d := range src {
			repos = append(repos, BsonToRepo(d.(primitive.M)))
		}
	}

	return repos
}

func BsonToRepo(src primitive.M) *model.Repo {
	repo := model.Repo{
		Id: 		cast.ToString(src["_id"]),
		GithubRepo: BsonToGithubRepo(utils.CheckIf(src["githubRepo"] != nil, src["githubRepo"], primitive.M{}).(primitive.M)),
		Likes: 		cast.ToInt64(src["likes"]),
		LikedBy: 	utils.ToStringSlice(utils.CheckIf(src["likedBy"] != nil, src["likedBy"], primitive.A{}).(primitive.A)),
		CreatedAt:	utils.ToFormattedTime(src["createdAt"]),
		UpdatedAt:  utils.ToFormattedTime(src["updatedAt"]),
	}

	return &repo
}

func BsonToGithubRepo(src primitive.M) *model.GithubRepo {
	githubRepo := model.GithubRepo{
		Id: 			cast.ToInt64(src["id"]),
		FullName:		cast.ToString(src["fullName"]),
		Name:			cast.ToString(src["name"]),
		Description:	cast.ToString(src["description"]),
		Homepage:		cast.ToString(src["homepage"]),
		DefaultBranch:	cast.ToString(src["defaultBranch"]),
		PushedAt:  		utils.ToFormattedTime(src["pushedAt"]),
		CloneUrl:		cast.ToString(src["cloneUrl"]),
		HtmlUrl:		cast.ToString(src["htmlUrl"]),
		Forks: 			cast.ToInt64(src["forks"]),
		Stars: 			cast.ToInt64(src["stars"]),
	}

	return &githubRepo
}