/**La plataforma UTN-eLearning brinda capacitaciones virtuales, las cuales son dictadas por
profesores de las distintas regionales de la UTN. Para el seguimiento de los mismos registra
las inscripciones a los cursos en un archivo llamado INSCRIPCIONES. El archivo está
ordenado por Facultad, Tema, Curso, Turno, fecha inicio e ID curso. Además, contiene la
duración del curso en horas (es un entero), el cupo (cantidad de alumnos que puede tener
como máximo) y la cantidad de inscriptos actuales.
*/

|- Facultad -|- Tema -|- Curso -|- Turno -|- fecha_inicio -|- ID_curso -| Duración | Cupo | Cant Inscriptos |

/**
*! CONSIGNAS
*/
1. Generar el siguiente archivo, con el siguiente formato, considerando solo los cursos
que comienzan en el siguiente mes o después. Puede usar una función
fecha_sistema() para obtener la fecha actual

| Facultad | Tema | Cant_Cursos |

2. Obtener totales parciales por curso de la cantidad de inscriptos, solo de los cursos
que ya han comenzado

3. Mostrar cuál es el tema que tenga la mayor cantidad de inscriptos.

4. Mostrar el porcentaje de cursos que son de un tema particular indicado por el usuario.

5. Obtener total general de inscriptos, discriminando por inscriptos a cursos que ya
comenzaron y cursos que aún no comienzan

/**
** PSEUDOCODIGO
*/
Accion consigna4 Es
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
      Turno: {"M", "T", "N"}
      fecha_inicio: FECHA
      ID_curso: AN(20)
    FinRegistro
      Duración: AN(20)
      Cupo: AN(3)
      Cant_Inscriptos: AN(6)
  FinRegistro

  Archivo1: Archivo inscripciones ordena por Clave
  reg1: INSCRIPCIONES

  SALIDA = REGISTRO
    Tema: AN(10)
	  ID_curso: AN(20)
	  Cant_Cursos: AN(20)
  FinRegistro

  Archivo2: Archivo salida ordena por Tema, ID_curso, Cant_Inscriptos;
  reg2: SALIDA
  
  tot_gen_iniciados: entero
  tot_gen_no_iniciados: entero
  tot_tema: entero
  tot_tema_iniciados: entero
  tot_curso_iniciados: entero
  


  res_fac: AN(20)
  res_tema: AN(10)
  res_curso: AN(10)
  res_cant_com: N(6)

  Subaccion corte_curso Es
    Esc("La cantidad de inscriptos por curso que ya han comenzado", res_curso "Total", tot_curso_com);
    tot_tema := tot_tema + tot_curso_com
    res_curso := reg1.Clave.Curso
    tot_curso_com := 0
  FinSub
  
  Subaccion corte_tema Es
    corte_curso()
      Esc("La cantidad de inscriptos por tema cuya duración es mayor a 40 horas", res_tema "Total", tot_tema);
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

  