package transporte

import "internal/parada"

type Linea struct {
    tipoMedio string
    nombre string
    numLinea string
    paradas []*Parada
}