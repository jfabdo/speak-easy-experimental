package api

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/mediocregopher/radix/v3"
)

var ctx = context.Background()

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func getClient() *radix.Pool {
	pool, err := radix.NewPool("tcp", "172.17.0.1:6379", 10)
	if err != nil {
		println(err)
		panic(err)
	} else {
		return pool
	}
}

// CheckMessages checks the incoming list of messages and makes sure that
func CheckMessages(w http.ResponseWriter, usermap map[string]string) {
	var rdb = getClient()

	var canonicallist []string
	err := rdb.Do(radix.Cmd(&canonicallist, "lrange", "chat"+"message"+channel+usermap["name"], "0", "10"))
	if err != nil {
		println(err)
		panic(err)
	}
	print(canonicallist)
	for i := range canonicallist {
		println(canonicallist)
		if usermap[canonicallist[i]] == "" {
			fmt.Fprintf(w, canonicallist[i])

		}
	}
}
