create table attendance (
    id serial not null,
    name varchar(255),
    attendance_time timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);