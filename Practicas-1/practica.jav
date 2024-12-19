/**Se cuenta con una secuencia de caracteres, con información de ALUMNOS, de acuerdo a la siguiente estructura:
-Nro Comisión: 1 caracter, indicar el número de comisión 1…6
-Información de los horarios de clase (siempre 12 caracteres, 6 para indicar teoría y los siguientes 6 para indicar práctica)
  Tipo: (T - Teoria, P- Practica)
  Hora: 2 Caracteres (00 … 23)
  Minutos: 2 Caracteres (00..59)
  Dia de la semana: “L”, “M”, “X”, “J”, “V”, “S”
  
  Información de los alumnos inscriptos, finaliza con carácter “%”
  Fecha Ingreso: 6 Caracteres
  Dia: 2 caracteres
  Mes: 2 caracteres
  Año: 2 caracteres
  Apellido y Nombre. Finaliza con el carácter “!”
  Legajo: 4 caracteres

  Se requiere que informe:
    a)Por cada comisión, cantidad de alumnos inscriptos
    b)Listado de todos los alumnos (apellido y nombre) que ingresaron en el año 2024

*/

Accion secuencia_alumnos Es
Ambiente
  sec: secuencia de caracter
  v: caracter
  comision: caracter;
  cont_alumnos: entero;
  año1, año2: entero;

Proceso
  ARR(sec);
  AVZ(sec,v);
  cont_alumnos := 0;
  Mientras NFDS (sec) Hacer
    cont_alumnos := cont_alumnos + 1;
    comision := v;
    AVZ(sec,v)

    Para I := 1 a 12 Hacer  
      AVZ(sec,v);
    Finpara

    Mientras v <> "%" Hacer
      cont_alumnos := cont_alumnos + 1;

      Para I := 1 hasta 4 Hacer
        AVZ(sec,v);
      Finpara  

      año1 := v;
      AVZ(sec,v);
      año2 := v;
      AVZ(sec,v); // avanzo hasta la primer letra del nombre

      Si (año1 = "2") ^ (año2 = "4") Entonces
        Mientras v <> "!" Hacer
          Esc(v);
          AVZ(sec,v);
        FinMientras
      sino
        Mientras v <> "!" Hacer
          AVZ(sec,v);
        FinMientras          
      FinSi

      AVZ(sec,v);

      Para I := 1 hasta 4 Hacer
        AVZ(sec,v);
      Finpara
    FinMientras
    Esc("Comision:", comision, "Cantidad de alumnos:", cont_alumnos);
  FinMientras
  cerrar (sec)
FinAccion    