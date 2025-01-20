package models

type Entity struct {
	BomRef string `json:"bom-ref,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
}

type Hash struct {
	Alg     string `json:"alg"`
	Content string `json:"content"`
}

type Text struct {
	ContentType string `json:"contentType,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Content     string `json:"content"`
}

type Swid struct {
	TagId      string `json:"tagId"`
	Name       string `json:"name"`
	Version    string `json:"version,omitempty"`
	TagVersion string `json:"tagVersion,omitempty"`
	Patch      string `json:"patch,omitempty"`
	Text       Text   `json:"text,,omitempty"`
	Url        string `json:"url,omitempty"`
}

type Property struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Component struct {
	Type        string `json:"type"`
	MimeType    string `json:"mime-type,omitempty"`
	BomRef      string `json:"bom-ref,omitempty"`
	Supplier    Entity `json:"supplier,omitempty"`
	Author      string `json:"author,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
	Group       string `json:"group,omitempty"`
	Name        string `json:"name"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	Scope       string `json:"scope,omitempty"`
	Hashes      []Hash `json:"hashes,omitempty"`
	// Licenses []License `json:"licenses,omitempty"` TODO: Implement dynamic parse logic
	Copyright string `json:"copyright,omitempty"`
	Cpe       string `json:"cpe,omitempty"`
	Purl      string `json:"purl,omitempty"`
	Swid      Swid   `json:"swid,omitempty"`
	Modified  string `json:"modified,omitempty"` // Deprecated use Pedigree instead
	// Pedigree  Pedigree `json:"pedigree,omitempty"` TODO
	// ExternalReferences []ExternalReferences `json:"externalReferences,omitempty"` TODO
	Components []Component `json:"components,omitempty"`
	// Evidence Evidence `json:"evidence,omitempty"` TODO
	// ReleaseNotes ReleaseNotes `json:"releaseNotes,omitempty"` TODO
	// ModelCard ModelCard `json:"modelCard,omitempty"` TODO
	// Data Data `json:"data,omitempty"` TODO
	Properties []Property `json:"properties,omitempty"`
	// Signature Signature `json:"signature,omitempty"` TODO
}

type Dependency struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn,omitempty"`
}

// type Metadata struct {
// 	Timestamp string `json:"timestamp,omitempty"`
// 	// Lifecycles []Lifecycle `json:"lifecycles,omitempty"` TODO: Implement dynamic parse logic
// 	// Tools   []Tool   `json:"tool,omitempty"` TODO: Write Tools models
// 	Authors   []Author  `json:"authors,omitempty"`
// 	Component Component `json:"component"`
// }

type CycloneDX struct {
	BomFormat   string `json:"bomFormat"`
	SpecVersion string `json:"specVersion"`
	// SerialNumber string   `json:"serialNumber,omitempty"`
	// Version      string   `json:"version,omitempty"`
	// Metadata     Metadata `json:"metadata,omitempty"`
	Components []Component `json:"components,omitempty"`
	// Services   []Service   `json:"services,omitempty"` TODO
	// ExternalReferences []ExternalReferences `json:"externalReferences,omitempty"` TODO
	Dependencies []Dependency `json:"dependencies,omitempty"`
	// Compositions []Composition `json:"compositions,omitempty"` TODO
	// Vulnerabilities []Vulnerability `json:"vulnerabilities,omitempty"` TODO
	// Annotations []Annotation `json:"annotations,omitempty"` TODO
	// Formulation []Formulation `json:"formulation,omitempty"` TODO
	Properties []Property `json:"properties,omitempty"`
	// Signature Signature `json:"signature,omitempty"` TODO
}
