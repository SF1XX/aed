/**
La empresa Ideas Argentinas S.A. posee datos de sus empleados en dos secuencias de caracteres; la primera
secuencia (Sec1) formada por los nombres (uno por persona) de los empleados separados por blancos y la 
segunda secuencia (Sec2) posee los números de documento de cada uno de los empleados 
(palabras de 8 dígitos numéricos). Ambas secuencias poseen la misma cantidad de datos, 
es decir al primer nombre de la primera secuencia le corresponde el primer número de documento de la 
segunda secuencia y así sucesivamente. La secuencia de números de documentos no posee espacios en blanco 
ni ningún otro tipo de caracter que separe un número de documento de otro.

A la empresa le interesa tener en una nueva secuencia (Sec3) los datos de todas aquellas personas que 
cumplan la condición de que el nombre se encuentre en una posición impar. La nueva secuencia debe 
generarse de la siguiente manera: el número de documento seguido (sin espacios) por una coma y luego 
(sin espacios) por el nombre y seguido al nombre un #.
*/

Accion 2.1.22 Es
Ambiente
  Sec1: Secuencia de caracteres  // empleados separados por blancos
  Sec2: Secuencia de caracteres  // DNI
  Sec3: Secuencia de caracteres  // Salida

  v1, v2: caracter
  cont, I: entero

Proceso
  ARR(Sec1)
  ARR(Sec2)
  CREAR(Sec3)
  Avz(Sec1,v1)
  Avz(Sec2,v2)
  cont := 0

  Mientras NFDS(Sec1) ^ NFDS(Sec2) Hacer
    cont := cont + 1

    SI cont MOD 2 <> 0 Entonces
      PARA I := 8 Hacer
	    ESCRIBIR(Sec3,v3)
	    Avz(Sec2,v2)
	  FinPara
	  ESCRIBIR(Sec3,",")
	 
	  Mientras (v1 <> " ") Hacer
	    Avz(Sec1,v1)
      FinMientras

	  Mientras (V1 <> " ")
	    ESCRIBIR(Sec3,v3)
	    Avz(Sec1,v1)
	  FinMientras
	  ESCRIBIR(Sec3,"#")

    SINO 

      PARA I := 8 Hacer
	    Avz(Sec2,v2)
	  FinPara
	 
	  Mientras (v1 <> " ") Hacer
	    Avz(Sec1,v1)
      FinMientras

	  Mientras (V1 <> " ")
	    ESCRIBIR(Sec1,v1)
	  FinMientras
	  	 
    FinSi
  FinMientras
  CERRAR(Sec1)
  CERRAR(Sec2)
  CERRAR(Sec3) 	 
  Esc('La secuencia de salida ha sido creada con exito!')
FinAccion  



