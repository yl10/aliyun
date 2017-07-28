package sts

import (
	"bytes"
	"net/url"
	"sort"
	"strings"
)

// 按照字母顺序重排map并进行加密
func paramsSortAndEncodeToString(pm map[string]string) string {
	// 获取map的key排序
	keys := make([]string, 0)
	for k := range pm {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		vs := pm[k]
		prefix := percentEncode(k) + "="
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		buf.WriteString(percentEncode(vs))
	}
	return buf.String()
}

// url 加密
func percentEncode(s string) string {
	newStr := url.QueryEscape(s)
	return strings.Replace(newStr, "+", "%20", -1)
}
