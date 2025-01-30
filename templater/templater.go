package templater

// type template struct {
// 	text string
// 	envs map[string]string
// }

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

// ParseTemplate uses Go's text/template package for parsing
func ParseTemplate(text string, envs map[string]string) (string, error) {
	tmpl, err := template.New("template").Parse(text)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, envs)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ParseRegex uses regex to parse simple template patterns
func ParseRegex(text string, envs map[string]string) (string, error) {
	// Define regex to match {template_name}, {templateName}, or {template.name}
	re := regexp.MustCompile(`\{([a-zA-Z0-9._]+)\}`)

	// Replace all matches with corresponding values from envs
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		key := strings.Trim(match, "{}")
		if val, ok := envs[key]; ok {
			return val
		}
		return match
	})

	return result, nil
}
