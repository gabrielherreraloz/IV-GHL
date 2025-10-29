package transporte

type Transporte struct {
    ID   string `json:"id"`
    Nombre string `json:"nombre"`
    coordenadaX float32 `json:"coordenadaX"`
    coordenadaY float32 `json:"coordenadaY"`
    vaSentado bool `json:"vaSentado"`
}