/**La plataforma UTN-eLearning brinda capacitaciones virtuales, las cuales son dictadas por
profesores de las distintas regionales de la UTN. Para el seguimiento de los mismos registra
las inscripciones a los cursos en un archivo llamado INSCRIPCIONES. El archivo está
ordenado por Facultad, Tema, Curso, Turno, fecha inicio e ID curso. Además, contiene la
duración del curso en horas (es un entero), el cupo (cantidad de alumnos que puede tener
como máximo) y la cantidad de inscriptos actuales.
*/

|- Facultad -|- Tema -|- Curso -|- Turno -|- fecha_inicio -|- ID_curso -| Duración | Cupo | Cant Inscriptos |

1. Generar un archivo con el siguiente formato
|- Tema -|- ID_curso -| Cant Inscriptos |

2. Obtener totales parciales por Tema, Curso y Turno de la cantidad de cursos que ya
completaron el cupo máximo de estudiantes
3. Mostrar el total general de inscriptos
4. Mostrar el total general de inscriptos a cursos que ya completaron el cupo

Accion consigna3 Es
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
      Cant_Inscriptos: AN(20)
  FinRegistro

  Archivo1: Archivo inscripciones ordena por Clave
  reg1: INSCRIPCIONES

  SALIDA = REGISTRO
    Tema: AN(10)
  	ID_curso: AN(20)
  	Cant_Inscriptos: AN(20)
  FinRegistro

  Archivo2: Archivo salida ordena por Tema, ID_curso, Cant_Inscriptos;
  reg2: SALIDA

   /**
  *? Totalizadores
  */
  tot_tema, tot_curso, tot_turno: entero
  tot_global, tot_global_cupo: entero
  
  /**
  *? Resguardos
  */
  res_fac: AN(20)
  res_tema: AN(10)
  res_curso: AN(10)
  res_turno: {"M", "T", "N"}
  res_cupo: N(3)

  /**
  *! Subacciones de corte
  */
  Subaccion corte_turno Es
    Esc("La cantidad de inscriptos por turno de cursos que ya completaron el cupo máximo de estudiantes", res_turno "Total", tot_turno);
    tot_curso := tot_curso + tot_turno
    res_turno := reg1.Clave.Turno
    tot_turno := 0
  FinSub

  Subaccion corte_curso Es
    corte_turno()
    Esc("La cantidad de inscriptos por curso de cursos que ya completaron el cupo máximo de estudiantes", res_curso "Total", tot_curso);
    tot_tema := tot_tema + tot_curso
    res_curso := reg1.Clave.Curso
    tot_curso := 0
  FinSub
  
  Subaccion corte_tema Es
    corte_curso()
      Esc("La cantidad de inscriptos por tema de cursos que ya completaron el cupo máximo de estudiantes", res_tema "Total", tot_tema); 
      tot_fac := tot_fac + tot_global
      res_tema := reg1.Clave.Tema
      tot_global := 0
  FinSub

  Subaccion corte_facultad Es
    corte_tema()
    res_fac := reg1.Clave.Facultad
  FinSub

Proceso
  Abrir E/(Archivo1)
  Abrir S/(Archivo2)
  Leer(Archivo1,reg1)

  tot_turno := 0
  tot_tema := 0 
  tot_curso := 0
  tot_global := 0
  tot_global_cupo := 0

  res_fac := reg1.Clave.Facultad
  res_tema := reg1.Clave.Tema
  res_curso := reg1.Clave.Curso  
  res_turno := reg1.Clave.Turno

  Mientras NFDA (Archivo1) Hacer
    Si (res_fac <> reg1.Clave.Facultad) Entonces
      corte_facultad()
        sino
          Si (res_tema <> reg1.Clave.Tema) Entonces
            corte_tema()
              sino
                Si (res_curso <> reg1.Clave.Curso) Entonces
                    corte_curso()
                      Si (res_turno <> reg1.Clave.Turno) Entonces
                        corte_turno()
                      sino
                FinSi 
          FinSi           
    FinSi        

    // Tratar registro
    reg2.Tema := reg1.Clave.Tema
    reg2.ID_curso := reg1.Clave.ID_curso
    reg2.Cant_Inscriptos := reg1.Clave.Cant_Inscriptos
    Grabar(Archivo2, reg2)

    Si (reg1.Cant_Inscriptos <> reg1.Cupo) Entonces
      tot_global_cupo := tot_global_cupo + Cant_Inscriptos
      tot_turno := tot_turno + 1
    FinSi
    tot_global := tot_global + reg1.ClaveCant_Inscriptos  
    Leer(Archivo1,reg1)
  FinMientras
  
  corte_facultad()
  Esc("El total de alumnos inscriptos a cursos con el cupo completo es:", tot_global_cupo)
  Esc("El total de cursos que han completado el cupo es:", tot_global_cupo)
  Esc("El total de inscriptos actuales es:", tot_global)
  Cerrar(Archivo1)
  Cerrar(Archivo2)
FinAccion  
