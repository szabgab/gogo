package cmd

type Language struct {
	Name string `yaml:"Name"`
	Code string `yaml:"IETF BCP 47"`
}

type License struct {
	Name      string `yaml:"Name"`
	ShortName string `yaml:"Short name"`
	Link      string `yaml:"Link"`
}

type CourseData struct {
	Language    Language `yaml:"Language"`        // TargetLanguage
	ForSpeakers Language `yaml:"For speakers of"` // SourceLanguage
	License     License  `yaml:"License"`
	Repository  string   `yaml:"Repository"`
	Characters  []string `yaml:"Special characters"`
}

type CourseFile struct {
	Course  CourseData `yaml:"Course"`
	Modules []string   `yaml:"Modules"`
}

type ModuleFile struct {
	Module Module   `yaml:"Module"`
	Skills []string `yaml:"Skills"`
}

type Module struct {
	Name string `yaml:"Name"`
}

type SkillMeta struct {
	Name       string   `yaml:"Name"`
	Id         int      `yaml:"Id"`
	Thumbnails []string `yaml:"Thumbnails"`
}

type Word struct {
	Word        string   `yaml:"Word"`
	Translation string   `yaml:"Translation"`
	Images      []string `yaml:"Images"`
}

type Phrase struct {
	Phrase      string `yaml:"Phrase"`
	Translation string `yaml:"Translation"`
}

type MiniDictionary map[string]interface{}

type Skill struct {
	Meta       SkillMeta                   `yaml:"Skill"`
	Words      []Word                      `yaml:"New words"`
	Phrases    []Phrase                    `yaml:"Phrases"`
	Dictionary map[string][]MiniDictionary `yaml:"Mini-dictionary"`
}
