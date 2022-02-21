package tdtl

type TentacleConfig struct {
	SourceEntity string
	PropertyKeys []string
}

type TQLConfig struct { // nolint
	TargetEntity   string
	SourceEntities []string
	Tentacles      []TentacleConfig
}
