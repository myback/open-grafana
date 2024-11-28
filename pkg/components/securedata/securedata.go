package securedata

import (
	"github.com/myback/open-grafana/pkg/setting"
	"github.com/myback/open-grafana/pkg/util"
)

type SecureData []byte

func Encrypt(data []byte) (SecureData, error) {
	return util.Encrypt(data, setting.SecretKey)
}

func (s SecureData) Decrypt() ([]byte, error) {
	return util.Decrypt(s, setting.SecretKey)
}
