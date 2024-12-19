Accion 2.17 Es
Ambiente

	FECHA = REGISTRO
		aa: N(4)
		mm: N(2)
		dia: N(2)
	Finregistro

	ASPIRANTES = REGISTRO
		DNI: N(8);
		ApeyNom: AN(30);
		Carrera: {"ISI", "IEM", "IQ", "LAR"}
		F_Nac: FECHA
		Email: AN(30);
		ColegioSec: AN(30);
		FechaInscripcion: FECHA
		Aprobado: {"SI", "NO"}
	Finregistro

		
	Arch_Ago, Arch_Feb: Archivos ordenados po DNI
	R_AGO, R_FEB, R_S: ASPIRANTES

	SEGUIMIENTO = REGISTRO
		DNI: N(8);
		ApeyNom: AN(30);
		Email: AN(30);
		ColegioSec: AN(30);
	Finprocedimiento

	A_SEG: Archivo de siguimiento ordenado por DNI
	R_S: SEGUIMIENTO

	HV: entero;
	cont: entero;

	Procedure Leer_Ago Es
		Leer(Arch_Ago);

		Si FDA(Arch_Ago) Hacer
			R_AGO:=HV
		Finsi
	Finprocedimiento

	Procedure Leer_Feb Es
		Leer(Arch_Feb);

		Si FDA(Arch_Feb) Hacer
			R_FEB:=HV
		Finsi
	Finprocedimiento

Proceso
	Abrir E/(Arch_Ago);	 
	Leer(Arch_Ago,R_AGO);

	Abrir E/(Arch_Feb);	 
	Leer(Arch_Feb,R_FEB);

	Abrir S/(Salida);

	Mientras (R_AG.DNI <> HV) ^ (R_Feb) Hacer
		Si R_Feb.DNI < R_AGO.DNI Entonces
			Leer_Feb;
		Sino
			Si 	R_FEB.DNI > R_AGO.DNI Entonces
				Leer_Ago;;
			Sino
				Si R_FEB.DNI = R_AGO.DNI Entonces
					Si (R_FEB.Aprobado = "NO") ^ (R_AGO.Aprobado = "NO") Entonces
						R_S:=R_FEB.DNI
						R_S:=R_FEB.ApeyNom
						R_S:=R_FEB.Email
						R_S:=R_FEB.ColegioSec
						GRABAR(Arch_SEG,R_S);
						cont:=cont + 1;
					Finsi
					Leer(Arch_Ago,R_AGO);	
					Leer(Arch_Feb,R_FEB);
				Finsi
			Finsi			
		Finsi
	Finmientras
	CERRA(Arch_Ago);
	CERRA(Arch_Feb);
	CERRA(A_SEG);
	Esc("La cantidad de aspirantes que se grabaron en el archivo SEGUIMIENTO es: ", cont);
FinAccion.			