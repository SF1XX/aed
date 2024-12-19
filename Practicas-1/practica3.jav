Accion ventas Es
Ambiente
  sec: secuencia de caracter;
  salida: secuencia de entero
  v: caracter;
  mes_par: {"0", "2", "4", "6", "8"}
  resguardo: caracter;

  cont_vendidas: entero;


Proceso
  ARR(sec);
  AVZ(sec,v);
  CREAR(salida);
  cont_vendidas := 0;

  Mientras NFDS(sec) Hacer
    Para I := 1 a 3 Hacer
      AVZ(sec,v);
    Finpara
    resguardo := v;

    Si resguardo En mes_par Entonces
      Segun v Hacer
        ="1":GABAR(salida,1)
        ="2":GABAR(salida,2)
        ="3":GABAR(salida,3)
        ="4":GABAR(salida,4)
        ="5":GABAR(salida,5)
        ="6":GABAR(salida,6)
        ="7":GABAR(salida,7)
        ="8":GABAR(salida,8)
        ="9":GABAR(salida,9)
        ="0":GABAR(salida,0)
      Finsegun
    Finsi
    AVZ(sec,v);
  Finmientras
  CERRAR(sec);
  CERRAR(salida);
FinAccion    


    