package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/raffy-io/testdeploy/ui/layout"
	"github.com/raffy-io/testdeploy/ui/pages"
)

func GetHome(w http.ResponseWriter, r *http.Request){
	component := pages.Home()
	pageLayout := layout.BaseLayout("Welcome",component)
	templ.Handler(pageLayout).ServeHTTP(w,r)
}