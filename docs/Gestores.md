# Decisiones tomadas en la elección de gestores enfocados en el uso de tests

## Gestor de Dependencias
Para este proyecto escrito en Go, se ha seleccionado como gestor de dependencias "Go Modules", el cual ha sido elegido ya que viene directamente integrado con Go y no es necesario instalar nada adicional, lo cual facilita la portabilidad del proyecto. Además, al ser el gestor nativo del lenguaje, lo hace perfectamente compatible con este.

## Gestor de tareas
Para la elección del gestor de tareas se presentaban distintas opciones bastante interesantes, de las cuales he querido recalcar tres:

 - gotaskr: Recomendada por algunos usuarios en reddit, es una buena opción en cuanto a sintaxis se refiere, con una estructura sencilla, moderna, y  muy bien integrada en Go. Sin embargo, no viene preinstalado en ningún sistema operativo y su comunidad es menor que las de otras opciones, es por ello por lo que he descartado esta opción.

 - Goyek: También muy recomenada en reddit, se trata de una opción muy utilizada y que presenta una gran portabilidad entre distintos sistemas operativos, sin embargo, he decidido descartarla también debido a su complejidad y gran curva incial de aprendizaje en cuanto a su sintaxis en comparación a las demás opciones.

 - makefile: Finalemente el gestor elegido, muy trabajado a lo largo de la carrera, lo que aporta gran familiaridad con la sintaxis, aunque no venga preinstalado en Windows, si está instalado de forma predeterminada en Linux, y es muy fácil de ejecutar. Además presenta gran compatibilidad con Go y trae tras de sí una amplia comunidad muy consolidada.

    ### Comandos de interés de makefile
        - "make build": compila el pryecto.
        - "make install deps": instala las dependencias necesarias con Go Modules.
        - "make test": ejecuta los tests del proyecto.
        - "make check": comprueba la sintaxis del código fuente.
        - "make clean": elimina los archivos generados por compilaciones pasadas.