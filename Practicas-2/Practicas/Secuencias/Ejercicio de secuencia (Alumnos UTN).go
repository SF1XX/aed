Se cuenta con una secuencia de caracteres, con información de ALUMNOS, de acuerdo a la siguiente estructura:
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

Accion Alumnos Es
Ambiente
  Sec: Secuencia de caracteres;
  v:caracter;
  comision: caracter; 
  cont_alumno, i: Entero
  aa_ingreso, aa_ingreso2: caracter;
Proceso
  ARR(Sec)
  AVZ(Sec,v)  
  Mientras NFDS (Sec) Hacer
    comision:= v;
    cont_alumno:= 0;
    AVZ(Sec,v)
    PARA i hasta 12 Hacer
      AVZ(Sec,v)
    Fin_Para
    Mientras v <> "%" Hacer
      cont_alumno:= cont_alumno + 1;
      PARA i:= 1 a 4 Hacer
        AVZ(Sec,v)
       Fin_Para

       aa_ingreso:= v
       AVZ(Sec,v)
       aa_ingreso2:= v
       AVZ(Sec,v)
       SI (aa_ingreso = "2")^(aa_ingreso2 = "4")
         Mientras v <> "!" Hacer
           ESC(v)
           AVZ(Sec,v)
         Fin_Mientras
       SINO
         Mientras v <> "!" Hacer  
           AVZ(Sec,v)
         Fin_Mientras
       Fin_Si
       AVZ(Sec,v)
       Para i:= 1 a 4 Hacer
         AVZ(Sec,v)
       Fin_Para
  Fin_Mientras
  ESC("Comision:", comision, "Cantidad de alumnos:", cont_alumnos)
  CERRAR(Sec)
Fin_Accion             
            
    

      
        

    
