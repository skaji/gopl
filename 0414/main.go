// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"html/template"

	"github.com/skaji/gopl/github"
)

//!+template

func onlyTitle(m *github.Milestone) string {
	if m == nil {
		return "N/A"
	}
	return m.Title
}

var issueList = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"onlyTitle": onlyTitle}).
	Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Milestone</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td>{{.Milestone | onlyTitle}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

//!-template

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	w := io.Writer(&b)
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Length", strconv.Itoa(b.Len()))
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write(b.Bytes())
	})
	fmt.Println("server :8080")
	http.ListenAndServe(":8080", nil)

}

//!-
