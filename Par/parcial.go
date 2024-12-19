|- Delegacion -|- Deporte -|- NroCredencial -|- Apellido y Nombre -| Participaciones | Fecha_nacim | 

Accion corte_olimpiadas Es	
Ambiente

	FECHA = REGISTRO
		aa: N(4)
		mm: N(2)
		dia: N(2)
	Finregistro

	ATLETAS = REGISTRO
		CLAVE = REGISTRO	
			Delegacion: AN(20)
			Deporte: AN(3)
			NroCreden: N(10)
			AyN: AN(30)
		Finregistro
		Participaciones: N(3)
		Fecha_nacim: FECHA	
	Finregistro	

	arch: archivo ordenado por CLAVE
	reg: ATLETAS

	tot_delegacion: entero
	tot_deport: entero
	tot_mas: entero
	tot_men: entero
	tot_global: entero

	res_delegacion: AN(20)
	res_deporte. An(3)

	Procedimiento corte_deporte Es
		Esc("La cantidad de participantes en el deporte", res_deporte, "es de: ", tot_deport);
		tot_delegacion := tot_delegacion + tot_deport;
		res_deporte := reg.CLAVE.Deporte
		tot_deport := 0;
	Finprocedimiento
	
	Procedimiento corte_delegacion Es
		Esc("La cantidad de participantes en la delegacion", res_delegacion, "es de: ", tot_delegacion);
		tot_global := tot_global + tot_delegacion;
		res_delegacion := reg.CLAVE.Delegacion
		tot_global := 0;
	Finprocedimiento
		