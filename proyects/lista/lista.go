//esta Lista esta hecha en go. Simple pero esperemos que funciona

package lista

import(
	"math"
)

type Nodo struct {
    Valor float64
    Siguiente *Nodo
}

type Lista struct {
    Primero *Nodo
    Cantidad int
}

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
	}
}

func (l *Lista) elementos() []float64{
	tmp := l.Primero
	elementos := []float64{}
	elementos = nil
	i := 0
	for i < l.Cantidad{
		elementos = append(elementos, tmp.Valor)
		tmp = tmp.Siguiente
		i ++
	}
	//elementos = append(elementos, tmp.Valor)
	return elementos

}

func (l *Lista) Init() (*Lista){
	l.Primero = nil
	l.Cantidad = 0
	return l
}

func (l *Lista) Media() float64{
	suma := 0.0
	valores := l.elementos()
	for index, _ := range l.elementos(){
		suma = suma + valores[index]
	}
	return suma/float64(l.Cantidad)
}

func (l *Lista) DesviacionEst() float64{
	//fmt.Print(l.elementos())
	sumar := 0.0
	valoresr := l.elementos()
	media := l.Media()
	for index, _ := range l.elementos(){
		cuad := math.Pow((valoresr[index] - media),2)
		sumar = sumar + cuad
	}
	desvEst := math.Pow((sumar/float64(l.Cantidad-1)),.5)
	//fmt.Print(desvEst)
	return desvEst
}
