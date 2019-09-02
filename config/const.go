package config

// DO NOT CHANGE THIS FILE
// This file contains mapping of constant to environment variable name

// ON OFF toggle
const (
	// OFF off
	OFF = "off"
	// ON on
	ON = "on"
)

// Environment
const (
	// ENV_DEVELOPMENT development
	ENV_DEVELOPMENT = "development"
	// ENV_STAGING staging
	ENV_STAGING = "staging"
	// ENV_PRODUCTION production
	ENV_PRODUCTION = "production"
	// ENV_TESTING testing
	ENV_TESTING = "testing"
)

const (
	// ENV ENV
	ENV = "ENV"
	// DB_HOST DB_HOST
	DB_HOST = "DB_HOST"
	// DB_PORT DB_PORT
	DB_PORT = "DB_PORT"
	// DB_NAME DB_NAME
	DB_NAME = "DB_NAME"
	// DB_USER DB_USER
	DB_USER = "DB_USER"
	// DB_PASS DB_PASS
	DB_PASS = "DB_PASS"
	// DB_MIGRATE_USER DB_MIGRATE_USER
	DB_MIGRATE_USER = "DB_MIGRATE_USER"
	// DB_MIGRATE_PASS DB_MIGRATE_PASS
	DB_MIGRATE_PASS = "DB_MIGRATE_PASS"
	// REDIS_HOST REDIS_HOST
	REDIS_HOST = "REDIS_HOST"
	// REDIS_PASS REDIS_PASS
	REDIS_PASS = "REDIS_PASS"
	// PORT PORT
	PORT = "PORT"
	// BASE_URL BASE_URL
	BASE_URL = "BASE_URL"
	// THROTTLE_PREFIX THROTTLE_PREFIX
	THROTTLE_PREFIX = "THROTTLE_PREFIX"
	// THROTTLE_TRIGGER THROTTLE_TRIGGER
	THROTTLE_TRIGGER = "THROTTLE_TRIGGER"
	// FORCE_MIGRATE FORCE_MIGRATE
	FORCE_MIGRATE = "FORCE_MIGRATE"
	// DEBUG_LOG DEBUG_LOG
	DEBUG_LOG = "DEBUG_LOG"
	// MAX_OPEN_CONNS MAX_OPEN_CONNS
	MAX_OPEN_CONNS = "MAX_OPEN_CONNS"
	// MAX_OPEN_CONNS MAX_OPEN_CONNS
	MAX_IDLE_CONNS = "MAX_IDLE_CONNS"
)

const (
	// TEST_DB_HOST TEST_DB_HOST
	TEST_DB_HOST = "TEST_DB_HOST"
	// TEST_DB_PORT TEST_DB_PORT
	TEST_DB_PORT = "TEST_DB_PORT"
	// TEST_DB_NAME TEST_DB_NAME
	TEST_DB_NAME = "TEST_DB_NAME"
	// TEST_DB_USER TEST_DB_USER
	TEST_DB_USER = "TEST_DB_USER"
	// TEST_DB_PASS TEST_DB_PASS
	TEST_DB_PASS = "TEST_DB_PASS"
	// TEST_DB_MIGRATE_USER TEST_DB_MIGRATE_USER
	TEST_DB_MIGRATE_USER = "TEST_DB_MIGRATE_USER"
	// TEST_DB_MIGRATE_PASS TEST_DB_MIGRATE_PASS
	TEST_DB_MIGRATE_PASS = "TEST_DB_MIGRATE_PASS"
	// TEST_REDIS_HOST TEST_REDIS_HOST
	TEST_REDIS_HOST = "TEST_REDIS_HOST"
	// TEST_REDIS_PASS TEST_REDIS_PASS
	TEST_REDIS_PASS = "TEST_REDIS_PASS"
)

//noinspection ALL
const (
	// I_QUERY_CONFIG QUERY_CONFIG
	I_QUERY_CONFIG = "QUERY_CONFIG"
	// I_READ_TIMEOUT READ_TIMEOUT
	I_READ_TIMEOUT = "READ_TIMEOUT"
	// I_WRITE_TIMEOUT WRITE_TIMEOUT
	I_WRITE_TIMEOUT = "WRITE_TIMEOUT"
	// I_IDLE_TIMEOUT IDLE_TIMEOUT
	I_IDLE_TIMEOUT = "IDLE_TIMEOUT"
	// I_WAIT_SHUTDOWN WAIT_SHUTDOWN
	I_WAIT_SHUTDOWN = "WAIT_SHUTDOWN"
	// I_CREDENTIAL_EXPIRE CREDENTIAL_EXPIRE
	I_CREDENTIAL_EXPIRE = "CREDENTIAL_EXPIRE"
	// I_REDIS_SESSION_EXPIRE REDIS_SESSION_EXPIRE
	I_REDIS_SESSION_EXPIRE = "REDIS_SESSION_EXPIRE"
	// I_THROTTLE_LIMIT THROTTLE_LIMIT
	I_THROTTLE_LIMIT = "THROTTLE_LIMIT"
	// I_THROTTLE_DURATION THROTTLE_DURATION
	I_THROTTLE_DURATION = "THROTTLE_DURATION"
	// I_THROTTLE_TYPE THROTTLE_TYPE
	I_THROTTLE_TYPE = "THROTTLE_TYPE"
)