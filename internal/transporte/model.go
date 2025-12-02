package transporte


type Linea struct {
    tipoMedio string
    nombre string
    numLinea string
    paradas []*Parada
}

type Parada struct {
    nombre string
    correspondenciaLineas map[string]*Linea
}