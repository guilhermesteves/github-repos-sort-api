package repos

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guilhermesteves/github-repos-sort-api/internal/api"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"github.com/spf13/cast"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	language := request.QueryStringParameters["lang"]
	page := cast.ToInt(request.QueryStringParameters["page"])

	body, err := api.ListRepos(language, page)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:	utils.ParseError(err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}