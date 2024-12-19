Accion Ejercicios_de_secuencia Es
Ambeinte
	Fecha: secuencia de caracter
	Vehiculos: secuencia de caracter
	Salida: secuencia de caracter

	f: caracter
	v: caracter
	sal: caracter

	cont_vehiculos: entero
	descuento: logico
	pat1, pat2, pat3, pat4, pat5: caracter;

	mes1_us, mes2_us: caracter
	m1, m2: caracter




Proceso
	ARR(Fecha);
	AVZ(Fecha,f); // Estoy en el primer caracter del dia
	ARR(Vehiculos);
	AVZ(Vehiculos,v); // Estoy en categoria
	CREAR(Salida);

	cont_vehiculos := 0;

	Esc("Ingrese el primer caracter del mes");
	Leer(mes1_us);
	Esc("Ingrese el segundo caracter del mes");
	Leer(mes2_us);

	Mientras NFDS(Fecha) ^ NFDS(Vehiculos);
		// Trato la fecha
		AVZ(Fecha,f);
		AVZ(Fecha,f); // Estoy en el primer caracter del mes
		m1 := f;
		AVZ(Fecha,f); // Estoy en el segundo caracter del mes
		m2 := f;
	
		AVZ(Fecha,f); // Estoy en el primer caracter del año

		Para I := 1 hasta 6 Hacer
			AVZ(Fecha,f);
		FinPara

		Mientras v <> "-" Hacer

