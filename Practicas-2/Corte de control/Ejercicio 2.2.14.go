/**Dados el siguiente fichero con datos de un censo de la ciudad de Resistencia:

CENSO Ordenado por Radio, Manzana, Nro_Vivienda

|- Radio -|- Manzana -|- Nro_Vivienda -| Condicion_Vivienda | Cantidad_Habitantes |

Condición puede ser : Muy Buena, Buena o Mala. Obtener los siguientes totales de personas que 
habitan en viviendas cuya condición es muy buena: total en la ciudad y totales por Radio y Manzana).
*/

Accion 2.2.14 Es 
Ambiente
  CENSO = REGISTRO
    Clave = REGISTRO
	  Radio: N(6)
	  Manzana: N(4)
	  Nro_Vivienda: N(4)
	FinRegistro
	Condicion: {"Muy Buena", "Buena", "Mala"}
	Cantidad_H: N(6)
  FinRegistro

  Archivo1: Archivo ordenado por Clave
  Reg_Censo: CENSO

  tot_MB: entero
  tot_ciudad, tot_radio, tot_manzana: entero

  res_radio: N(6)
  res_manzana: N(4)

  Subaccion corte_manzana Es
     Esc("La cantidad de habitantes por manzana ",res_manzana "Total:", tot_manzana)
	 tot_radio := tot_radio + tot_manzana
	 res_manzana := Reg_Censo.Clave.Manzana
	 tot_manzana := 0
   FinSubaccion 	 

   Subaccion corte_radio Es
     corte_manzana()
	 Esc("La cantidad de habitantes por radio",res_radio "Total:", tot_radio)
	 tot_ciudad := tot_ciudad + tot_radio
	 res_radio := Reg_Censo.Clave.Radio
	 tot_radio := 0
  FinSubaccion	 

Proceso
  Abrir E/(Archivo1)
  Leer(Archivo1,Reg_Censo)
   
  tot_MB := 0
  tot_ciudad := 0 
  tot_radio := 0 
  tot_manzana := 0

  res_radio := Reg_Censo.Clave.Radio
  res_manzana := Reg_Censo.Clave.Manzana

  Mientras NFDA(Archivo1) Hacer
    Si (res_radio <> Reg_Censo.Clave.Radio) Entonces
	  corte_radio()
	    sino
		  Si (res_manzana <> Reg_Censo.Clave.Manzana) Entonces
		    corte_manzana()
		 FinSi	
	FinSi
	
	Si (Reg_Censo.Clave.Condicion = Muy Buena) Entonces
	  tot_manzana := tot_manzana + Reg_Censo.Cantidad_H
	FinSi 
	Leer(Archivo1,Reg_Censo)
  FinMienras
  corte_radio()
  Cerrar(Archivo1)
  Esc("Total en la ciudad: ", tot_ciudad)

FinAccion	







