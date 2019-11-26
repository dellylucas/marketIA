<?php

$Identificador = $_GET["Document"];

echo "Esto es Identificador:". $Identificador;


// Create the context for the request
$context = stream_context_create(array(
  'http' => array(
      'method' => 'DELETE',
      'header' => "Content-Type: application/json\r\n",
  )
));

// Send the request
$RespuestaApi = file_get_contents('http://52.229.9.122:8085/v1/user/'.$Identificador, FALSE , $context);

echo $RespuestaApi;

?>