package social

import (
	"context"
	"fmt"
	"mm-ndj/config"
	"mm-ndj/pkg/xhttp"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/user/userlookup"
	Utype "github.com/michimani/gotwi/user/userlookup/types"
	"go.uber.org/zap"
)

// const ( 此配置已经放入配置文件中
// 	tweetsEndpoint  = "https://api.twitter.com/2/tweets/"
// 	usersEndpoint   = "https://api.twitter.com/2/users/"
// 	TaskRetweetedBy = "/retweeted_by"
// 	TaskLikingUsers = "/liking_users"
// 	TaskFollowers   = "/followers"
// )

func TwitterClientInit() (*gotwi.Client, error) {
	in := &gotwi.NewClientWithAccessTokenInput{
		AccessToken: config.Conf.TwitterC.AccessKey,
	}
	return gotwi.NewClientWithAccessToken(in)

}

func GetTwitterTaskUserList(taskSuffix, twitterId, nextToken string) (response *RetweetRes, err error) {
	endpoint := config.Conf.TwitterC.TweetsEnpoint
	if taskSuffix == config.Conf.TwitterC.TaskFollowers {
		endpoint = config.Conf.TwitterC.UserEndpoints
	}
	client := xhttp.NewDefaultClient()
	header := map[string]string{
		"Content-Type":  "application/json;charset=UTF-8",
		"Authorization": fmt.Sprintf("Bearer %s", config.Conf.TwitterC.AccessKey),
	}
	req, err := client.GetRequest(xhttp.MethodGet, endpoint+twitterId+taskSuffix, header, nil)
	if err != nil {
		return
	}
	if nextToken != "" {
		params := req.URL.Query()
		params.Add("pagination_token", nextToken)
		req.URL.RawQuery = params.Encode()
	}
	_, err = client.CallWithRequest(req, &response)
	if err != nil {
		return
	}
	return
}

func GetTwitterIdByName(userName string) string {
	twitterClient, err := TwitterClientInit()
	if err != nil {
		return ""
	}
	userInput := &Utype.GetByUsernameInput{
		Username: userName,
	}
	config.Logger.Info("", zap.Any("userInput", userInput))

	//二 拿到推特name后到推特API找到这个人
	userOutput, err := userlookup.GetByUsername(context.Background(), twitterClient, userInput)
	if err != nil {
		config.Logger.Error("userlookup.GetByUsername", zap.Error(err))
		return ""
	}
	//三 拿项目方的推特id 封装一下 作为下一步的参数
	projectTwitterId := userOutput.Data.ID
	return *projectTwitterId

}
