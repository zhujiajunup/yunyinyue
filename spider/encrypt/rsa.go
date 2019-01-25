package encrypt

import (
	"math/rand"
	"time"
)

func Random(size int)(result []byte){
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(s)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i ++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return
}
