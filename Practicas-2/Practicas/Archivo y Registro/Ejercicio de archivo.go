Accion Multas Es
Ambiente

  FECHA = REGISTRO
    aa: N(4)
    mm: 1..12
    dd: N(2)
  FinRegistro

  MULTA = REGISTRO
    IDMulta: N(5)
    FECHAM: FECHA
    DNI: N(8)
    MOTIVO: AN(100)
    MONTO: N(7.2)
  FinRegistro
  SMULTA, AMULTA: Archivo de multas ordenado por IDMulta
  SRMULTA, RMULTA: MULTA

   MasInfo = REGISTRO
     IDMulta: N(5)
     FECHAM: FECHA
     DNI: N(8)
     LUGAR: AN(200)
     HORA: N(2)
     PAGADO: (S-N)
   FinRegistro
   OTROS: Archivos de MasInfo ordenado por IDMulta
   ROTROS: MasInfo

Proceso
  ABRIR E/(AMULTA), ABRIR S/(SMULTA);
  LEER(AMULTA,RMULTA);
  MIENTRAS NFDA(AMULTA) HACER
    SI RMULTA.aa = 2023 ^ RMULTA.mm = 12 ENTONCES
      SMULTA := RMULTA
      ESC(SMULTA,SRMULTA)

      Esc("Ingrese lugar y hora donde se labro la multa")
      Leer(ROTROS.LUGAR)
      Leer(ROTROS.HORA)
      Esc("Se ha pagado la multa?")   
      Leer(ROTROS.PAGADO)
      ROTROS.IDMulta:=RMULTA.IDMulta
      ROTROS.FECHAM:=RMULTA.FECHAM
      ROTROS.DNI:=RMULTA.DNI

      ESCRIBIR(OTROS,ROTROS)
    FinSi
    LEER(AMULTA,RMULTA)  
  FinMientras
  CERRAR(AMULTA)  
  CERRAR(OTROS)
FinACCION