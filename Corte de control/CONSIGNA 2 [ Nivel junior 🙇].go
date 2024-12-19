/**La plataforma UTN-eLearning brinda capacitaciones virtuales, las cuales son dictadas por
profesores de las distintas regionales de la UTN. Para el seguimiento de los mismos registra
las inscripciones a los cursos en un archivo llamado INSCRIPCIONES. El archivo está
ordenado por Facultad, Tema, Curso, Turno, fecha inicio e ID curso. Además, contiene la
duración del curso en horas (es un entero), el cupo (cantidad de alumnos que puede tener
como máximo) y la cantidad de inscriptos actuales.
*/
|- Facultad -|- Tema -|- Curso -|- Turno -|- fecha_inicio -|- ID_curso -| Duración | Cupo | Cant Inscriptos |

1. Obtener totales parciales de cantidad de inscriptos por Facultad, Curso y Turno, solo
de los cursos cuya duración sea mayor a 40 horas.
2. Obtener total de inscriptos por Tema, solo si hay menos de 500 inscriptos
actualmente en el tema
3. Mostrar total general de inscriptos

Accion consigna2 Es
Ambiente

  FECHA = REGISTRO
    aa: N(4)
    mm: 1..12
    dd: 1.31
  FinRegistro

  INSCRIPCIONES = REGISTRO
    Clave = REGISTRO
      Facultad: AN(20)
      Tema: AN(10)
      Curso: AN(10)
      Turno: AN(10)
      fecha_inicio: FECHA
      ID_curso: AN(20)
    FinRegistro
      Duración: AN(20)
      Cupo: AN(20)
      Cant_Inscriptos: AN(20)
  FinRegistro

  Archivo1: Archivo inscripciones ordena por Clave
  reg1: INSCRIPCIONES

  tot_fac, tot_tema, tot_curso, tot_turno: entero
  tot_global: entero

  res_fac: AN(20)
  res_tema: AN(10)
  res_curso: AN(10)
  res_turno: AN(10)

  Subaccion corte_turno Es
    Esc("La cantidad de inscriptos por turno cuya duración es mayor a 40 horas", res_turno "Total", tot_turno);
    tot_curso := tot_curso + tot_turno
    res_turno := reg1.Clave.Turno
    tot_turno := 0
  FinSub

  Subaccion corte_curso Es
    corte_turno()
    Esc("La cantidad de inscriptos por curso cuya duración es mayor a 40 horas", res_curso "Total", tot_curso);
    tot_tema := tot_tema + tot_curso
    res_curso := reg1.Clave.Curso
    tot_curso := 0
  FinSub
  
  Subaccion corte_tema Es
    corte_curso()
	Si (tot_tema < 500) Entonces
      Esc("La cantidad de inscriptos por tema cuya duración es mayor a 40 horas", res_tema "Total", tot_tema);
	FinSi  
      tot_fac := tot_fac + tot_global
      res_tema := reg1.Clave.Tema
      tot_global := 0
  FinSub

  Subaccion corte_facultad Es
    corte_tema()
    Esc("La cantidad de inscriptos por facultad cuya duración es mayor a 40 horas", res_fac "Total", tot_fac);
    tot_global := tot_global + tot_fac
    res_fac := reg1.Clave.Facultad
    tot_fac := 0
  FinSub

Proceso
  Abrir E/(Archivo1)
  Leer(Archivo1,reg1)
  
  tot_turno := 0
  tot_fac := 0 
  tot_tema := 0 
  tot_curso := 0
  tot_global := 0

  res_fac := reg1.Clave.Facultad
  res_tema := reg1.Clave.Tema
  res_curso := reg1.Clave.Curso  
  res_turno := reg1.Clave.Turno

  Mientras NFDA(Archivo1) Hacer
    Si (res_fac <> reg1.Clave.Facultad) Entonces
      corte_facultad()
        sino
          Si (res_tema <> reg1.Clave.Tema) Entonces
            corte_tema()
              sino
                Si (res_curso <> reg1.Clave.Curso) Entonces
                  corte_curso()
				            sino
				              Si (res_turno <> reg1.Clave.Turno) Entonces
				                corte_turno()
				              FinSi
			          FinSi
		      FinSi
   	FinSi	  	  	  	

	  Si (reg1.Duración = 40) Entonces
      tot_turno := tot_turno + reg1.Cant_Inscriptos
	  FinSi
	  Leer(Archivo1,reg1)
  FinMientras
  corte_facultad()
  Esc("El total general de inscriptos", tot_global)
  Cerra(Archivo1)

FinAccion  		  