# IV-GHL
Repositorio dedicado a los objetivos de la asignatura Infraestructura Virtual de UGR
Gabriel Herrera Lozano

## Problema a resolver
Como estudiante en la facultad de ingeniería informática en la UGR, tengo serios problemas a la hora de desplazarme de casa a la universidad ya que vivo en Alhendín, un pueblo del área metropolitana, ya que no tengo coche y el transporte público no me viene especialmente cerca. Es por ello por lo que para ir a la facultad y perder el menor tiempo posible en ida y vuelta, tengo que elegir la combinación más óptima. 

Me gustaría tener una solución que incluya los horarios de las paradas de mi zona y la facultad para calcular en cada momento la combinación andando, en autobús, en metro... necesaria para tardar el menor tiempo posible. Toda esta información es pública y se puede extraer de archivos docx alojados en la web.

## Gestor de Dependencias
Para este proyecto escrito en Go, se ha seleccionado como gestor de dependencias "Go Modules", el cual ha sido elegido ya que viene directamente integrado con Go y no es necesario instalar nada adicional, lo cual facilita la portabilidad del proyecto. Además, al ser el gestor nativo del lenguaje, lo hace perfectamente compatible con este.

## Gestor de tareas
Se ha seleccionado como gestor de tareas "Makefile", el cual además de ser muy versátil para el uso con Go, posee una sintaxis sencilla de aprender, clara de entender, y fácil de ejecutar. Además, está instalado por defecto en linux lo que también facilita la portabilidad del proyecto.

## Comandos de interés
    - "make build": compila el pryecto.
    - "make install deps": instala las dependencias necesarias con Go Modules.
    - "make test": ejecuta los tests del proyecto.
    - "make check": comprueba la sintaxis del código fuente.
    - "make clean": elimina los archivos generados por compilaciones pasadas.

## Clave SSH
[Clave SSH](assets/Conf-GIT.png)

## Configuración de GIT
[Configuración de GIT](assets/Clave-SSH.png)

## Historias de usuario
[Historias de usuario](docs/HU.md)

## User Journeys
[User Journeys](docs/UserJourneys.md)

## Milestones
[Milestones](docs/Milestones.md)

