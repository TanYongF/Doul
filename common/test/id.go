package main

import (
	"fmt"
	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"log"
	"math/rand"
	"time"
)

// To run:
// go run main.go

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
}

func genKsuid() {
	ksuid.New()
	//fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
}

func genBetterGUID() {
	_ = betterguid.New()
	//fmt.Printf("github.com/kjk/betterguid:   %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	ulid.MustNew(ulid.Timestamp(t), entropy)
	//fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	_, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	//fmt.P/**/rintf("github.com/sony/sonyflake:   %x\n", id)
}

func genSid() {
	_ = sid.Id()
	//fmt.Printf("github.com/chilts/sid:       %s\n", id)
}

func genUUIDv4() {
	var _ = uuid.NewV4()
	//fmt.Printf("github.com/satori/go.uuid:   %s\n", id)
}

func main() {
	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSid()
	genUUIDv4()
}
