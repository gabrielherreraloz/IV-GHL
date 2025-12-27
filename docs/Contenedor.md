# Selección de imágenes para la contenerización
En este documento se analizan las distintas opciones para la elección de la imagen base que se utilizará para la contenerización de este proyecto escrito en Go.

Las distintas opciones oficiales se han extraido del siguiente sitio web: https://hub.docker.com/_/golang/ 

El tamaño de las imagenes se ha comprobado de la siguiente forma: [docker](assets/docker-images.png)

## Imagen base

### Criterios de selección
  **Tamaño mínimo**: Cuanto menos ocupe la imagen, menor será el almacenamiento y el coste de transferencia y descarga.

  **Soporte de gestor de tareas**: Capacidad para instalar y ejecutar make, el gestor de tareas de este proyecto.

  **Estabilidad**: Debe seleccionarse una versión estable para garantizar su correcto funcionamiento, quedan descartadas versiones en fase testing o en RC(Realease candidate). 

### Opciones presentadas
  **Oficial: Basada en Debian** (golang:1.25.5-bookworm): [Repositorio](https://github.com/docker-library/golang/blob/5ac8de688220e63940b9df4c27614898aae9e3fc/1.25/bookworm/Dockerfile)
  Es la última versión estable de Debian, incluye la herramienta Make instalada por defecto, sin embargo, su uso del almacenamiento es mucho mayor, ocupando aproximadamente unos 300MB de forma comprimida, y al expandirse en el disco llega a pesar 1,2GB. 

  **Oficial: Basada en Alpine** (golang:1.25.5-alpine3.23): [Repositorio](https://github.com/docker-library/golang/blob/adf0ff9a10a766d7dc1b394d52544f3a1b1da4e2/1.25/alpine3.23/Dockerfile)
  Cuenta con un tamaño mucho menor que la versión de Debian, pesando aproximadamente unos 64,5MB de forma comprimida, llegando a pesar 329MB al expandirse en el disco, por otro lado, no incluye Make instalada por defecto, aunque puede instalarse en un paso posterior sin apenas aumentar el uso de almacenamiento, ni el tiempo de lanzamiento del contenedor.

  **No ofical: Scratch**: [Documentación](https://hub.docker.com/_/scratch)
  Se trata de una imagen vacía sin nada instalado, es la opción más ligera y segura pero esto dificulta mucho su uso al no tener instalado ni siquiera shell.

### Elección final
Finalmente se ha seleccionado Alpine (golang:1.25.5-alpine3.23) debido a su reducido tamaño con respecto a las otras opciones tenidas en cuenta, además, aunque no cuente con make instalado por defecto, su instalación en su entorno es muy sencilla y ligera.

## Resumen de elecciones
Imagen base seleccionada: Alpine (golang:1.25.5-alpine3.23)