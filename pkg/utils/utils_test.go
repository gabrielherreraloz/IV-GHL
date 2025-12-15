package utils

import (
	"testing"
	"strings"
    "IV-GHL/pkg/scrapper"
)

// Bloque para procesar número de línea en GetNumLinea
const mockHTMLTitulo = `
    <div class="sub-nav-titl"><h2>0100 - Granada - Jun - A.Humeya - Víznar </h2></div>
    <div class="migas">
        <span>0100 - Granada - Jun - A.Humeya - Víznar</span>
    </div>
`

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA EXTRAER PARADAS
/////////////////////////////////////////////////////////////////////////////////////

func TestExtraerNombresParadas(t *testing.T) {
    t.Run("Happy_Path_Estructura_Normal", func(t *testing.T) {
        html := `
            <table>
                <tr>
                    <th>Parada A</th>
                    <th>Parada B</th>
                    <th><a href="/link">Parada C</a></th>
                </tr>
            </table>
        `
        nombres := scrapper.ExtraerNombresParadas(html)
        
        expected := []string{"Parada A", "Parada B", "Parada C"}
        if len(nombres) != len(expected) {
            t.Fatalf("Longitud incorrecta. Se esperaba %d, se obtuvo %d", len(expected), len(nombres))
        }
        for i, name := range nombres {
            if name != expected[i] {
                t.Errorf("Fallo en índice %d. Se esperaba '%s', se obtuvo '%s'", i, expected[i], name)
            }
        }
    })

    t.Run("Sad_Path_Sin_Paradas", func(t *testing.T) {
        html := `<table><tr></tr></table>`
        nombres := scrapper.ExtraerNombresParadas(html)

        if len(nombres) != 0 {
            t.Errorf("Se esperaba lista vacía para HTML sin paradas, se obtuvo %v", nombres)
        }
    })
    
    t.Run("Sad_Path_Headers_Vacios_o_Espacios", func(t *testing.T) {
        html := `
            <table>
                <tr>
                    <th>  </th>
                    <th>&nbsp;</th>
                    <th><br></th>
                </tr>
            </table>
        `
        nombres := scrapper.ExtraerNombresParadas(html)
        
        expected := []string{"---", "---", "---"} 
        if len(nombres) != len(expected) {
            t.Fatalf("Fallo de conteo")
        }
        for _, name := range nombres {
            if name != "---" {
                t.Errorf("La limpieza de nombres vacíos o con solo tags falló. Obtenido: %s", name)
            }
        }
    })
}

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA EXTRAER NÚMERO DE LÍNEA
/////////////////////////////////////////////////////////////////////////////////////

func TestGetNumLinea(t *testing.T) {
    t.Run("Happy_Path_Titulo_Estandar", func(t *testing.T) {
        numLinea, error := scrapper.GetNumLinea(mockHTMLTitulo)
        if (error != nil) {
            t.Errorf("ERROR: Se obtuvo: %v", error)
        }

        expected := "0100"
        if numLinea != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, numLinea)
        }
    })

    t.Run("Happy_Path_Titulo_Con_Espacios_Adicionales", func(t *testing.T) {
        html := `<h2>   0200   -    OTRA RUTA   </h2>`
        numLinea, error := scrapper.GetNumLinea(html)
        if (error != nil) {
            t.Errorf("ERROR: Se obtuvo: %v", error)
        }

        expected := "0200"
        if numLinea != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, numLinea)
        }
    })

    t.Run("Sad_Path_Sin_Separador", func(t *testing.T) {
        html := `<h2> LINEA SIN GUION </h2>`
        numLinea, err := scrapper.GetNumLinea(html)

        if err != nil {
            t.Errorf("ERROR: No se esperaba un error en este caso, se obtuvo: %v", err)
            return
        }

        expected := "LINEA SIN GUION" 
        if numLinea != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, numLinea)
        }
    })
    
    t.Run("Sad_Path_Sin_Titulo_H2", func(t *testing.T) {
        html := `<h1>Otro Titulo</h1>`
        
        // Ahora GetNumLinea devuelve "" si no hay match
        numLinea, error := scrapper.GetNumLinea(html)
        if (error == nil) {
            t.Errorf("ERROR: Se esperaba un error pero no se detectó")
        }
        expected := ""
        
        if numLinea != expected {
            t.Errorf("Se esperaba una cadena vacía ('%s') al no encontrar H2, Se obtuvo '%s'", expected, numLinea)
        }
    })
}

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA EXTRAER Y LIMPIAR TEXTO DE CELDA (ExtraerTexto ya estaba exportada)
/////////////////////////////////////////////////////////////////////////////////////

func TestExtraerTexto(t *testing.T) {
    t.Run("Happy_Path_Texto_Estandar", func(t *testing.T) {
        input := "   Texto Limpio  "
        expected := "Texto Limpio"
        if result := scrapper.ExtraerTexto(input); result != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, result)
        }
    })

    t.Run("Sad_Path_Solo_Tags_o_Espacios", func(t *testing.T) {
        inputs := map[string]string{
            "   ":          "---",
            "<strong></strong>": "---",
            "&nbsp;":       "---", 
            "-":            "---", 
        }
        
        for input, expected := range inputs {
            if result := scrapper.ExtraerTexto(input); result != expected {
                t.Errorf("Input: '%s'. Se esperaba '%s', Se obtuvo '%s'", strings.ReplaceAll(input, "\n", "\\n"), expected, result)
            }
        }
    })
}