/**
Realice un algoritmo para el enunciado del ejercicio 1.22, pero considerando que los datos que se deben 
copiar en la Sec3 son los de aquellas personas que cumplan la condición de que: el número de documento 
comienza con un número impar.
*/

Accion 2.1.23 Es
Ambiente
Sec1: Secuencia de caracteres  // empleados separados por blancos
  Sec2: Secuencia de caracteres  // DNI
  Sec3: Secuencia de caracteres  // Salida

  v1, v2: caracter
  pc, I: entero
  
  FUNCION convertir(v2:caracter): entero
    SEGUN v2 Hacer
      ="1": convertir := 1
	  ="2": convertir := 2
	  ="3": convertir := 3
	  ="4": convertir := 4
	  ="5": convertir := 5
	  ="6": convertir := 6
	  ="7": convertir := 7
	  ="8": convertir := 8
	  ="9": convertir := 9
	  ="0": convertir := 0 
	FinSegun   
  FinFuncion  

Proceso
  ARR(Sec1)
  ARR(Sec2)
  CREAR(Sec3)
  Avz(Sec1,v1)
  Avz(Sec2,v2)

  Mientras NFDS(Sec1) ^ NFDS(Sec2) Hacer
    SEGUN (v2) Hacer
      ="1": convertir := 1
	  ="2": convertir := 2
	  ="3": convertir := 3
	  ="4": convertir := 4
	  ="5": convertir := 5
	  ="6": convertir := 6
	  ="7": convertir := 7
	  ="8": convertir := 8
	  ="9": convertir := 9
	  ="0": convertir := 0 
	FinSegun  

    SI pc MOD 2 <> 0 Entonces
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