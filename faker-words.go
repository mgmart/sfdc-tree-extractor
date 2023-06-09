//
//  faker-words.go
//  sfdcTreeExtractor
//
//  Created by Mario Martelli on 18.03.23.
//  Copyright © 2023 Mario Martelli. All rights reserved.
//
//  This file is part of EverOrg.
//
//  sfdcTreeextractor is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  EverOrg is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with sfdcTreeExtractor. If not, see <http://www.gnu.org/licenses/>.

package sfdcTreeExtractor

var adjectives = []string{"adäquat", "affektiert", "agil", "akribisch", "antagonistisch", "apathisch", "arriviert", "autokratisch", "banal", "brachial", "Contenance", "designiert", "desolat", "dediziert", "definitiv", "dezidiert", "diabolisch", "diametral", "differenziert", "diffizil", "diffus", "diskutabel", "distinguiert", "effektiv", "effizient", "elanvoll", "eloquent", "eminent", "essenziell", "evident", "exorbitant", "explizit", "expressiv", "fulminant", "generös", "gravierend", "heterogen", "homogen", "ikonisch", "illustrativ", "impraktikabel", "inadäquat", "inakzeptabel", "indiskutabel", "infernalisch", "informell", "initial", "irrelevant", "komplex", "kongenial", "konsistent", "konsterniert", "kontinuierlich", "konträr", "kurios", "lapidar", "legitim", "lethargisch", "loyal", "lukrativ", "maliziös", "maniriert", "marginal", "martialisch", "medioker", "melodramatisch", "morbid", "nebulös", "neuralgisch", "normativ", "obligatorisch", "obsolet", "omnipotent", "opportun", "opulent", "pekuniär", "penibel", "perfide", "pittoresk", "pointiert", "prädestiniert", "prägnant", "präsent", "prätentiös", "prekär", "prosaisch", "redundant", "relevant", "renitent", "renommiert", "respektabel", "restriktiv", "rudimentär", "sakrosankt", "satanisch", "saturiert", "servil", "skurril", "stringent", "subsidiär", "subtil", "substanziell", "superb", "theatralisch", "titanisch", "tolerabel", "tradiert", "trist", "trivial", "vakant", "vehement", "versiert"}

var substantives = []string{"Absenz", "Agilität", "Agonie", "Akribie", "Ambition", "Ambivalenz", "Analyse", "Antipathie", "Antithese", "Apathie", "Approximation", "Aspiration", "Assoziation", "Attitüde", "Attraktion", "Autarkie", "Authentizität", "Aversion", "Contenance", "Dedikation", "Dependance", "Desavouierung", "Destination", "Dezenz", "Dignität", "Diffamierung", "Differenz", "Direktive", "Diskrepanz", "Diskrimination", "Disproportionalität", "Distinktion", "Divergenz", "Doktrin", "Dualität", "Elementarität", "Eloquenz", "Euphorie", "Exkulpation", "Familiarität", "Fiktion", "Fluktuation", "Fortune", "Gravität", "Hypothese", "Imagination", "Implikation", "Imponderabilie", "Indifferenz", "Inkonsequenz", "Innovation", "Insistenz", "Integration", "Intervention", "Intimität", "Inspiration", "Invektive", "Präsenz", "Kohärenz", "Kohäsion", "Komposition", "Kontroverse", "Kontemplation", "Konvergenz", "Konversation", "Konzilianz", "Kreativität", "Krux", "Lakonie", "Lethargie", "Malaise", "Manie", "Maxime", "Mission", "Modifikation", "Modifizierung", "Obskurität", "Operation", "Perseveranz", "Phantasmagorie", "Phase", "Prämisse", "Präferenz", "Präpotenz", "Präsumtion", "Prävalenz", "Präzision", "Punktualität", "Ratio", "Reflexion", "Relation", "Relevanz", "Renitenz", "Reputation", "Retrospektive", "Rigorosität", "Schimäre", "Sentenz", "Servilität", "Signifikanz", "Suada", "Supposition", "Temperenz", "Tirade", "Transzendenz", "Usance", "Variante", "Virtualität", "Vita", "Zivilität"}

var companyForm = []string{
	//Personengesellschaften
	"GbR", "KG", "AG & Co. KG", "GmbH & Co. KG", "Limited & Co. KG",
	// "Stiftung & Co. KG", "Stiftung GmbH & Co. KG",
	"UG (haftungsbeschränkt) & Co. KG", "OHG", "GmbH & Co. OHG", "AG & Co. OHG",
	// "Partenreederei", "PartG", "PartG mbB", "Stille Gesellschaft",
	// Kapitalgesellschaften:
	"AG", "gAG", "GmbH", "gGmbH", "InvAG", "KGaA", "AG & Co. KGaA", "SE & Co. KGaA", "GmbH & Co. KGaA", "Stiftung & Co. KGaA",
	// "REIT-AG",
	"UG (haftungsbeschränkt)",
	// Sonstige Rechtsformen:
	// "AöR", "eG", "Eigenbetrieb", "Einzelunternehmen", "e. V.", "KöR", "Regiebetrieb", "Stiftung", "VVaG"
}

