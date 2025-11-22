# Decisiones tomadas en la elección de gestores

## Gestor de Dependencias

Para este proyecto escrito en Go, se ha seleccionado como gestor de dependencias "Go Modules", es realemnete la unica posibilidad que se plantea ya que "Go Modules" es la única versión realista y actualizada a día de hoy, puesto que el resto de posibilidades están obsoletas y todos proyectos de Go en la actualidad usan esta herramienta, siendo además "Go Modules" el gestor de dependencias oficial de Go.

## Gestor de tareas

 ### Criterios de selección
  - **Requisito de instalación**: Es preferible que el gestor de tareas elegido esté preinstalado de forma automática en el sistema para así facilitar y mejorar su portabilidad entre sistemas.

  - **Lenguaje de ejecución**: Si el lenguaje usado por el gestor coincide con el lenguaje elegido para el proyecto, facilita en gran medida su utilización.

 ### Opciones posibles
Para la elección del gestor de tareas se presentaban distintas opciones bastante interesantes, de las cuales he querido recalcar tres:

 - **GoTask**: Se ha tenido en cuenta por haber sido referenciada por algunos usuarios de Reddit. Cumplimiento de criterios:
   
   - **Requisito de instalación**: No viene instalado de forma predeterminada en ningún sistema operativo por lo que es necesario la instalación del binario task.
   
   - **Lenguaje de ejecución**: Shell, comandos de ordenes del sistema.
 

 - **Goyek**: También muy referenciado y más recomendado que gotaskr en Reddit. Cumplimiento de criterios:
   
   - **Requisito de instalación**: No viene instalado predeterminadamente en ningún sistema operativo, requiere compilación del mismo o instalación.
   
   - **Lenguaje de ejecución**: En Go, lenguaje nativo del proyecto, esto facilita mucho la tareas de programador.
 

 - **makefile**: Uno de los más conocidos y muy trabajado en la carrera de Ingeniería Informática. Cumplimiento de criterios:

   - **Requisito de instalación**: Está instalado en Linux y MacOS de manera predeterminada por lo que no es necesaria su instalación independiente para su uso si es en uno de estos sistemas, cosa que no se cumple si lo hacemos en Windows.

   - **Lenguaje de ejecución**: Shell, comandos de ordenes del sistema.

  ### Elección final
  
  Finalemente, en base a los criterios mencionados, se ha decididio usar Makefile puesto que aporta mayor portabilidad al estar instalado de forma predeterminada en muchos sistemas operativos, además su largo bagaje con años de mantenimiento elimina una posible deuda técnica generada a largo plazo. Si bien es cierto que Goyek se escribe en Go al igual que el proyecto, no justifica su elección por delante del requisito de instalación, que prevalece en importancia al marcar una diferencia mucho más significativa.
 

### Comandos de interés de makefile
  - "make build": compila el pryecto.
  - "make install deps": instala las dependencias necesarias con Go Modules.
  - "make test": ejecuta los tests del proyecto.
  - "make check": comprueba la sintaxis del código fuente.
  - "make clean": elimina los archivos generados por compilaciones pasadas.