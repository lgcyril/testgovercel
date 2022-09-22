package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lgcyril/projet-collectif---mobile-front-daswifties/GO/api"
)

type urbex struct {
	ID          string  `json:"ID"`
	Name        string  `json:"Name"`
	City        string  `json:"City"`
	Description string  `json:"Description"`
	ImageName   string  `json:"ImageName"`
	Longitude   float64 `json:"Longitude"`
	Latitude    float64 `json:"Latitude"`
}

type allUrbexSpots []urbex

var urbexSpots = allUrbexSpots{
	{
		Name:        "Villa Bela Kiss",
		City:        "Oise",
		ID:          "1001",
		Description: "Digne d’un épisode d'American Horror Story, cette ancienne maison de campagne a tout d’une maison hantée ! Il se murmure qu’une jeune fille, Camille, serait morte à l’âge de 18 ans dans de terribles circonstances au sein de ce petit château de la Renaissance… Bonjour les frissons me direz-vous ! Mais outre l’histoire de cette maison, c’est son architecture figée dans le temps, ses papiers peints décrépis et ses meubles d’origine prenant la poussière qui attirent les photographes. Bien entendu, l’adresse de ce lieu restera secrète. Pour vous y rendre, à vous de trouver !",
		ImageName:   "villabelakiss",
		Longitude:   2.0823355,
		Latitude:    49.4300997,
	},
	{
		Name:        "Hôpital Earle Nelson",
		City:        "Paris",
		ID:          "1003",
		Description: "Cet immense hôpital abandonné se situe en plein cœur de Paris dans le 14e arrondissement ! Mais vous n’en saurez pas plus, le principe même de l’urbex – l’exploration urbaine – étant justement de garder les adresses secrètes. Contrairement à la plupart des spots abandonnés, quelques bâtiments de cet hôpital sont encore occupés, même si la plupart des pièces sont le nouveau terrain de jeu des graffeurs !",
		ImageName:   "hopitalearlenelson",
		Longitude:   2.331406,
		Latitude:    48.8371047,
	},
	{
		Name:        "Le fort de Vaujours",
		City:        "Vaujours",
		ID:          "1004",
		Description: "Construit à la fin du 19ème pour défendre Paris contre l’invasion de l’Armée Prussienne, cet ancien fort de l’Est parisien, situé entre la Seine-et-Marne (77) et la Seine-Saint-Denis (93), est sûrement l’un des sites d’urbex les plus dangereux de la région. De fait, entre 1955 et 1997, il fut occupé par le Commissariat à l’énergie atomique (CEA) comme centre de recherches. C’est notamment ici que fut mis au point le détonateur de la première bombe atomique française. Hautement surveillé, ce site est aujourd’hui encore marquée par de l’uranium sur près de 50 hectares. C’est pourquoi, les clichés de cet endroit se font si rares : s’y aventurent seulement les explorateurs les mieux équipés. À l’intérieur, les traces d’explosion, encore largement visibles, font réellement froid dans le dos !",
		ImageName:   "fortvaujours",
		Longitude:   2.5691921,
		Latitude:    48.9302563,
	},
	{
		Name:        "La plage",
		City:        "Médan",
		ID:          "1005",
		Description: "Un complexe aquatique oublié sur une île près de Médan. Grands toboggans et piscines vides. La principale difficulté est d'y accéder car il faut traverser la Seine. Une fois qu'on y est arrivé, la promenade est très agréable.",
		ImageName:   "laplage",
		Longitude:   1.99503,
		Latitude:    48.95275,
	},
	{
		Name:        "Station de métro Croix-Rouge",
		City:        "Paris",
		ID:          "1006",
		Description: "Nombreux d’entre vous marchent au-dessus d’elle sans savoir ce qui s’y trame. Cette station de la ligne 10 a été mise en service en 1923 et fermée en 1939 avant la 2nde Guerre Mondiale. Elle n’a jamais été réutilisée en raison de sa trop grande proximité avec les stations les plus proches, Sèvres-Babylone et Mabillon.",
		ImageName:   "croixrouge",
		Longitude:   2.3349394,
		Latitude:    48.8528857,
	},
	{
		Name:        "Le Sanatorium d’Aincourt",
		City:        "Aincourt",
		ID:          "1007",
		Description: "Partez de Paris direction Rouen en suivant la Seine et à mi-chemin vous passerez devant cet immense bâtiment. Un lieu déjà foulé par de nombreux aventuriers du dimanche mais qui attire toujours autant de monde, et on comprend pourquoi quand vu le nombre de pièces à explorer !",
		ImageName:   "sanatoriuMaincourt",
		Longitude:   1.771976,
		Latitude:    49.0719037,
	},
	{
		Name:        "La cave a vin",
		City:        "Marseille",
		ID:          "1008",
		Description: "On trouve encore là-bas tout un tas de trésors en tout genre allant d’une vieille carcasse de Berliet à des documents administratifs et autres journaux d’époque. Perdue dans la pampa, cette ancienne cave à vin contient encore beaucoup d’éléments des années 40 à 80, il ne tient qu’a vous de les retrouver. . . et de les préserver.",
		ImageName:   "caveavin",
		Longitude:   5.3699525,
		Latitude:    43.2961743,
	},
	{
		Name:        "Le bateau pirate",
		City:        "Toulouse",
		ID:          "1009",
		Description: "Dans le genre bizarre, cet ancien bateau/restaurant envoie du pâté ! Sur le pont, on se croirait vraiment dans un vieux galion usé par les tempêtes, mais toujours plus ou moins prêt à braver les océans. Si l’envie vous prend d’aller jouer les pirates ! Attention toutefois à ne pas finir dans l’eau, ça reste un lieu abandonné en bois, donc fragile !",
		ImageName:   "bateaupirate",
		Longitude:   1.457240104675293,
		Latitude:    43.601661682128906,
	},
	{
		Name:        "La barrée – La prison Sainte-Anne",
		City:        "Avignon",
		ID:          "1010",
		Description: "La prison Sainte-Anne d’Avignon fût abandonnée en 2003 et ouverte temporairement à une exposition. La ville envisage depuis quelques années de la convertir en hôtel de luxe, mais cela semble au point mort. Vous pouvez retrouver de nombreux clichés qui témoignent du passé chargé de ce lieu abandonné.",
		ImageName:   "prisonsainteanne",
		Longitude:   4.8059012,
		Latitude:    43.9492493,
	},
	{
		Name:        "L’hôpital de la marine de Rochefort",
		City:        "Rochefort",
		ID:          "1011",
		Description: "On ne dirait pas comme ça mais cet hôpital très ancien accueillait autrefois les marins, militaires et charpentiers de la marine abîmés par leur métier. Ce lieu abandonné depuis plus de 30 ans cache derrière cette belle verrière de vieilles cellules pour les internés psychiatriques, et autres pièces bien moins accueillantes encore équipées ! Il s’agit désormais d’un lieu privé qui sera bientôt réhabilité, le gardien vous laissera peut être entrer si vous respectez les lieux !",
		ImageName:   "hopitalrochefort",
		Longitude:   -0.9687519,
		Latitude:    45.9438412,
	},
	{
		Name:        "Le village Martyr",
		City:        "Oradour-sur-Glane",
		ID:          "1012",
		Description: "Le 10 juin 1944, cette ville est devenue fantôme. C’est à cette date que les SS ont entièrement détruit le village et massacré la quasi totalité des habitants. Aujourd’hui, une partie du village a été conservée à l’état de ruine pour témoigner de l’horreur vécue par sa population.",
		ImageName:   "villagemartyr",
		Longitude:   1.0325,
		Latitude:    45.932778,
	},
}

func GetAllUrbexSpots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(urbexSpots)
}
