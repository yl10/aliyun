package aliyun

import (
	"github.com/yl10/aliyun/sts"
)

type AliClient struct {
	STS *sts.StsClient
}

func NewAliClient(serv, region, endpoint, accID, accSec string) (*AliClient, error) {
	sts, err := sts.NewStsClient(serv, region, endpoint, accID, accSec)
	if err != nil {
		return nil, err
	}

	cli := &AliClient{
		STS: sts,
	}

	return cli, nil
}
