# Choose Your Own Adventure
Choose Your Own Adventure exercise from https://gophercises.com/

story json file from https://github.com/gophercises/cyoa

## v1 
CLI based, print story to terminal and present text menu for options 

## v2
HTML Verision using `html/template` to create HTML pages, and `http.Handler` to handle the requests

## v3
Added flags to toggle CLI vs WEB and provide custom story input file

## v4
updates/changes while watching course solution videos
- use seperate commands to start CLI vs WEB isntead of flags
- set http port with flag
- change to use http.Handler interface instead of http.HandlerFunc
- globally declare template instead of loading it in Handler
- allow custom http options to be passed in
- created a http server mux