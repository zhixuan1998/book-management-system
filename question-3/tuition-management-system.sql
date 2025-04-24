-- Question a.i)
SELECT
    d.id AS Department_ID,
    d.code AS Department_code,
    d.description AS Department_description,
    COUNT(e.id) AS Employee_count
FROM
    Department d
LEFT JOIN
    Employee e ON d.id = e.departmentId
GROUP BY
    d.id, d.code, d.description

-- Question a.ii)
SELECT
    e.id AS Employee_ID,
    e.code AS Employee_code,
    e.name AS Employee_name,
    e.salary AS Employee_salary,
    e.departmentId AS Department_ID,
    d.code AS Department_code,
    d.description AS Department_description
FROM
    Employee e
LEFT JOIN
    Department d ON e.departmentId = d.id
WHERE
    e.salary >= 3000 AND e.salary <= 4000
ORDER BY
    d.code ASC, e.name ASC

-- Question b)
CREATE TABLE Subject (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
)

CREATE TABLE Semester (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    is_active BIT DEFAULT 0
)

CREATE TABLE Lecturer (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    phone_number varchar(255) DEFAULT NULL,
    joined_date DATE NOT NULL,
)

CREATE TABLE Class (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    subject_id INT FOREIGN KEY REFERENCES Subject(id),
    semester_id INT FOREIGN KEY REFERENCES Semester(id),
    lecturer_id INT FOREIGN KEY REFERENCES Lecturer(id),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
)

CREATE TABLE Student (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    code VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    enrolled_date DATE not null,
    graduated_date DATE DEFAULT NULL,
)

CREATE TABLE Student_class (
    id INT NOT NULL IDENTITY(1,1) PRIMARY KEY,
    student_id INT FOREIGN KEY REFERENCES Student(id)
)
