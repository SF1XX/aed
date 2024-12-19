Accion fuchibol Es
Ambiente
  sec: secuencia de caracter
  sec2: secuencia de entero
  v, v2: caracter

  n = 26
  edad: entero;
  edad_may: 1..100;
  edad_men: 100..1:

  edad_juagadores: entero;
  promedio: real;

  equipo: caracter;

Mientras NFDS(sec) Hacer  
  Repetir
    AVZ(sec,v);
  Hasta que v = "#"  

  edad_juagadores := 0;

  Para I := 1 a n Hacer
   si edad > edad_may Entonces
     edad_juagadores := edad_juagadores + 1
   sino 
    edad  


    