 <?php
$usuario = $_POST['Usuario'];
$contraseña = $_POST['Contraseña'];

// The data to send to the API
$postData = array(

      'correo'=> $usuario,
      'clave'=> $contraseña
      
);

// Create the context for the request
$context = stream_context_create(array(
    'http' => array(
        'method' => 'POST',
        'header' => "Content-Type: application/json\r\n",
        'content' => json_encode($postData)
    )
));

// Send the request
//$RespuestaApi = file_get_contents('http://52.229.9.122:8085/v1/user/login', FALSE, $context);

// Check for errors
if($RespuestaApi === FALSE){
    die('Error');
}

// Decode the RespuestaApi
$RespuestaApiData = json_decode($RespuestaApi, TRUE);

// Print the date from the RespuestaApi
//echo $RespuestaApiData;

$RespuestaApiData = "insert success!";

if($RespuestaApiData == "insert success!")

{ 
     header("Location: http://52.229.9.122/storeIAFront/tables.php");

}

}
else
{

  
  header("Location: http://52.229.9.122/storeIAFront/index.php");

}

?>