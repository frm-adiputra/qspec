CREATE TABLE tb_blog (ID TEXT PRIMARY KEY NOT NULL, Title TEXT NOT NULL, Content TEXT NOT NULL);


INSERT INTO tb_blog(ID, Title, Content)
VALUES("first",
       "My First Blog",
       "My First Blog content");


INSERT INTO tb_blog(ID, Title, Content)
VALUES("second",
       "My Second Blog",
       "My Second Blog content");


INSERT INTO tb_blog(ID, Title, Content)
VALUES("third",
       "My Third Blog",
       "My Third Blog content");


CREATE TABLE tb_comment (ID TEXT PRIMARY KEY NOT NULL, Username TEXT NOT NULL, COMMENT TEXT NOT NULL);


INSERT INTO tb_comment(ID, Username, COMMENT)
VALUES("1",
       "Username1",
       "Comment1");


INSERT INTO tb_comment(ID, Username, COMMENT)
VALUES("2",
       "Username2",
       "Comment2");
