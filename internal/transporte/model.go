package transporte


type Linea struct {
    tipoMedio string
    nombre string
    numLinea string
    paradas []*Parada
}

type Parada struct {
    nombre string
    // Referencia a la linea
    lineasIds map[uint]*Linea
}

// Almacenamos las lineas a partir de un ID único.
lineas = map[uint]Linea

