package parada

import "internal/transporte"

type Parada struct {
    nombre string
    correspondenciaLineas map[string]*Linea
}
