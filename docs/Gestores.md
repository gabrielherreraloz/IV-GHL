# Decisiones tomadas en la elección de gestores

## Gestor de Dependencias

Criterio principal: Que esté instalado de forma predeterminada junto con el lenguaje.

Para este proyecto escrito en Go, se ha seleccionado como gestor de dependencias **Go Modules**, cumpliendo el requisito principal de que esté instalado de forma automática junto con el modelo de lenguaje, siendo además el gestor de dependencias oficial de Go. 

Actualemente es la única opción tener en cuenta ya que el resto de opciones no están mantenidas a día de hoy, resultando en su obsolescencia y desuso.

## Gestor de tareas

 ### Criterios de selección
  - **Requisito de instalación**: Es preferible que el gestor de tareas elegido esté preinstalado de forma automática en el sistema para así mejorar su portabilidad entre sistemas.

  - **Lenguaje de ejecución**: Si el lenguaje usado por el gestor coincide con el lenguaje elegido para el proyecto, permite a este trabajar sobre las funciones, clases o librerias del código directamente. Además, al usar el mismo lenguaje, no son necesarios intérpretes externos para este, simplificando la instalación y ejecución.

 ### Opciones posibles
Para la elección del gestor de tareas se presentaban distintas opciones bastante interesantes, de las cuales he querido recalcar tres:

 - [**GoTask**](https://github.com/go-task/task): Cumplimiento de criterios:
   
   - **Requisito de instalación**: No viene instalado de forma predeterminada en ningún sistema operativo por lo que es necesario la instalación del binario task.
   
   - **Lenguaje de ejecución**: Shell, comandos de ordenes del sistema. En la mayoría de sistemas no supone un problema pero en algunos entornos como Windows podría ser necesario instalar un intérprete adicional.
 

 - [**Goyek**](https://github.com/goyek/goyek): Cumplimiento de criterios:
   
   - **Requisito de instalación**: No viene instalado predeterminadamente en ningún sistema operativo, requiere compilación del mismo o instalación.
   
   - **Lenguaje de ejecución**: En Go, lenguaje nativo del proyecto, permite a este trabajar sobre las funciones, clases o librerias del código directamente.
 

 - [**makefile**](https://www.gnu.org/software/make/manual/make.html): Cumplimiento de criterios:

   - **Requisito de instalación**: Está instalado en Linux y MacOS de manera predeterminada por lo que no es necesaria su instalación independiente para su uso si es en uno de estos sistemas, cosa que no se cumple si lo hacemos en Windows.

   - **Lenguaje de ejecución**: Shell, comandos de ordenes del sistema. En la mayoría de sistemas no supone un problema pero en algunos entornos como Windows podría ser necesario instalar un intérprete adicional.

  ### Elección final
  
  Finalemente, en base a los criterios mencionados, se ha decididio usar Makefile puesto que aporta mayor portabilidad al estar instalado de forma predeterminada en muchos sistemas operativos. Si bien es cierto que Goyek se escribe en Go al igual que el proyecto, no justifica su elección por delante del requisito de instalación, que prevalece en importancia al marcar una diferencia mucho más significativa.
 

### Comandos de interés de makefile
  - "make build": compila el pryecto.
  - "make install deps": instala las dependencias necesarias con Go Modules.
  - "make test": ejecuta los tests del proyecto.
  - "make check": comprueba la sintaxis del código fuente.
  - "make clean": elimina los archivos generados por compilaciones pasadas.
