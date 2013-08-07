///---Header:
///---*******************Header*******************************/
///---     BY: Alfonso Alvarez Sanchez
///---  Reuse: func leeArch(nombreArch string){}
///---   Lang: GoLang
///---   Desc: Line counter for files.
///---   More: Made for Psp course
///---         Program2
///---******************FinHeader****************************/
package main

///***Imports
import(
    "io"
    "fmt"
    "bufio"
    "os"
"strings"
     )

///***Funcion
///...Recibe el nombre del archivo y lo lee
func leeArch(nombreArch string){
	archivo, _ := os.Open(nombreArch)
	lector := bufio.NewReader(archivo)
	m := make(map[string]int)
	_, mapa := leeLinea(lector, m)
	fmt.Print("Here we go: \n")
	//fmt.Println(mapa)
	fmt.Print("Lineas Totales: ", mapa["Lineas"], "\n")
	fmt.Print("Lineas Logicas: ", mapa["::::::LINEAS LOGICAS:::::"], "\n")
	fmt.Print("Comentarios: ", mapa["Comments"], "\n")
	fmt.Print("Funciones: ", mapa["///***Funcion"], "\n")
	fmt.Print("Structs: ", mapa["///***Struct"], "\n")
	fmt.Print("Blank lines: ", mapa["Blank Lines"], "\n")
	fmt.Print("Asignaciones: ", mapa["Asignacion"], "\n")
	fmt.Print("Modificadas: ", mapa["Modificadas"], "\n")
}///...FinFuncion

///***Funcion
///... Lee linea por linea del Reader que recibe. El reader trae el archivo
func leeLinea(lector *bufio.Reader, m map[string]int)([]uint8, map[string]int){
        linea, _, err := lector.ReadLine()
        if err != io.EOF{
		///IniciaIf: Evalua Linea en Blank
		if strings.EqualFold(string(linea),""){
			m["Blank Lines"] = m["Blank Lines"]+1
		///...Cuando no es linea en blank
		}else if strings.Contains(string(linea), "///"){
			m[":::::::::::LINEAS ANTI LOGICAS:::::::::"] = m[":::::::::::LINEAS ANTI LOGICAS:::::::::"]+1
			if strings.Contains(string(linea), "///..."){
				m["Comments"] = m["Comments"]+1
			}else if strings.Contains(string(linea), "///---"){
				m["Lineas de Header"] = m["Lineas de Header"]+1
			}else{
				m[string(linea)] = m[string(linea)]+1
			}
		///...Else si no es linea en blanco ni comentario
		}else{
			m["::::::LINEAS LOGICAS:::::"] = m["::::::LINEAS LOGICAS:::::"]+1
			if strings.Contains(string(linea), ":="){ ///...Asignacion
				m["Asignacion"] = m["Asignacion"]+1
			}///...FinIF de asignaciones
		}///...FinIf de comments
		m["Lineas"] = m["Lineas"]+1
		if strings.Contains(string(linea), "//mm"){
			m["Modificadas"] = m["Modificadas"] + 1
		}else if strings.Contains(string(linea), "//dd"){
			m["Eliminadas"] = m["Eliminadas"] + 1
		}
		leeLinea(lector, m)
        }///...FinIf
	//fmt.Println(m)
	m["::::::LINEAS LOGICAS:::::"] = m["Lineas"] - (m[":::::::::::LINEAS ANTI LOGICAS:::::::::"] + m["Blank Lines"])
	return linea, m
}///...FinLeeLinea

///***Funcion
///... Main. Corre el programa enviando el nombre del archivo
func main(){
    leeArch("listaMejorada.go")
}///...FinMain
