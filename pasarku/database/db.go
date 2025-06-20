package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"pasarku/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gopkg.in/yaml.v2"
)

var DB *gorm.DB

type DBConfig struct {
	Driver   string `yaml:"driver"`   // harus "postgres"
	Host     string `yaml:"host"`     // biasanya "localhost"
	Port     string `yaml:"port"`     // default PostgreSQL: "5432"
	User     string `yaml:"user"`     // user PostgreSQL
	Password string `yaml:"password"` // password PostgreSQL
	DBName   string `yaml:"dbname"`   // nama database
	SSLMode  string `yaml:"sslmode"`  // "disable" jika lokal
}

// Load konfigurasi dari dbconfig.yml
func LoadDBConfig() (*DBConfig, error) {
	data, err := ioutil.ReadFile("dbconfig.yml")
	if err != nil {
		return nil, err
	}

	var config DBConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// InitDB menginisialisasi koneksi database PostgreSQL
func InitDB() {
	config, err := LoadDBConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load DB config: %v", err)
	}

	if config.Driver != "postgres" {
		log.Fatalf("❌ Unsupported DB driver: %s. Use 'postgres'", config.Driver)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB = db
	fmt.Println("✅ Connected to PostgreSQL database!")
}
