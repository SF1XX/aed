Se tiene información de todas las ediciones de la copa américa hasta la actual, con los
respectivos estadios en un archivo secuencial con el siguiente formato:
Estadios (Ordenado por País edición, año edición, ciudad sede y Nombre estadio)

|- País Edición -|- Año Edición -|- Ciudad Sede -|- Nombre estadio -| capacidad | cant partidos jugados |

Se le solicita:
1. Total de partidos jugados por país y ciudad sede en estadios con capacidad superior
a 40mil personas

2. Generar un archivo de salida con el siguiente formato
| País Edición | Año edicion | Cant partidos jugados |

3. Informar por estadio, el país, año de edición, ciudad, la cantidad de partidos jugados.

Accion corte_copa Es
Ambiente
  ESTADIOS = REGISTRO
    CLAVE = REGISTRO
      Pedicion: AN(30)
      Añoedicion: N(4)
      Ciudadsede: AN(30)
      NomEstadio: AN(30)
    Finregistro
    capacidad: N(5)
    Cantpartidos: N(3)
  Finregistro

  arch: archivo Ordenado por CLAVE
  reg: ESTADIOS

  SALIDA = REGISTRO
    Pedicion: AN(30)
    Añoedicion: N(4)
    Cantpartidos: N(3)
  Finregistro

  arch2: archivo de salida ordenado por Pedicion, Añoedicion, Cantpartidos;  
  reg2: SALIDA


  tot_partidos_pais: entero
  tot_partidos_ciudad: entero    
  tot_partidos_estadios: entero
  tot_general_partidos: entero
  tot_partidos_año: entero


  res_pais: AN(30);
  res_ciudad: AN(30);
  res_año: N(4);
  res_estadio: AN(30);
  

  Procedimiento corte_estadio Es
    Esc("La cantidad de partidos juagados en el estadio", res_estadio, "es de: ", tot_partidos_estadios);
    