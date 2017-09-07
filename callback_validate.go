package weixin

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

//微信回调后台时验证接口有效性的方法
func ValidateSignature(token, timestamp, nonce, signature string) bool {
	strSlices := []string{token, timestamp, nonce}
	sort.Strings(strSlices)

	str := strings.Join(strSlices, "")
	sigBytes := sha1.Sum([]byte(str))

	return fmt.Sprintf("%x", sigBytes) == signature
}
