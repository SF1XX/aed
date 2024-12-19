Accion supermercado Es
Ambiente
  sec: secuencia de caracter
  sec2: secuencia de caracter
  sal: secuencia de caracter

  v: caracter
  v2: caracter

  u_mes1, u_mes2: caracter  // caracteres que ingresa el usuario
  mes1, mes2: caracter  // resguardos
  r_envio: caracter
  stock: entero

  cant_entregadas_en_suc: entero  // Cant. unid entregadas en suc
  cant_enviadas_a_domicilio: entero  // Cant. unid enviadas a domicilio
  cant_vendidas: entero

  Funcion converti_a_entero(v:caracter) Hacer
    segun v Hacer
      = "0":converti_a_entero := 0
      = "1":converti_a_entero := 1
      = "2":converti_a_entero := 2
      = "3":converti_a_entero := 3
      = "4":converti_a_entero := 4
      = "5":converti_a_entero := 5
      = "6":converti_a_entero := 6
      = "7":converti_a_entero := 7
      = "8":converti_a_entero := 8
      = "0":converti_a_entero := 9
    Finsegun
  Finfuncion  

Proceso
  ARR(sec);
  AVZ(sec,v);
  ARR(sec2);
  AVZ(sec2,v2);
  CREAR(sal);

  Esc("ingrese el primer caracter del mes");
  Leer(u_mes1);
  Esc("ingrese el segundo caracter del mes");
  Leer(u_mes2);

  Mientras NFDS(sec) ^ NFDS(sec2) Hacer
    Para I := 1 a 7 Hacer
      AVZ(sec);
    Finpara

    stock := 0;

    Para I = 2 a 0 Hacer
      stock := stock * 10 + converti_a_entero(v);
      AVZ(sec,v);
    Finpara  

    cant_entregadas_en_suc := 0;
    cant_enviadas_a_domicilio := 0;

    Mientras v <> "#" Hacer
      AVZ(sec2,v2); // segundo caracter del dia
      AVZ(sec2,v2); // primer caracter del mes
      mes1 := v2;
      AVZ(sec2,v2);
      mes2 := v2;
      AVZ(sec2,v2);
      AVZ(sec2,v2);
      r_envio := v2;
      cant_vendidas := 0;

      Para I := 1 a 0 Hacer
        cant_vendidas := cant_vendidas * 10 + converti_a_entero(v2);
        AVZ(sec2,v2);
      Finpara
        

