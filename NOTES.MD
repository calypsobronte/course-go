# Introducción al Desarrollo en Go: Una Guía Práctica
#### ¿Qué es y por qué usar Go?
Go, también conocido como Golang, es un lenguaje de programación de código abierto desarrollado por Google. Se destaca por ser rápido, eficiente, seguro y fácil de usar. Es una opción popular para construir aplicaciones de servidor, herramientas de línea de comandos y aplicaciones de alto rendimiento.

#### Playground de Go
El Playground de Go es una herramienta en línea que permite a los usuarios escribir, ejecutar y compartir código de Go sin necesidad de instalar nada en su propia computadora. Es una excelente manera de probar el lenguaje y experimentar con código rápidamente.

#### Instalar y Configurar Go
Para comenzar a desarrollar en Go, es necesario instalarlo en tu computadora. Aquí te explicamos cómo:

**Descargar Go:** Ve al sitio oficial de Go y descarga el instalador adecuado para tu sistema operativo.

**Instalar Go:** Sigue las instrucciones del instalador para completar la instalación.

**Configurar GOPATH:** Configura la variable de entorno GOPATH, que es el directorio donde se guardarán tus proyectos y dependencias de Go. En la mayoría de los casos, también necesitarás agregar el binario de Go (go) a tu variable de entorno PATH.

### Hola Mundo con Go
Un primer paso MUY clásico en cualquier lenguaje de programación es escribir un "Hola Mundo". Aquí te muestro cómo hacerlo en Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola Mundo.")
}
```

Guarda este código en un archivo con la extensión `.go`, por ejemplo, `hello-world.go`. Luego, abre una terminal, navega hasta el directorio donde guardaste el archivo y ejecuta el siguiente comando:

```sh
$ go run hello-world.go
$ Hola Mundo.
```

### Paquetes de Terceros
Go facilita la incorporación de paquetes de terceros en tus proyectos. Aquí te explicamos cómo:

- **Agregar Paquetes:** Usa el comando go get para descargar e instalar paquetes de terceros. Por ejemplo:
    ```sh
    $ go get rsc.io/quote
    ```
- **Administrar Dependencias:** Go utiliza dos archivos, `go.mod` y `go.sum`, para gestionar las dependencias del proyecto. El archivo [go.mod](go.mod) define el módulo y las versiones de los paquetes necesarios, mientras que [go.sum](go.sum) almacena los `hashes` de verificación de los paquetes.

## Conclusión
Estos temas proporcionan una introducción útil al desarrollo de aplicaciones en Go y permiten a los programadores comenzar a escribir y ejecutar programas en Go en su propia computadora. Con Go, puedes construir aplicaciones rápidas y eficientes, aprovechar la simplicidad del lenguaje y beneficiarte de una comunidad activa y recursos en línea. ¡Empieza a explorar Go hoy mismo y descubre lo que puedes crear!