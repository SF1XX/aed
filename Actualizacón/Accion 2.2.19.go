Accion 2.2.19 Es
Ambiente
	FECHA = REGISTRO
		aa: N(4);
		mm: N(2);
		dd: N(2);
	Finregistro
		
	REG_CLAVE = REGISTRO
		Farmacia: AN(30)
		Medicamento: AN(30)
	Finregistro
	
	MAE_REMEDIOS = REGISTRO
		CLAVE = REG_CLAVE
		Cant_Actual: N(4)
		Fecha_vencimietno: FECHA
	Finregistro
	
	MOVIMIENTOS = REGISTRO	
		CLAVE = REG_CLAVE
		Cod_Mov: ("1", "2", "3")
		Cant_Recibida: N(4)
	Finregistro
	
	REG_VEN = REGISTRO
		Medicamento: AN(30)
		Cant_Vencida: N(4)
	Finregistro
	
	Procedure Leer_Movimiento Es
		Leer(Arch_Mov, reg_mov);
		Si FDA(Arch_Mov) Entonces
			reg_mov:= HV
		Finsi	
	Finprocedure
	
	Procedure Leer_Maestro Es	
		Leer(Arch_Mae, reg_mae);
		Si FDA(Arch_Mae) Entonces 
			reg_mae:= HV
		FinSi	
	Finprocedure	

	// Variables de archivos y registros

	Arch_Mae: Archivo ordenado por CLAVE;	
	Arch_Mov: Archivo ordenado por CLAVE, Cod_Mov;
	Arch_Sal: Archivo de salida ordenado por Medicamento, Cant_Vencida;
	reg_sal: REG_VEN;
	reg_mae: MAE_REMEDIOS;
	reg_mov: MOVIMIENTOS;

Proceso
	Abrir E/(Arch_Mae); Leer_Maestro(Arch_Mae,reg_mae);
	Abrir E/(Arch_Mov);	Leer_Movimiento(Arch_Mov,reg_mov);
	
	Mientras (reg_mae.CLAVE <> HV ^ reg_mov.CLAVE <> HV)Entonces
		Si (reg_mae.CLAVE < reg_mov.CLAVE) Entonces
			GRABAR(Arch_Sal, reg_mae);
			Leer_Maestro()
		Sino 	
			Si (reg_mae.CLAVE = reg_mov.CLAVE) Entonces
				Si reg_mov.Cod_Mov = 1 Entonces
					Esc("Error, No se puede dar de alta un archivo existente");
					GRABAR(Arch_Sal, reg_mae);
				Sino
					Si reg_mov.Cod_Mov = 2 Entonces
						reg_sal.Medicamento:= reg_mov.Medicamento
						reg_sal.Cant_Vencida:= reg_mae.Cant_Actual
						GRABAR(Arch_Sal, reg_sal);
					Sino 
						reg_mae.Cant_Actual:= reg_mov.Cant_Recibida	
						GRABAR(Arch_Mae_Act, reg_mae);
					FinSi	
				Finsi
				Leer_Maestro();
				Leer_Movimiento();
			Sino
				Si reg_mov.Cod_Mov = 1 Entonces
					reg_mae_act.CLAVE.Farmacia:= reg_mov.CLAVE.Farmacia
					reg_mae_act.CLAVE.Medicamento:= reg_mov.CLAVE.Medicamento
					reg_mae_act.CLAVE.Cant_Actual:= reg_mov.Cant_Recibida
					reg_mae_act.Fecha_vencimietno:= sumar_dias(reg_mae_act.FECHA(), 30)
					GRABAR(Arch_Mae_Act, reg_mae_act);
				Sino
					Esc("Error");	
				FinSi
				Leer_Movimiento();
			Finsi		
		FinSi
	FinMientras
	CERRAR(Arch_Mae);
	CERRAR(Arch_Mov);
	CERRAR(Arch_Sal);
	CERRAR(Arch_Mae_Act);
FinAccion		
					

