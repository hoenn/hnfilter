CREATE TABLE posts (
    By varchar(255),
    Descendants integer,
    ID integer,
    Kids integer[],
    Score integer,
    Time timestamp,
    Title varchar(255),
    Url varchar(255)
)
