package bootstrap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/variable"
)

const (
	configJSONPath = "config.json"
)

func Start(m *middleware.M) {

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

	{
		registerRoute(m)
	}
}
