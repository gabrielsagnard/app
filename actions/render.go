package actions

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),

		// Add template helpers here:
		Helpers: render.Helpers{
			"assetPath": func(file string) template.HTML {
				manifest, err := assetsBox.MustString("manifest.json")
				if err != nil {
					log.Println("[INFO] didn't find manifest, using raw path to assets")

					return template.HTML(fmt.Sprintf("assets/%v", file))
				}

				var manifestData map[string]string
				err = json.Unmarshal([]byte(manifest), &manifestData)

				if err != nil {
					log.Println("[Warning] seems your manifest is not correct")
					return ""
				}

				if file == "application.css" {
					file = "main.css"
				}

				if file == "application.js" {
					file = "main.js"
				}

				return template.HTML(fmt.Sprintf("assets/%v", manifestData[file]))
			},
		},
	})
}
