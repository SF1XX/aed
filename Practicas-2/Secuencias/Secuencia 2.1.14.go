/**
*! CONSIGNA
*/
Se dispone de una secuencia de caracteres y se desea saber la cantidad de 
caracteres (incluidos los espacios) que existen entre una coma y la siguiente. 

Se debe considerar que puede haber más de un par de comas, y que las subsecuencias inicial y final 
deben descartarse por no cumplir la condición enunciada. Supóngase que las comas son siempre contiguas 
al último caracter de la palabra.

Accion 2.1.14 Es
Ambiente
	sec: secuencia de caracteres
	v: caracter
	cont: entero

Proceso
	ARR(sec)
	AVZ(sec,v)
	Mientras NFDS(sec) Hacer
		cont := 0
		Mientras (v <> ",")
			cont := cont + 1;
			AVZ(sec,v)
		FinMientras
	FinMientras	
	Esc("en este bloque hay:", cont)
	CERRA(sec)
FinAccion	


	

		


