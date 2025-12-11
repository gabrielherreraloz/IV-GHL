# Herramientas para uso de Tests
En este archivo se especifican los criterios de selección de TestRunner sobre el proyecto y se especifica la elección final y el motivo de esta. También se detalla la elección de la biblioteca de aserciones.


## TestRunner

### Criterios de selección
   - Que no sea necesaria su instalación mediante dependencias externas.
   - Velocidad de ejecución máxima.
   - No debe escribir en el directorio en el que está el código fuente.

### Opciones presentadas
   **Go Test**: https://pkg.go.dev/testing
   La herramienta nativa de Go para testing, está instalado nativamente junto con el propio lenguaje, por lo que no es necesaria su instalación aparte. Al estar integrado de forma nativa, su velocidad de ejecución es máxima. No escribe en el directorio del código fuente.

   **gotestsum**: https://pkg.go.dev/github.com/IstrateM/gotestsum/pkg/gotestsum 
   Muy parecido a GoTest, pero muestra la información más legible, sin embargo requiere la instalación de alguna herramienta adicional. No escribe en el directorio del código fuente.

   **ginkgo**: https://pkg.go.dev/github.com/onsi/ginkgo/ginkgo
   Al ser más pesado que las otras opciones tiene un mayor tiempo de ejecución, además de no estar instalado de forma nativa junto con el lenguaje, por lo que inclumple los dos principales criterios tenidos en cuenta. Tiene riesgo de presentar casos en los que sí escribe en el directorio de código fuente.

### Elección final
La herramienta seleccionada como TestRunner es Go Test, ya que cumple a la perfección los dos criterios principales, está integrado de manera nativa con Go por lo que no es necesaria la instalación de dependencias adicionales, y esto a su vez, le aporta una gran velocidad de ejecución, mejor que las demás opciones.


## Biblioteca de aserciones

### Criterios de selección
   - Que no sea necesario instalarlo con dependencias externas, sería ideal si viene integrado junto con la herramienta de testing seleccionada.
   - Sintaxis más simplificada, aunque predomina el criterio anterior por su importancia.

### Opciones presentadas
   **Estandar de Go**: 
   Instalada nativamente junto con el lenguaje, sin embargo, su sintaxis es objetivamente más elaborada que las demás opciones.

   **Testify**: https://github.com/stretchr/testify 
   Cuenta con una sintaxis mucho más simplificada que la estandar de Go, es decir, se pueden expresar ordenes más complejas con menos código, por el contrario, requiere de su instalación adicional ya que no está integrado en Go.

### Elección final
La elección definitiva es la bibilioteca estandar de Go, ya que ofrece la herramienta directamente integrada en el lenguaje sin necesidad de instalaciones adicionales, y sacrificamos la simplicidad de la sintaxis ya que no es un factor tan determinante en el producto final como lo es el criterio anteriormente mencionado.

## Herramienta CLI de ejecución de tests
Se usará la propia del lenguaje ya que es la forma más optima de realizar esta tarea.

## Resumen de elecciones
TestRunner: Go Test
Biblioteca de aserciones: Estandar de Go
Herramienta CLI: Go