Ejercicio 1
La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:
Argentina#Brasil#Canada#Peru# … # Colombia#FDS

Además se encuentra con información individual de cada jugador: dorsal y edad en una secuencia de enteros con el siguiente formato
10|36|11|36|19|...|33|18|FDS
Esto significa que, el primer jugador de la Selección Argentina que usará el dorsal 10 tiene 36 años (🐐 🏆), el siguiente usará el dorsal 11 y 
también tiene 36 años, y así sucesivamente.

La cantidad de jugadores que componen un plantel es de 26 jugadores.

Se le solicita:
1. Informar el promedio de edad de los jugadores de un determinado dorsal, que lo indicará el usuario por única vez.
2. ¿Cuál es la edad del jugador más joven por equipo? ¿Y la edad del más longevo de cada equipo?


Accion EJ_N1 Es
Ambiente
	sec: secuencia de caracteres;  // secuencia de paises 
	sec2: secuencia de enteros;  // secuencia de jugadores
	v: caracter
	v2: entero
	I: entero
	D: entero

	dorsal: entero
	mas_longevo: entero
	mas_joven: entero

	promedio: real
	edad: entero

Proceso
	ARR(sec)
	ARR(sec2)
	AVZ(sec,v)
	AVZ(sec2,v2)
	edad := 0
	mas_longevo := 100
	mas_joven := 0

	Esc("Ingrese los dorsales")
	Leer(dorsal)	
	Mientras NFDS(sec) ^ NFDS(sec2) Hacer
		Mientras(v <> #) Hacer
			Esc(v)
			AVZ(sec,v)
		FinMientras	
		
		Para I := 1 hasta 26 Hacer
			Si (D = dorsal) Entonces
				