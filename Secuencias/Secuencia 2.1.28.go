1. Generar una secuencia de salida con información de los artistas Renacentistas. 
La secuencia debe contener el nombre del artista, su estilo de arte, seguido de sus obras 
(nombre y año de creación). 

SECUENCIA ARTISTAS: M+RENE BARTOL+ROSARIO+34+2?R+JUAN B JUSTO+NEUQUEN+61+5?……..

SECUENCIAS OBRAS: SOL Y PARANA,1997,$913,V/GRITO DE ESPERANZA,2003,$235,R/PENAS,1997,$781,V/………

2. Al final del proceso informar:
a) la mayor cantidad de obras vendidas por un artista.
b) el porcentaje de obras de artistas "renacentistas" sobre el total de obras.

Accion 2.1.28 Es
Ambiente
  ARTISTAS: Secuencia de caracteres
  OBRAS: Secuencia de caracteres
  SALIDA: Secuencia de caracteres

  v1, v2, v3: caracteres
  cantidad: Entero
  porcentaje: real

  FUNCION conv_entero(v:caracter):entero
    SEGUN v Hacer
	   FUNCION convertir(v:caracter): entero
    SEGUN v Hacer
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
  ARR(ARTISTAS)
  ARR(OBRAS)
  CREAR(SALIDA)
  AVZ(ARTISTAS,v1)
  AVZ(OBRAS,V2)

  num_obras,cont_vend,mayor : entero


	

	  



