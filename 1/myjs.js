//alert("Tu as fais quoooi ?");

var nbclic=0 // Initialisation à 0 du nombre de clic
function CompteClic() // Fonction appelée par le bouton
{
    nbclic++; // nbclic+1
    if (nbclic > 1) // Plus de 1 clic
    {
	alert("Vous avez déjà cliqué ce bouton.\nLe formulaire est en cours de traitement... Patience");
    }
    else // 1 seul clic
    {
	alert("Premier Clic.");
    }
}

