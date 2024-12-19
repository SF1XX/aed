Se posee una secuencia con los mensajes internos enviados entre los empleados de una empresa, la secuencia está formada de la siguiente 
manera: 
3 caracteres para el destino, 
3 caracteres para el origen, 
3 caracteres que indican la cantidad de caracteres que tiene el mensaje y el 
mensaje. No existe ningún carácter que separe un mensaje de otro. se pide:

1- Escribir una secuencia de salida con todos los mensajes (incluyendo origen y destino) que van hacia el área de Mantenimiento 
(que destino comience con 1).

2- Contar la cantidad de mensajes que se dirigen hacia el área de sistemas (que destino comience con 23).


Accion 2.1.32 Es 
Ambiente
	sec: secuencia de caracteres // secuencia prencipal
	sal: secuencia de caracteres // secuencia de salida
	v: caracter
	v2: caracter
	c1: caracter // primer caracter
	c2: caracter //segundo caracter
	c3: caracter // tercer caracter

	/**
	*! Resguardos
	*/
	destino: caracter 
	origen: caracter
	cantidad: caracter
	mensaje: caracter

	contDes: entero // contador de mensajes con sierto destino
	contMa: entero // contador de mansejes que van a matenimiento
	contM: entero // contador de mensejes general

Proceso
	ARR(sec)
	AVZ(sec,v)
	CREAR(sal)
	

	Mientras NFDS (sec) Hacer  // mientras no finalice la secuencia principal
		contMa := 0  // cunento los mensejes que van a mantinimiento
		Si NFDS (sec) Entonces
			destino := v  // resguardo el destino
			Mientras (v <> ".") Entonces
				destino := c1  // resguardo el primer caracter
				AVZ(sec,v)  // avanzo al segundo caracter
				destino := c2  // resguardo el segundo caracter
				AVZ(sec,v)  // avanzo al tercer caracter
				destino := c3  // resguardo el tercer caracter
				Si (c1 = "1") Entonces  // si el primer caracter es igual a 1 entonces cuento el menseje con destino 1
					contDes := contDes + 1;
					GRABAR(sal,v2)
					AVZ(sec,v)
					Para origen := 1 hasta 3 Hacer // grabo el origen
						contaOri: contaOri + 1;
						GRABAR(sal,v2)
						AVZ(sec,v)
					FinPara	
				FinSi	
			FinMientras	
					




