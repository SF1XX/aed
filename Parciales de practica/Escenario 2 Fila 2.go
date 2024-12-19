Ejercicio 2

Los datos historicos de cada vehiculo que paso por el peaje se encuentran en un archivo con el siguiente
formato:
| Año | Mes | Dia | Categoria | Patente | Origen ("Rcia" o "Ctes") |

Se le solicita:

1. Cantidad total por año y mes, discriminados por ciudad de origen (Resistencia o Corrientes) y total 
   general.

2. Emitir un informe con el siguiente formato considerando solo los meses donde hubo mas de 100.000 pases,
   desde ambas ciudades.

| Año | Mes | Cantidad desde Resistencia | Cantidad desde Corrientes |


Accion Corte_control Es
Ambiente

	FECHA = REGISTRO
		aa: N(4);
		mm: N(2);
		dd: N(2);
	FinReg


	PEAJE = REGISTRO
		CLAVE = REGISTRO
			Año: N(4);
			Mes: N(2);
			Dia: N(2)
			Categoria: N(10);
			Patente: AN(7);
		FinReg
			Origen: ("rcia", "ctes");
	FinReg
	
	Arch: Archivo de peaje ordenado por CLAVE
	Reg: PEAJE

	res_año: N(4);
	res_mes: N(2);

	tot_gral_rcia: entero
	tot_gral_ctes: entero
	tot_rcia: entero
	tot_ctes: entero
	tot_mes_ctes: entero
	tot_mes_rcia: entero
	tot_año_ctes: entero
	tot_año_rcia: entero


	procedimiento corte_mes Es
		SI (tot_año_ctes + tot_año_rcia) ENTONCES // CONSIGNA 2
			ESC("Los siguientes datos corresponden a mas de 100.000 pases en el mes")
			ESC("| Año | Mes | Cantidad desde Resistencia | Cantidad desde Corrientes")
			ESC(res_año, " | ", res_mes, " | ", tot_mes_rcia, " | ", tot_mes_ctes)
		SINO
			ESC("En el mes ", res_mes, "la cantidad de pases desde Resistencia", tot_mes_rcia)
			ESC("En el mes ", res_mes, "la cantidad de pases desde Corrientes", tot_mes_ctes)
		FINSI
		tot_año_ctes := tot_año_ctes + tot_mes_ctes
		tot_año_rcia := tot_año_rcia + tot_mes_rcia
		tot_mes_ctes := 0
		tot_mse_rcia := 0
		res_mes := reg.CLAVE.Mes
	FinProcedimiento
	
	procedimiento corte_año Es
		corte_mes();
		ESC("En el año ", res_año, "la cantidad de pases desde Resistencia", tot_año_rcia)
		ESC("En el año ", res_año, "la cantidad de pases desde Corrientes", tot_año_ctes)
		tot_gral_ctes := tot_gral_ctes + tot_año_ctes
		tot_gral_rcia := tot_gral_rcia + tot_año_rcia
		tot_año_ctes := 0
		tot_año_rcia := 0
		res_año := Reg.CLAVE.Año
	FinProcedimiento

Proceso
	Abrir E/(Arch);
	Leer(Arch,Reg);
	
	tot_gral := 0;
	tot_rcia := 0;
	tot_ctes := 0;
	tot_mes_ctes := 0;
	tot_mes_rcia := 0;
	tot_año_ctes := 0;
	tot_año_rcia := 0;

	res_año: Reg.CLAVE.Año;
	res_mes: Reg.CLAVE.Mes;

	Mientras NFDA(Arch) Hacer
		Si (res_año <> reg.CLAVE.Año) Entonces
			corte_año();
				sino
					Si (res_mes <> reg.CLAVE.Mes) Entonces
						corte_mes();
					FinSi
		FinSi
		
		// Tratar registro
		
		Si (Reg.Origen = "Rcia") ENTONCES
			tot_mes_rcia := tot_mes_rcia + 1;
		sino
			tot_mes_ctes := tot_mes_ctes + 1;
		FinSi
		
		Leer(Arch,Reg);
	FinMientras
	
	corte_año();
	ESC("el total de pases desde Resistencia es: ", tot_gral_rcia)
	ESC("el total de pases desde Corrientes es: ", tot_gral_ctes)
	cerrar(PEAJE);
finaccion
