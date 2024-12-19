Accion corte_n1 Es
Ambiente
	fecha = registro
		aa: N(4);
		mm: N(2);
		dd: N(2);
	finreg

	stock = registro
		clave = registro	
			codsuc: N(2);
			rubro: AN(20);
			codart: N(5);
		finreg
		fechaultrep: fecha
		stockdeseg: 1..100;
		stockactual: 1..100;
	finreg

	arch: archivo ordenado por clave
	reg: stock
	
	salida = registro
		codsuc: N(2);
		rubro: AN(20);
		stocklart: 1.100;
	finreg
	
	arch2: archivo de salida
	reg2: salida

	res_sucursal: N(2);
	res_rubro: AN(20);
	res_articulo: N(5);

	tot_sucursal: entero;
	tot_rubro: entero;
	tot_gral: entero;
	tot_repos: entero;

	procedimiento corte_rubro Es
		Esc("Los articulos que estan por debajo del stock de seguridad son: ", tot_rubro);

		sal.codsuc := res_sucursal;
		sal.rubro := res_rubro;
		sal.stocklart := tot_rubro;
		grabar(arch2,sal);

		tot_sucursal := tot_sucursal + tot_rubro;
		res_rubro := reg.clave.rubro;
		tot_rubro := 0;
	finprocedimiento
	
	procedimiento corte_sucursal Es	
		corte_rubro();
		Esc("Los articulos que entan por debajo del stock de seguridad son: ", tot_sucursal);
		tot_gral := tot_gral + tot_sucursal;
		res_sucursal := reg.clave.codsuc;
		tot_sucursal := 0,
	finprocedimiento
	
Proceso
	abrir E/(arch);
	leer(arch,reg);
	abrir S/(arch2);	

	res_sucursal:= reg.clave.codsuc;
	res_rubro:= reg.clave.rubro;
	res_articulo:= reg.clave.codart;

	tot_sucursal := 0;
	tot_rubro := 0;
	tot_gral := 0;
	tot_repos := 0;

	Mientras NFDA(arch) hacer
		si (res_sucursal<>reg.clave.codsuc) entonces
			corte_sucursal();
		sino
			si (res_rubro<>reg.clave.rubro) entonces
				corte_rubro();
			finsi
		finsi

		si (reg.stockactual <> reg.stockdeseg) entonces
			tot_rubro := tot_rubro + 1;
			Esc("Articulo:", reg.clave.codart, "Sucursal:", reg.clave.codsuc, "Rubro:", res.clave.rubro);
		finsi
		
		leer(arch,reg);
	finmientras
	Esc("El total general de articulos con stock actual por debajo del de seguridad es ", tot_gral);
	cerrar(arch);
	cerrar(arch2);
FinAccion	


