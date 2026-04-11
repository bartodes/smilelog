package database

// appointmentsStatus Enum as Check
var tableSchema = []string{
	`CREATE TABLE  IF NOT EXISTS "patients" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL,
		"last_name" TEXT,
		"email" TEXT UNIQUE,
		"phone_number" TEXT UNIQUE,
		CONSTRAINT "check_null_email_or_telephone_number" CHECK (email IS NOT NULL OR phone_number IS NOT NULL)
	);`,
	"PRAGMA foreign_keys = ON;",
	`CREATE TABLE IF NOT EXISTS "appointments" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"patient_id" INTEGER NOT NULL,
		"status" TEXT 
			NOT NULL 
			DEFAULT 'CREATED' 
			CHECK (status in ('CREATED','COMPLETED','CANCELED','NO_SHOW')),
		"scheduled_for" TEXT NOT NULL,
		"duration_minutes" INTEGER NOT NULL DEFAULT '60',
		FOREIGN KEY (patient_id) REFERENCES patients(id)
	);`,
	`CREATE TABLE IF NOT EXISTS "visits" ( 
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"appointment_id" INTEGER UNIQUE NOT NULL,
		"notes" TEXT,
		FOREIGN KEY (appointment_id) REFERENCES appointments(id)
	);`,
	"CREATE INDEX IF NOT EXISTS 'idx_appointment_scheduled' ON 'appointments' ('scheduled_for');",
	"CREATE INDEX IF NOT EXISTS 'idx_comp_appointment_patient' ON 'appointments' ('patient_id', 'scheduled_for');",
}
