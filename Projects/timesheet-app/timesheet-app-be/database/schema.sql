CREATE TABLE Employee (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Type VARCHAR(255),
    Username VARCHAR(255),
    Password VARCHAR(255)
);

CREATE TABLE Projects (
    ProjectID INT AUTO_INCREMENT PRIMARY KEY,
    ProjectName VARCHAR(255)
);

CREATE TABLE Subprojects (
    SubProjectID INT AUTO_INCREMENT PRIMARY KEY,
    SubProjectName VARCHAR(255),
    ProjectID INT,
    FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID)
);

CREATE TABLE Timesheets (
    TimesheetID INT AUTO_INCREMENT PRIMARY KEY,
    EmployeeID INT,
    ProjectID INT,
    SubProjectID INT,
    JiraSnowID VARCHAR(255),
    TaskDescription TEXT,
    HoursSpent FLOAT,
    Comments TEXT,
    CreatedAt DATETIME,
    FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID),
    FOREIGN KEY (SubProjectID) REFERENCES Subprojects(SubProjectID),
    FOREIGN KEY (EmployeeID) REFERENCES Employee(EmployeeID)
);