package persona

type Persona struct {
    ID   string `json:"id"`
    Nombre string `json:"nombre"`
    Edad int `json:"edad"`
    PesoMochila float32 `json:"pesoMochila"`
}