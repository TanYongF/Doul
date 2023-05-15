package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/conf"
	"go_code/Doul/app/video/cmd/task/config"
	"time"
)

var c config.Config

var SqlConn *sql.DB
var RedisConn *redis.Client
var SyncThousandKeys = 1 // every SyncDuration update the 1000 * SyncKeys counter
var cursor = 0

var configFile = flag.String("f", "/etc/task.yaml", "the config file")

func Init() {
	flag.Parse()
	conf.MustLoad(*configFile, &c)
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       0,
	})
	SqlConn, _ = sql.Open("mysql", c.DataSourceName)
	SqlConn.SetConnMaxLifetime(100)
	SqlConn.SetMaxIdleConns(10)
}

// task sync logic
func task() {
	for i := 0; i < SyncThousandKeys; i++ {
		counterKey := fmt.Sprintf("doul:counter:video:likes:%d", cursor+i)
		cmds := RedisConn.HGetAll(context.Background(), counterKey)
		if len(cmds.Val()) == 0 {
			cursor = 0
			fmt.Printf("no values to sync in this time\n")
			return
		}

		sql := fmt.Sprintf("insert into %s(video_id, `like`, comment) values (?, ?, ?) ON DUPLICATE KEY UPDATE `like` = ?, `comment` = ?", "dy_video_counter")
		for k, v := range cmds.Val() {
			_, err := SqlConn.Exec(sql, k, v, 0, v, 0)
			if err != nil {
				fmt.Printf("insert error %+v", err)
				return
			}
		}
	}
	cursor += SyncThousandKeys
	fmt.Printf("%s sync success!\n", time.Now().Format("2006-01-02 15:04:01"))
}

func main() {

	Init()
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	//定时任务
	spec := "*/5 * * * * ?" //cron表达式，每五秒一次
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()

	defer SqlConn.Close()
	defer RedisConn.Close()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} //阻塞主线程停止

}
