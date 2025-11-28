package parada

import "internal/transporte"

type Parada struct {
    nombre string
    correspondenciaTransportes map[string]*Trasnporte
}
