package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateKey(t *testing.T) {
	rand.Seed(time.Now().Unix())
	key := GenerateKey(4)
	fmt.Println(key)
}
