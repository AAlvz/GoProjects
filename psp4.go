package main

import( "./proyects/listaMejorada"
	"fmt"
	"./proyects/round"
	"math"
)

func main(){
	eps := []float64{130, 650, 99, 150, 128, 302, 95, 945, 368, 961}
	Pam := []float64{163, 765, 141, 166, 137, 355, 136, 1206, 433, 1130}
	Aam := []float64{186, 699, 132, 272, 291, 331, 199, 1890, 788, 1601}
	Adt := []float64{15.0, 69.9, 6.5, 22.4, 28.4, 65.9, 19.4, 198.7, 38.8, 138.2}
	Loc := []float64{18, 18, 25, 31, 37, 82, 82, 87, 89, 230, 85, 87, 558}
	NMe := []float64{3, 3, 3, 3, 3, 5, 4, 4, 4, 10, 3, 3, 10}
	cha := []float64{7, 12, 10, 12, 10, 12, 12, 12, 12, 8, 8, 8, 20, 14, 18, 12}
	///////////////////
	////   AGREGANDO DATOS A LISTAS
	//////////////////
	var l1 listaMejorada.Lista
	var l2 listaMejorada.Lista
	var l3 listaMejorada.Lista
	var l4 listaMejorada.Lista
	//fmt.Print(columna1.Length)
	for index, _ := range eps{
		l1.Agrega(&listaMejorada.Nodo{ValorX:eps[index], ValorY:Aam[index]})
		l2.Agrega(&listaMejorada.Nodo{ValorX:eps[index], ValorY:Adt[index]})
		l3.Agrega(&listaMejorada.Nodo{ValorX:Pam[index], ValorY:Aam[index]})
		l4.Agrega(&listaMejorada.Nodo{ValorX:Pam[index], ValorY:Adt[index]})
		//fmt.Print(columna1[index], ",")
	}

	fmt.Print("Beta0 de Test1 es: ", l1.GetBeta0(), "\n")
	fmt.Print("Beta1 de Test1 es: ", l1.GetBeta1(), "\n")
	fmt.Print("r de Test1 es: ", round.Round(l1.GetR(), 4), "\n")
	fmt.Print("r cuadrada de Test1 es: ", round.Round(l1.GetRCuad(),4), "\n")
	fmt.Print("YK de Test1 es: ", l1.GetYK(386.0), "\n")

	fmt.Print("Beta0 de Test2 es: ", l2.GetBeta0(), "\n")
	fmt.Print("Beta1 de Test2 es: ", l2.GetBeta1(), "\n")
	fmt.Print("r de Test2 es: ", round.Round(l2.GetR(),4), "\n")
	fmt.Print("r cuadrada de Test2 es: ", round.Round(l2.GetRCuad(),4), "\n")
	fmt.Print("YK de Test2 es: ", l2.GetYK(386.0), "\n")

	fmt.Print("Beta0 de Test3 es: ", l3.GetBeta0(), "\n")
	fmt.Print("Beta1 de Test3 es: ", l3.GetBeta1(), "\n")
	fmt.Print("r de Test3 es: ", round.Round(l3.GetR(), 4), "\n")
	fmt.Print("r cuadrada de Test3 es: ", round.Round(l3.GetRCuad(),4), "\n")
	fmt.Print("YK de Test3 es: ", l3.GetYK(386.0), "\n")

	fmt.Print("Beta0 de Test4 es: ", round.Round(l4.GetBeta0(), 3), "\n")
	fmt.Print("Beta1 de Test4 es: ", l4.GetBeta1(), "\n")
	fmt.Print("r de Test4 es: ", round.Round(l4.GetR(), 5), "\n")
	fmt.Print("r cuadrada de Test4 es: ", round.Round(l4.GetRCuad(),4), "\n")
	fmt.Print("YK de Test4 es: ", round.Round(l4.GetYK(386.0), 4), "\n")

	//Agregando valores a lista
	var l5, l6 listaMejorada.Lista

	for index, _ := range Loc{
		l5.Agrega(&listaMejorada.Nodo{ValorX:(Loc[index]/NMe[index]), ValorY:cha[index]})
	}

	for index, _ := range cha{
		l6.Agrega(&listaMejorada.Nodo{ValorX:(cha[index]), ValorY:cha[index]})
	}

	valoresX, _ := l5.Elementos()
	avg := l5.AvgLog(valoresX)
	sigma := math.Pow(l5.VarLog(valoresX), .5)
	VS := math.Exp(avg - (2*sigma))
	S := math.Exp(avg - sigma)
	M := math.Exp(avg)
	L := math.Exp(avg + sigma)
	VL := math.Exp(avg + (2*sigma))

	fmt.Print("VS: ", VS, "\n")
	fmt.Print("S: ", S, "\n")
	fmt.Print("M: ", M, "\n")
	fmt.Print("L: ", L, "\n")
	fmt.Print("VL: ", VL, "\n")
	fmt.Print("sigma: ", sigma, "\n")

	valoresX, _ = l6.Elementos()
	avg = l6.AvgLog(valoresX)
	sigma = math.Pow(l6.VarLog(valoresX), .5)
	VS = math.Exp(avg - (2*sigma))
	S = math.Exp(avg - sigma)
	M = math.Exp(avg)
	L = math.Exp(avg + sigma)
	VL = math.Exp(avg + (2*sigma))

	fmt.Print("VS: ", VS, "\n")
	fmt.Print("S: ", S, "\n")
	fmt.Print("M: ", M, "\n")
	fmt.Print("L: ", L, "\n")
	fmt.Print("VL: ", VL, "\n")
	fmt.Print("sigma: ", sigma, "\n")
}
