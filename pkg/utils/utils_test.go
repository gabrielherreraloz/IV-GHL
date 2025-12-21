package utils

import (
    "testing"
    "os"
    "path/filepath"
    "IV-GHL/pkg/scraper"
    "IV-GHL/internal"
    "errors"
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

            htmlContent, err := os.ReadFile(filePath)
            if err != nil {
                t.Fatalf("No se pudo leer el archivo %s: %v", fileName, err)
            }
            
            htmlString := string(htmlContent)
            linea, err := scraper.ExtraerLinea(htmlString)

            if err != nil {
                t.Errorf("No se esperaba error en línea %s y se obtuvo: %v", err, linea)
                return
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
        nombres, err := scraper.ExtraerNombresParadas(html)

        if !errors.Is(err, scraper.ErrNoParadas){
            t.Errorf("Se esperaba ErrNoParadas, se obtuvo: %v", err)
        }
        if (len(nombres) != 0){
            t.Errorf("Se esperaba lista vacía, se obtuvieron %d elementos", len(nombres))
        }
    })
}

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA EXTRAER NÚMERO DE LÍNEA (SAD PATH)
/////////////////////////////////////////////////////////////////////////////////////

func TestGetNumLinea(t *testing.T) {
    t.Run("Sad_Path_Sin_Separador", func(t *testing.T) {
        html := `<h2> LINEA SIN GUION </h2>`
        numLinea, err := scraper.GetNumLinea(html)

        if !errors.Is(err, scraper.ErrNoSeparador){
            t.Errorf("Se esperaba error %v, pero se obtuvo: %v", scraper.ErrNoSeparador, err)
        }

        if (numLinea != ""){
            t.Errorf("Se esperaba string vacío al haber error, se obtuvo '%s'", numLinea)
        }
    })
    
    t.Run("Sad_Path_Sin_Titulo_H2", func(t *testing.T) {
        html := `<h1>Otro Titulo</h1>`
        numLinea, err := scraper.GetNumLinea(html)

        if !errors.Is(err, scraper.ErrNoH2){
            t.Errorf("Se esperaba el error específico ErrH2NotFound, se obtuvo: %v", err)
        }
        
        if (numLinea != ""){
            t.Errorf("Se esperaba string vacío, se obtuvo '%s'", numLinea)
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
        
        for input := range inputs {
            result, err := scraper.ExtraerTexto(input)
            
            if !errors.Is(err, scraper.ErrCeldaVacia){
                t.Errorf("Input '%s': Se esperaba ErrCeldaVacia, se obtuvo: %v", input, err)
            }
            
            if (result != "---"){
                t.Errorf("Input '%s': Se esperaba '---', se obtuvo '%s'", input, result)
            }
        }
    })
}

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA ÚNICA FUENTE DE VERDAD
/////////////////////////////////////////////////////////////////////////////////////

func Testalmacenado(t *testing.T){
    t.Run("Test_Líneas_Repetidas", func(t *testing.T){
        alm := &internal.Almacen{
        Lineas: make(map[uint]*internal.Linea),
        }
        linea1 := &internal.Linea{
            NumLinea: "L1",
            TipoMedio: internal.AUTOBUS,
        }
        linea2 := &internal.Linea{
            NumLinea: "L1",
            TipoMedio: internal.AUTOBUS,
        }

        // 3. Primer intento: Debe ser exitoso (Happy Path)
        err := alm.AlmacenarLinea(1, linea1)
        if err != nil {
            t.Fatalf("No se esperaba error al insertar la primera vez, pero se obtuvo: %v", err)
        }

        // 4. Segundo intento: Debe fallar (Sad Path)
        err = alm.AlmacenarLinea(2, linea2)
        if !errors.Is(err, internal.ErrLíneaExistente) {
            t.Errorf("Se esperaba el error internal.ErrLineaExistente, pero se obtuvo: %v", err)
        }
        if len(alm.Lineas) != 1 {
            t.Errorf("El almacén debería tener solo 1 línea guardada, pero tiene %d", len(alm.Lineas))
        }
    })
}