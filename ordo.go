package main

import "fmt"

type reading struct {
  /*
  Struct for an individual reading in the ordo. 
  */
  verse string // the verse/source of the reading, e.g Psalm 1:1
  latin string // the Latin of the reading
  vernacular map[string]string // the vernacular translation(s) of the reading, in language:translation map, e.g "english":"hello world"
}

type propers struct {
  /*
  Struct containing all of the propers for a Mass.
  */
  id string // id allows the proper to be easily referenced, for example for commemorations
  introit *reading
  collect *reading
  epistle *reading
  gradual *reading
  sequence *reading //optional
  gospel *reading
  offertory *reading
  secret *reading
  preface *reading
  communion *reading
  postcommunion *reading
}

func (p *propers) printSetFields(languages ...string){
  /*
  Print all of the fields that are set in the propers struct.
  */
  fields := []struct {
    name string
    value *reading
  }{
    {"Introit", p.introit},
		{"Collect", p.collect},
		{"Epistle", p.epistle},
		{"Gradual", p.gradual},
		{"Sequence", p.sequence},
		{"Gospel", p.gospel},
		{"Offertory", p.offertory},
		{"Secret", p.secret},
		{"Preface", p.preface},
		{"Communion", p.communion},
		{"Postcommunion", p.postcommunion},
  }
  for _, reading := range fields {
    if reading.value != nil {
      fmt.Println(reading.name)
      printReading(*reading.value, languages...)
    }
  }
}

func printReading(r reading, languages ...string){
  /*
  Print a single reading in the requested languages
  */
  if (r.verse != ""){
    fmt.Println(r.verse)
  }
  for _, l := range languages{
    if l == "Latin" || l == "latin"{
      fmt.Printf("Latin: \n %v \n", r.latin)
    } else if t, key := r.vernacular[l]; key {
      fmt.Printf("%v: \n %v \n", l, t)
    } else {
      fmt.Println(l, "translation not found")
    }
  }
}

func (r *reading) getReading() (string, map[string]string) {
  /*
  Get a single reading in all available languages
  Receiver: p *reading - a reading struct
  Inputs: language - the language to return the reading in
  Returns:  verse - the verse of the reading
            reading - a map of the reading in all available languages 
  */
  // Initialize the map to be returned
  reading := make(map[string]string)
  verse := r.verse
  latin := r.latin
  reading["latin"] = latin
  for l, t := range r.vernacular {
    reading[l] = t
  }
  return verse, reading
}

func (p *propers) getPropers()map[string]map[string]string{
  /*
Get all the texts for a given Mass' propers in the requested language, returned as a map of string slices, e.g:
{
  "introit": 
  {
    "verse": "Ps 1:1",
    "latin": "latin text",
    "english": "english text"
  },
  "collect": 
  {
    "verse": "", 
    "latin": "latin text",
    "english": "english text"
  },
  etc
}
  */
  // initialising the map
  propersMap := make(map[string]map[string]string)
  // get each text from the proper
  readings := make(map[string]*reading)
  readings["introit"] = p.introit
  readings["collect"] = p.collect
  readings["epistle"] = p.epistle
  readings["gradual"] = p.gradual
  readings["sequence"] = p.sequence
  readings["gospel"] = p.gospel
  readings["offertory"] = p.offertory
  readings["secret"] = p.secret
  readings["preface"] = p.preface
  readings["communion"] = p.communion
  readings["postcommunion"] = p.postcommunion
  // iterate through the readings and add them to the propers map
  for name, reading := range readings {
    // skip if the reading is nil
    if reading == nil {
      continue
    }
    verse := reading.verse
    latin := reading.latin
    toInsert := map[string]string{
      "verse": verse,
      "latin": latin,
    }
    for l, t := range reading.vernacular {
      // add each translation to the map
      toInsert[l] = t
    }
    propersMap[name] = toInsert
  }
  return propersMap
}

func main() {
	// For Testing only
	propersInstance := propers{
    id: "test",
		introit: &reading{
			verse: "Psalm 23:1",
			latin: "Dominus regit me, nihil mihi deerit.",
			vernacular: map[string]string{
				"english": "The Lord is my shepherd; I shall not want.",
        "spanish": "El Señor es mi pastor; nada me faltará.",
			},
		},
		gospel: &reading{
			verse: "John 3:16",
			latin: "Sic enim dilexit Deus mundum ut Filium suum unigenitum daret.",
			vernacular: map[string]string{
				"english": "For God so loved the world, that he gave his only Son.",
        "japanese": "神はこうして世を愛され、ご自身のひとり子を賜わった。",
			},
		},
		// Set other fields as needed
	}
  p := propersInstance.getPropers()
  fmt.Println(p)
}
