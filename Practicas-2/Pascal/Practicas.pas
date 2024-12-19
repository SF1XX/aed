// program Practicas;
// begin
//     WriteLn(15+2);
//     WriteLn(15-2);
//     WriteLn(15/2);
//     WriteLn(15*2);
//     WriteLn(15 mod 2);
//     WriteLn(15 div 2);
// end.


// program op_usuario;
// var
//     n1 : Integer;
//     n2 : Integer;
// begin
//     WriteLn('ingrese un numero');
//     ReadLn(n1);
//     WriteLn('ingrese otro numero');
//     ReadLn(n2);
//     WriteLn('la suma de los dos numeros es:');
//     WriteLn(n1 + n2);
// end.


// program condicional_if;
// var
//     n1 : Integer;
//     n2 : Integer;
// begin
//     writeln('ingrese un numero');
//     ReadLn(n1);
//     writeln('ingrese otro numero');
//     ReadLn(n2);
//     if n1 > n2 then WriteLn('piola');
// end.


// program condicional_else;
// var
//     n1 : Integer;
//     n2 : Integer;
// begin
//     writeln('ingrese un numero');
//     readln(n1);
//     writeln('ingrese otro numero');
//     readln(n2);
//     if n1 > n2 then
//         begin
//             writeln('es mayor');
//         end
//     else
//         begin
//             writeln('es menor');
//         end    
// end.


program corte_de_control;
var
    archivo: textfile;
    calificaciones: Integer;
    suma, promedio: real;
     numerodecalificaciones, contadorsuperior: Integer;

begin
    assign(archivo, 'C:\Users\Luciano\Desktop\cortepascal.txt');
    reset(archivo);

    suma := 0;
    numerodecalificaciones := 0;

    while not Eof(archivo) do
    begin
        readln(archivo, calificaciones);
        suma := suma + calificaciones;
        numerodecalificaciones := numerodecalificaciones + 1;
    end;

    if numerodecalificaciones > 0 then
        promedio := suma / numerodecalificaciones
    else
        promedio := 0;

    reset(archivo);
    contadorsuperior := 0;

    while not Eof(archivo) do
    begin
        readln(archivo,calificaciones);
        if calificaciones >= promedio then        
            contadorsuperior := contadorsuperior + 1;
    end;

    close(archivo);
    writeln('promedio de calificaciones:', promedio);
    writeln('numero de estudiantes con calificacion mayor o igual al promedio:', contadorsuperior);
    writeln('presiones cualquier tecla para salir...');       



end.