var mans = []string{"Alfred", "Arthur", "Artur", "Bruno", "Carl", "Claus", "Curt", "Erich", "Ernst", "Franz", "Friedrich", "Fritz", "Georg", "Gerhard", "Günther", "Günter", "Hans", "Harry", "Heinz", "Hellmut", "Helmuth", "Herbert", "Hermann", "Horst", "Joachim", "Karl", "Karlheinz", "Klaus", "Kurt", "Manfred", "Max", "Otto", "Paul", "Richard", "Rudolf", "Walter", "Werner", "Wilhelm", "Willi/Willy", "Wolfgang"}

var women = []string{"Anna", "Anneliese", "Berta", "Bertha", "Charlotte", "Clara", "Klara", "Edith", "Elfriede", "Elisabeth", "Ella", "Else", "Emma", "Erika", "Erna", "Frieda", "Frida", "Gerda", "Gertrud", "Gisela", "Hedwig", "Helene", "Helga", "Herta", "Hertha", "Hildegard", "Ida", "Ilse", "Ingeborg", "Irmgard", "Johanna", "Käte", "Käthe", "Lieselotte", "Liselotte", "Louise", "Luise", "Margarethe", "Margarete", "Margot", "Maria", "Marie", "Marta", "Martha", "Ruth", "Ursula", "Waltraud", "Waltraut"}

var title = []string{"Dr.", "Professor"}

var lastName = []string{"Müller", "Schmidt", "Schneider", "Fischer", "Weber", "Meyer", "Wagner", "Becker", "Schulz", "Hoffmann", "Schäfer", "Bauer", "Koch", "Richter", "Klein", "Wolf", "Schröder", "Neumann", "Schwarz", "Braun", "Hofmann", "Zimmermann", "Schmitt", "Hartmann", "Krüger", "Schmid", "Werner", "Lange", "Schmitz", "Meier", "Krause", "Maier", "Lehmann", "Huber", "Mayer", "Herrmann", "Köhler", "Walter", "König", "Schulze", "Fuchs", "Kaiser", "Lang", "Weiß", "Peters", "Scholz", "Jung", "Möller", "Hahn", "Keller", "Vogel", "Schubert", "Roth", "Frank", "Friedrich", "Beck", "Günther", "Berger", "Winkler", "Lorenz", "Baumann", "Schuster", "Kraus", "Böhm", "Simon", "Franke", "Albrecht", "Winter", "Ludwig", "Martin", "Krämer", "Schumacher", "Vogt", "Jäger", "Stein", "Otto", "Groß", "Sommer", "Haas", "Graf", "Heinrich", "Seidel", "Schreiber", "Ziegler", "Brandt", "Kuhn", "Schulte", "Dietrich", "Kühn", "Engel", "Pohl", "Horn", "Sauer", "Arnold", "Thomas", "Bergmann", "Busch", "Pfeiffer", "Voigt", "Götz", "Seifert", "Lindner", "Ernst", "Hübner", "Kramer", "Franz", "Beyer", "Wolff", "Peter", "Jansen", "Kern", "Barth", "Wenzel", "Hermann", "Ott", "Paul", "Riedel", "Wilhelm", "Hansen", "Nagel", "Grimm", "Lenz", "Ritter", "Bock", "Langer", "Kaufmann", "Mohr", "Förster", "Zimmer", "Haase", "Lutz", "Kruse", "Jahn", "Schumann", "Fiedler", "Thiel", "Hoppe", "Kraft", "Michel", "Marx", "Fritz", "Arndt", "Eckert", "Schütz", "Walther", "Petersen", "Berg", "Schindler", "Kunz", "Reuter", "Sander", "Schilling", "Reinhardt", "Frey", "Ebert", "Böttcher", "Thiele", "Gruber", "Schramm", "Hein", "Bayer", "Fröhlich", "Voß", "Herzog", "Hesse", "Maurer", "Rudolph", "Nowak", "Geiger", "Beckmann", "Kunze", "Seitz", "Stephan", "Büttner", "Bender", "Gärtner", "Bachmann", "Behrens", "Scherer", "Adam", "Stahl", "Steiner", "Kurz", "Dietz", "Brunner", "Witt", "Moser", "Fink", "Ullrich", "Kirchner", "Löffler", "Heinz", "Schultz", "Ulrich", "Reichert", "Schwab", "Breuer", "Gerlach", "Brinkmann", "Göbel"}

var emailExtension = []string{"abwesend.de", "addcom.de", "alpenjodel.de", "alphafrau.de", "ama-trade.de", "anonmails.de", "antispam.de", "antispam24.d", "antispammail.d", "aol.de", "arcor.de", "berlin.de", "bigfoot.de", "bin-wieder-da.de", "bleib-bei-mir.de", "centermail.de", "cheatmail.de", "epost.de", "gmx.de", "hab-verschlafen.de", "home.de", "hotmail.de", "kommespaeter.de", "kurzepost.de", "online.de", "outlook.de", "quantentunnel.de", "secretemail.de", "sinnlos-mail.de", "sueddeutsche.de", "t-online.de", "web.de", "yahoo.de", "alpenjodel.de", "alphafrau.de", "aol.de"}
