package util

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_GenerateToken(t *testing.T) {
	token, err := GenerateToken("123", map[string]interface{}{
		"id":      "fefwq",
		"fewe1":   "fwwwwr",
		"website": "a1",
	},
		time.Now().Add(time.Second*10))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(token)
}

func Test_ParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA1MDE5OTE5NjcsImlkIjoiZmVmd3EiLCJ3ZWJzaXRlIjoiYTEifQ.vuk5slR5SbArrFCKMX0UXdQ1GeHXaOu6eZXZPdAAQ4I"
	data, err := ParseToken("123", token)
	if err != nil {
		log.Fatalln(err)
	}

	js, _ := json.Marshal(data)
	fmt.Println(string(js))
}
