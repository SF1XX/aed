Ejercicio 2
Se tiene la información de todos los jugadores de cada uno de los planteles en un archivo secuencial con 
el siguiente formato

Jugadores (Ordenado por selección, Posición y Dorsal)
Seleccion | Posicion (AR, DF, ME, DE, CT) | Dorsal | Apellido y Nombre| Edad | Equipo 

Donde: Posicion (AR(Arquero), DF(Defensor), ME (Mediocampista), DE (Delantero), CT (Cuerpo Técnico)
Se le solicita:
1. Obtener total de personas que integran cada Selección, por nombre de selección y posición, y total 
general. Excluyendo a los integrantes de la comisión técnica.

2. Por selección indicar, promedio de edad de los jugadores.

Accion Ejercicio_N2 Es
Ambiente
	jugadores = REGISTRO
		CLAVE = REGISTRO
			Seleccion: AN(10)
			Posición: {"AR", "DF", "ME", "DE", "CT"}
		Fin_Reg
			Dorsal: N(2)
			Apellido_Nombre: AN(20)
			Edad: N(2)
			Equipo: AN(20)
	Fin_Reg
	
	Arch: archivo de jugador ordenado por CLAVE;
	reg: jugador

	res_seleccion: AN(10)
	res_posicion: {"AR", "DF", "ME", "DE", "CT"}
	res_dorsal: N(2)
	res_AyP: AN(20)
	res_edad: AN(2)
	res_equipo: AN(20)

	tot_general: entero
	tot_seleccion: entero
	tot_posicion: entero
	tot_edad_may: entero
	tot_edad_men: entero
	
	promedio_edad: real

	cont_edad_may: entero
	cont_edad_men: entero
	cont_posicion: entero
	cont_seleccion: entero

	Subaccion corte_edad Es
		Esc("La cantidad de juagadores menores",cont_edad_men "de 20 es:", tot_edad_men);
		Esc("La cantidad de juagadores mayores",cont_edad_may "de 20 es:", tot_edad_may);
		cont_posicion := cont_posicion + cont_edad_men + cont_edad_may
		res_edad := reg.Edad;
		cont_edad_may := 0
		cont_edad_men := 0
	FinSub

	Subaccion corte_posicon Es	
		corte_edad()
		Esc("La cantidad de jugadores por posicion",res_posicion "es de:"cont_posicion);
		Mientras (res_posicion <> "CT") Hacer
			cont_seleccion := cont_seleccion +  cont_posicion;
			res_posicion := reg.CLAVE.Posicion;
			cont_posicion := 0;
		FinMientras	
	FinSub
	
	Subaccion corte_seleccion Es
		corte_posicon()
		Esc("La cantidad de jugadores por seleccion", res_seleccion "es de:"cont_seleccion)
		tot_general := tot_general + cont_seleccion;
		res_seleccion := reg.CLAVE.Seleccion
		cont_seleccion := 0
	FinSub

Proceso
	Abrir E/(Arch)
	Leer(Arch,reg)

	res_seleccion := reg.CLAVE.Seleccion
	res_posicion := reg.CLAVE.Posicion

	tot_general := 0
	tot_seleccion := 0
	tot_posicion := 0

	cont_posicion := 0
	cont_seleccion := 0	

	ientras NFDA(Arch) Hacer
		Si (res_seleccion <> reg.CLAVE.Seleccion) Entonces
			corte_seleccion()
			sino
			Si (res_posicion <> reg.CLAVE.Posicion) Entonces
				corte_posicon()
				sino
				Si (res_edad <> reg.Edad) Entonces
					corte_edad()
				FinSi	
			FinSi
		FinSi
		
		