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

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:user_id; primaryKey"`
	Name          string `json:"name,omitempty" gorm:"column:name"`
	PassWord      string `gorm:"column:password" json:"-"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:follow_count"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:is_follow"`
	Salt          string `json:"-" gorm:"column:salt"`
	Token         string `gorm:"-" json:"token,omitempty"`
}
type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"column:video_id"`
	Author        User   `json:"author" gorm:"-" sql:"-" ` //不会入库，只做返回给前端的时候结构体
	UserId        int64  `json:"user_id" gorm:"user_id"`   //入库，实际存入的是userId
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"-" sql:"-"`
	Title         string `json:"title,omitempty"`
}

func main() {
	//dsn := "doul:Tyf136212.@tcp(81.68.239.206:3306)/douyin?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   "dy_", // table name prefix, table for `User` would be `t_users`
	//		SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
	//		NoLowerCase:   true,  // skip the snake_casing of names
	//	},
	//})
	//db.Model(Video{}).Preload()

}
