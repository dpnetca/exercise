package cyoa

const StoryTpl = `
<html>
  <head>
    <title> {{.Title}} </title>
  </head>
  <body>
	<h1>{{.Title}}</h1>
	{{range .Story}}
	<p>{{.}}</p>
	{{end}}
	<ul>
	{{range .Options}}
	  <li> <a href="/?arc={{ .Arc}}">{{ .Arc}}</a>  - {{ .Text}}</li>
	{{end}}
	</ul>
  </body>
`
