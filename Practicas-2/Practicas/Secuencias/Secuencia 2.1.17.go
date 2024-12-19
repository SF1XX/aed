Se desea calcular el costo de un telegrama, que se determina en función del número de palabras (que vale V1 cada una), salvo que el promedio de letras por palabra supere las cinco letras, caso en que cada palabra vale V2.

Accion 2.1.17 Es
Ambiente
  Sec: Secuencia de caracteres;
  v: Caracter
  V1: Real;
  V2: Real;
  CL,CP: Entero;
  Promedio: Real;

Proceso
  Esc("Ingrese valor de V1")
  Leer(V1);
  Esc("Ingrese valor de V2")
  Leer(V2);

  ARR(Telegrama);
  AVZ(Telegrama,Palabra);
  CL:= 0
  CP:= 0
    Mientras NFDS (Sec) Hacer
      Mientras v = " " Hacer
        AVZ(Sec,v)
      Fin_Mientras
      CP:= CP + 1
      Mientra NFDS (Sec) ^ (v <> "") Hacer
        CL:= CL + 1
        AVZ(Sec,v);
      Fin_Mientra
      Promedio:= (CL/CP);
      SI Promedio > 5 Entonces
        Esc("El valor del telegrama es:",CP*V2)
      SINO
        Esc("El valor del telegrama es:",CP*V1)     
      FS;
      Cerrar(Sec);
FIN_ACCION      
        