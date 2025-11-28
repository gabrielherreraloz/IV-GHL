package parada

import "internal/transporte"

type Parada struct {
    nombre string
    latitud float64
    longitud string
    correspondenciaTransportes map[string]*Trasnporte
}
