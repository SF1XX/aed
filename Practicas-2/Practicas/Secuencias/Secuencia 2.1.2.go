Dada una secuencia de letras del alfabeto que finaliza con la letra "Z", contar cuantas consonantes hay en la secuencia.

Accion 2.1.2 Es
Ambiente
     Sec: Secuencia de caracteres
     v: Caracter
     cont: Entero
     Vocal= ("A","E","I","O","U")
Proceso
  ARR(Sec)
  AVZ(Sec,v)
  Cont:=0
  Mientras v <> "Z" Hacer
    Si v no es Vocal Entoces
      cont:= cont + 1
    Fin_Si
    AVZ(Sec,v)
  Fin_Mientras
  Esc("La cantidad de letras consonantes es:"cont);
  Cerrar(Sec)
Fin_Accion      
