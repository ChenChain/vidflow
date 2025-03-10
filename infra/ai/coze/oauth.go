package coze

import (
	"time"
	"vidflow/infra/util"
)

// pub e0QEnUBk2f3ySWH57o8c48GljV6xITmQxOzulEpodFs
var key = `
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCoTSLDt/B2vHiv
R+AooEOKPb9NrkVdpASMDF/kIHNC2IZR7qjY/zTIdDEGw8r7NQz6iRQDJXysfdZF
iI2w3YvchbxWxrpSIb6ON1CRSX3yXx3phsT0c7b1znfNcXmbS+Ch9ygYxpzb+OHm
cNX7/ZV4D80rHjDXUOMw06Ct8MGgdWjwlAJtNRm55yd6dVMtE+IroisrId1j0k+0
IDtDxxkY5QpkZrLOK6nCr6H3lNtrHyfFNWyo4Vf8bz85sPKfgDQ75eR3weWiyI3+
l9VO/2kpBZ1TSat6MyvgbsFepjt3JcvAYriAmSIeN/jT26F4hY6SiHCaJtdItc9f
B1ALmqAVAgMBAAECggEADCO4Ng7T+PSM7fO3bvwUXb8ox9p3UBL5wvz3g7XK1TpX
V0c+2vFu5hCUzWHyXQlNZ4WGBy4SiUHwYoaVgHmSfN8RSom+EcOJAlVqFHi3wsjq
mELA0nMa8i2o2uLsUcUjgvtCuKVxgLUPMkyBmtIGkH5EnxoSjXgHqH4d+73pxtWN
sTy6tgpT5EPnRdWbmXQS9BYJgiaAG4hU5wzFgR2dJXnuVGIUwey6YlBszM6LTagT
PGC2hTHxQMIUtlsP2GODD1OrxxSDA6rf4J2qZq6CI/sFydoKcGkB3uamlZ9kYmjT
LTioIyvVM/ncDpekjw1T7137/PTizLErM6SOJsBjgQKBgQDq0aOMNVcaM9YULymV
xZhhLruiG3u6W83AxpnIClgNNkt5cwf+3yBzmbkdKTfNsLnwpMnOyJ159pZPFyF+
VjCGLLplnKzr6KycR4qOj4Hw1nKt1DDc9b+4EMgOaA/ctgj6qrIxHfR2RL51Kp1L
HW6ztWD0QEsS7hkCUuY8H0aowQKBgQC3e36ag7x/dcJSpIXVch1QsOKofzwo0vnp
TQ3oMkUW0dEO/WNe4geA1mn2T8Q0OwLd1iZEhCGd9KBGf2mUd1d9dG1n1dcdxDlg
AuloH/RHHT/deooAipB7eqK132/2/GapWbejzIk7An5lPyJG3a5q5+Gdkb6JVrld
TwF9T3KYVQKBgCexRBoBMjYFqRxEVJ0yh211/tWKG8IGnqMFbx03Umb5VIy0+xcE
FHI1++qH+xzT0LpywIIpuyTQn9vCpzC51P35NZDTiQ9fhz3rgepTK73QIhZsFc5j
5AJzI737rSK7yozEqdZPn/EV8bPQfkIiOYI7GKZw77/Fa4jPBogHKWSBAoGAY6M/
hB8HzNug9An914RJoRj9bOzzOWQgbG34oGA0HolAEvjM7qil1HQbRLPaY1asXtU9
ILX0H3fJVZ85MXOBYlJIWzvHvpVhZt8N6wp3N+sNVHOH33Vfsn5NP6Cfh6tXAJ4E
2IFpAE+BXe/j6EHXxpw3a77KjoA21xHhSDBNlaUCgYEA15qWzMbPL4EUEjirv4/c
aWxmcTnU5nqmY/ZmGN+0L5CGxtJ/azC3HMLjAzoK8nYED5JW2lapIMNsXeMtgW7N
2lEi6whOnpS2xe4i+qfon27NDJCdDMJ0sJy5J1B4fvmgYsaAp2pTKkO3vRUU2sNr
/Iu5S5xPqPDqwbUfw22sGq4=
-----END PRIVATE KEY-----
`

func GenJwt(issId, sessionName string) (string, int64, error) {
	privateKey, err := util.LoadPrivateKeyFromString(key)
	if err != nil {
		return "", 0, err
	}
	now := time.Now()
	exp := now.Add(1 * time.Hour).Unix()
	// 生成 JWT
	token, err := util.GenerateJWT(privateKey, map[string]interface{}{
		"iss":          issId,
		"aud":          "api.coze.cn",
		"exp":          exp,
		"iat":          now.Unix(),
		"jti":          "jti-aaa",
		"session_name": sessionName,
	})
	if err != nil {
		return "", 0, err
	}
	return token, exp, nil
}
