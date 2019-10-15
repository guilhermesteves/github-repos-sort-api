package transform

import (
	"github.com/guilhermesteves/github-repos-sort-api/internal/model"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ReposToBson(repos []*model.Repo) primitive.A {
	arr := primitive.A{}

	for _, r := range repos {
		arr = append(arr, RepoToBson(r))
	}

	return arr
}

func RepoToBson(repo *model.Repo) primitive.M {
	gr := repo.GithubRepo
	githubRepo := primitive.M{
		"id":				cast.ToInt64(gr.Id),
		"fullName":			gr.FullName,
		"name":				gr.Name,
		"description":		gr.Description,
		"homepage":			gr.Homepage,
		"defaultBranch":	gr.DefaultBranch,
		"updatedAt": 		utils.CheckIf(gr.PushedAt != "", cast.ToTime(gr.PushedAt), time.Now()),
		"cloneUrl":			gr.CloneUrl,
		"htmlUrl":			gr.HtmlUrl,
		"forks":			cast.ToInt64(gr.Forks),
		"stars":			cast.ToInt64(gr.Stars),
	}

	return primitive.M{
		"_id":        	repo.Id,
		"githubRepo":	githubRepo,
		"likes":		cast.ToInt64(repo.Likes),
		"likedBy":		utils.CheckIf(repo.LikedBy != nil, utils.ToStringSlice(repo.LikedBy), []string{}),
		"createdAt": 	utils.CheckIf(repo.CreatedAt != "", cast.ToTime(repo.CreatedAt), time.Now()),
		"updatedAt": 	utils.CheckIf(repo.UpdatedAt != "", cast.ToTime(repo.UpdatedAt), time.Now()),
	}
}