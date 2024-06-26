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

CREATE TYPE "attend_type" AS ENUM (
  'attended',
  'absent',
  'late'
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

-- IS DONE

-- CREATE TABLE IF NOT EXISTS "schedules" (
--   "id" UUID PRIMARY KEY,
--   "group_id" UUID,
--   "group_type" course_types,
--   "number_of_students" INT DEFAULT 0,
--   "start_time" TIMESTAMPTZ NOT NULL,
--   "end_time" TIMESTAMPTZ NOT NULL,
--   "branch_id" UUID,
--   "teacher_id" UUID,
--   "support_teacher_id" UUID,
--   "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
-- );

-- CREATE TABLE IF NOT EXISTS "lessons" (
--   "id" UUID PRIMARY KEY,
--   "schedule_id" UUID,
--   "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
-- );

-- CREATE TABLE IF NOT EXISTS "tasks" (
--   "id" UUID PRIMARY KEY,
--   "lesson_id" UUID,
--   "label" VARCHAR(255) NOT NULL,
--   "deadline" TIMESTAMPTZ NOT NULL,
--   "score" NUMERIC(3,1) DEFAULT 0,
--   "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
-- );

-- CREATE TABLE IF NOT EXISTS "attendances" (
--   "id" UUID PRIMARY KEY,
--   "lesson_id" UUID,
--   "event_id" UUID UNIQUE,
--   "student_id" UUID UNIQUE,
--   "status" attend_type,
--   "late_minute" INT DEFAULT 0
-- );
-- -- ALTER TABLE "branches" ADD FOREIGN KEY ("id") ;

-- -- ALTER TABLE "teachers" ADD FOREIGN KEY ("support_teacher_id") ;

-- -- ALTER TABLE "branches" ADD FOREIGN KEY ("id") ;

-- ALTER TABLE "branches" ADD FOREIGN KEY ("id") REFERENCES "managers" ("branch_id");

-- ALTER TABLE "branches" ADD FOREIGN KEY ("id") REFERENCES "administrators" ("branch_id");

-- ALTER TABLE "groups" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");

-- ALTER TABLE "groups" ADD FOREIGN KEY ("suppurt_teacher_id") REFERENCES "support_teachers" ("id");

-- ALTER TABLE "students" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

-- ALTER TABLE "students" ADD FOREIGN KEY ("branch_id") REFERENCES "groups" ("id");

-- ALTER TABLE "schedules" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

-- ALTER TABLE "schedules" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

-- ALTER TABLE "schedules" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");

-- ALTER TABLE "schedules" ADD FOREIGN KEY ("support_teacher_id") REFERENCES "support_teachers" ("id");

-- ALTER TABLE "lessons" ADD FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("id");

-- ALTER TABLE "tasks" ADD FOREIGN KEY ("lesson_id") REFERENCES "lessons" ("id");

-- ALTER TABLE "events" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

-- ALTER TABLE "students" ADD FOREIGN KEY ("id") REFERENCES "events" ("student_id");

-- ALTER TABLE "attendances" ADD FOREIGN KEY ("lesson_id") REFERENCES "lessons" ("id");

-- ALTER TABLE "attendances" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

-- ALTER TABLE "attendances" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");