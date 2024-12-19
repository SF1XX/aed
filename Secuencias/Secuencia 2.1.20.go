/**
*! CONSIGNA
*/
Se dispone de dos secuencias texto formadas por oraciones bimembres (sujeto y predicado, separados por 
comas ‘,’). Se desea una secuencia texto de salida formada por oraciones bimembres, de la siguiente 
forma:

El sujeto, de la secuencia 1, y el predicado, de la secuencia 2.

Al finalizar informar cuantas oraciones tiene cada secuencia.

Accion 2.2.20 Es
Ambiente
	sec1: secuencia de caracteres
	sec2: secuencia de caracteres
	sal: secuencia de caracteres
	v1: caracter
	v2: caracter

	cont_oracion1: entero
	cont_oracion2: entero

Proceso
	ARR(sec1)
	ARR(sec2)
	CREAR(sal)
	AVZ(sec1)
	AVZ(sec2)
	
	cont_oracion1 := 0
	cont_oracion2 := 0

	Mientras NFDS(sec1) ^ NFDS(sec2) Hacer
		Mientras (v1 <> ",") Hacer
			GRABAR(sal,v1)
			AVZ(sec1,v1)
		FinMientras
		Mientras (v1 <> ".") Hacer
			AVZ(sec1,v1)
		FinMientras
		cont_oracion1 := cont_oracion1 + 1;
		Mientras (v2 <> ".") Hacer
			GRABAR(sal,v2)
			AVZ(sec2,v2)
		FinMientras
		cont_oracion2 := cont_oracion2 + 1;
		GRABAR(sal,v2)
		AVZ(sec2,v2)
		AVZ(sec1,v1)
	FinMientras
	Esc("La cantidad de oraciones de la secuencia 1 es: ", cont_orac1)
    Esc("La cantidad de oraciones de la secuencia 2 es: ", cont_orac2)
    Cerrar(sec1)
    Cerrar(sec2)
    Cerrar(sal)
FinAccion	
	
