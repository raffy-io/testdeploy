package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/raffy-io/testdeploy/ui/layout"
	"github.com/raffy-io/testdeploy/ui/pages"
)

func GetAbout(w http.ResponseWriter, r *http.Request){
	component := pages.About()
	pageLayout := layout.BaseLayout("About Us",component)
	templ.Handler(pageLayout).ServeHTTP(w,r)
}