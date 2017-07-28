package sts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/yl10/kit/random"
	"github.com/yl10/kit/request"
)

// GetCallerIdentity sts 接口
// https://help.aliyun.com/document_detail/43767.html?spm=5176.doc28763.6.681.itUL25
func (s *StsClient) GetCallerIdentity() (*CallerIdentity, error) {
	s.SetAction(ACTION_GETCALLERIDENTITY).
		SetSignatureNonce(random.GenRandAphla(16)).
		SetTimestamp(time.Now().UTC().Format("2006-01-02T15:04:05Z"))

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

	var callerIdentity CallerIdentity
	err = json.Unmarshal(resp, &callerIdentity)
	if err != nil {
		return nil, err
	}

	if callerIdentity.Code != "" {
		return nil, fmt.Errorf("[%v] %v", callerIdentity.Code, callerIdentity.Message)
	}

	return &callerIdentity, nil
}
