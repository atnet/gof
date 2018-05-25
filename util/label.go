/**
 * Copyright 2015 @ z3q.net.
 * name : label
 * author : jarryliu
 * date : 2015-07-27 09:28
 * description :
 * history :
 */
package util

import "regexp"

var reg = regexp.MustCompile("\\{([^\\}]+)\\}")

// 翻译标签
func Transplate(c string, m map[string]interface{}) string {
	return reg.ReplaceAllStringFunc(c, func(k string) string {
		key := k[1 : len(k)-1]
		if v, ok := m[key]; ok {
			return Str(v)
		}
		return k
	})
}
