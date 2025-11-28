package parada

import "internal/transporte"

type Parada struct {
    nombre string
    latitud float64
    longitud string
    lineasCorrespondencia map[string]Trasnporte
}
