Se tiene una secuencia de enteros que contiene los números de teléfono de los clientes de una empresa,cada número tiene 10 dígitos. La empresa necesita organizar los clientes de acuerdo al último dígito de su número telefónico, generar una secuencia con los números de las personas que termine con 5, 6 o 7.

Accion Ejercicio1 Es
Ambiente
  Sec: Secuencia de enteros;
  Sec2: Secuencia de enteros;
  v: Entero;
  ult_digito: Entero
  Valore = {"5","6","7"};

Proceso
  ARR(Sec)
  CREAR(Sec2)
  AVZ(Sec,v)
  Mientras NFDS (Sec) Hacer
     ult_digito:= v MOD 10
     SI ult_digito en Valores Entoces
       GRABAR(Sec2,v)
     FS
     AVZ(Sec,v)
  Fm    
  CERRAR(Sec);
  CERRAR(Sec2);
Fin_Accion  


Se tiene una secuencia de caracteres que contiene los numeros de telefono de los clientes de una empresa,cada número tiene 10 dígitos. La empresa necesita organizar los clientes de acuerdo al último dígito de su número telefónico, generar una secuencia con los números de las personas que inician en un número ingresado por usuario.

Accion Ejercicio2 Es
Ambiente
  Sec: Secuencia de caracteres;
  Sec2: Secuencia de enteros;
  v: Caracter; num: Caracter;
  i: Enetero;
  
Proceso
  ARR(Sec);
  CREAR(Sec2);
  AVZ(Sec,v):
  Esc("Ingrese un numero");
  Leer(num)
  Mientras NFDS (Sec) Hacer
    SI v = num Entoncec
      PARA i:= 1 a 10 Hacer
        GRABAR(Sec2,v)
        AVZ(Sec,v)
      Fin_Para
    SINO
      PARA i:= 1 a 10 Hacer       
        AVZ(Sec,v)
      Fin_Para
    Fin_SI  
  Fin_Mientras
  CERRAR(Sec);
  CERRAR(Sec2);
Fin_Accion      


