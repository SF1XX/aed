/**
*! CONSIGNA
*/

Ejercicio 1

La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:
Argentina#Brasil#Canada#Peru# … # Colombia#FDS

Y la edad de cada uno de los jugadores que componen el plantel, se encuentra disponible en una secuencia de enteros
32|22|27|19|...|33|18|FDS
La cantidad de jugadores que componen un plantel es de 26 jugadores.

Se le solicita:
1- Informar el promedio de edad de los jugadores que participaran este año en la Copa.
2- ¿Cuál es la edad del jugador más joven que participará en la copa? ¿Y la edad del más longevo ?
3- Informar también por cada equipo (a y b).

Accion Ej1_N1 Es
Ambiente
	sec: secuencia de caracteres
	sec2: secuencia de enteros
	v: caracter
	v2: entero
	cont_May: entero
	cont_Men: entero
	promedio: real
Proceso
	ARR(sec)
	ARR(sec2)
	AVZ(sec,v)
	AVZ(sec2,v2)
	promedio := 0

	Mientras NFDS(sec) ^ NFDS(sec2) Hacer
		Para I := 1 Hasta 16 Hacer
			cont_May := 0
			cont_Men := 0
			Mientras (v <> "#") Hacer
				AVZ(sec,v)
			FinMientras
			Si ()	


/**
*! CONSIGNA
*/

Ejercicio 2
Se tiene la información de todos los jugadores de cada uno de los planteles en un archivo secuencial con el siguiente formato

Jugadores (Ordenado por selección, Posición y Dorsal)
| Seleccion | Posicion (AR, DF, ME, DE, CT) | Dorsal | Apellido y Nombre | Equipo | 

Donde: Posicion (AR (Arquero), DF (Defensor), ME (Mediocampista), DE (Delantero), CT (Cuerpo Técnico)
Se le solicita:

1- Obtener total de personas que integran cada Selección, por nombre de selección y posición, y total general.


Accion Ej1_N2 Es
Ambiente
	Jugadores = REGISTRO
		CLAVE = REGISTRO
			selección: AN(14) 
			posición = {"AR", "DF", "ME", "DE", "CT"}
			dorsal: N(2)
		FinRegistro	
			AyP: AN(20)
			Equipo: AN(20)
	FinRegistro

	archivo: archivo ordenado por CLAVE
	registro: Jugadores

	tot_general: entero
	tot_seleccion: entero
	tot_posicion: entero

	res_sel: AN(14)
	res_posicion: {"AR", "DF", "ME", "DE", "CT"}

	Subaccion corte_posicion Es
		Esc("la cantidad de jugadores por posición",res_posicion " es de:",tot_posicion )
		tot_seleccion := tot_seleccion + tot_posicion;
		
