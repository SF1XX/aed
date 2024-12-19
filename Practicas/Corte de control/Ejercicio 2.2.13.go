/**Un Casino de la región ha notado un incremento en los costos de las reparaciones de tragamonedas
en sus sucursales. Por ello solicitó un informe con la cantidad de reparaciones y sus costos, 
discriminados según tragamonedas, modelo, marca, sucursal y total general.

Se dispone de un archivo REPARACIONES, con el siguiente formato. Cada registro representa la 
reparación de un tragamonedas, tener en cuenta que puede existir más de una reparación por 
tragamonedas.

REPARACIONES Ordenado por Cod_Sucursal, Marca, Modelo, Cod_Tragamonedas

|- Cod_Sucursal -|- Marca -|- Modelo -|- Cod_Tragamonedas -| Fecha_Reparacion | Costo_Reparacion |

Realice el algoritmo para emitir el informe con los totales solicitados */

Accion 2.2.13 Es
Ambiente
  
  FECHA = REGISTRO  
    aa: N(4)
	  mm: 1..12
	  dd: N(2)
  FinRegistro

  REPARACIONES = REGISTRO 
    Clave = REGISTRO
      Cod_Sucursal: AN(20)
	    Marca: AN(10)
	    Modelo: AN(10)
	    Cod_Tragamonedas: AN(20)
  	FinRegistro
     	Fecha_Reparacion: FECHA
    	Costo_Reparacion: N(6)
  FinRegistro	  
  
  ArchivoR: Archivo de reparaciones ordenado por Clave;
  R_Reparacion: REPARACIONES 

  Fecha_r: FECHA 

  /**
  *? resgurdos
   */
  res_sucursal: AN(20)
  res_marca: AN(10)
  res_modelo: AN(10)
  res_codtragamonedas: AN(20)

  /**
  *? tatelales en cantidad
  */
  tot_codtragamonedas: entero
  tot_modelo: entero
  tot_sucursal: entero
  tot_marca: entero
  tot_gca: entero  // ** total general

  /**
  *? tatelales en costos
  */
  costo_codtragamonedas: entero
  costo_modelo: entero
  costo_sucursal: entero
  costo_marca: entero
  costo_gco: entero  // ** total general

  Subaccion corte_tragamonedas Es
    Esc("Para el tragamonedas de codigo", res_codtragamonedas)
    Esc("La cantidad de reparaciones es de ",tot_codtragamonedas )
    Esc("El costo de reparación es de ", costo_codtragamonedas)
    tot_modelo := tot_modelo + tot_codtragamonedas +  1
    costo_modelo := costo_modelo + costo_codtragamonedas
    res_codtragamonedas := R_Reparacion.Clave.Cod_Tragamonedas
    tot_codtragamonedas := 0
    costo_codtragamonedas := 0
  FinSubaccion
  
  Subaccion corte_modelo Es
    corte_tragamonedas()
    Esc("La cantidad de reparaciones para el modelo", res_modelo"es de", tot_modelo)
    Esc("El costo de reparacion es de ", costo_modelo)
    tot_marca := tot_marca + tot_modelo
    costo_marca := costo_marca + costo_modelo
    res_modelo := R_Reparacion.clave.Modelo
    tot_modelo := 0
    costo_modelo = 0
  FinSubaccion
  
  Subaccion corte_marca Es
    corte_modelo()
    Esc("La cantidad de reparaciones de la marca", res_marca"es de", tot_marca)
    Esc("El costo de reparacion es de ", costo_marca) 
    tot_sucursal := tot_sucursal + tot_marca
    costo_marca := costo_marca + costo_sucursal
    res_marca := R_Reparacion.Clave.Marca
    tot_marca := 0
    costo_sucursal := 0
  FinSubaccion
  
  Subaccion corte_sucursal Es
    corte_marca()
    Esc("La cantidad de reparaciones de la sucursal", res_sucursal"es de", tot_sucursal)
    Esc("El costo de reparacion es de ", costo_sucursal)
    tot_gca := tot_gca + tot_sucursal
    costo_gco := costo_gco + costo_sucursal
    res_sucursal := R_Reparacion.Clave.Cod_Sucursal
    tot_sucursal := 0
    costo_sucursal := 0
  FinSubaccion  
  
Proceso
  Abrir E/(ArchivoR)
  Leer (ArchivoR,R_Reparacion)
  tot_codtragamonedas := 0
  tot_modelo := 0
  tot_sucursal := 0
  tot_marca := 0
  tot_gca := 0  

  costo_codtragamonedas := 0
  costo_modelo := 0
  costo_sucursal := 0
  costo_marca := 0
  costo_gco := 0

  res_sucursal := R_Reparacion.Clave.Cod_Sucursal
  res_marca := R_Reparacion.Clave.Marca
  res_modelo := R_Reparacion.Clave.Modelo
  res_codtragamonedas := R_Reparacion.Clave.Cod_Tragamonedas

  Mientras NDFA(ArchivoR) Hacer
    Si (R_Reparacion.Clave.Cod_Sucursal <> res_sucursal ) Entonces
      corte_sucursal()
        sino
          Si (R_Reparacion.Clave.Marca <> res_marca) Entonces
            corte_marca()
              sino
                Si (R_Reparacion.Clave.Modelo <> res_modelo ) Entonces
                  corte_modelo()
                FinSi
          FinSi        
    FinSi
    /**
     *! Tratar registro
    */      
    tot_codtragamonedas := tot_codtragamonedas + 1
    costo_codtragamonedas := costo_codtragamonedas + R_Reparacion.Costo_Reparacion
    Leer(ArchivoR,R_Reparacion)
  FinMierntras
  corte_sucursal() 
  Esc("Total general de reparaciones: ", tot_gca)
  Esc("Total general de costos: ", costo_gco) 
  Cerrar(ArchivoR)
  
FinAccion  


    



  



