package metadata1

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"math"
	"strings"
)

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

func tea(payload string, material []int) []byte {
	chunks := int(math.Ceil(float64(len(payload)) / float64(4)))
	o := []int32{}

	for i := 0; i < chunks; i++ {
		o = append(o, 255&charCodeAt(payload, 4*i)+(255&charCodeAt(payload, 4*i+1))<<8+(255&charCodeAt(payload, 4*i+2))<<16+(255&charCodeAt(payload, 4*i+3))<<24)
	}

	n := math.Floor(6 + 52/float64(chunks))

	c := o[chunks-1]

	for d := 0; n > 0; n-- {
		d += 2654435769
		h := int((uint32(d) >> 2) & 3)

		for u := 0; u < chunks; u++ {
			a := o[(u+1)%chunks]
			o[u] += int32((int(uint32(c)>>5)^int(a<<2))+
				(int(uint32(a)>>3)^int(c<<4))) ^ ((int32(d)^a)+
				(int32(material[(3&u)^h])^c))
			c = o[u]
		}
	}

	g := []byte{}
	for s := 0; s < chunks; s++ {
		g = append(g, []byte{byte(255 & o[s]), byte((uint32(o[s]) >> 8) & 255), byte((uint32(o[s]) >> 16) & 255), byte((uint32(o[s]) >> 24) & 255)}...)
	}

	return g
}

func toTwosComplementHex(num uint32) string {
	//size := 8
	if num >= 0 {
		return fmt.Sprintf("%08x", num) //08x is the size as declared
	}
	return ""
}

var material = []int{1888420705, 2576816180, 2347232058, 874813317}
var identifier = "ECdITeCs"

func Encrypt(data string) string {
	checksum := crc32.Checksum([]byte(data), crc32.IEEETable)
	crcsum := toTwosComplementHex(checksum)
	crcsum = strings.ToUpper(crcsum)
	payload := crcsum + "#" + data
	bytes := tea(payload, material)
	b64 := base64.StdEncoding.EncodeToString(bytes)
	return identifier + ":" + b64
}