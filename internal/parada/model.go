package parada

import "internal/transporte"

type Parada struct {
    nombre string `json:"nombre"`
    latitud float64 `json:"latitud"`
    longitud string `json:"longitud"`
    lineasCorrespondencia map[string]Trasnporte `json:"lineasCorrespondencia"`
}
