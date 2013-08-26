package main

import( "./proyects/gauss"
	"./proyects/listaMejorada"
	"./proyects/integra"
	"math"
	"fmt"
)

func main(){

	w2 := []float64{345.0, 168.0, 94.0, 187.0, 621.0, 255.0}
	x2 := []float64{65.0, 18.0, 0.0, 185.0, 87.0, 0.0}
	y2 := []float64{23.0, 18.0, 0.0, 98.0, 10.0, 0.0}
	z2 := []float64{31.4, 14.6, 6.4, 28.3, 42.1, 15.3}

	w := []float64{1142.0, 863.0, 1065.0, 554.0, 983.0, 256.0}
	x := []float64{1060.0, 995.0, 3205.0, 120.0, 2896.0, 485.0}
	y := []float64{325.0, 98.0, 23.0, 0.0, 120.0, 88.0}
	z := []float64{201.0, 98.0, 162.0, 54.0, 138.0, 61.0}
	wk := 650.0
	xk := 3000.0
	yk := 155.0
	wk2 := 185.0
	xk2 := 150.0
	yk2 := 45.0

	/****************   Lista **************************/
	var lwx listaMejorada.Lista
	var lyz listaMejorada.Lista

	var lwx2 listaMejorada.Lista
	var lyz2 listaMejorada.Lista
	//fmt.Print(columna1.Length)
	for index, _ := range w{
		lwx.Agrega(&listaMejorada.Nodo{ValorX:w[index], ValorY:x[index]})
		lyz.Agrega(&listaMejorada.Nodo{ValorX:y[index], ValorY:z[index]})
	}
	for index, _ := range w{
		lwx2.Agrega(&listaMejorada.Nodo{ValorX:w2[index], ValorY:x2[index]})
		lyz2.Agrega(&listaMejorada.Nodo{ValorX:y2[index], ValorY:z2[index]})
	}

	matriz := gauss.LlenarMatriz(w, x, y, z)
	matriz2 := gauss.LlenarMatriz(w2, x2, y2, z2)
	fmt.Print("Aqui se muestra la matriz antes de betas \n", matriz, "\n")
	r := make([]float64,4)
	fmt.Print("Betas: \n", gauss.Gauss(matriz, r), "\n")
	te := GetZ(wk, xk, yk, r[0], r[1], r[2], r[3])
	ra := GetRango(w, x, y, z, wk, xk, yk)
	fmt.Print("Tiempo estimado: \n", te, "\n")
	fmt.Print("Rango: \n", ra, "\n")
	fmt.Print("UPI: \n", te + ra, "\n")
	fmt.Print("LPI: \n", te - ra, "\n")


	fmt.Print("Aqui se muestra la matriz antes de betas \n", matriz2, "\n")
	r2 := make([]float64,4)
	fmt.Print("Betas: \n", gauss.Gauss(matriz2, r2), "\n")
	te2 := GetZ(wk2, xk2, yk2, r2[0], r2[1], r2[2], r2[3])
	ra2 := GetRango(w2, x2, y2, z2, wk2, xk2, yk2)
	fmt.Print("Tiempo estimado: \n", te2, "\n")
	fmt.Print("Rango: \n", ra2, "\n")
	fmt.Print("UPI: \n", te2 + ra2, "\n")
	fmt.Print("LPI: \n", te2 - ra2, "\n")

}

func GetZ(Wk, Xk, Yk, B0, B1, B2, B3 float64) float64{
	return B0 + (Wk * B1) + (Xk * B2) + (Yk * B3)
}

func GetRango(w, x, y, z []float64, Wk, Xk, Yk float64) float64{
	var lwx listaMejorada.Lista
	var lyz listaMejorada.Lista
	//fmt.Print(columna1.Length)
	for index, _ := range w{
		lwx.Agrega(&listaMejorada.Nodo{ValorX:w[index], ValorY:x[index]})
		lyz.Agrega(&listaMejorada.Nodo{ValorX:y[index], ValorY:z[index]})
	}
	n := len(w)
	t := integra.AdivinaPconX(10.0, .35, float64(n)-4.0)
	fmt.Print("Aqui la T: ", t, "\n")
	matriz := gauss.LlenarMatriz(w, x, y, z)
	r := make([]float64, 4)
	gauss.Gauss(matriz,r)
	sigma := GetSigma(w, x, y, z, r[0], r[1], r[2], r[3] )
	fmt.Print("Sigma: ", sigma, "\n")
	wab, xab, yab := 0.0, 0.0, 0.0
	Wavg, Xavg := lwx.Media()
	Yavg, _ := lyz.Media()
	for index, _ := range w{
		wab = wab + math.Pow( (w[index] - Wavg), 2)
		xab = xab + math.Pow( (x[index] - Xavg), 2)
		yab = yab + math.Pow( (y[index] - Yavg), 2)
	}
	war := math.Pow( (Wk - Wavg), 2)
	xar := math.Pow( (Xk - Xavg), 2)
	yar := math.Pow( (Yk - Yavg), 2)
	fmt.Print("n :", n, "\n")
	tercero := math.Pow(1 + (1.0/float64(n)) + ( war/wab ) + ( xar/xab ) + ( yar/yab ), .5)
	fmt.Print("Rango: \n", t * sigma * tercero, "\n")
	return t * sigma * tercero
}

func GetSigma(w, x, y, z []float64, B0, B1, B2, B3 float64) float64{
	sum := 0.0
	for index, _ := range w{
		sum = sum + math.Pow( ( z[index] - B0 - (B1 * w[index]) - (B2 * x[index]) - (B3 * y[index]) ), 2)
	}
	return math.Pow((1/(float64(len(w)) - 4.0)) * sum, .5)
}
