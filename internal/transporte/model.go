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
    paradas [uint]*Parada
}

type Parada struct {
    nombre string
    // Referencia a la linea
    lineasIds map[uint]*Linea
}

// Almacenamos las lineas a partir de un ID único.
var paradas = map[uint]*Parada{}
var lineas = map[uint]*Linea{}

