/**El organismo del cual dependen las escuelas de un distrito lleva un archivo con los registros 
para todos los alumnos de nivel secundario de ese distrito:

ALUMNOS Ordenado por Escuela, Año, División

| -Escuela- | -Año- | -División- | Nombre | Cant_Inasistencias |

Se introduce como dato el número de distrito y la cantidad de días de clase dictados en el año. 
El archivo está ordenado por número de escuela, año y división. Los alumnos que superan el 20% de 
las inasistencias están en situación de LIBRES, de lo contrario son REGULARES.

Se desea conocer:
Para cada división:
- Cantidad de alumnos
Para cada año:
- Cantidad total de alumnos libres
- Cantidad de alumnos regulares
- Cantidad total de alumnos
Todas las escuelas:
- Cantidad total de alumnos
- Porcentaje de alumnos libres. - Promedio de inasistencias por alumnos.
*/

Accion Corte2215 Es
Ambiente
	ALUMNOS = REGISTRO
		CLAVE = REGISTRO
			Escuela: N(3);
			Año: N(4);
			Division: N(2);
		FinRegistro
		Nombre: AN(15);
		Cant_inas: N(2);
	FinRegistro
	
	Arch: Archivo de alumnos ordenado por CLAVE;
	Reg: ALUMNOS;

	Ndistrito: entero;
	Cant_clases: entero;

	Cant_dias: entero;
	Cant_alum_div: entero;
	Cant_alum_libres_año: entero;
	Cant_alum_regulares_año: entero;
	Cant_alum_total_año: entero;
	Cant_alum_total_gen: entero;
	total_gen: entero;
	Alum_libres: entero;

	res_escuela: N(3);
	res_año: N(4);
	res_division: N(2);

	Procedimiento corte_division Es
		Esc("La cantidad de alumnos en la division", res_division, "es: " Cant_alum_div);
		Cant_alum_total_año := Cant_alum_total_año + Cant_alum_div;
		res_division := Reg.CLAVE.Division;
		Cant_alum_div := 0;
	FinProcedimiento
	
	Procedimiento corte_año Es
		corte_division();
		Esc("La cantidad de alumnos libres en el año", res_año, "es de: ", Cant_alum_libres_año);
		Esc("La cantidad de alumnos regulares en el año", res_año "es de: " Cant_alum_regulares_año);
		Esc("La cantidad de total de alumnos en el año", res_año, "es de: " Cant_alum_total_año);
		Cant_alum_total_gen := Cant_alum_total_gen + Cant_alum_total_año;
		Alum_libres := Alum_libres + Cant_alum_libres_año;
		res_año := Reg.CLAVE.Año;
		Cant_alum_libres_año := 0;
		Cant_alum_regulares_año := 0;
		Cant_alum_total_año := 0;
	FinProcedimiento

	Procedimiento corte_escuela Es
		corte_año();
		Esc("La cantidad de alumnos es la escule", res_escuela, "es de: " Cant_alum_total_gen);
		total_gen :total_gen + Cant_alum_total_gen;
		res_año := Reg.CLAVE.Escuela;
		Cant_alum_total_gen := 0
	FinProcedimiento	

Proceso
	Abrir E/(ALUMNOS);
	Leer(Arch,Reg);	

	
	Cant_alum_div := 0;
	Cant_alum_libres_año := 0;
	Cant_alum_regulares_año := 0;
	Cant_alum_total_año := 0;
	Cant_alum_total_gen := 0;
	total_gen := 0;
	Alum_libres := 0;

	res_escuela := Reg.CLAVE.Escuela;
	res_año := Reg.CLAVE.Año;
	res_division := Reg.CLAVE.Division;

	Esc("Ingrese el numero de distrito")
	Leer(Ndistrito);
	Esc("Ingrese la cantidad de días de clase dictados en el año");
	Leer(Cant_clases);

	Mientras NDFA(ALUMNOS) Hacer
		Si (res_escuela <> Reg.CLAVE.Escuela) Entonces
			corte_escuela();
		sino
			Si (res_año <> Reg.CLAVE.Año) Entonces
				corte_año();
			sino
				Si (res_division <> Reg.CLAVE.Division) Entonces
					corte_division()
				FinSi
			FinSi
		FinSi
		
		// Tratar registro

		Cant_alum_div := Cant_alum_div + 1;
		Cant_Inasistencias := Cant_Inasistencias + 1;
		Si((Reg.Cant_inas *100)/ Cant_dias) > 20 Entonces
			Cant_alum_libres_año := Cant_alum_libres_año + 1;
		sino
			Cant_alum_regulares_año := Cant_alum_regulares_año + 1;
		FinSi
		 Leer(Arch,Reg);
	FinMienras

	corte_escuela();
	Esc("El total de alumnos de todas las escuelas es de:", total_gen);
  Esc("El porcentaje de alumnos libres es de:", (Alum_libres *100)/total_gen);
  Esc("El promedio de inasistencias por alumnos es de :", Cant_Inasistencias/total_gen);
  Cerrar(Arch);	 	
FinAccion	