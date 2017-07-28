package sts

const (
	ACTION_ASSUMEROLE        = "AssumeRole"
	ACTION_GETCALLERIDENTITY = "GetCallerIdentity"
)

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      string
}

type AssumedRoleUser struct {
	Arn               string `json:"arn"`
	AssumedRoleUserId string
}

type RespError struct {
	RequestId string
	HostId    string
	Code      string
	Message   string
}

// 这里冗余了 错误、Region、Bucket、Endpoint
type AssumeRole struct {
	RespError
	Credentials     `json:"Credentials"`
	AssumedRoleUser `json:"AssumedRoleUser"`
	RequestId       string
	Region          string
	Bucket          string
	Endpoint        string
}

type CallerIdentity struct {
	RespError
	AccountId string
	UserId    string
	Arn       string
	RequestId string
}

// NewStsClient 初始化一个 sts 客户端
func NewStsClient(serv, region, endpoint, accID, accSec string) (*StsClient, error) {
	client := &StsClient{
		StsServer:        serv,
		Region:           region,
		Endpoint:         endpoint,
		AccessKeyID:      accID,
		AccessKeySecret:  accSec,
		Format:           "json",
		Version:          "2015-04-01",
		SignatureMethod:  "HMAC-SHA1",
		SignatureVersion: "1.0",
		DurationSeconds:  3600,
	}

	return client, nil
}
