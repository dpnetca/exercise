package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	mapHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// can combine as `if destination, ok := pathsToUrls[path]; ok { ...
		destination, ok := pathsToUrls[path]
		if ok {
			http.Redirect(w, r, destination, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
	return mapHandlerFunc

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yamlData)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJson(jsonData)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
}

type pathData struct {
	Path string
	Url  string
}

func parseYaml(yamlData []byte) ([]pathData, error) {
	var data []pathData
	err := yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseJson(jsonData []byte) ([]pathData, error) {
	var data []pathData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func buildMap(data []pathData) map[string]string {
	mapData := make(map[string]string)
	for _, d := range data {
		mapData[d.Path] = d.Url
	}
	return mapData
}
