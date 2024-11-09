data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "postgres", 
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres?search_path=public"
  dev-url = "docker://postgres?search_path=public"
  url = "postgres://mailcast:password@:5432/mailcast?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}