/*
Petit programme pour extraire les amendements pour les textes législatifs.

Auteur : Ollivier Robert pour le Projet Arcadie
Licence : BSD 2-clauses.

*/

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/atotto/encoding/csv"
)

func realmain(args []string) int {

	if len(args) == 0 {
		log.Fatalf("Besoin du num d'un fichier\n")
	}

	var Textes Everything

	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Printf("lecture %s: %v", args[0], err)
	}

	err = json.Unmarshal(buf, &Textes)
	if err != nil {
		log.Printf("unmarshal: %v", err)
	}

	var (
		anum  int
		allam []Output
	)

	log.Printf("taille %d octets\n", len(buf))

	for _, t := range Textes.TextesEtAmendements.Texteleg {
		anum += len(t.Amendements.Amendement)
		for _, a := range t.Amendements.Amendement {
			am := Output{
				UID:                a.UID,
				Numero:             a.NumeroLong,
				RefTexteLegislatif: a.Identifiant.Saisine.RefTexteLegislatif,
				NumeroPartiePLF:    a.Identifiant.Saisine.NumeroPartiePLF,
				Etat:               a.Etat,
				Auteur:             a.Signataires.Auteur.ActeurRef,
				GroupePolitiqueRef: a.Signataires.Auteur.GroupePolitiqueRef,
				Cosignataires:      strings.Join(a.Signataires.Cosignataires.ActeurRef, ","),
				ExposeSommaire:     a.Corps.ExposeSommaire,
				DateSaisie:         a.Sort.DateSaisie,
				SortEnSeance:       a.Sort.SortEnSeance,
			}
			allam = append(allam, am)
		}
	}
	log.Printf("%d textes lus, %d amendements", len(Textes.TextesEtAmendements.Texteleg), anum)

	out, err := os.Create("out.csv")
	if err != nil {
		log.Printf("erreur création")
		return 1
	}
	wh := csv.NewWriter(out)

	if err := wh.WriteStructHeader(allam[0]); err != nil {
		log.Printf("erreur entête")
		return 2
	}

	if err := wh.WriteStructAll(allam); err != nil {
		log.Printf("erreur données")
		return 3
	}

	return 0
}

func main() {
	flag.Parse()

	_ = realmain(flag.Args())
}
