package internal

type TipoMedio int
const AUTOBUS TipoMedio = 0
const METRO TipoMedio = 1

type Linea struct {
	TipoMedio              TipoMedio
	NumLinea               string
	Horario_Paradas_Ida    map[string][]string
	Horario_Paradas_Vuelta map[string][]string
}

type Parada struct {
	Nombre string
	Lineas_ida []*Linea
	Lineas_vuelta []*Linea
}

type Almacen struct {
	lineas map[uint]Linea
	paradas map[uint]Parada
}

