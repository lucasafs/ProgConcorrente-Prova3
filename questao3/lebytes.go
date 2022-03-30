package main
 
import (
   "os"
   "log"
   "fmt"
   "io/ioutil"
)
 
func main() {

   file, erro := os.Open("/Digite o diretório aqui")

   if erro != nil {

      log.Fatal(erroo)
   }
 
   data, erro := ioutil.ReadAll(file)

   if erro != nil {

      log.Fatal(erro)
   }

     if (len(data) % 2) == 0 {

        fmt.Println("Número de bytes lidos:", len(data))

     }
     
}