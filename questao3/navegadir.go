package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func getFiles(initial_dir string) []string {

    lista := make([]string, 0, 10)

    erro := filepath.Walk(initial_dir, func(path string, info os.FileInfo, erro error) error {

        if info.IsDir() {
            return nil
        }

        lista = append(lista, path)
        return nil
      
    })

    if erro != nil {
        fmt.Printf("erro: [%v]\n", erro)
    }

    return lista
}


func main() {

    path := "/Digite o diretorio aqui"

    lista := getFiles(path)

    for _, p := range lista {
		
        fmt.Printf("nome do arquivo: [%s]\n", p)
    }


}