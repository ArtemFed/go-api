CREATE TABLE students
(
    id        serial      NOT NULL PRIMARY KEY,
    name      varchar(60) NOT NULL,
    age       serial      NOT NULL,
    specialty varchar(60) NOT NULL
);

CREATE TABLE grades
(
    id           serial       NOT NULL PRIMARY KEY,
    grade        serial       NOT NULL,
    subject_name varchar(60)  NOT NULL,
    student_id   serial       NOT NULL,
    publish_date timestamp(0) NOT NULL DEFAULT NOW(),
    FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE CASCADE
);

CREATE TABLE students_grades
(
    student_id serial NOT NULL references students (id) ON DELETE CASCADE,
    grade_id   serial NOT NULL references grades (id) ON DELETE CASCADE,
    PRIMARY KEY (student_id, grade_id)
);
