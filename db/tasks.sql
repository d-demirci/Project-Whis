create table tasks ( 
    id int not null auto_increment primary key,
    Executions varchar(255), 
    RandomId varchar(255),
    TaskName varchar(255),
    DateIssued varchar(255),
    CommandName varchar(255),
    MaxExecutions varchar(255),
    TaskTimeout varchar(255)
     );