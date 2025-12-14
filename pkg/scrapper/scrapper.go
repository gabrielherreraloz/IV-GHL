package scrapper

import (
	"fmt"
	"strings"
	"IV-GHL/internal"
	"regexp"
)

func ExtraerLinea(htmlContent string) *internal.Linea {
	linea := &internal.Linea{
		TipoMedio:              internal.AUTOBUS,
		Horario_Paradas_Ida:    make(map[string][]string),
		Horario_Paradas_Vuelta: make(map[string][]string),
	}

	// Extracción número de la línea
	linea.NumLinea = GetNumLinea(htmlContent)

	// Bloques ida y vuelta
	idxIda := strings.Index(htmlContent, "<strong>Ida</strong>")
	idxVuelta := strings.Index(htmlContent, "<strong>Vuelta</strong>")
	idxFinVuelta := strings.Index(htmlContent, `<div class="leyendas">`)

	if idxIda == -1 || idxVuelta == -1 || idxFinVuelta == -1 {
		fmt.Println("Error: No se encontraron los marcadores Ida/Vuelta/Leyendas.")
		return linea
	}

	bloqueIda := htmlContent[idxIda:idxVuelta]
	bloqueVuelta := htmlContent[idxVuelta:idxFinVuelta]

	// Extracción de nombres de paradas y horarios de ida y vuelta
	nombresIda := ExtraerNombresParadas(bloqueIda)
	horariosMatrizIda := ProcesarBloqueTabla(bloqueIda)
	for i := 0; i < len(horariosMatrizIda); i++ {
		for j := 0; j < len(nombresIda); j++ {
			if j < len(horariosMatrizIda[i]) {
				linea.Horario_Paradas_Ida[nombresIda[j]] = append(linea.Horario_Paradas_Ida[nombresIda[j]], horariosMatrizIda[i][j])
			}
		}
	}

	nombresVuelta := ExtraerNombresParadas(bloqueVuelta)
	horariosMatrizVuelta := ProcesarBloqueTabla(bloqueVuelta)
	for i := 0; i < len(horariosMatrizVuelta); i++ {
		for j := 0; j < len(nombresVuelta); j++ {
			if j < len(horariosMatrizVuelta[i]) {
				linea.Horario_Paradas_Vuelta[nombresVuelta[j]] = append(linea.Horario_Paradas_Vuelta[nombresVuelta[j]], horariosMatrizVuelta[i][j])
			}
		}
	}

	return linea
}

func GetNumLinea(htmlContent string) string {
    re := regexp.MustCompile(`<h2>\s*(.*?)\s*</h2>`)
    match := re.FindStringSubmatch(htmlContent)

    if len(match) < 2 {
        return ""
    }

    tituloCompleto := strings.TrimSpace(match[1])
    partes := strings.SplitN(tituloCompleto, " - ", 2)

    return strings.TrimSpace(partes[0])
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


func ProcesarBloqueTabla(bloqueHTML string) [][]string{
    var datos [][]string
    
    reTBody := regexp.MustCompile(`(?s)<tbody>(.*?)</tbody>`)
    matchTBody := reTBody.FindStringSubmatch(bloqueHTML)

    contenidoTabla := bloqueHTML
    if len(matchTBody) >= 2 {
        contenidoTabla = matchTBody[1] 
    }

    reFila := regexp.MustCompile(`(?s)<tr>(.*?)</tr>`)
    reCelda := regexp.MustCompile(`(?s)<td>(.*?)</td>`)
    filas := reFila.FindAllStringSubmatch(contenidoTabla, -1) // Aplicamos búsqueda al <tbody>

		for _, f := range filas {
		contenidoFila := f[1] // Aquí asumes que f[1] existe.
		
		var datosFila []string
		celdas := reCelda.FindAllStringSubmatch(contenidoFila, -1)
		
		for _, c := range celdas {
			datosFila = append(datosFila, ExtraerTexto(c[1]))
		}
		
		if len(datosFila) > 0 {
			datos = append(datos, datosFila)
		}
	}
	return datos
}