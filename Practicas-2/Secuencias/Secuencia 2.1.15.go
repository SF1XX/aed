Se desea saber la cantidad promedio de palabras que contienen las oraciones de una secuencia de 
caracteres. Supóngase que los puntos son siempre contiguos al último caracter de la palabra. 
La secuencia finaliza con una marca.


Accion 2.1.15 Es
Ambiente
  sec: Secuencia de caracteres
  v: caracter
  cont: entero  // contador de palabras
  cont2: entero  // contador de oraciones
  promedio: real

Proceso
  ARR(sec)
  AVZ(sec,v)
  cont := 0
  cont2 := 0
  
  MIENTRAS NFDS (sec) HACER
    MIENTRAS (v <> ".") HACER   
      MIENTRAS (v <> " ") HACER
	    AVZ(sec,v)
	  FinMientras
	  cont := cont + 1;
	  MIENTRAS (v <> ".") ^ (v <> " ") HACER
	    AVZ(sec,v)
	  FinMientras
	FinMientras
	cont2 := cont2 + 1;
	AVZ(sec,v)  	
  FinMientras
  CERRAR(sec)	
  Esc('cantidad de palabras:'cont)
  Esc('cantidad de oraciones:'cont2)
  Esc('promedio de palabras:'(cont/cont2))	  
FinAccion  
