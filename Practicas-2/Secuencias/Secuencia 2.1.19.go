/**
*! CONSIGNA
*/
1. Dada una secuencia de caracteres, generar otra secuencia con todas las palabras de tres caracteres.
2. Informar el porcentaje de las mismas, sobre el total de palabras de la secuencia de entrada. 
3. Considerar que puede haber más de un blanco entre palabras.



 Accion 2.1.19 Es
 Ambiente

   sec1: Secuencia de caracteres
   sal: Secuencia de caracteres
   x: caracter
   v1: caracter 
   v2: caracter
   v3: caracter

   contP: entero
   contL: entero
   contP3: entero
   porcentaje: entero


Proceso
   ARR(sec)
   ARR(sal)
   CREAR(sal)
   AVZ(sec,v1)

   contP := 0
   contP3 := 0
   porcentaje := 0

   Mientras NFDS (sec) Hacer 
     contL := 0
     Mientras (x <> " ") Hacer
	     AVZ(sec,x)
	   FinMientras
     contP := contP + 1;
     AVZ(sec,x) 
     CL := v1
     contL := contL + 1;
     AVZ(sec,x)
     CL := v2
     contL := contL + 1;
     AVZ(sec,x)
     CL := v3
     contL := contL + 1;
     AVZ(sec,x)
     Si contL = 3 Entonces
       GRABAR(sal,v1)
       GRABAR(sal,v2)
       GRABAR(sal,v3)
       contP := contP + 1;
       GRABAR(sal," ")
     FinSi  
   FinMientras              
     CERRAR(sec)  
     CERRAR(sal)
     Esc("porcentaje de palabras con 3 letras:", porcentaje(contP3*100)/contP)
FinAccion       

     
       

     

    
	 

