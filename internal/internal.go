package internal

import (
	"io"
	"io/ioutil"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// DefaultTemplate is the default template for rendering mrconfig files
const DefaultTemplate = `##
# {{ .Name }} repositories
#
{{- $prefix := .Prefix }}
{{- $owner := .Name }}
{{ range .ReposList }}
[{{ $prefix }}/{{ $owner }}/{{ .Name }}]
checkout = git clone '{{ .GitURL }}' '{{ .Name }}'
{{ end }}
`

// Repo is a simple struct to record the details of a repository
type Repo struct {
	Name   string
	GitURL string
}

// Mrconfig is a container type to hold the necessary information
//  to render a mrconfig file
type Mrconfig struct {
	Prefix    string
	Name      string
	ReposList []Repo
}

// getTemplate checks if a user supplied template was specified, if so
//  it checks it's valid, else it returns the default template
func getTemplate(file string) string {
	if file == "" {
		return DefaultTemplate
	}

	tpl, err := ioutil.ReadFile(file)
	if err != nil {
		log.WithFields(log.Fields{
			"file":  file,
			"error": err,
		}).Debug("error reading template file")

		return DefaultTemplate
	}

	tplString := string(tpl)

	if _, err := template.New("tpl").Parse(tplString); err != nil {
		log.WithFields(log.Fields{
			"file":  file,
			"error": err,
		}).Debug("error parsing template file")

		return DefaultTemplate

	}

	return tplString
}

// RenderTemplate is responsible for
func RenderTemplate(prefix, name string, repos []Repo, tmpl string, output io.Writer) {

	mrconfig := Mrconfig{
		Prefix:    prefix,
		Name:      name,
		ReposList: repos,
	}

	// Create a new template and parse the mrconfig into it.
	t := template.Must(template.New("mrconfig").Parse(getTemplate(tmpl)))

	err := t.Execute(output, mrconfig)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("error executing template")
	}
}
