package scraper

import (
	"errors"
)

var (
	ErrH2NotFound = errors.New("no se encontró la etiqueta <h2> para el título")
	ErrMarcadorIda = errors.New("no se encontró el marcador de 'Ida'")
	ErrMarcadorVuelta = errors.New("no se encontró el marcador de 'Vuelta'")
	ErrFormatoInvalido = errors.New("el formato del HTML ha cambiado o es inválido")
	ErrCeldaVacia = errors.New("la celda no contiene texto válido tras la limpieza del HTML")
)