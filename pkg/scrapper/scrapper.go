package scrapper

import (
	"strings"
	"regexp"
)

func getNumLinea(htmlContent string) string {
	re := regexp.MustCompile(`<h2>\s*(.*?)\s*</h2>`)
    match := re.FindStringSubmatch(htmlContent)
	tituloCompleto := strings.TrimSpace(match[1])
	partes := strings.SplitN(tituloCompleto, " - ", 2)

	return strings.TrimSpace(partes[0])
}
