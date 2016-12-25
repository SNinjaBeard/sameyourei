-- DROP TABLE sameyourei.todo
CREATE TABLE sameyourei.todo 
(
	id BIGSERIAL PRIMARY KEY,
	name CHARACTER VARYING(100) NOT NULL,
	completed boolean NOT NULL,
	due DATE,
	created DATE
)

INSERT INTO sameyourei.todo (name, completed, due, created)
	VALUES ('Do this', false, CURRENT_DATE, CURRENT_DATE);
INSERT INTO sameyourei.todo (name, completed, due, created)
	VALUES ('Do that', false, CURRENT_DATE, CURRENT_DATE);
INSERT INTO sameyourei.todo (name, completed, due, created)
	VALUES ('Do nothing', true, CURRENT_DATE, CURRENT_DATE);