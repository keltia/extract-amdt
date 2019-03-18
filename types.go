package main

//"encoding/xml"

type Everything struct {
	TextesEtAmendements struct {
		Texteleg []struct {
			Schema             string `json:"@xmlns:xsi"`
			RefTexteLegislatif string `json:"refTexteLegislatif"`
			Amendements        struct {
				Amendement []struct {
					UID         string `json:"uid"`
					Identifiant struct {
						Legislature string `json:"legislature"`
						Numero      string `json:"numero"`
						NumRect     string `json:"numRect"`
						Saisine     struct {
							RefTexteLegislatif         string `json:"refTexteLegislatif"`
							NumeroPartiePLF            string `json:"numeroPartiePLF"`
							OrganeExamen               string `json:"organeExamen"`
							MentionSecondeDeliberation string `json:"mentionSecondeDeliberation"`
						} `json:"saisine"`
					} `json:"identifiant"`
					NumeroLong               string      `json:"numeroLong"`
					EtapeTexte               string      `json:"etapeTexte"`
					TriAmendement            string      `json:"triAmendement"`
					CardinaliteAmdtMultiples string      `json:"cardinaliteAmdtMultiples"`
					AmendementParent         interface{} `json:"amendementParent"`
					Etat                     string      `json:"etat"`
					Signataires              struct {
						Auteur struct {
							TypeAuteur         string      `json:"typeAuteur"`
							ActeurRef          string      `json:"acteurRef"`
							OrganeRef          interface{} `json:"organeRef"`
							GroupePolitiqueRef string      `json:"groupePolitiqueRef"`
						} `json:"auteur"`
						Cosignataires struct {
							ActeurRef []string `json:"acteurRef"`
						} `json:"cosignataires"`
						TexteAffichable string `json:"texteAffichable"`
					} `json:"signataires"`
					PointeurFragmentTexte struct {
						MissionVisee interface{} `json:"missionVisee"`
						Division     struct {
							Titre                    string      `json:"titre"`
							ArticleDesignationCourte string      `json:"articleDesignationCourte"`
							Type                     string      `json:"type"`
							AvantAApres              string      `json:"avant_A_Apres"`
							DivisionRattachee        string      `json:"divisionRattachee"`
							ArticleAdditionnel       string      `json:"articleAdditionnel"`
							ChapitreAdditionnel      interface{} `json:"chapitreAdditionnel"`
							URLDivisionTexteVise     string      `json:"urlDivisionTexteVise"`
						} `json:"division"`
						Alinea interface{} `json:"alinea"`
					} `json:"pointeurFragmentTexte"`
					Corps struct {
						Dispositif           string      `json:"dispositif"`
						ExposeSommaire       string      `json:"exposeSommaire"`
						AnnexeExposeSommaire interface{} `json:"annexeExposeSommaire"`
					} `json:"corps"`
					Representations struct {
						Representation struct {
							Nom      string `json:"nom"`
							TypeMime struct {
								Type    string `json:"type"`
								SubType string `json:"subType"`
							} `json:"typeMime"`
							StatutRepresentation struct {
								Verbatim       string `json:"verbatim"`
								Canonique      string `json:"canonique"`
								Officielle     string `json:"officielle"`
								Transcription  string `json:"transcription"`
								Enregistrement string `json:"enregistrement"`
							} `json:"statutRepresentation"`
							RepSource interface{} `json:"repSource"`
							Offset    interface{} `json:"offset"`
							Contenu   struct {
								DocumentURI string `json:"documentURI"`
							} `json:"contenu"`
							DateDispoRepresentation interface{} `json:"dateDispoRepresentation"`
						} `json:"representation"`
					} `json:"representations"`
					SeanceDiscussion interface{} `json:"seanceDiscussion"`
					Sort             struct {
						DateSaisie   string `json:"dateSaisie"`
						SortEnSeance string `json:"sortEnSeance"`
					} `json:"sort"`
					DateDepot        string `json:"dateDepot"`
					DateDistribution string `json:"dateDistribution"`
					Article99        string `json:"article99"`
					LoiReference     struct {
						CodeLoi         string `json:"codeLoi"`
						DivisionCodeLoi string `json:"divisionCodeLoi"`
					} `json:"loiReference"`
				}
			} `json:"amendements"`
		} `json:"texteleg"`
	} `json:"textesEtAmendements"`
}

type Output struct {
	UID                string `json:"uid"`
	Numero             string `json:"numero"`
	RefTexteLegislatif string `json:"refTexteLegislatif"`
	NumeroPartiePLF    string `json:"numeroPartiePLF"`
	Etat               string `json:"etat"`
	Auteur             string `json:"acteurRef"`
	GroupePolitiqueRef string `json:"groupePolitiqueRef"`
	Cosignataires      string
	ExposeSommaire     string `json:"exposeSommaire"`
	DateSaisie         string `json:"dateSaisie"`
	SortEnSeance       string `json:"sortEnSeance"`
}
