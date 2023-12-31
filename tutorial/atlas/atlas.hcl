data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./main.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}