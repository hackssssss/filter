package util

import (
	"github.com/pkg/errors"
	"strings"

)

type filter struct {
	data        string
	maxLimitLen int
}

func NewKeywordFilter(str string, maxLimitLen int) (*filter, error) {
	if len(str) > maxLimitLen {
		return nil, errors.Errorf("长度：%d，不能超过：%d", len(str), maxLimitLen)
	}
	return &filter{
		data:        str,
		maxLimitLen: maxLimitLen,
	}, nil
}

func (f *filter) GetData() string {
	return f.data
}

func (f *filter) FilterKeywords(keywords map[string]bool) (err error) {
	if keywords == nil {
		return
	}
	for i := 0; i < len(f.data); i++ {
		for j := i + 1; j <= len(f.data); j++ {
			subStr := f.data[i:j]
			if _, found := keywords[subStr]; found {
				err = errors.Errorf("昵称违规，建议修改")
			}
		}
	}
	return
}

func (f *filter) TrimAllCharset(ch []string) (err error) {
	if ch == nil {
		return
	}
	for _, c := range ch {
		f.data = strings.Replace(f.data, c, "", -1)
	}
	if len(f.data) == 0 {
		err = errors.New("剔除相关转移字符后，数据长度为0.")
		return
	}
	return
}
