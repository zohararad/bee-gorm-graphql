package db

var models []interface{}
var migrations []string

func RegisterModel(model interface{}) {
  models = append(models, model)
}

func RegisterMigration(migration string) {
  migrations = append(migrations, migration)
}

func AutoMigrate() {
  for _, migration := range migrations {
    Conn.Exec(migration)
  }
  Conn.AutoMigrate(models...)
}
