package utils // <-- Paquete de pruebas externo

import (
	"testing"
	"strings"
    "IV-GHL/pkg/scrapper" // <-- Necesitas importar el paquete scrapper
)

/////////////////////////////////////////////////////////////////////////////////////
// Configuración de Mocks y Errores
/////////////////////////////////////////////////////////////////////////////////////

// Bloque para procesar horarios en ProcesarBloqueTabla
const mockHTMLTabla = `
    <strong>Ida</strong>
    <table border="0" class="table tabla_horario">
        <thead>
            <tr>
                <th valign="bottom" style="background:#ffffbf;" data-toggle="true"><div id="nucleo_0">Granada</div></th>
                <th valign="bottom" style="background:#ffedae;"><div id="nucleo_1">Jun (San Jerónimo)</div></th>
                <th valign="bottom" style="background:#ffedae;"><div id="nucleo_5">Víznar</div></th>
                <th valign="bottom" id="dias">Frecuencia</th>
            </tr>
        </thead>
        <tbody>
            <tr><td>08:20</td><td>08:30</td><td>08:45</td><td>LMXJV---</td><td>&nbsp;</td></tr>
            <tr><td>09:30</td><td>09:40</td><td>10:00</td><td>LMXJV---</td><td>&nbsp;</td></tr>
            <tr><td>10:30</td><td>10:40</td><td>--</td><td>LMXJV---</td><td>&nbsp;</td></tr>
        </tbody>
    </table>
`
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
        numLinea := scrapper.GetNumLinea(mockHTMLTitulo) // Llamada a función exportada
        expected := "0100"
        if numLinea != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, numLinea)
        }
    })

    t.Run("Happy_Path_Titulo_Con_Espacios_Adicionales", func(t *testing.T) {
        html := `<h2>   0200   -    OTRA RUTA   </h2>`
        numLinea := scrapper.GetNumLinea(html) // Llamada a función exportada
        expected := "0200"
        if numLinea != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, numLinea)
        }
    })

    t.Run("Sad_Path_Sin_Separador", func(t *testing.T) {
        html := `<h2> LINEA SIN GUION </h2>`
        numLinea := scrapper.GetNumLinea(html) // Llamada a función exportada
        expected := "LINEA SIN GUION" 
        if numLinea != expected {
            t.Errorf("Se esperaba '%s', Se obtuvo '%s'", expected, numLinea)
        }
    })
    
    t.Run("Sad_Path_Sin_Titulo_H2", func(t *testing.T) {
        html := `<h1>Otro Titulo</h1>`
        
        // Ahora GetNumLinea devuelve "" si no hay match
        numLinea := scrapper.GetNumLinea(html) 
        expected := ""
        
        if numLinea != expected {
            t.Errorf("Se esperaba una cadena vacía ('%s') al no encontrar H2, Se obtuvo '%s'", expected, numLinea)
        }
    })
}

/////////////////////////////////////////////////////////////////////////////////////
// TESTS PARA PROCESAR TABLAS (HORARIOS)
/////////////////////////////////////////////////////////////////////////////////////

func TestProcesarBloqueTabla(t *testing.T) {
    t.Run("Happy_Path_Tabla_Estandar", func(t *testing.T) {
        datos := scrapper.ProcesarBloqueTabla(mockHTMLTabla) // Llamada a función exportada
        
        if len(datos) != 3 {
            t.Fatalf("Se esperaba 3 filas de datos (horarios), Se obtuvo %d", len(datos))
        }
        if len(datos[0]) != 5 {
            t.Fatalf("Se esperaba 5 columnas de datos, Se obtuvo %d", len(datos[0]))
        }
        
        expected := "08:30" 
        if datos[0][1] != expected { 
            t.Errorf("Valor incorrecto. Se esperaba '%s', Se obtuvo '%s'", expected, datos[0][1])
        }
    })
    
    t.Run("Sad_Path_Tabla_Vacia", func(t *testing.T) {
        html := `<table></table>`
        datos := scrapper.ProcesarBloqueTabla(html) // Llamada a función exportada
        
        if len(datos) != 0 {
            t.Errorf("Se esperaba matriz vacía, Se obtuvo %v", datos)
        }
    })
    
    t.Run("Sad_Path_Filas_con_Celdas_Faltantes", func(t *testing.T) {
        html := `
            <table>
                <tr><td>07:00</td><td>07:05</td></tr>
                <tr><td>08:00</td></tr>
            </table>
        `
        datos := scrapper.ProcesarBloqueTabla(html) // Llamada a función exportada
        
        if len(datos) != 2 {
            t.Fatalf("Debe procesar ambas filas, obtuve %d", len(datos))
        }
        
        if len(datos[1]) != 1 {
            t.Errorf("La fila incompleta no se procesó correctamente. Se esperaba 1 celda, Se obtuvo %d", len(datos[1]))
        }
        if datos[1][0] != "08:00" {
            t.Errorf("Valor incorrecto en la fila incompleta.")
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