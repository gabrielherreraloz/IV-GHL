package transporte

import "internal/parada"

type Transporte struct {
    tipoMedio string `json:"tipoMedio"`
    nombre string `json:"nombre"`
    numLinea string `json:"numLinea"`
    paradas []Parada `json:"paradas"`
}