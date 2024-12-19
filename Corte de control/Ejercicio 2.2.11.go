Se dispone de un archivo con datos de los alumnos de la U.T.N. con la siguiente información 
para cada alumno:

PRODUCTOS Ordenado por Sexo, Carrera, Nro_Legajo

|- Sexo -|- Carrera -|- Nro_Legajo -| Fecha_Ingreso | Total_Materias_Aprobadas |  

Se desea un listado con el siguiente detalle por sexo, carrera y total general

1.Total de alumnos que ingresaron en el año 2009 con más de 20 materias aprobadas.
2.Total de alumnos que ingresaron en el año 2009 con menos de 20 materias aprobadas


Accion 2.2.11 Es
Ambiente

  Fecha = REGISTRO
    aa: N(4)
    mm: 1..12
    dd: N(2)
  FinRegistro

  Facultad = REGISTRO
    Clave = REGISTRO
      Sexo: {"F","M"}
      Carrera: {"ISI","IEM","IQ","LAR"}
      Nro_Legajo: N(6)
      Fecha_Ingreso: Fecha
      Tot_Mat_Aprob: N(2)
    FinRegistro
  FinRegistro
  Arch_Fac: Archivo de Facultad ordenado por Clave;
  Reg_Fac: Facultad

  c_totmas_car, c_totmen_car: entero  // total por carrera
  c_totmas_sex, c_totmen_sex: entero  // total por sexo  
  tg_totmas, tg_totmen: entero  // total general

  res_carrera: {"ISI","IEM","IQ","LAR"}
  res_sexo: {"F","M"}

  Subaccion corte_sexo Es
    Esc("Para el sexo:",res_sexo);
    Esc("El total de alumnos que ingresaron en el 2009 con mas de 20 materias aprobadas es", c_totmas_sex)
    Esc("El total de alumnos que ingresaron en el 2009 con mas de 20 materias aprobadas es", c_totmen_sex)
    tg_totmas := c_totmas_sex
    tg_totmen := c_totmen_sex
    res_sexo: Reg_Fac.Clave.Sexo
    c_totmas_sex := 0
    c_totmen_sex := 0
  FinSub

  Subaccion corte_carrera Es
    corte_sexo();
    Esc("Para la carrera",res_carrera)
    Esc("El total de alumnos que ingresaron en el 2009 con mas de 20 materias aprobadas es", c_totmas_car)
    Esc("El total de alumnos que ingresaron en el 2009 con menos de 20 materias aprobadas es", c_totmen_car)
    c_totmas_sex := c_totmas_sex + c_totmas_car 
    c_totmen_sex := c_totmen_sex + c_totmen_car
    res_carrera: Reg_Fac.Clave.Carrera
    c_totmas_car := 0
    c_totmen_car := 0
  FinSub

Proceso
  Abrir E/(Arch_Fac)
  Leer (Arch_Fac,Reg_Fac)
  res_carrera := Reg_Fac.Clave.Carrera
  res_sexo := Reg_Fac.Clave.Sexo
  c_totmas_car := 0
  c_totmen_car := 0
  c_totmas_sex := 0
  c_totmen_sex := 0
  tg_totmas := 0
  tg_totmen := 0

  MIENTRAS NFDA(Arch_Fac) Hacer
    SI (res_sexo <> Reg_Fac.Clave.Sexo) Entonces
      corte_sexo
    SINO 
      SI (res_carrera <> Reg_Fac.Clave.Carrera) Entonces
        corte_carrera
      FinSI
    FinSI  
    
    SI Reg_Fac.Fecha_Ingreso.aa = 2009
      SI Reg_Fac.Tot_Mat_Aprob > 20
        tg_totmas := tg_totmas + 1
      SINO 
        tg_totmen := tg_totmen +1
      FinSI
    FinSI
    Leer(Arch_Fac,Reg_Fac)      
  FinMientras
  corte_sexo()
  Esc("El total de alumnos que ingresaron en el 2009 con mas de 20 materias aprobadas es", tg_totmas)
  Esc("El total de alumnos que ingresaron en el 2009 con menos de 20 materias aprobadas es", tg_totmen)
  Cerrar(Arch_Fac)
FinAccion

    
    
    




