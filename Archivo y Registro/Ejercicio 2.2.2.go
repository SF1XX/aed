Se dispone de una secuencia de facturas con los siguientes datos:

-Nº de Factura
-Fecha
-Total

Se desea un listado de las facturas cuyas fechas sean posteriores a una fecha dada, y cuyos
importes totales no alcancen los $1000, con el siguiente formato:

| Nro_Factura | Fecha | Importe_Total |

Accion 2.2.2 Es
Ambiente

  FACTURAS = REGISTRO
    Num_Fac: N(9)
    Total: N(10)
    Fecha: FECHAS
  FinRegistro

  FECHAS = REGISTRO
      aa: N(4)
      mm: 1..12
      dd: N(2)
  FinRegistro

  ARCHIVO_FAC: Archivo de facturas
  REG_FAC: FACTURA
  FECHA_DADA: FECHA

Proceso
  ABRIR E/(ARCHIVO_FAC)
  LEER (ARCHIVO_FAC,REG_FAC)
  ESCRIBIR("Ingrese fecha con el formato aa/mm/dd")
  LEER(FECHA_DADA.aa, FECHA_DADA.mm, FECHA_DADA.dd)
  ESCRIBIR("| Nro Factura | Fecha | Importe Total |")

  MIENTRAS NFDA(ARCHIVO_FAC) HACER
    SI REG_FAC.FECHA > FECHA_DADA ENTONCES
      SI REG_FAC.Total < 1000 ENTONCES
        ESCRIBIR(REG_FAC.Num_Fac, REG_FAC.Fecha.aa, REG_FAC.Fecha.mm, REG_FAC.Fecha.dd, REG_FAC.Total)
      FinSi 
    FinSi     
    LEER (ARCHIVO_FAC,REG_FAC)
  FinMientras
  CERRAR(ARCHIVO_FAC)  
  
FinAccion    
