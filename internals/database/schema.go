package database

// REMPLAZAR ENUM POR CHECK CONSTRAINT
var tableSchema = `
CREATE TABLE  IF NOT EXISTS "patients" (
	"id" INTEGER PRIMARY KEY,
	"name" TEXT NOT NULL,
	"last_name" TEXT,
	"email" TEXT UNIQUE,
	"telephone_number" TEXT UNIQUE,
	CONSTRAINT "check_null_email_or_telephone_number" CHECK (email IS NOT NULL OR telephone_number IS NOT NULL)
);

PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS "appointments" (
	"id" INTEGER PRIMARY KEY,
	"patient_id" INTEGER NOT NULL,
	"status" TEXT 
		NOT NULL 
		DEFAULT 'CREATED' 
		CHECK (status in ('CREATED','CONFIRMED','CANCELED','NO_SHOW')),
	"created_on" TEXT NOT NULL,
	"scheduled_for" TEXT NOT NULL,
	"duration" TEXT NOT NULL DEFAULT '1h',
	FOREIGN KEY (patient_id) REFERENCES patients(id)
);

CREATE TABLE IF NOT EXISTS "visits" ( 
	"id" INTEGER PRIMARY KEY,
	"appointment_id" INTEGER UNIQUE NOT NULL,
	"created_on" TEXT NOT NULL,
	"notes" TEXT,
	FOREIGN KEY (appointment_id) REFERENCES appointments(id)
);

CREATE INDEX IF NOT EXISTS "idx_appointment_scheduled" ON "appointments" ("scheduled_for");

CREATE INDEX IF NOT EXISTS "idx_comp_appointment_patient" ON "appointments" ("patient_id", "scheduled_for");
`
