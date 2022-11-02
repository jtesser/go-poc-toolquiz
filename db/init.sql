CREATE DATABASE toolquiz;

\connect toolquiz;

CREATE TABLE tool (
    question_num integer,
    image_url varchar(512),
    hint text,
    answer varchar(256),
    PRIMARY KEY (question_num)
);

insert into tool values (1,
                         'https://image.shutterstock.com/image-photo/power-drill-cordless-screwdriver-battery-600w-1773580895.jpg',
                         'Uses an electrical motor that rotates a replaceable drill bit',
                         'Drill'
                         )