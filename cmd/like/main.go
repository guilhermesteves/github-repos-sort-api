package like

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guilhermesteves/github-repos-sort-api/internal/api"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"strings"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ip := strings.Split(request.Headers["X-Forwarded-For"], ",")[0]
	repoId := request.PathParameters["id"]

	body, err := api.LikeRepo(ip, repoId)

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