package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// SixRandomNums 6位随机数
func SixRandomNums() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// //main测试
// func main() {
// 	fmt.Println(RandomNums())
// }
