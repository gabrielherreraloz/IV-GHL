# Decisiones tomadas en la elección de gestores enfocados en el uso de tests

## Gestor de Dependencias
Para este proyecto escrito en Go, se ha seleccionado como gestor de dependencias "Go Modules", el cual ha sido elegido ya que está directamente integrado en las herramientas base de Go y no es necesario instalar nada adicional, lo cual facilita la portabilidad del proyecto. Además, al ser el gestor nativo del lenguaje, lo hace perfectamente compatible con este. Además, Go Modules es sin duda el gestor de dependncias más popular y utilizado en proyectos escritos en Go.

## Gestor de tareas

 ### Criterios de selección
  - **Requisito de instalación**: Es preferible que el gestor de tareas elegido esté preinstalado de forma automática en el sistema para así facilitar y aumentar su escalabilidad.

  - **Lenguaje de ejecución**: Si el lenguaje usado por el gestor coincide con el lenguaje elegido para el proyecto, facilita en gran medida su utilización.

  - **Compilación propia**: Si el gestor de tareas incluye un compilador propio, esto ayuda a la detección de errores antes de la ejecución del código, permite la reutilización directa del código del lenguaje nativo del proyecto, en este caso Go, además se evita la necesidad de utilizar herramientas externas.
  
  - **Integración con el lenguaje**: Una alta integración permite tareas más complejas que hacerlas en shell ccomo interactuar con librerías o estructuras de datos del proyecto.

 ### Opciones posibles
Para la elección del gestor de tareas se presentaban distintas opciones bastante interesantes, de las cuales he querido recalcar tres:

 - **GoTask**: Se ha tenido en cuenta por haber sido referenciada por algunos usuarios de Reddit. Cumplimiento de criterios:
   
   - **Requisito de instalación**: No viene instalado de forma predeterminada en ningún sistema operativo por lo que es necesario la instalación del binario task.
   
   - **Lenguaje de ejecución**: Shell, comandos de ordenes del sistema.

   - **Compilación propia**: No incluida.
   
   - **Integración con el lenguaje**: Es baja puesto que solo realiza llamadas a comandos externos de Go.
 

 - **Goyek**: También muy referenciado y más recomendado que gotaskr en Reddit. Cumplimiento de criterios:
   
   - **Requisito de instalación**: No viene instalado predeterminadamente en ningún sistema operativo, requiere compilación del mismo o instalación.
   
   - **Lenguaje de ejecución**: En Go, lenguaje nativo del proyecto, esto facilita mucho la tareas de programador.

   - **Compilación propia**: Inlcuida, al estar integrado directamente con Go.
   
   - **Integración con el lenguaje**: Muy alta, al estar escrito en el mismo lenguaje, tiene compatibilidad total con el código nativo.
 

 - **makefile**: Uno de los más conocidos y muy trabajado en la carrera de Ingeniería Informática. Cumplimiento de criterios:

   - **Requisito de instalación**: Está instalado en Linux y MacOS de manera predeterminada por lo que no es necesaria su instalación independiente para su uso.

   - **Lenguaje de ejecución**: Shell, comandos de ordenes del sistema.

   - **Compilación propia**: No incluida.

   - **Integración con el lenguaje**: Es baja puesto que solo realiza llamadas a comandos externos de Go.

  ### Elección final
  
  Finalemente, en base a los criterios mencionados, se ha decididio usar Makefile puesto que aporta mayor universalidad al estar instalado de forma predeterminada en muchos sistemas operativos, además su largo bagaje con años de mantenimiento elimina una posible deuda técnica generada a largo plazo. Sin bien es cierto que Goyek aporta mayor integración con Go en programación más avanzada, para las tareas que se van a realizar make creo que es más que suficiente en este caso.
 

### Comandos de interés de makefile
  - "make build": compila el pryecto.
  - "make install deps": instala las dependencias necesarias con Go Modules.
  - "make test": ejecuta los tests del proyecto.
  - "make check": comprueba la sintaxis del código fuente.
  - "make clean": elimina los archivos generados por compilaciones pasadas.