package util

import (
"math/rand"
"time"
)

func RandomString(i int) string {
	var letters = []byte("assefjksdfiajkgljkalsjSDKGJAJOWNVMFNDFVLKJJSFSLKFJGOIR")
	result := make([]byte, i)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
