package transporte


type Linea struct {
    // Identifica a la linea
    id uint
    tipoMedio string
    nombre string
    numLinea string
    paradas []*Parada
}

type Parada struct {
    nombre string
    // Lista de ids únicos para identificar la misma linea
    lineasIds []uint 
}

