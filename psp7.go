package main

import( "./proyects/listaMejorada"
	"fmt"
	"./proyects/round"
	"./proyects/integra"
	//"math"
)

func main(){
	eps := []float64{130, 650, 99, 150, 128, 302, 95, 945, 368, 961}
	Aam := []float64{186, 699, 132, 272, 291, 331, 199, 1890, 788, 1601}
	Adt := []float64{15.0, 69.9, 6.5, 22.4, 28.4, 65.9, 19.4, 198.7, 38.8, 138.2}
	epsMias := []float64{56.1, 9.2, 26.6, 11.3}
	AamMias := []float64{56, 15, 29, 18}
	AdtMias := []float64{168.0, 54.0, 111.0, 52.8}
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
		//fmt.Print(columna1[index], ",")
	}

	for index, _ := range epsMias{
		l3.Agrega(&listaMejorada.Nodo{ValorX:epsMias[index], ValorY:AamMias[index]})
		l4.Agrega(&listaMejorada.Nodo{ValorX:epsMias[index], ValorY:AdtMias[index]})
	}

	Xk := 386.0


	fmt.Print("Beta0 de Test1 es: ", l1.GetBeta0(), "\n")
	fmt.Print("Beta1 de Test1 es: ", l1.GetBeta1(), "\n")
	fmt.Print("r de Test1 es: ", l1.GetR(), "\n")
	fmt.Print("r cuadrada de Test1 es: ", l1.GetRCuad(), "\n")
	x := l1.GetNuevaX()
	rango := l1.GetRango(Xk)
	Yk := l1.GetYK(Xk)
	fmt.Print("YK de Test1 es: ", Yk, "\n")
	fmt.Print("Cola de Test1 es: ", integra.GetCola(x, float64(l1.Cantidad - 2)), "\n")
	fmt.Print("Rango de Test1 es: ", rango, "\n")
	fmt.Print("UPI de Test1 es: ", Yk + rango, "\n")
	fmt.Print("LPI de Test1 es: ", Yk - rango, "\n")


	fmt.Print("Beta0 de Test2 es: ", l2.GetBeta0(), "\n")
	fmt.Print("Beta1 de Test2 es: ", l2.GetBeta1(), "\n")
	fmt.Print("r de Test2 es: ", l2.GetR(), "\n")
	fmt.Print("r cuadrada de Test2 es: ", l2.GetRCuad(), "\n")
	x = l2.GetNuevaX()
	rango = l2.GetRango(Xk)
	Yk = l2.GetYK(Xk)
	fmt.Print("YK de Test2 es: ", Yk, "\n")
	fmt.Print("Cola de Test2 es: ", integra.GetCola(x, float64(l2.Cantidad - 2)), "\n")
	fmt.Print("Rango de Test2 es: ", rango, "\n")
	fmt.Print("UPI de Test2 es: ", Yk + rango, "\n")
	fmt.Print("LPI de Test2 es: ", Yk - rango, "\n")


	Xk = 32.3


	fmt.Print("Beta0 de Test3 es: ", l3.GetBeta0(), "\n")
	fmt.Print("Beta1 de Test3 es: ", l3.GetBeta1(), "\n")
	fmt.Print("r de Test3 es: ", round.Round(l3.GetR(), 4), "\n")
	fmt.Print("r cuadrada de Test3 es: ", round.Round(l3.GetRCuad(),4), "\n")
	x = l3.GetNuevaX()
	rango = l3.GetRango(Xk)
	Yk = l3.GetYK(Xk)
	fmt.Print("YK de Test3 es: ", Yk, "\n")
	fmt.Print("Cola de Test3 es: ", integra.GetCola(x, float64(l3.Cantidad - 2)), "\n")
	fmt.Print("Rango de Test3 es: ", rango, "\n")
	fmt.Print("UPI de Test3 es: ", Yk + rango, "\n")
	fmt.Print("LPI de Test3 es: ", Yk - rango, "\n")

	fmt.Print("Beta0 de Test4 es: ", round.Round(l4.GetBeta0(), 3), "\n")
	fmt.Print("Beta1 de Test4 es: ", l4.GetBeta1(), "\n")
	fmt.Print("r de Test4 es: ", round.Round(l4.GetR(), 5), "\n")
	fmt.Print("r cuadrada de Test4 es: ", round.Round(l4.GetRCuad(),4), "\n")
	x = l4.GetNuevaX()
	rango = l4.GetRango(Xk)
	Yk = l4.GetYK(Xk)
	fmt.Print("YK de Test4 es: ", Yk, "\n")
	fmt.Print("Cola de Test4 es: ", integra.GetCola(x, float64(l4.Cantidad - 2)), "\n")
	fmt.Print("Rango de Test4 es: ", rango, "\n")
	fmt.Print("UPI de Test4 es: ", Yk + rango, "\n")
	fmt.Print("LPI de Test4 es: ", Yk - rango, "\n")
}
