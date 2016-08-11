<?php

if((isset($_SESSION['identifiant'])) && (isset($_SESSION['mot_de_passe'])))
{
  if(($_SESSION['identifiant'] == $Identifiant)
  	&& ($_SESSION['mot_de_passe'] == $Mot_de_passe))
   {
      header('Location: ./accueil.php');
      exit();
   }
}
?>