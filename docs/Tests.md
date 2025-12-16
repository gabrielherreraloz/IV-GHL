# Herramientas para uso de Tests
En este archivo se especifican los criterios de selección de TestRunner sobre el proyecto y se especifica la elección final y el motivo de esta. También se detalla la elección de la biblioteca de aserciones.


## Biblioteca de aserciones

### Criterios de selección
   - Que no sea necesario instalarlo con dependencias externas, sería ideal si viene integrado junto con la herramienta de testing seleccionada.

### Opciones presentadas
Se han tenido en cuenta las dos herramientas más usadas en proyectos con Go:

   **Estandar de Go**: 
   Instalada de forma nativa junto con el lenguaje.

   **Testify**: https://github.com/stretchr/testify 
   Requiere de su instalación adicional ya que no está integrado en Go.

### Elección final
La elección definitiva es la bibilioteca estandar de Go, ya que ofrece la herramienta directamente integrada en el lenguaje sin necesidad de instalaciones adicionales. 

Cabe destacar, solo como apunte extra sin entrar en los criterior tomados, que mientras muchas bibliotecas hacen uso de assert(), Go promueve el uso de if() y t.Errorf para reportar errores en los tests, lo cual permite el lanzamiento de errores más personalizados y descriptivos para cada ocasión.


## TestRunner

### Criterios de selección
   - Que no sea necesaria su instalación mediante dependencias externas.
   - No debe escribir en el directorio en el que está el código fuente.

### Opciones presentadas
   **Go Test**: https://pkg.go.dev/testing
   La herramienta nativa de Go para testing, está instalado nativamente junto con el propio lenguaje, por lo que no es necesaria su instalación aparte. No escribe en el directorio del código fuente.

   **gotestsum**: https://pkg.go.dev/github.com/IstrateM/gotestsum/pkg/gotestsum 
   Muy parecido a GoTest, pero muestra la información más legible, sin embargo requiere la instalación de alguna herramienta adicional. No escribe en el directorio del código fuente.

   **ginkgo**: https://pkg.go.dev/github.com/onsi/ginkgo/ginkgo
   No está instalado de forma nativa junto con el lenguaje, por lo que inclumple el principal criterio además de tener riesgo de presentar casos en los que sí escriba en el directorio de código fuente, como a la hora de generar archivos de reporte de errores, o en la generación de archivos base de tests.

### Elección final
La herramienta seleccionada como TestRunner es Go Test, ya que cumple a la perfección los dos criterios principales, está integrado de manera nativa con Go por lo que no es necesaria la instalación de dependencias adicionales, y además no escribe en el directorio principal del código fuente. Se ha seleccionado Go Test no solo como Test Runner, sino como herramienta CLI, como se indica más abajo en la documentación, ejecutando esta herramienta ambas funciones, manteniendo la consistencia durante todo el test.


## Herramienta CLI de ejecución de tests
Se usará la propia del lenguaje ya que es la forma más optima de realizar esta tarea.

## Resumen de elecciones
TestRunner: Go Test
Biblioteca de aserciones: Estandar de Go
Herramienta CLI: Go