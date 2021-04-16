package main

import (
	"crypto/md5"
	"encoding/hex"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func validImgExtension(extension string) bool {
	for _, ext := range validImageExtensions {
		if extension == ext {
			return true
		}
	}
	return false
}
