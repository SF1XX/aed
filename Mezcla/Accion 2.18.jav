Accion 2.18 Es
Ambiente
	ARTICULOS = REGISTO
		Cod_Prod: N(6);
		Tipo: 1..5
		Marca: AN(20);
		Descripción: AN(50);
		Cant_Unidades: N(3);
	Finregistro
	
	A_SUC1, A_SUC2: Archivos de articulos
	R_SUC1, R_SUC2. ARTICULOS

	SALIDA = REGISTRO
		Cod_Prod: N(6);	
		Tipo: 1..5	
		Marca: AN(20);	
		Descripción: AN(50);	
		Cant_Suc_1: N(3);	
		Cant_Suc_2: N(3);	
		Total_Unidades: N(3);	
	Finregistro	

	A_SAL: Archivo de salida
	R_SAL: SALIDA

	Procedimiento Leer_SUC1 Es
		Leer(A_SUC1,R_SUC1);

		Si FDA(A_SUC1) Entonces
			R_SUC1:=HV;
		Finsi
	Finprocedimiento

	Procedimiento Leer_SUC2 Es
		Leer(A_SUC2,R_SUC2);

		Si FDA(A_SUC2) Entonces
			R_SUC2:=HV;
		Finsi
	Finprocedimiento		

	HV: entero;
	Total_Unidades: entero;

Proceso
	Abrir E/(A_SUC1);	
	Leer(A_SUC1,R_SUC1);

	Abrir E/(A_SUC2);	
	Leer(A_SUC2,R_SUC2);

	Abrir S/(A_SAL);

	| Cod_Prod | Tipo | Marca | Descripción | Cant_Suc_1 | Cant_Suc_2 | Total_Unidades |
	Mientras (R_SUC1.Cod_Prod <> HV) ^ (R_SUC2.Cod_Prod <> HV) Hacer
		Si R_SUC1.Cod_Prod < R_SUC2.Cod_Prod Entonces
			R_SAL:=R_SUC1;
			GRABAR(R_SUC1.Cod_Prod, R_SUC1.Tipo, R_SUC1.Descripción, R_SUC1.Cant_Suc_1, R_SUC1.Total_Unidades);
			Leer_SUC2;
		sino
			Si 	(R_SUC1.Cod_Prod > R_SUC2.Cod_Prod) Entonces
				R_SAL:=R_SUC2;
				GRABAR(R_SUC2.Cod_Prod, R_SUC2.Tipo, R_SUC2.Descripción, R_SUC2.Cant_Suc_2, R_SUC2.Total_Unidades);
				Leer_SUC1;
			sino
				Si R_SUC1.Cod_Prod = R_SUC2.Cod_Prod Entonces
					R_SAL:=R_SUC1;
					cont:= Cant_Suc_1 + Cant_Suc_2;
					GRABAR(R_SUC1.Cod_Prod, R_SUC1.Tipo, R_SUC1.Descripción, R_SUC1.Cant_Suc_1, R_SUC2.Cant_Suc_2, cont);
				Finsi	
				Leer(A_SUC1,R_SUC1);
				Leer(A_SUC2,R_SUC2);
			Finsi
		Finsi
		GRABAR(A_SAL,R_SAL);
	Finmientras	
	CERRAR(A_SUC1);	
	CERRAR(A_SUC2);
	CERRAR(A_SAL);
FinAccion	




