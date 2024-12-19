Dada una secuencia de letras del alfabeto que finaliza con una marca '*', contar cuantas letras "A" hay en la secuencia.

Accion 2.1.1 Es
Ambiente
     Sec: Secuencia de caracteres;
     cont: entero;
     v: Caracter;
Proceso
     ARR(Sec)
     AVZ(Sec,v)
     cont:=0
     Minetras v <> "*" Hacer
       Si v = "A" Entonces
         Cont:=Cont + 1
       Fin_Si
       AVZ(Sec,v)  
     Fin_Mientras
     Esc("La cantidad de latras A que hay en la secuencia es de:" cont)
     Cerrar(Sec)     
Fin_Accion     

