module rewarding-harvest-backend

go 1.22

toolchain go1.22.10

require (
	// Core Frameworks
	github.com/gin-gonic/gin v1.10.0
	github.com/golang-jwt/jwt/v5 v5.2.1

	// Database and SSL Libraries
	github.com/jackc/pgx/v5 v5.7.1
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
)

require (
	// ByteDance Libraries for Advanced JSON Parsing
	github.com/bytedance/sonic v1.12.5 // indirect
	github.com/bytedance/sonic/loader v0.2.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect

	// MIME Type and File Parsing
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect

	// Validation and Localization
	github.com/go-playground/validator/v10 v10.23.0 // indirect

	// JSON and Parsing Libraries
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/kr/text v0.2.0 // indirect

	// Utility Libraries
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect

	// Concurrent and Reflection Libraries
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect

	// Encoding and Compression Libraries
	github.com/ugorji/go/codec v1.2.12 // indirect

	// System Architecture and Assembly
	golang.org/x/arch v0.12.0 // indirect

	// Cryptography, Networking, and Sync Tools
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.35.2 // indirect

	// YAML Support
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/joho/godotenv v1.5.1 // indirect
