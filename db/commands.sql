drop table commands;
create table commands (
    id int not null auto_increment primary key,
    task int, 
    UID varchar(128), 
    DAT varchar(128), 
    Command varchar(1024), 
    Parameters varchar(128), 
    Status varchar(128), 
    DateIssued varchar(128), 
    Timeout varchar(128)
);