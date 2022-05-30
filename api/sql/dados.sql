insert into usuarios (nome, nick, email, senha)
values
("Amilcar Ferreira", "Amilcar", "amilcar@mail.com", "$2a$10$uJHIRNOdQ9/7FoVr.hH6lO1zsWyfPWGtg0xV8tkCdbi.eJM4o/ICK"),
("Gabriel Ferreira", "Ga", "ga@mail.com", "$2a$10$uJHIRNOdQ9/7FoVr.hH6lO1zsWyfPWGtg0xV8tkCdbi.eJM4o/ICK"),
("Flora Luz Barbosa", "Flora", "flora@mail.com", "$2a$10$uJHIRNOdQ9/7FoVr.hH6lO1zsWyfPWGtg0xV8tkCdbi.eJM4o/ICK");

insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values 
("Publicação do Usuário 1", "Essa é a publicação do usuário 1", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3", 3);