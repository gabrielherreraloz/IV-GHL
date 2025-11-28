package transporte

import "internal/parada"

type Transporte struct {
    tipoMedio string
    nombre string
    numLinea string
    paradas []Parada
}