# Selección de imágenes para la contenerización
En este documento se analizan las distintas opciones para la elección de la imagen base que se utilizará para la contenerización de este proyecto escrito en Go.

Generalmente, en entornos de ejecución de proyectos escritos en Go se suelen utilizar dos imágenes, la primera sería la imagen base o de construcción, y la encargada de instalar las dependencias y gestores necesarios, y la segunda llevaría a cabo la tarea de compilado y ejecución. Esta técnica se conoce como "Multi-stage". Sin embargo, en esta fase del proyecto unicamente vamos a lanzar la ejecución de tests, por lo que veo innecesario su uso, y una mejor opción para este caso es usar una única imagen base que realice todas la tareas.
Más información sobre el "Multi-stage": https://docs.docker.com/build/building/multi-stage/

Las distintas opciones oficiales se han extraido del siguiente sitio web: https://hub.docker.com/_/golang/ 

El tamaño de las imagenes se ha comprobado de la siguiente forma: [docker](assets/docker-images.png)

## Imagen de construcción

### Criterios de selección
  **Tamaño mínimo**: Cuanto menos ocupe la imagen, menor será el almacenamiento y el coste de transferencia y descarga.

  **Soporte de gestor de tareas**: Capacidad para instalar y ejecutar make, el gestor de tareas de este proyecto.

  **Fuente fiable**: Si la imagen no es oficial, la fuente de esta debe ser fiable, para que esto sea un criterio objetivo, establecemos el límite de la fiabilidad en tener verificado de publicador en la web de docke: [Sobre el verificado en docker](https://docs.docker.com/docker-hub/repos/manage/trusted-content/dvp-program/)

  **Estabilidad**: La imagen debe de ser una imagen totalmente estable, quedan descartadas versiones experimentales.

### Opciones presentadas
  **Oficial: Basada en Debian**: [Repositorio](https://github.com/docker-library/golang/blob/5ac8de688220e63940b9df4c27614898aae9e3fc/1.25/bookworm/Dockerfile)
  Es la última versión estable de Debian, incluye la herramienta Make instalada por defecto, sin embargo, su uso del almacenamiento es mucho mayor, ocupando aproximadamente unos 300MB de forma comprimida, y al expandirse en el disco llega a pesar 1,2GB.

  **Oficial: Basada en Alpine**: [Repositorio](https://github.com/docker-library/golang/blob/adf0ff9a10a766d7dc1b394d52544f3a1b1da4e2/1.25/alpine3.23/Dockerfile)
  Cuenta con un tamaño mucho menor que la versión de Debian, pesando aproximadamente unos 64,5MB de forma comprimida, llegando a pesar 329MB al expandirse en el disco, por otro lado, no incluye Make instalada por defecto, aunque puede instalarse en un paso posterior sin apenas aumentar el uso de almacenamiento, ni el tiempo de lanzamiento del contenedor. Sin embargo, en la documentación oficial de Go, se incluye la siguiente aclariación: "Esta variante es altamente experimental y no cuenta con el respaldo oficial del proyecto Go." lo cual descarta automaticamente ya que también especifica que no estan disponibles muchas herramientas básicas en esta versión.

  **Oficial: Scratch**: [Documentación](https://hub.docker.com/_/scratch)
  Se trata de una imagen vacía sin nada instalado, es la opción más ligera y segura, sería una excelente opción en caso de usar la técnica de multi-stage, anteriormente mencionada, pero en este caso, una imagen tan vacía no es la mejor opción, ya que al no incluir make, scratch no tiene ninguna forma de instalarlo pues solo trabaja con binarios directamente, lo cual descarta esta opción.

  **No Oficial: GoogleDistroless**: [Documentación](https://github.com/GoogleContainerTools/distroless)
  Es una opción muy parecida a Scratch pero algo menos vacía, tiene un peso de 830kB de manera comprimida, con un peso final de 6,23MB. Es ultrasegura y muy recomendada para entornos que utilizan multistage, pero en este caso no es la opción más recomendada por los mismos motivos que se han comentado con Scratch.

  **No Oficial: Bitnami**: [Documentación](https://hub.docker.com/r/bitnami/golang)
  Se trata de una opción no oficial pero mantenida por la ya conocida empresa VMware, que cuenta con verificado de publicador en la web oficial de docker, lo que aporta cierta fiabilidad a esta imagen a pesar de ello. Se trata de una versión reducida de debian y muy segura, ya que no utiliza root por defecto. Al igual que alpine, no trae make instalado por defecto pero se puede instalar de forma manual, e incluso de manera más simplificada que en la imagen estandar de debian. Su peso de forma comprimida es de 231MB, con un tamaño final de 991MB.

### Elección final
Finalmente se ha seleccionado Bitnami debido a su reducido tamaño con respecto a las otras opciones tenidas en cuenta, además, aunque no cuente con make instalado por defecto, su instalación en su entorno es muy sencilla y ligera, además de ser una versión estable a diferencia de, por ejemplo, alpine, que es más experimental.

## Resumen de elecciones
Imagen base seleccionada: bitnami/golang