package scraper

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
	// <h2> : Etiqueta de apertura
    // \s* : Salta espacios en blanco iniciales
    // (.*?) : Captura todos los caracteres que encuentra hasta el cierre de etiqueta, especificado detrás.
    // \s* : Salta espacios en blanco finales
    // </h2> : Etiqueta de cierre
    re := regexp.MustCompile(`<h2>\s*(.*?)\s*</h2>`)
    match := re.FindStringSubmatch(htmlContent)

    if len(match) < 2 {
        return "", fmt.Errorf("el contenido HTML no contiene el encabezado <h2> necesario para el título")
    }
	
	// match[1] contiene unicamente el nombre de la línea completo.
    tituloCompleto := strings.TrimSpace(match[1])
    partes := strings.SplitN(tituloCompleto, " - ", 2)

	// Devuelve unicamente el número de la línea, separado del nombre completo.
    return strings.TrimSpace(partes[0]), nil
}

func ExtraerNombresParadas(bloqueHTML string) []string {
    var nombres []string

	// Explicación:
	// <th[^>]*> : Busca la apertura de la etiqueta th y cualquier atributo.
	// (.*?) : Captura todos los caracteres que encuentra hasta el cierre de etiqueta, especificado detrás.
	// </th> : Busca el cierre de la etiqueta.
    re := regexp.MustCompile(`<th[^>]*>(.*?)</th>`)
    matches := re.FindAllStringSubmatch(bloqueHTML, -1)

    for _, match := range matches {
		// match[1] contiene unicamente el nombre de la parada
		nombres = append(nombres, ExtraerTexto(match[1]))
    }

    return nombres
}

func ExtraerTexto(celda string) string {
	// < : Apertura de etiqueta.
	// [^>]* : Cualquier contenido en medio que no sea el cierre.
	// > : Cierre de etiqueta.
	reTags := regexp.MustCompile(`<[^>]*>`)

	// ELimina todas las etiquetas y espacios que no formen parte del contenido deseado de la captura
    s := reTags.ReplaceAllString(celda, "")

	// Sustituye espacios en blanco (&nbsp) por tres guiones
    s = strings.ReplaceAll(s, "&nbsp;", "---")

	// Elimina los espacios en blanco en los extremos del contenido
    s = strings.TrimSpace(s)

	// Elimina saltos de línea y retornos
    s = strings.ReplaceAll(s, "\n", "")
    s = strings.ReplaceAll(s, "\r", "")

    if s == "" || s == "-" {
        return "---"
    }
    return s
}