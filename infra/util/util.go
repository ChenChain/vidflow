package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func LoadPrivateKeyFromString(keyString string) (*rsa.PrivateKey, error) {
	// 解码 PEM 格式的私钥字符串
	block, _ := pem.Decode([]byte(keyString))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block")
	}

	// 解析 PKCS8 格式的私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// 类型断言为 *rsa.PrivateKey
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return rsaPrivateKey, nil
}

func GenerateJWT(privateKey *rsa.PrivateKey, claims map[string]interface{}) (string, error) {
	// 定义 JWT Payload
	vv := jwt.MapClaims{}
	for k, v := range claims {
		vv[k] = v
	}

	// 创建 Token 结构体
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, vv)

	// 使用私钥对 Token 进行签名
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return signedToken, nil
}
