package bc

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/wechat/pkg/bc"
	"github.com/twiglab/crm/wechat/pkg/bc/low"
)

type PointsSync struct {
	TransactionID   string
	OpenID          string
	IncreasedPoints int
}

type WxBcCli struct {
	Client graphql.Client
}

func (c *WxBcCli) Sync(ctx context.Context, pointsSync PointsSync) error {
	_, err := low.BcPointSync(ctx, c.Client, bc.BusinessCirclePointsSync{
		TransactionID:    pointsSync.TransactionID,
		OpenID:           pointsSync.OpenID,
		EarnPoints:       true,
		IncreasedPoints:  pointsSync.IncreasedPoints,
		PointsUpdateTime: time.Now(),
	})
	return err
}
