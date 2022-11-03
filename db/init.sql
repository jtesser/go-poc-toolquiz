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
                         'drill'
                         );
insert into tool values (2,
                   'https://image.shutterstock.com/image-vector/saw-icon-isolated-on-white-600w-486498283.jpg',
                   'A tool used to cut through material',
                   'saw'
                   );
insert into tool values (3,
                   'https://image.shutterstock.com/image-photo/wrench-isolated-on-white-600w-785772973.jpg',
                   'A tool used to provide grip and mechanical advantage in applying torque to turn objects',
                   'wrench'
                   );