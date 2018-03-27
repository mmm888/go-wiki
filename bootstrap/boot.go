package bootstrap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/markdown"
	"github.com/mmm888/go-wiki/middleware/templates"
	"github.com/mmm888/go-wiki/middleware/variable"
)

const (
	configJSONPath = "config.json"
)

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
		tmpls.Route("show", "show.tmpl")
		tmpls.Route("edit", "edit.tmpl")
		tmpls.Route("diff", "diff.tmpl")
		tmpls.Route("config", "config.tmpl")

		m.Templates = tmpls
	}

	// markdown初期化
	m.Markdown = markdown.NewBlackfriday()

	// ルーティング設定
	registerRoute(m)
}
