Ejercicio 1

La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:
Argentina#Brasil#Canada#Peru# … # Colombia#FDS

Y la edad de cada uno de los jugadores que componen el plantel, se encuentra disponible en una secuencia de enteros
32|22|27|19|...|33|18|FDS
La cantidad de jugadores que componen un plantel es de 26 jugadores.

Se le solicita:
1. Informar el promedio de edad de los jugadores que participaran este año en la Copa.
2. ¿Cuál es la edad del jugador más joven que participará en la copa? ¿Y la edad del más longevo ?
3. Informar también por cada equipo (a y b).


Accion EJ_N1 Es
Ambiente
	sec1: secuencia de caracteres;
	sec2: secuencia de enteros;

	I, v1 : caracter
	v2: entero

	edad: entero
	jugador_mas_joven: enteros
	jugador_mas_viejo: enteros

	cont_mas: entero
	cont_men: entero

Proceso
	Mientras NFDS(sec1) Hacer
		Repetir
			AVZ(sec1,v1)
		Hasta que v = "#"

		