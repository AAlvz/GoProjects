///---Header:
///---*******************Header*******************************/
///---     BY: Alfonso Alvarez Sanchez
///---  Reuse: listaMejorada.go and its methods
///---   Lang: GoLang
///---   Desc: List that carries 2 values per node. We can calculate proxy size til here.
///---   More: Made for Psp course
///---         Program4
///---******************FinHeader****************************/
package listaMejorada

///***Imports
import(
	"math"
	"../integra"
)///...FinImport

///***Struct
type Nodo struct {
    ValorX float64 //mm
    ValorY float64 //mm
    Siguiente *Nodo
}///...FinStruct

///***Struct
type Lista struct {
    Primero *Nodo
    Cantidad float64
}///...FinStruct

///***Funcion
///...Agrega nuevo nodo a lista
func (l *Lista) Agrega(n *Nodo){
	if l.Primero == nil{
		l.Primero = n
		l.Cantidad = l.Cantidad + 1

	}else if l.Cantidad == 1{
		n.Siguiente = l.Primero
		l.Primero = n
		l.Cantidad = l.Cantidad + 1
	}else{
		tmp := l.Primero
		n.Siguiente = tmp
		l.Primero = n
		l.Cantidad = l.Cantidad + 1
	}///...FinIf
}///...FinFuncion

///***Funcion
///...Regresa los elementos de la lista
func (l *Lista) Elementos() (elementosX, elementosY []float64){  //mm
	tmp := l.Primero
	elementosX, elementosY = []float64{}, []float64{}  //mm
	elementosX, elementosY = nil, nil   //mm
	i := 0.0
	for i < l.Cantidad{
		elementosX = append(elementosX, tmp.ValorX) //mm
		elementosY = append(elementosY, tmp.ValorY) //mm
		tmp = tmp.Siguiente
		i ++
	}///...FinFor
	return

}///...FinFunc

///***Funcion
///... Inicializa la lista vacia
func (l *Lista) Init() (*Lista){
	l.Primero = nil
	l.Cantidad = 0
	return l
}///...FinFunc

///***Funcion
///... Regresa la sumatoria de la lista de los elementosX
func (l *Lista) SumX() (sumX float64){
	sumX = 0.0
	valoresX, _ := l.Elementos()
	for index, _ := range valoresX{
		sumX = sumX + valoresX[index]
	}///...FinFor
	return
}///...FinFunc

///***Funcion
///... Regresa la sumatoria de la lista de los elementosY
func (l *Lista) SumY() (sumY float64){
	sumY = 0.0
	_, valoresY := l.Elementos()
	for index, _ := range valoresY{
		sumY = sumY + valoresY[index]
	}///...FinFor
	return
}///...FinFor

///***Funcion
///... Regresa la sumatoria de la lista de los elementos recibidos al cuadrado cada uno
func (l *Lista) SumatoriaCuadrada(elementos []float64) float64{
	sum := 0.0
	for index, _ := range elementos{
		sum = sum + math.Pow(elementos[index], 2)
	}///...FinFor
	return sum
}///...FinFunc

///***Funcion
///... Regresa la sumatoria de la lista de los elementosX por los elementosY al cuadrado
func (l *Lista) SumXporY() float64{
	valoresX, valoresY := l.Elementos()
	sum := 0.0
	for index, _ := range valoresX{
		sum = sum + (valoresX[index] * valoresY[index])
	}///...FinFunc
	return sum
}///...FinFunc

///***Funcion
///... Regresa la media de los elementos en x y en Y
func (l *Lista) Media() (mediaX, mediaY float64){
	mediaX = l.SumX() / float64(l.Cantidad)   //mm
	mediaY = l.SumY() / float64(l.Cantidad)   //mm
	return
}///...FinFunc

///***Funcion
///... Regresa la desv estandar de los elementos en X y en Y
func (l *Lista) DesviacionEst() (desvEX, desvEY float64){
	valoresX, valoresY := l.Elementos()      //mm
	desvEX = math.Pow((l.SumatoriaCuadrada(valoresX)/float64(l.Cantidad-1)), .5)  //mm
	desvEY = math.Pow((l.SumatoriaCuadrada(valoresY)/float64(l.Cantidad-1)), .5)  //mm
	return
}///...FinFunc

