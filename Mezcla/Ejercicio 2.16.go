/**Accion 2.16 Es
Ambiente

	FECHA = REGISTRO
		aa: N(4)
		mm: N(2)
		dia: N(2)
	Finregistro

	PELICULA = REGISTRO
		Nro_pel: N(3)
		Titulo: AN(30)
		Genero: AN(15)
		Cant_copias: N(3)
		Fecha_Estreno: FECHA
	Finregistro
	
	Arch_act, Peliculas, P_nuevas: Archivo de Peliculas
	reg_a, reg_p, reg_pn: PELICULAS

Proceso
	Abrir E/(PELICULAS);
	Leer(PELICULAS,reg_p);	
	Abrir E/(P_nuevas);
	Leer(PELICULAS,reg_pn);
	Abrir S/(Arch_act);

	Mientras NFDA(PELICULAS) ^ NFDA(P_nuevas) Hacer
		Si reg_p.Nro_pel < reg_pn.Nro_pel Entonces
			Mientras reg_p.Nro_pel < reg_pn.Nro_pel Hacer
				reg_a := reg_p;
				Esc(Arch_act,reg_a);
				Leer(PELICULAS,reg_p);
			Finmientras
		Sino
			Mientras reg_p.Nro_pel > reg_pn.Nro_pel hacer
				reg_a:=reg_pn; 
				Escribir(archaAct, reg_a)
				Leer(P_nuevas, reg_pn)
			Finmientras		
		Finsi
	Finmientras
	Mientras NFDA(Peliculas) hacer
		reg_a:=reg_p; 
		Escribir(arch_act, reg_a)
		Leer(Peliculas, reg_p)
	Fin Mientras
	Mientras NFDA(P_nuevas) hacer
		reg_a:=reg_pn; 
		Escribir(arch_act_act, reg_a)
		Leer(P_nuevas, reg_pn)
	Fin Mientras
	Cerrar(Peliculas); 
	Cerrar(P_Nuevas); 
	Cerrar(Arch_Act);
FinAccion		
*/


// Excluyente
Accion 2.16 Es
Ambiente

	FECHA = REGISTRO
		aa: N(4)
		mm: N(2)
		dia: N(2)
	Finregistro

	PELICULA = REGISTRO
		Nro_pel: N(3)
		Titulo: AN(30)
		Genero: AN(15)
		Cant_copias: N(3)
		Fecha_Estreno: FECHA
	Finregistro

	Peliculas, P_Nuevas, Arch_Act: archivo Peliculas
	R_P, R_A, R_PN: PELICULA

	Proceso
		Abrir E/(Peliculas);
		Leer(Peliculas,R_P);

		Abrir E/(P_Nuevas);
		Leer(P_Nuevas, R_PN);

		Abrir S/(Arch_Act);

		Mientras NFDA(Peliculas) ^ NFDA(P_Nuevas) Hacer
			Si R_P.Nro_pel < R_PN.Nro_pel Entonces	
				Grabar(Arch_Act,R_P);
				Leer(Peliculas,R_P);
			Sino
				Grabar(Arch_Act,R_PN);
				Leer(Peliculas,R_PN)
			FinSi
				
		Finmientras
			
		Mientras NFDA(Peliculas) Hacer
			Grabar(Arch_Act, R_A);
			Leer(Peliculas,R_P);
		Finmientras
		
		Mientras NFDA(P_Nuevas) Hacer
			Grabar(Arch_Act,R_PN);
			Leer(P_Nuevas,R_PN);
		Finmientras
		
		Cerrar(Peliculas);
		Cerrar(P_Nuevas);
		Cerrar(Arch_Act);
	FinAccion	


// Incluyente

Accion 2.16 Es
Ambiente
	
	FECHA = REGISTRO
		aa: N(4)
		mm: N(2)
		dia: N(2)
	Finregistro

	PELICULA = REGISTRO
		Nro_pel: N(3)
		Titulo: AN(30)
		Genero: AN(15)
		Cant_copias: N(3)
		Fecha_Estreno: FECHA
	Finregistro

	Peliculas, P_Nuevas, Arch_Act: archivo Peliculas
	R_P, R_A, R_PN: PELICULA


	Hv: Entero;

	Procedimiento leer_peliculas Es
		Leer(Peliculas,R_P);

		Si FDA(Peliculas) Entonces
			R_P.Nro_pel:=HV
		FinSi
	Finprocedimiento

	Procedimiento leer_peliculasN Es
		Leer(P_Nuevas,R_PN);

		Si FDA(P_Nuevas) Entonces
			R_PN.Nro_pel:=HV
		FinSi
	Finprocedimiento

Proceso
	Abrir E/(Peliculas);
	Leer(Peliculas,R_P);

	Abrir E/(P_Nuevas);
	Leer(P_Nuevas, R_PN);

	Abrir S/(Arch_Act);		

	Mientras (R_P<>HV) ^ (R_PN<>HV) Hacer
		Si R_P.Nro_pel < R_PN.Nro_pel Entonces	
			Grabar(Arch_Act,R_P);
			Leer(Peliculas,R_P);
		Sino
			Grabar(Arch_Act,R_PN);
			Leer(Peliculas,R_PN)
		FinSi
	Finmientras
	
	Cerrar(Peliculas);
	Cerrar(P_Nuevas);
	Cerrar(Arch_Act);	

FinAccion	