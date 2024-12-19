Se dispone de una secuencia con los datos de los alumnos de la facultad:

-Apellido y Nombre
-Carrera (ISI - IEM - IQ)
-Nº de Legajo
-Fecha de Nacimiento
-Fecha de Ingreso
-D.N.I.
-Sexo (M o F)
-Fecha de último examen aprobado
-Nota
Se desea un listado de los mismos, con el siguiente formato:

|Nro_Legajo|Apellido_Nombre|Documento|Carrera|

Accion 2.2.1 Es
Ambiente

  FECHA = REGISTRO
    aa: N(4)
    MM: 1..12
    DD: N(2)
  FinRegistro

  FACULTAD = REGISTRO
    ApellidoyNombre: AN(50)
    Carrera: {"ISI","IEM","IQ"}
    NLegajo: N(6)
    FechadeNac: FECHA;
    FechadeIng: FECHA;
    DNI: N(8);
    Sexo {"M","F"}
    Fecultiexaaprob: FECHA;
    Nota: N(2)
  FinRegistro

  ARCH_FAC: Archivo de facultad
  Reg_FAC: FACULTAD

Proceso
  ABRIR E/(ARCH_FAC);
  LEER (ARCH_FAC,Reg_FAC);  
  ESCRIBIR("|Nro_Legajo|Apellido_Nombre|Documento|Carrera|");
  Mientras NFDA(ARCH_FAC) Hacer
    Esc(Reg_FAC.NLegajo,Reg_FAC.ApellidoyNombre,Reg_FAC.DNI,Reg_FAC.Carrera);
    LEER(ARCH_FAC,Reg_FAC);
  FinMientras
  CERRAR(ARCH_FAC);
FinACCION  
  