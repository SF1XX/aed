sec1: CodArtCodRubroStockNombreArticulo&CodArtCodRubroStockNombreArticulo&CodArtCodRubroStockNombreArticulo&FDS
12345L789Detergente Magistral&23456F078Jamon Iberico& [...] &FDS
sec2: DiaMesFPFEUVDiaMesFPFEUVDiaMesFPFEUV#DiaMesFPFEUVDiaMesFPFEUVDiaMesFPFEUV# [...] #


Accion eje1 Es
Ambiente
	sec1: secuencia de caracteres;
	sec2: secuencia de caracteres;
	sal: secuencia de caracteres;

	v1: caracter;
	v2: caracter;
	v3: caracter;
	m1: caracter;
	m2: caracter;

	funcion convertir (v:caracter): entero
		segun v hacer
			="1":convertir:= 1
			="2":convertir:= 2
			="3":convertir:= 3
			="4":convertir:= 4
			="5":convertir:= 5
			="6":convertir:= 6
			="7":convertir:= 7
			="8":convertir:= 8
			="9":convertir:= 9
			="0":convertir:= 0
		FinSegun
	FinFuncion		

Proceso
	ARR(sec1);
	ARR(sec2);
	AVZ(sec1,v1);
	AVZ(sec2,v2);
	CREAR(sal);
	
	Esc("Ingrese un mes para generar un informe, separado por 2 digitos");
	Leer(m1,m2);

	Mientras NFDS(sec1);
		Para i := 1 a 5 hacer
			AVZ(sec1,v);
		FinPara
		
		AVZ(sec1,v);

		Para i := 

