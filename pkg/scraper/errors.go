package scraper

import (
	"errors"
)

var (
	ErrNoH2 = errors.New("no se encontró la etiqueta <h2> para el título")
	ErrNoMarcadorIda = errors.New("no se encontró el marcador de 'Ida'")
	ErrNoMarcadorVuelta = errors.New("no se encontró el marcador de 'Vuelta'")
	ErrFormatoInvalido = errors.New("el formato del HTML ha cambiado o es inválido")
	ErrCeldaVacia = errors.New("la celda no contiene texto válido tras la limpieza del HTML")
    ErrNoParadas  = errors.New("no se encontraron etiquetas <th> en el bloque")
	ErrNoSeparador = errors.New("el título de la línea no contiene el separador ' - '")
)