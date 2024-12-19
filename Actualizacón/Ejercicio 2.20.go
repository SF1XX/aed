Accion 2.20 Es
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
  
  EXAMENES = REGISTRO
    Nro_Legajo: N(6)
    Carrera: {"ISI","IEM","IQ"}
    Cod_Materia: N(10)
    Fecha_Examen: FECHA
    Nota: N(2)
  Finregistro
  
  Procedure Leer_Movimiento Es
    Leer(Arch_Ex, reg_ex);
    Si FDA(Arch_Ex) Entonces
      reg_ex:= HV
    Finsi 
  Finprocedure
  
  Procedure Leer_Maestro Es 
    Leer(Arch_Fac, reg_Fac);
    Si FDA(Arch_Fac) Entonces 
      reg_fac:= HV
    FinSi 
  Finprocedure 

  Arch_Fac: Archivo facultad ordenado por NLegajo;
  Arch_Ex: Archivo de movimiento ordenado por NLegajo;
  