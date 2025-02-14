package crypto

// 私钥解密数据和签名(防止别人冒充数据), 公钥加密.根据私钥可以推导出公钥

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// MD5
func Md5(b []byte) string {
	c := md5.New()
	c.Write(b)
	return hex.EncodeToString(c.Sum(nil))
}
func Sha1(b []byte) string {
	c := sha1.New()
	c.Write(b)
	return hex.EncodeToString(c.Sum(nil))
}

// HmacSha256
func HmacSha256(data, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// 使用私钥Sha1WithRSA签名
func Sha1WithRSA(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	h := sha1.New()
	h.Write(data)
	digest := h.Sum(nil)

	bytes, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
