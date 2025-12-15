package scrapper

import (
	"fmt"
	"strings"
	"IV-GHL/internal"
	"regexp"
)

func ExtraerLinea(htmlContent string) (*internal.Linea, error) {
	linea := &internal.Linea{
		TipoMedio:              internal.AUTOBUS,
		Horario_Paradas_Ida:    make(map[string][]string),
		Horario_Paradas_Vuelta: make(map[string][]string),
	}

	// Extracción número de la línea
	numLinea, error := GetNumLinea(htmlContent)
	if(error != nil){
		return linea, error
	}

	linea.NumLinea = numLinea

	// Bloques ida y vuelta
	idxIda := strings.Index(htmlContent, "<strong>Ida</strong>")
	idxVuelta := strings.Index(htmlContent, "<strong>Vuelta</strong>")
	idxFinVuelta := strings.Index(htmlContent, `<div class="leyendas">`)

	if idxIda == -1 {
		error = fmt.Errorf("no se encontraron los marcadores Ida")
		return linea, error
	}
	if idxVuelta == -1 {
		error = fmt.Errorf("no se encontraron los marcadores Vuelta")
		return linea, error
	}
	if idxFinVuelta == -1 {
		error = fmt.Errorf("no se encontraron los marcadores FinVuelta")
		return linea, error
	}

	bloqueIda := htmlContent[idxIda:idxVuelta]
	bloqueVuelta := htmlContent[idxVuelta:idxFinVuelta]

	// Extracción de nombres de paradas y horarios de ida y vuelta
	nombresIda := ExtraerNombresParadas(bloqueIda)
	for i := 0; i < len(nombresIda); i++ {
		linea.Horario_Paradas_Ida[nombresIda[i]] = append(linea.Horario_Paradas_Ida[nombresIda[i]], "")
	}

	nombresVuelta := ExtraerNombresParadas(bloqueVuelta)
	for i := 0; i < len(nombresVuelta); i++ {
		linea.Horario_Paradas_Vuelta[nombresVuelta[i]] = append(linea.Horario_Paradas_Vuelta[nombresVuelta[i]], "")
	}

	return linea, error
}

func GetNumLinea(htmlContent string) (string, error) {
    re := regexp.MustCompile(`<h2>\s*(.*?)\s*</h2>`)
    match := re.FindStringSubmatch(htmlContent)

    if len(match) < 2 {
        return "", fmt.Errorf("el contenido HTML no contiene el encabezado <h2> necesario para el título")
    }

    tituloCompleto := strings.TrimSpace(match[1])
    partes := strings.SplitN(tituloCompleto, " - ", 2)

    return strings.TrimSpace(partes[0]), nil
}

func ExtraerNombresParadas(bloqueHTML string) []string {
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