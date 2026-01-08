package intermediate

import (
	"fmt"
	"time"
)

func epoch() {

	// 00:00:00 UTC on Jan 1 1970

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println(unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)

}
