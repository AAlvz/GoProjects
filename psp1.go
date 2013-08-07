package main

import( "./proyects/lista"
	"fmt"
)

func main(){
	columna1 := []float64{160,591,114,229,230,270,128,1657,624,1503}
	columna2 := []float64{15.0,69.9,6.5,22.4,28.4,65.9,19.4,198.7,38.8,138.2}
	//columna3 := []float64{186, 699, 132, 272, 291, 331, 199, 1890, 788, 1601}
	///////////////////
	////   AGREGANDO DATOS A LISTA 1
	//////////////////
	var l lista.Lista
	//fmt.Print(columna1.Length)
	for index, _ := range columna1{
		l.Agrega(&lista.Nodo{Valor:columna1[index]})
		//fmt.Print(columna1[index], ",")
	}

	//l.agrega(&Nodo{Valor:25})
	//l.agrega(&Nodo{Valor:30})
	fmt.Print("La media de C1 es: ", l.Media(), "\n")
	fmt.Print("La DesvEst de C1 es: ", l.DesviacionEst(), "\n")

	///////////////////
	////   AGREGANDO DATOS A LISTA 1
	//////////////////
	var lis lista.Lista
	for index, _ := range columna2{
		lis.Agrega(&lista.Nodo{Valor:columna2[index]})
		//fmt.Print(columna1[index], ",")
	}
	//fmt.Print(lis.elementos())
	//fmt.Print("La media es: ", lis.media())
	fmt.Print("La media de C2 es: ", lis.Media(), "\n")
	fmt.Print("La DesvEst de C2 es: ", lis.DesviacionEst(), "\n")

}
