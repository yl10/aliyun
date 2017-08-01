package sts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/yl10/kit/random"
	"github.com/yl10/kit/request"
)

// AssumeRole sts 接口
// https://help.aliyun.com/document_detail/28763.html?spm=5176.doc28761.6.680.DeKZzx
func (s *StsClient) AssumeRole(roleArn, roleSessionName, policy string) (*AssumeRole, error) {
	s.SetAction(ACTION_ASSUMEROLE).
		SetSignatureNonce(random.GenRandAphla(16)).
		SetTimestamp(time.Now().UTC().Format("2006-01-02T15:04:05Z")).
		SetRoleArn(roleArn).
		SetRoleSessionName(roleSessionName).
		SetPolicy(policy)

	var err error
	if err = s.genSignature(ACTION_ASSUMEROLE); err != nil {
		return nil, err
	}

	query, err := s.getQueryString(ACTION_ASSUMEROLE)
	if err != nil {
		return nil, err
	}

	url := s.StsServer + "?" + query
	resp, err := request.Request(url, true)
	if err != nil {
		return nil, err
	}

	var assumeRole AssumeRole
	err = json.Unmarshal(resp, &assumeRole)
	if err != nil {
		return nil, err
	}

	if assumeRole.Code != "" {
		return nil, fmt.Errorf("[%v] %v", assumeRole.Code, assumeRole.Message)
	}

	assumeRole.Bucket = s.Bucket
	assumeRole.Region = s.Region
	assumeRole.Endpoint = s.Endpoint

	return &assumeRole, nil
}
