// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package utils

import "utils/security"

//Cookie编码
func CookieEncode(src string) string {
	//aes编码
	s, _ := security.AesEncode([]byte(src))
	//base64编码
	return security.Base64Encode(s)
}

//Cookie解码
func CookieDecode(src string) string {
	//base64解码
	b := security.Base64Decode(src)
	//aes解码
	s, _ := security.AesDecode([]byte(b))

	return string(s)
}
