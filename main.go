package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	AWS struct {
		Region          string `yaml:"region"`
		S3Bucket        string `yaml:"s3_bucket"`
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
	} `yaml:"aws"`
	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
	Security struct {
		EnableTLS    bool   `yaml:"enable_tls"`
		TLSCertFile  string `yaml:"tls_cert_file"`
		TLSKeyFile   string `yaml:"tls_key_file"`
	} `yaml:"security"`
	Monitoring struct {
		CloudwatchNamespace  string `yaml:"cloudwatch_namespace"`
		EnableCustomMetrics  bool   `yaml:"enable_custom_metrics"`
	} `yaml:"monitoring"`
}

var config Config

func loadConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Load configuration based on environment variable or default to production.yaml
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/production.yaml"
	}

	err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Starting server on port %d\n", config.Server.Port)

	// Initialize blockchain
	blockchain := NewBlockchain()

	// Initialize marketplace service
	marketplace := NewMarketplace(blockchain)

	// Initialize smart contract with marketplace
	smartContract := NewSmartContract(marketplace)
	smartContract.InitializeServices()

	// Assign smart contract to marketplace for reference if needed
	marketplace.SmartContract = smartContract

	// Initialize governance module
	governance := NewGovernance(blockchain)

	// Setup HTTP server and routes
	router := SetupRouter(marketplace, governance)

	// Register docs route
	handlers.RegisterDocsRoutes(router)

	// Initialize OPA evaluator for security policies
	opaEvaluator, err := middleware.NewOPAEvaluator("pkg/security/policies.rego")
	if err != nil {
		log.Fatalf("Failed to initialize OPA evaluator: %v", err)
	}

	// Wrap router with OPA authorization middleware
	securedRouter := middleware.OPAAuthMiddleware(opaEvaluator)(router)

	addr := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Starting Blockchain Logistics Marketplace server on %s", addr)
	if err := http.ListenAndServe(addr, securedRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
