<?php

$Document = $_POST['Documento'];




?>


<!DOCTYPE html>
<html lang="en">

<head>

  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta name="description" content="">
  <meta name="author" content="">

  <title>Marketplace - Register</title>

  <!-- Custom fonts -->
  <link href="vendor/fontawesome-free/css/all.min.css" rel="stylesheet" type="text/css">
  <link href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i" rel="stylesheet">

  <!-- Custom styles -->
  <link href="css/sb-admin-2.min.css" rel="stylesheet">

</head>

<body class="bg-gradient-primary">

  <div class="container">

    <div class="card o-hidden border-0 shadow-lg my-5">
      <div class="card-body p-0">
        <!-- Nested Row within Card Body -->
        <div class="row">
          <div class="col-lg-5 d-none d-lg-block bg--image">
          <img  src="https://render.fineartamerica.com/images/rendered/default/poster/10/8/break/images/artworkimages/medium/2/pug-with-big-eyes-photography-by-daniel-hans-peter-christensen.jpg"  width="470" height="450" >
          </div>
          <div class="col-lg-7">
            <div class="p-5">
              <div class="text-center">
                <h1 class="h4 text-gray-900 mb-4">Update a Service</h1>
              </div>
              <form class="user" method="POST" action="UpdateRegister.php">
                <div class="form-group row">
                  <div class="col-sm-6 mb-3 mb-sm-0">
                    <input type="text" class="form-control input-lg" name="Documento" placeholder="Identification">
                    <br>

                  <input  type="submit" value="Buscar" class="btn btn-primary btn-user btn-block">
                  <br>
                </div>
                
                <table width="100" class="table-bordered" id="dataTable"  cellspacing="0">
                  <thead>
                    <tr>
                      <th>Id</th>
                      <th>Nombre</th>
                      <th>Apellido</th>
                      <th>Celular</th>
                      <th>Correo</th>
                      <th>Documento</th>
                    </tr>
                    
                    <?php
                    if($Document == null)
                    {
                        $data = file_get_contents("http://52.229.9.122:8085/v1/user/");
                        $products = json_decode($data, true);
                       
                        
                        foreach ($products as $product) {?>
                  
                  
                  <tr>
                              <th><?php print_r($product['id']);  ?></th>
                              <th><?php print_r($product['nombre']);  ?></th>
                              <th><?php print_r($product['apellido']); ?></th>
                              <th><?php print_r($product['celular']);  ?></th>
                              <th><?php print_r($product['correo']); ?></th>
                              <th><?php print_r($product['documento']); }?></th>
                            </tr>
                        <?php
                    } else{
                        $data = file_get_contents("http://52.229.9.122:8085/v1/user/".$Document);
                        $products = json_decode($data, true);
                        ?>
                        <td><?php  print_r($products['id']);        ?></td>
                        <td><?php print_r($products['nombre']);     ?></td>
                        <td><?php print_r($products['apellido']);   ?></td>
                        <td><?php print_r($products['celular']);    ?></td>
                        <td><?php print_r($products['correo']);     ?></td>
                        <td><?php print_r($products['documento']);} ?></td> 
                     
                            
                         

               
                    </tfoot>
                </form> 
                </table>
                
                <br>

             
              <hr>
          

        
              <a href="UpdateUser.php?".$Document class="btn btn-primary btn-user btn-block">Actualizar</a>
             
              <a href="DeleteUser.php?$Document" class="btn btn-primary btn-user btn-block">Borrar</a>

             
              
              </form>
              <br>
              
           
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>

  <!-- Bootstrap core JavaScript-->
  <script src="vendor/jquery/jquery.min.js"></script>
  <script src="vendor/bootstrap/js/bootstrap.bundle.min.js"></script>

  <!-- Core plugin JavaScript-->
  <script src="vendor/jquery-easing/jquery.easing.min.js"></script>

  <!-- Custom scripts for all pages-->
  <script src="js/sb-admin-2.min.js"></script>

</body>

</html>