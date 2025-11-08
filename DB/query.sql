-- Application Functions
-- 1. 
SELECT * FROM students;

--2. 
INSERT INTO students (first_name, last_name, email, enrollment_date) VALUES
($1, $2, $3, $4);

-- 3. 
UPDATE students SET email = $1 WHERE student_id = $2;

-- 4. 
DELETE FROM students WHERE student_id = $1;