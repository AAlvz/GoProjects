///---Header:
///---*******************Header*******************************/
///---     BY: Alfonso Alvarez Sanchez
///---  Reuse: listaMejorada.go and its methods
///---   Lang: GoLang
///---   Desc: List that carries 2 values per node.
///---   More: Made for Psp course
///---         Program3
///---******************FinHeader****************************/
package listaMejorada

///***Imports
import(
	"math"
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
func (l *Lista) elementos() (elementosX, elementosY []float64){  //mm
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
	valoresX, _ := l.elementos()
	for index, _ := range valoresX{
		sumX = sumX + valoresX[index]
	}///...FinFor
	return
}///...FinFunc

///***Funcion
///... Regresa la sumatoria de la lista de los elementosY
func (l *Lista) SumY() (sumY float64){
	sumY = 0.0
	_, valoresY := l.elementos()
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
	valoresX, valoresY := l.elementos()
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
	valoresX, valoresY := l.elementos()      //mm
	desvEX = math.Pow((l.SumatoriaCuadrada(valoresX)/float64(l.Cantidad-1)), .5)  //mm
	desvEY = math.Pow((l.SumatoriaCuadrada(valoresY)/float64(l.Cantidad-1)), .5)  //mm
	return
}///...FinFunc

///***Funcion
///... Regresa Beta1 de los elemenos de la lista
func (l *Lista) GetBeta1() (beta1 float64){
	mediaX, mediaY := l.Media()
	valoresX, _ := l.elementos()
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
	valoresX, valoresY := l.elementos()
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
