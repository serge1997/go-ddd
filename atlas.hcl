data "external_schema" "gorm" {
    program = [
        "go",
        "run",
        "-mod=mod",
        "ariga.io/atlas-provider-gorm",
        "load",
        "--path", "./internal/infrastructure/model",
        "--dialect", "mysql",
    ]
}

env "gorm" {
    src = data.external_schema.gorm.url
    url = "mysql://root:@localhost:3306/goddd"
    dev = "docker://mysql/8/dev"
    migration {
    dir = "file://./internal/database/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}