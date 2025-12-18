package internal
import "errors"

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
	Lineas map[uint]*Linea
	Paradas map[uint]*Parada
}

var ErrLíneaExistente = errors.New("la línea introducida ya se había guardado anteriormente")

func (alm *Almacen) AlmacenarLinea(id uint, nuevaLinea *Linea) error {
    if alm.Lineas == nil {
        alm.Lineas = make(map[uint]*Linea)
    }

    for _, l := range alm.Lineas {
        if l.NumLinea == nuevaLinea.NumLinea {
            return ErrLíneaExistente
        }
    }

    alm.Lineas[id] = nuevaLinea
    return nil
}

