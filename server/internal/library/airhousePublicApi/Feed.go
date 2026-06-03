package airhousePublicApi

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"net/http"
	"time"
)

var (
	Feed        = "/feeds"
	CommitFeeds = "/commit_feeds"
)

// GetFeed 主动拉取更新信息
func GetFeed(ctx context.Context) (FeedResponseData FeedJSONData, err error) {
	var (
		httpStatus  int
		responseStr string
	)
	if responseStr, err, httpStatus = CurlAirHostPublicApiGet(ctx, Feed, nil); err != nil {
		return
	}
	if httpStatus != http.StatusOK {
		err = gerror.New(responseStr)
		return
	}
	if err = json.Unmarshal([]byte(responseStr), &FeedResponseData); err != nil {
		return
	}
	return
}

// CommitFeed 消费提交拉取的消息 ID
func CommitFeed(ctx context.Context, FeedId string) (err error) {
	var (
		httpStatus       int
		commitFeedParams CommitFeedReq
	)
	commitFeedParams.FeedId = FeedId
	if _, err, httpStatus = CurlAirHostPublicApiPost(ctx, CommitFeeds, commitFeedParams); err != nil {
		return
	}
	if httpStatus != http.StatusNoContent {
		err = gerror.New("消费失败")
		return
	}
	return
}

type CommitFeedReq struct {
	FeedId string `json:"feed_id"`
}

type FeedJSONData struct {
	ObjectType string     `json:"object_type"`
	Data       []FeedData `json:"data"`
}

type FeedData struct {
	ID         string      `json:"id"`
	ObjectType string      `json:"object_type"`
	Event      string      `json:"event"`
	Property   Property    `json:"property"`
	Body       interface{} `json:"body"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
