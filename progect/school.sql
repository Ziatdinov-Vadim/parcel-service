CREATE TABLE classes (
    class_id INTEGER PRIMARY KEY AUTOINCREMENT,
    grade_year INTEGER NOT NULL CHECK (grade_year BETWEEN 1 AND 11),
    letter TEXT NOT NULL,
    CONSTRAINT unique_class UNIQUE (grade_year, letter)
);

CREATE TABLE students (
    student_id INTEGER PRIMARY KEY AUTOINCREMENT,
    full_name TEXT NOT NULL,
    parent_phone TEXT NOT NULL,
    birth_date TEXT NOT NULL,
    home_address TEXT NOT NULL,
    class_id INTEGER NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes(class_id) ON DELETE RESTRICT
);

INSERT INTO classes (grade_year, letter) VALUES (5, 'А');

INSERT INTO students (full_name, parent_phone, birth_date, home_address, class_id) 
VALUES ('Сидоров Алексей Петрович', '+79991234567', '2015-05-12', 'ул. Ленина, д. 10', 1);
