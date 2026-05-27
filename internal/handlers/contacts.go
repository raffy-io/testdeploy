package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/raffy-io/testdeploy/ui/layout"
	"github.com/raffy-io/testdeploy/ui/pages"
)

func GetContacts(w http.ResponseWriter, r *http.Request){
	component := pages.Contacts()
	pageLayout := layout.BaseLayout("Contact Us",component)
	templ.Handler(pageLayout).ServeHTTP(w,r)
}
