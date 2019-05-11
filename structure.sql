DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    author varchar(255),
    id integer,
    kids integer[],
    parent integer,
    body text,
    time timestamp,
    PRIMARY KEY (id)
);
