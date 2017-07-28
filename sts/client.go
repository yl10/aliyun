package sts

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/yl10/kit/encrypt"
)

type StsClient struct {
	StsServer string
	Region    string
	Endpoint  string
	Bucket    string

	Format           string `sts:"Format"`
	Version          string `sts:"Version"`
	AccessKeyID      string `sts:"AccessKeyId"`
	AccessKeySecret  string
	Signature        string
	SignatureMethod  string `sts:"SignatureMethod"`
	SignatureVersion string `sts:"SignatureVersion"`
	SignatureNonce   string `sts:"SignatureNonce"`
	Timestamp        string `sts:"Timestamp"`

	Action          string `AssumeRole:"Action"`
	RoleArn         string `AssumeRole:"RoleArn"`
	RoleSessionName string `AssumeRole:"RoleSessionName"`
	Policy          string `AssumeRole:"Policy"`
	DurationSeconds int    `AssumeRole:"DurationSeconds"`
}

func (s *StsClient) SetRegion(v string) *StsClient {
	s.Region = v
	return s
}

func (s *StsClient) SetEndpoint(v string) *StsClient {
	s.Endpoint = v
	return s
}

func (s *StsClient) SetBucket(v string) *StsClient {
	s.Bucket = v
	return s
}

func (s *StsClient) SetFormat(v string) *StsClient {
	s.Format = v
	return s
}

func (s *StsClient) SetVersion(v string) *StsClient {
	s.Version = v
	return s
}

func (s *StsClient) SetAccessKeySecret(v string) *StsClient {
	s.AccessKeySecret = v
	return s
}

func (s *StsClient) SetAccessKeyID(v string) *StsClient {
	s.AccessKeyID = v
	return s
}

func (s *StsClient) SetSignature(v string) *StsClient {
	s.Signature = v
	return s
}

func (s *StsClient) SetSignatureMethod(v string) *StsClient {
	s.SignatureMethod = v
	return s
}

func (s *StsClient) SetSignatureVersion(v string) *StsClient {
	s.SignatureVersion = v
	return s
}

func (s *StsClient) SetSignatureNonce(v string) *StsClient {
	s.SignatureNonce = v
	return s
}

func (s *StsClient) SetTimestamp(v string) *StsClient {
	s.Timestamp = v
	return s
}

func (s *StsClient) SetAction(v string) *StsClient {
	s.Action = v
	return s
}

func (s *StsClient) SetRoleArn(v string) *StsClient {
	s.RoleArn = v
	return s
}

func (s *StsClient) SetRoleSessionName(v string) *StsClient {
	s.RoleSessionName = v
	return s
}

func (s *StsClient) SetPolicy(v string) *StsClient {
	s.Policy = v
	return s
}

func (s *StsClient) SetDurationSeconds(v int) *StsClient {
	s.DurationSeconds = v
	return s
}

func (s *StsClient) genSignature(act string, httpMethods ...string) error {
	params, err := s.getParamMap(act)
	if err != nil {
		return err
	}

	var httpMethod = "GET"
	if len(httpMethods) > 0 && httpMethods[0] != "" {
		httpMethod = httpMethods[0]
	}

	// map 排序并转码
	canonicalizedQueryString := paramsSortAndEncodeToString(params)
	stringToSign := httpMethod + "&%2F&" + percentEncode(canonicalizedQueryString)
	signature := encrypt.Base64Encode(encrypt.HmacSha1(stringToSign, s.AccessKeySecret+"&"))
	// signature = url.QueryEscape(signature)
	s.SetSignature(signature)
	return nil
}

func (s *StsClient) getParamMap(act string) (map[string]string, error) {
	m := make(map[string]string)

	t := reflect.TypeOf(*s)
	v := reflect.ValueOf(*s)

	for i := 0; i < t.NumField(); i++ {
		pubTag := t.Field(i).Tag.Get("sts")
		actTag := t.Field(i).Tag.Get(act)
		val := v.Field(i).Interface()

		var tag string
		if pubTag != "" {
			tag = pubTag
		}

		if actTag != "" {
			tag = actTag
		}

		if tag == "" {
			continue
		}

		var tmpVal string
		switch tv := val.(type) {
		default:
			return m, fmt.Errorf("转换 model 失败, 不支持的类型 %T\n", tv)
		case int:
			tmpVal = strconv.Itoa(tv)
		case string:
			tmpVal = tv
		}

		// 如果值为空, 则跳过
		if tmpVal == "" {
			continue
		}

		m[tag] = tmpVal
	}

	return m, nil
}

func (s *StsClient) getQueryString(act string) (string, error) {
	if s.Signature == "" {
		return "", errors.New("Signature 值为空, 请先调用 genSignature ")
	}

	params, err := s.getParamMap(act)
	if err != nil {
		return "", err
	}
	params["Signature"] = s.Signature

	query := make([]string, 0)
	for k, v := range params {
		str := url.QueryEscape(k) + "=" + url.QueryEscape(v)
		query = append(query, str)
	}

	return strings.Join(query, "&"), nil
}
