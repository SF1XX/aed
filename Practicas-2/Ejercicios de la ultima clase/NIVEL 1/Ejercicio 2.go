Ejercicio 2
Se tiene la información de todos los jugadores de cada uno de los planteles en un archivo secuencial con el siguiente formato

Jugadores (Ordenado por selección, Posición y Dorsal)

| Seleccion | Posicion (AR, DF, ME, DE, CT) | Dorsal | Apellido y Nombre | Equipo |

Donde: Posicion (AR (Arquero), DF (Defensor), ME (Mediocampista), DE (Delantero), CT (Cuerpo Técnico)
Se le solicita:
1. Obtener total de personas que integran cada Selección, por nombre de selección y posición, y total general.


Accion EJ_N2 Es
Ambiente
	jugadores = REGISTRO
		CLAVE = REGISTRO
			Seleccion: AN(10)
			Posición: {"AR", "DF", "ME", "DE", "CT"}
		Fin_Reg
			Dorsal: N(2)
			Apellido_Nombre: AN(20)
			Equipo: AN(20)
	Fin_Reg
	
	Arch: archivo de jugador ordenado por CLAVE;
	reg: jugador

	res_seleccion: AN(10)
	res_posicion: {"AR", "DF", "ME", "DE", "CT"}
	res_dorsal: N(2)
	res_AyP: AN(20)
	res_equipo: AN(20)

	tot_general: entero
	tot_seleccion: entero
	tot_posicion: entero

	cont_posicion: entero
	cont_seleccion: entero

	Subaccion corte_posicon Es	
		Esc("La cantidad de jugadores por posicion",res_posicion "es de:"cont_posicion);
		cont_seleccion := cont_seleccion +  cont_posicion;
		res_posicion := reg.CLAVE.Posicion;
		cont_posicion := 0;
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
	Leer (Arch,reg)

	res_seleccion := reg.CLAVE.Seleccion
	res_posicion := reg.CLAVE.Posicion

	cont_general := entero
	cont_posicion :=  entero
	cont_seleccion := entero
	
	Mientras NFDA(Arch) Hacer
		Si (res_seleccion <> reg.CLAVE.Seleccion) Entonces
			corte_seleccion()
		sino
			Si (res_posicion <> reg.CLAVE.Posicion) Entonces
				corte_posicon()
			FinSi
		FinSi				
	
	/**
	* ! Tratar registro
	*/
	
		cont_general := cont_general + 1
		cont_seleccionn := cont_seleccion + reg.CLAVE.Seleccion;
		Leer(Arch,reg)
	FinMierntras
	corte_seleccion()
	Esc("")
	CERRAR(Arch)
