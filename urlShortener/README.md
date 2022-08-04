# URL Shortener
URL Shortener exercise from https://gophercises.com/

create an http.Handler to redirec user to a new page like a url Shortener service

urlshort.go (from handler.go) and main.go (from main/main.go) starting frames from https://github.com/gophercises/urlshort

### Bonus objectives
- Update the main/main.go source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string.
- Build a JSONHandler that serves the same purpose, but reads from JSON data.
- Build a Handler that doesn’t read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.

## v1
- completed exercise (no bonus) 
- some items updated after watching solution to correct errors

## v2
- added flag to parse yaml file, if no file then use default 
- added json handler with flag to aprse json file if no file then use default