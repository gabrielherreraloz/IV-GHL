package transporte

type Parada struct {
    nombre string `json:"nombre"`
    latitud float64 `json:"latitud"`
    longitud string `json:"longitud"`
    lineasCorrespondencia []string `json:"lineasCorrespondencia"`
}