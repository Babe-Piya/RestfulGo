CREATE TABLE `todo`.`todo`
( `ID` INT(10) NOT NULL AUTO_INCREMENT , `TITLE` VARCHAR(255) NOT NULL ,
`CONTENT` VARCHAR(255) NOT NULL , `IS_DONE` BOOLEAN NOT NULL DEFAULT false ,
`CREATE_AT` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (`ID`));

INSERT INTO todo (TITLE,CONTENT) values('test1','test1');
INSERT INTO todo (TITLE,CONTENT) values('test2','test2');
INSERT INTO todo (TITLE,CONTENT) values('test3','test3');
INSERT INTO todo (TITLE,CONTENT) values('test4','test4');