///***Funcion
///... Regresa Beta1 de los elemenos de la lista
func (l *Lista) GetBeta1() (beta1 float64){
	mediaX, mediaY := l.Media()
	valoresX, _ := l.Elementos()
	dividendo := l.SumXporY() - (l.Cantidad * mediaX * mediaY)
	divisor := l.SumatoriaCuadrada(valoresX) - (l.Cantidad * math.Pow(mediaX, 2))
	beta1 = dividendo/divisor
	return
}///...FinFunc

///***Funcion
///... Regresa Beta0 de los elemenos de la lista
func (l *Lista) GetBeta0() (beta0 float64){
	mediaX, mediaY := l.Media()
	beta0 = mediaY - (l.GetBeta1() * mediaX)
	return
}///...FinFunc

///***Funcion
///... Regresa R de los elemenos de la lista
func (l *Lista) GetR() (r float64){
	valoresX, valoresY := l.Elementos()
	dividendo := (l.Cantidad * l.SumXporY()) - (l.SumX() * l.SumY())
	divisor := math.Pow(((l.Cantidad * l.SumatoriaCuadrada(valoresX)) - (l.SumX() * l.SumX())) * ((l.Cantidad * l.SumatoriaCuadrada(valoresY)) - (l.SumY() * l.SumY())), .5)
	r = dividendo/divisor
	return
}///...FinFunc

///***Funcion
///... Regresa R cuadrada de los elementos de la lista
func (l *Lista) GetRCuad() float64{
	return math.Pow(l.GetR(),2)
}///...FinFunc

///***Funcion
///... Regresa YK de los elemenos de la lista de acuerdo al XK que se tenga
func (l *Lista) GetYK(XK float64) float64{
	return l.GetBeta0() + (l.GetBeta1() * XK)
}///...FinFunc

///***Funcion
///... regresa la sumatoria logaritmica de los valores de la lista.
///...  Usa LOCs de proxy e items por cada uno, por lo que se deben dividir
func (l *Lista) sumLog(valoresX []float64) float64{
	sum := 0.0
	for index, _ := range valoresX{
		sum = sum + (math.Log(valoresX[index]))
	}///...FincFor
	return sum
}///...FinFunc

///***Funcion
///... Regresa el promedio de la sumatoria de logs
func (l *Lista) AvgLog(valoresX []float64) float64{
	return l.sumLog(valoresX) / l.Cantidad
}///...FinFunc

///***Funcion
///... regresa la varianza de los valores para calcular proxy size
func (l *Lista) VarLog(valoresX []float64) float64{
	sum := 0.0
	avg := l.AvgLog(valoresX)
	for index, _ := range valoresX{
		sum = sum + math.Pow((math.Log(valoresX[index]) - avg), 2)
	}///...FinFor
	return sum/(l.Cantidad -1)
}///...FinFunc

///***Funcion
///... Regresa el nuevo valor de X a obtener.
func (l *Lista) GetNuevaX() float64{
	r := l.GetR()
//	arriba := r * math.Pow( l.Cantidad - 2, .5)
//	abajo := math.Pow( 1 - math.Pow(r, 2), .5)
//	return arriba/abajo
	return ( (math.Abs(r) * math.Pow( ( l.Cantidad - 2.0 ), .5 ) ) / (math.Pow(  1 - math.Pow(r,2) , .5)) )
}///...FunFuncion

///***Funcion
///...obtiene el rango de prediccino de intervalo
func(l *Lista) GetRango(Xk float64) float64{
	n := float64(l.Cantidad)
	valX, _ := l.Elementos()
	t := integra.AdivinaPconX(10.0, .35, n-2.0)
	sigma := l.GetNuevoSigma()
	Xavg, _ := l.Media()
	sum := 0.0
	for index, _ := range valX{
		sum = sum + ( math.Pow( valX[index] - Xavg, 2 ))
	}
	tercero := math.Pow( ( 1 + (1.0/n) + ( (math.Pow((Xk - Xavg), 2)) / sum) ), .5)
	return (t * sigma * tercero)
}///...FinFunc

///***Funcion
///... Obtiene el nuevo sigma
func(l *Lista) GetNuevoSigma() float64{
	sum := 0.0
	B0 := l.GetBeta0()
	B1 := l.GetBeta1()
	n := float64(l.Cantidad)
	valX, valY := l.Elementos()
	for index, _ := range valX{
		sum = sum + ( math.Pow( (valY[index] - B0 - (B1 * valX[index]) ) , 2) )
	}
	prim := 1.0 / (n-2.0)
	sigma := math.Pow( prim * sum , .5)
	return sigma
	//return math.Pow( ((1.0/n-2.0) * sum ), .5 )
}
