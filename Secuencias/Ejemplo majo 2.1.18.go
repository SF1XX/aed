Acción pinturas es
AMBIENTE
    seca, seco, sal : secuencia de caracteres
    va,vo : caracter
    num_obras,cont_vend,mayor : entero
    función conv_entero(v: caracter): entero
        segun v hacer
                    "0":conv_entero:=0
                    "1":conv_entero:=1
                    "2":conv_entero:=2
                    "3":conv_entero:=3
                    "4":conv_entero:=4
                    "5":conv_entero:=5
                    "6":conv_entero:=6
                    "7":conv_entero:=7
                    "8":conv_entero:=8
                    "9":conv_entero:=9
        finsegun
    finfuncion
	funcion formar_numero(sec:secuencia de caracteres, v: caracter):entero ES
	AMBIENTE
		numero: entero
	PROCESO
	  numero:= 0
        MIENTRAS va <> "?" HACER
            numero := numero * 10 + conv_entero(v)
            AVZ(sec,v) // ?
        FM	
	formar_numero:= numero
    FINPROCEDIMIENTO
	
PROCESO
    mayor:=LV
    ARR(seca)
    ARR(seco)
    crear(sal)
    AVZ(seca,va)
    AVZ(seco,vo)

    MIENTRAS NFDS(seca) HACER
        SI va = “R” ENTONCES
            AVZ(seca,va)
            MIENTRAS va <> “+” HACER
                GRABAR(sal,va)
                AVZ(seca,va)
            FINMIENTRAS
            GRABAR(sal,”R”)

            MIENTRAS va <> “+” HACER //ubicacion
                AVZ(seca,va)
            FM
            AVZ(seca,va)

            MIENTRAS v <> “+” HACER //edad
                AVZ(seca,va)
            FM
            AVZ(seca,va) //quedo en cantidad de obras
            num_obras := formar_numero(seca,va)
            cont_vend := 0
            
            PARA i:= 1 hasta num_obras HACER
                renacentistas := renacentistas + 1
                obras := obras + 1
                REPETIR "," HACER
                    GRABAR(sal,vo)
                    AVZ(seco,vo)
                FM
                GRABAR(sal,vo)
                AVZ(seco,vo)
                MIENTRAS vo <> "," HACER
                    GRABAR(sal,vo)
                    AVZ(seco,vo)
                FM
                GRABAR(sal,vo)
                AVZ(seco,vo)
                MIENTRAS vo <> "," HACER
                    AVZ(seco,vo)
                FM
                AVZ(seco,vo)
                SI vo = "V" ENTONCES
                    cont_vend := cont_vend + 1
                FINSI
                AVZ(seco,vo) // estoy en "/"
                AVZ(seco,vo) // avanzo /
            FINPARA
           
        SINO //no es renacentista
            PARA i:= 1 a 4 HACER
                MIENTRAS v <>"+"
                    AVZ(seca,va)
                FM
                AVZ(seca,va)
            FPARA
            //estoy en cantidad de obras
            num_obras := formar_numero(seca,va)
            cont_vend := 0
            PARA i:= 1 hasta num_obras HACER
                obras := obras + 1
                PARA i := 1 a 3 HACER  
                    MIENTRAS v <> "," HACER
                        AVZ(seco,vo)
                    FM
                    AVZ(seco,vo)
                FPARA
                SI vo = "V" ENTONCES
                    cont_vend := cont_vend +1
                FINSI
                AVZ(seco,vo) //en /
                AVZ(seco,vo)


        FINSI
        SI cont_vend > mayor ENTONCES
                mayor := cont_vend
        FINSI
    FMientras
    ESC("La mayor cantidad de obras vendidas por un artista es", mayor)
    ESC("El procentaje de obras renacentistas es", cont_vend)
