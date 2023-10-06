package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Abre el fichero de texto
    file, err := os.Open("texto.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Crea un lector para leer el fichero
    reader := bufio.NewReader(file)

    // Inicializa una variable para almacenar las palabras
    palabras := []string{}

    // Lee el fichero, una palabra a la vez
    for {
        palabra, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println(err)
            return
        }

        // Elimina espacios en blanco de la palabra
        palabra = strings.TrimSpace(palabra)

        // Si la palabra tiene un tamaño mayor o igual a x, la añade a la lista
        if len(palabra) >= x {
            palabras = append(palabras, palabra)
        }
    }

    // Imprime la lista de palabras
    for _, palabra := range palabras {
        fmt.Println(palabra)
    }
}
