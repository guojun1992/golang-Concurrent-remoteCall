package lib

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCall(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	urls := []string{
		"http://safe.skin.uc.360.cn/index.php?method=holidaySkin.getSkin&app=safe&guid=qwe&appkey=safecommon&ver=12",
		"http://safe.skin.uc.360.cn/index.php?method=holidaySkin.getSkin&app=safe&guid=qwe&appkey=safecommon&ver=12",
	}

	s := Call(urls)
	// for i := 0; i < len(urls); i++ {
	// 	s := He(urls[i])
	// 	fmt.Println(s)
	// }

	fmt.Println(s)
}
