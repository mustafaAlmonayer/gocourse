package intermediate

import (
	"fmt"
	"time"
)

func timeFormatingParsing() {

	// Mon Jan 2 15:04:05 MST 2006

	layout := "2006-01-02-T15:04:05Z07:00"

	str := "2024-07-04-T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	fmt.Print(t)

}
