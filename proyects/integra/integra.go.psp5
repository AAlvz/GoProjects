///---Header:
///---*******************Header*******************************/
///---     BY: Alfonso Alvarez Sanchez
///---  Reuse: func Simpson(numeroDeSegmentos, X, gradosDeLibertad float64) float64
///---   Lang: GoLang
///---   Desc: Calcula la integral de cero a X
///---   More: Made for Psp course
///---         Program5
///---******************FinHeader****************************/
package integra

///***Imports
import("math"
//"fmt"
)

///***Funcion
///... calcula la funcion de una x necesaria para obtener simpson
func FunciondeX(x float64, dof float64) (float64){
	dividendo := math.Gamma((dof+1)/2)
	divisor := ( math.Pow( (dof*math.Pi), .5) * math.Gamma(dof/2) )
	multiplicador := math.Pow( (1 + ( math.Pow(x,2) / dof )), ( -(dof+1) / 2))
	//fmt.Print(dividendo/divisor*multiplicador, "+++ \n")
	return (dividendo/divisor)*multiplicador
}///...finFuncion

///***Funcion
///...Para los parametros pares e impares en simpson. Devuelve la sumatoria
func SumaFuncionDeX(numSeg, x, dof float64) (sumaPares, sumaImpares float64){
	sumaPares, sumaImpares = 0.0, 0.0
	index := 1.0
	for index <= numSeg-2.0{
		if int(index) % 2 == 0{
			sumaPares = sumaPares + ( 2 * FunciondeX( ( index * (x/numSeg) ), dof))
			//fmt.Print("Pares:", index, "\n")
		}///...Finif

		index++
	}///...FinFor
	index = 1.0
	for index <= numSeg-1.0{
		if int(index)%2 != 0{
			sumaImpares = sumaImpares + ( 4 * FunciondeX( ( index * (x/numSeg) ), dof))
			//fmt.Print("imares:", index, "\n")
		}///...FinIf
		index++
	}///...Finfor
	return
}///...finfuncion

///***Funcion
///...Simpson. Calcula la integral
func Simpson(numSeg, x, dof float64)(float64){
	W := x/numSeg
	sumatoriaPares, sumatoriaImpares := SumaFuncionDeX(numSeg, x, dof)
	return ( (W/3)*( FunciondeX(0.0,dof) + sumatoriaPares + sumatoriaImpares + FunciondeX(x,dof) ) )
}///...FinFuncion
