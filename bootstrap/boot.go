package bootstrap

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mmm888/go-wiki/middleware/worker"

	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/markdown"
	"github.com/mmm888/go-wiki/middleware/templates"
	"github.com/mmm888/go-wiki/middleware/variable"
)

const (
	configJSONPath = "config.json"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "address to bind")
	flag.Parse()
}

func Start(m *middleware.M) {

	// 設定ファイル読み込み
	{
		config := variable.CommonVars{
			ConfigPath: configJSONPath,
		}

		// ファイルが存在する場合
		if _, err := os.Stat(configJSONPath); err == nil {
			data, err := ioutil.ReadFile(configJSONPath)
			if err != nil {
				log.Print(err)
			}

			if err := json.Unmarshal(data, &config); err != nil {
				log.Print(err)
			}
		}

		m.CommonVars = &config
	}

	// templates初期化
	{
		tmpls := templates.NewTemplates("templates")
		tmpls.Route("show", "show.tmpl", "layout.tmpl")
		tmpls.Route("edit", "edit.tmpl", "layout.tmpl")
		tmpls.Route("diff", "diff.tmpl", "layout.tmpl")
		tmpls.Route("config", "config.tmpl", "layout.tmpl")
		tmpls.Route("tree", "tree.tmpl", "layout.tmpl")

		m.Templates = tmpls
	}

	// markdown初期化
	//m.Markdown = markdown.NewBlackfriday()
	m.Markdown = markdown.NewGithubMarkdown()

	// worker初期化
	m.JobQueue = worker.NewJobQueue(100)
	m.JobQueue.Start()
	defer m.JobQueue.Stop()

	// ルーティング設定
	registerRoute(m)

	// git初期設定
	gitSetting(m)

	// サーバはブロックするので別の goroutine で実行する
	srv := &http.Server{Addr: addr, Handler: m.Router}
	go func() {
		log.Printf("Start HTTP Server %v", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Error HTTP Server %v", err)
		}
	}()

	// シグナルを待つ
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh

	// シグナルを受け取ったらShutdown
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	log.Print("Shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Error Shutdown() %v", err)
	}
}
