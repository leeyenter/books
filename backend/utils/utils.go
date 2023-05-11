package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var ipSalt = GetEnv("IP_SALT", "82k%52G6")

// SaltedRemoteIPForTests is a magic value using "remote-ip" with salt "82k%52G6",
// sha256-ed and base64-encoded. Intended for easier/more consistent test setup.
const SaltedRemoteIPForTests = "wnxvHrLE4onzBk3C60swhsu+FgjFow/JJATc0pTba4k="

func GetRemoteIp(req *http.Request) string {
	remoteIp := req.Header.Get("X-Forwarded-For")
	if remoteIp == "" {
		remoteIp = strings.Split(req.RemoteAddr, ":")[0]
	}

	remoteIp += ipSalt

	h := sha256.New()
	h.Write([]byte(remoteIp))
	sumBytes := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sumBytes)
}

func WrapRequestForTest(req *http.Request, validJwt string) {
	req.RemoteAddr = "remote-ip"
	req.Header.Set("auth", validJwt)
}
