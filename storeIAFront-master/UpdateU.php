<?php

$Nombre = $_POST['Nombre'];
$Apellido = $_POST['Apellido'];
$Telefono = $_POST['Telefono'];
$Identification = $_POST['Identification'];
$Clave = $_POST['Password'];
$Email = $_POST['Email'];
$Imagen = $_POST['Url'];
$Admin = empty($_POST['Administrador']);


$postData = array(  
  "nombre" => $Nombre ,
  "apellido" => $Apellido  ,
  "celular" => $Telefono ,
  "documento" => $Identification ,
  "clave" => $Clave ,
  "correo" => $Email, 
  "admin" =>  $Admin,
  "imagen" => $Imagen
   
);

echo "La identificacion es".$Identification;

// Create the context for the request
$context = stream_context_create(array(
  'http' => array(
      'method' => 'PUT',
      'header' => "Content-Type: application/json\r\n",
      'content' => json_encode($postData)
      )
));

// Send the request
$RespuestaApi = file_get_contents('http://52.229.9.122:8085/v1/user/'.$Identification, FALSE , $context);



echo $RespuestaApi;

?>