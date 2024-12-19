Dada una secuencia con información de población de un PAIS:

POBLACION Ordenado por Provincia, Departamento, Ciudad, Barrio, ID_Casa

|- Provincia -|- Departamento -|- Ciudad -|- Barrio -|- ID_Casa -| Cantidad_Habitantes |

Generar una secuencia con información de los departamentos de esa provincia, cuyo registro tenga el
siguiente formato:
POBLACION_SALIDA
| Provincia | Departamento | Cantidad_Habitantes |

Accion 2.2.12 Es
Ambiente

  Poblacion = REGRISTRO
    Clave = REGRISTRO
      Provincia: AN(13)
	  Departamento: AN(20)
	  Ciudad: AN(20)
	  Barrio: AN(20)
	  ID_Casa: AN(4)
	FinRegistro
	Cantidad_H: N(7)
  FinRegistro
  
  Archivo1: Archivo de población ordenado por Clave
  r_poblacion: Poblacion

  Salida = REGISTRO
    Provincia: AN(13) 
	  Departamento: AN(20)
	  Cantidad_H: N(7)
  FinRegistro	

  Archivo2: Archivo de salida ordenado por Provincia, Departamento, Cantidad_H;
  r_salida: Salida

  res_Provincia: AN(13)
  res_Dep: AN(20)

  tot_pro, tot_dep: entero
  tot_g: entero

  Subaccion corte_departamento Es
    r_salida.Provincia := res_Provincia
	  r_salida.Departamento := res_Dep
	  r_salida.Cantidad_H := tot_g
	  Esc("Archivo2,r_salida") 
	  res_Dep := r_poblacion.Clave.Departamento
	  tot_g := 0 
  FinSub
  
  Subaccion corte_provincia Es
    corte_Departamento()
	  res_Provincia := r_poblacion.Clave.Provincia
  FinSub

Proceso
  Abrir E/(Poblacion)
  Abrir S/(Salida)
  Leer(Poblacion,r_poblacion)
  res_Dep := r_poblacion.Clave.Departamento
  res_Provincia := r_poblacion.Clave.Provincia
  tot_g := 0
  tot_pro := 0;
  tot_dep := 0;

  Mientras NFDA(Poblacion) Hacer
    SI (res_Provincia <> r_poblacion.Clave.Provincia) Entonces
	    corte_provincia
  	      SINO
	        SI (res_Dep <> r_poblacion.Clave.Departamento) Entonces
	          corte_departamento
	        FinSi	  
    FinSi
    
	tot_g := tot_g + r_poblacion.Clave.Cantidad_Habitantes
    Leer(Archivo1,r_poblacion)
  FinMientras
  corte_provincia()
  Cerrar(Archivo1)
  Cerrar(Archivo2)
  
FinAccion  

 
  
  










