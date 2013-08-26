///---Header:
///---*******************Header*******************************/
///---     BY: Alfonso Alvarez Sanchez
///---  Reuse: func Gauss(a [][]float64, r []float64) []float64{
///---   Lang: GoLang
///---   Desc: Calcula los valores de ecuaciones con el metodo de gauss
///---   More: Made for Psp course
///---         Program8
///---******************FinHeader****************************/
package gauss

///***Imports
import("math")

///***Funcion
///...Calcula el valor de las incognitas
func Gauss(a [][]float64, r []float64) []float64{
	var t, s float64
	var i, l, j, k, m, n int
	n = len(r) - 1
	m = n + 1
	for l = 0; l <= n - 1; l++ {
		j = l
		for k = l + 1; k <= n; k++ {
			if (!(math.Abs(a[j][l]) >= math.Abs(a[k][l]))){
				j = k
			}///...finIf
		}///...FinFor
		if (!(j == l)){
			for i = 0; i <= m; i++ {
				t = a[l][i]
				a[l][i] = a[j][i]
				a[j][i] = t
			}///...FinFor
		}///...FinIf
		for j = l + 1; j <= n; j++{
			t = (a[j][l] / a[l][l])
			for i = 0; i <= m; i++ {
				a[j][i] -= t * a[l][i]
			}///...finFor
		}///...FinFor
	}///...finFor
	r[n] = a[n][m] / a[n][n]
	for i = 0; i <= n - 1; i++ {
		j = n - i - 1
		s = 0
		for l = 0; l <= i; l++ {
			k = j + l + 1
			s += a[j][k] * r[k]
		}///...FInFor
		r[j] = ((a[j][m] - s) / a[j][j])
	}///...FinFor
	return r
}///...FinFunc


///***Funcion
///...para llenar la matriz con los valores de 4 slices
func LlenarMatriz(DatosW, DatosX, DatosY, DatosZ []float64) [][]float64{
	n := float64(len(DatosW))
	Wi, Xi, Yi, Zi := 0.0, 0.0, 0.0, 0.0
	for index, _ := range DatosW{
		Wi = Wi + DatosW[index]
		Xi = Xi + DatosX[index]
		Yi = Yi + DatosY[index]
		Zi = Zi + DatosZ[index]
	}///...FinFor
	Wi2, Xi2, Yi2 := 0.0, 0.0, 0.0
	for _, d := range DatosW{
		Wi2 = Wi2 + math.Pow(d, 2)
	}///...finFor
	for _, d := range DatosX{
		Xi2 = Xi2 + math.Pow(d, 2)
	}///...FinFor
	for _, d := range  DatosY{
		Yi2 = Yi2 + math.Pow(d, 2)
	}///...FinFor
	WiXi, WiYi, WiZi, XiYi, XiZi, YiZi := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	for i := 0; i < int(n); i++ {
		WiXi = WiXi + (DatosW[i] * DatosX[i])
		WiYi = WiYi + (DatosW[i] * DatosY[i])
		WiZi = WiZi + (DatosW[i] * DatosZ[i])
		XiYi = XiYi + (DatosX[i] * DatosY[i])
		XiZi = XiZi + (DatosX[i] * DatosZ[i])
		YiZi = YiZi + (DatosY[i] * DatosZ[i])
	}///...FinFor
	a := make([][]float64, 4)
	for i := range a{
		a[i] = make([]float64, 5)
	}///...FinFor
	a[0] = []float64{n, Wi, Xi, Yi, Zi}
	a[1] = []float64{Wi, Wi2, WiXi, WiYi, WiZi}
	a[2] = []float64{Xi, WiXi, Xi2, XiYi, XiZi}
	a[3] = []float64{Yi, WiYi, XiYi, Yi2, YiZi}
	return a
}///...FinFunc
