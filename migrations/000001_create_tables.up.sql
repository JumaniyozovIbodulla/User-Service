CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TYPE "course_types" AS ENUM (
  'beginner',
  'elementary',
  'pre_intermediate',
  'intermediate',
  'upper_intermediate',
  'ielts_level_one',
  'ielts_level_two'
);

CREATE TABLE IF NOT EXISTS "super_admins" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "fullname" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "password" VARCHAR NOT NULL,
  "salary" NUMERIC(10,2) NOT NULL,
  "months_worked" INT DEFAULT 1,
  "ielts_score" NUMERIC(2,1) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "branches" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "location" GEOMETRY(POINT, 4326) NOT NULL,
  "super_admin_id" UUID REFERENCES "super_admins" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "groups" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "name" VARCHAR(255) DEFAULT '',
  "level" course_types,
  "months" INT DEFAULT 3,
  "number_of_students" INT DEFAULT 0,
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "managers" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "fullname" VARCHAR NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "password" VARCHAR NOT NULL,
  "salary" NUMERIC(10,2) NOT NULL,
  "super_admin_id" UUID UNIQUE REFERENCES "super_admins" ("id"),
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "administrators" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "fullname" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "password" VARCHAR NOT NULL,
  "salary" NUMERIC(10,2) NOT NULL,
  "months_worked" INT DEFAULT 0,
  "ielts_score" NUMERIC(2,1) NOT NULL,
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "support_teachers" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "fullname" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "password" VARCHAR NOT NULL,
  "ielts_score" NUMERIC(2,1) NOT NULL,
  "ielts_attepms_count" INT DEFAULT 1,
  "salary" NUMERIC(10,2) NOT NULL,
  "months_worked" BIGINT DEFAULT 0,
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "group_id" UUID UNIQUE REFERENCES "groups" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "teachers" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "fullname" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "password" VARCHAR NOT NULL,
  "salary" NUMERIC(10,2) NOT NULL,
  "months_worked" INT DEFAULT 0,
  "ielts_score" NUMERIC(2,1) NOT NULL,
  "ielts_attepms_count" INT DEFAULT 1,
  "support_teacher_id" UUID REFERENCES "support_teachers" ("id"),
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "group_id" UUID UNIQUE REFERENCES "groups" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "students" (
  "id" UUID PRIMARY KEY,
  "extra_id" VARCHAR UNIQUE NOT NULL,
  "fullname" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "password" VARCHAR NOT NULL,
  "paid_sum" NUMERIC(10,2) DEFAULT 0,
  "course_count" INT DEFAULT 1,
  "total_sum" NUMERIC DEFAULT 0.0,
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "group_id" UUID UNIQUE REFERENCES "groups" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "events" (
  "id" UUID PRIMARY KEY,
  "topic" VARCHAR(255) NOT NULL,
  "descrioption" TEXT DEFAULT '',
  "day" VARCHAR(10) NOT NULL CHECK ("day" != 'sunday'),
  "start_time" TIMESTAMPTZ NOT NULL,
  "duration_hours" INT DEFAULT 1,
  "branch_id" UUID UNIQUE REFERENCES "branches" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "join_events" (
  "id" UUID PRIMARY KEY,
  "event_id" UUID REFERENCES "events" ("id"),
  "student_id" UUID REFERENCES "students" ("id"),
  "is_active" BIGINT DEFAULT 0,
  "joined_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);