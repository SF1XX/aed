/**
*! Martes 26/03
*/

1. Dada una secuencia de caracteres que posee las ventas realizadas de un
determinado producto. como máximo se pueden comprar 9. La información de
las ventas en la secuencia sigue el siguiente formato:
fechaVenta CantidadVendida
Ejemplo:
010220236210320242150220247230220241

La fecha siempre tiene 8 caracteres y un solo carácter para las unidades vendidas, no hay una marca
que separe cada venta.

FILA 1
- Generar una secuencia de salida de entero, con todas las cantidades vendidas, en
meses pares.


Accion Fila1 Es
Ambiente
	sec: secuencia de caracteres
	sal: secuencia de caracteres
	v: caracter
	v2: caracter
	resguardo: caracter
	cont: entero
	meses_pares = {"0","2","4","6","8"}

Proceso
	ARR(sec)
	AVZ(sec,v)
	CREAR(sal)
	Mientras NFDS(sec) Hacer
		Para I := 1 hasta 3 Hacer
			AVZ(sec,v)
		FinPara
		resguardo := v
		Para I := 1 hasta 4 Hacer
			AVZ(sec,v)
		FinPara	
		Si resguardo En meses_pares Entonces
			Segun v Hacer
				= "0" : GRABAR(sal,0)
				= "1" : GRABAR(sal,1)
				= "2" : GRABAR(sal,2)
				= "3" : GRABAR(sal,3)
				= "4" : GRABAR(sal,4)
				= "5" : GRABAR(sal,5)
				= "6" : GRABAR(sal,6)
				= "7" : GRABAR(sal,7)
				= "8" : GRABAR(sal,8)
				= "9" : GRABAR(sal,8)
			FinSegun
		FinSi		
		AVZ(sec,v)
	FinMientras
	CERRAR(sec)
	CERRAR(sal)
FinAccion	




