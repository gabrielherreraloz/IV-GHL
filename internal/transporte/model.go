package transporte


type TipoMedio int
const (
    AUTOBUS TipoMedio = iota
    TRANVIA
    TAXI
)

type Linea struct {
    tipoMedio TipoMedio
    nombre string
    numLinea string
    paradas map[uint]*Parada
}

type Parada struct {
    nombre string
    // Referencia a la linea
    lineasIds map[uint]*Linea
}

// Almacenamos las lineas a partir de un ID único.
type Almacen struct {
    paradas map[uint]*Parada
    lineas map[uint]*Linea
}

