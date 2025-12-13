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

func extraerNombresParadas(bloqueHTML string) []string {
    var nombres []string

    re := regexp.MustCompile(`(?s)<th[^>]*>(.*?)</th>`)
    matches := re.FindAllStringSubmatch(bloqueHTML, -1)

    for _, match := range matches {
		nombres = append(nombres, ExtraerTexto(match[1]))
    }

    return nombres
}

func ExtraerTexto(celda string) string {
	reTags := regexp.MustCompile(`<[^>]*>`)

    s := reTags.ReplaceAllString(celda, "")
    s = strings.ReplaceAll(s, "&nbsp;", "---")
    s = strings.TrimSpace(s)
    s = strings.ReplaceAll(s, "\n", "")
    s = strings.ReplaceAll(s, "\r", "")

    if s == "" || s == "-" {
        return "---"
    }
    return s
}
