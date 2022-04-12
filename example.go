package metadata1

import (
	"fmt"
	"time"
)

func example() {
	payload := NewPayload(BasePayload{
		Location:  "https://www.amazon.com/ap/signin?...",
		Referrer:  "https://www.amazon.com/ap/signin?...",
		Start:     int(time.Now().UnixNano() / 1e6),
		UserAgent: "...",
	})

	fmt.Println(Encrypt(payload.ToString()))
}