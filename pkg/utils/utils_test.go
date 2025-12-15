package utils

import (
    "testing"
    "strings"
    "os"
    "path/filepath"
    "IV-GHL/pkg/scrapper"
)

// Ruta de los HTML descargados de ejemplo estandar (Happy Path)
const testDataDir = "../../testdata" 

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA EXTRAER LÍNEA COMPLETA (HAPPY PATH COMPLETO)
/////////////////////////////////////////////////////////////////////////////////////

func TestExtraerLinea(t *testing.T) {
    files, err := filepath.Glob(filepath.Join(testDataDir, "*.html"))
    if err != nil {
        t.Fatalf("Error al buscar archivos HTML en %s: %v", testDataDir, err)
    }

    if len(files) == 0 {
        t.Fatalf("No se encontraron archivos HTML en la carpeta 'testdata'")
    }

    // Iteraración sobre cada HTML
    for _, filePath := range files {
        fileName := filepath.Base(filePath) 

        t.Run("Happy_Path_File_" + fileName, func(t *testing.T) {
            t.Parallel() 

            htmlContent, err := os.ReadFile(filePath)
            if err != nil {
                t.Fatalf("No se pudo leer el archivo %s: %v", fileName, err)
            }
            
            htmlString := string(htmlContent)
            linea, err := scrapper.ExtraerLinea(htmlString)

            if err != nil {
                t.Errorf("FAIL [%s]: Se obtuvo un error: %v", fileName, err)
                return
            }

            expectedNumLinea := strings.TrimPrefix(strings.TrimSuffix(fileName, ".html"), "linea_")

            if linea.NumLinea != expectedNumLinea {
                t.Errorf("FAIL [%s]: NumLinea incorrecto. Se esperaba '%s', se obtuvo '%s'", 
                    fileName, expectedNumLinea, linea.NumLinea)
            }
            
            if len(linea.Horario_Paradas_Ida) == 0 {
                t.Errorf("FAIL [%s]: No se extrajeron paradas de Ida.", fileName)
            }

            if len(linea.Horario_Paradas_Vuelta) == 0 {
                t.Errorf("FAIL [%s]: No se extrajeron paradas de Vuelta.", fileName)
            }
        })
    }
}

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA EXTRAER PARADAS (SAD PATH)
/////////////////////////////////////////////////////////////////////////////////////

func TestExtraerNombresParadas(t *testing.T) {
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
// TESTS PARA EXTRAER NÚMERO DE LÍNEA (SAD PATH)
/////////////////////////////////////////////////////////////////////////////////////

func TestGetNumLinea(t *testing.T) {
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
// TESTS PARA EXTRAER Y LIMPIAR TEXTO DE CELDA (SAD PATH)
/////////////////////////////////////////////////////////////////////////////////////

func TestExtraerTexto(t *testing.T) {
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