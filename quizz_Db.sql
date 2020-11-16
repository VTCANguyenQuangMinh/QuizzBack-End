drop database if exists Quizz;

create database Quizz;

use Quizz;

create table questions(
	question_id int auto_increment primary key,
    question varchar(255) not null unique,
    type_answer varchar(50) not null,
    
);

create table answers(
answer_id decimal(20,2) primary key,
question_id int not null,
answer varchar(255) not null,
correct boolean not null,
score decimal(20,2) default 1
constraint fk_Question_answer foreign key(question_id) references questions(question_id)
);

insert into questions(question,type_answer) values
	('Chu vi hình tròn có đường kính d= 2,5 cm là ?', 'single'),
    ('Cho số 4...9. Những số thích hợp viết vào chỗ chấm để được số chia hết cho 3 là:','multi'),
    ('Trong một mặt phẳng, hai đường thẳng không giao nhau thì hai đường thẳng đó song song.','y/n')
;
    
select * from questions;
    
insert into answers values
(1.1,1,'7.85 cm',true,1),
(1.2,1,'15.7 cm',false,0),
(1.3,1,'785 cm',false,0),
(1.4,1,'175 cm',false,0),
(2.1,2,'0',false,0),
(2.2,2,'5',true,0.5),
(2.3,2,'2',true,0.5),
(2.4,2,'7',false,0),
(2.5,2,'4',false,0),
(3.1,3,'Đúng',true,1),
(3.2,3,'Sai',false,0)
;

select * from answers;