CREATE DATABASE IF NOT EXISTS redesocialgo;

create user 'golang'@'localhost' IDENTIFIELD by 'golang';

grant all privileges on redesocialgo.* to 'golang'@'localhost';

DROP TABLE seguidores;
DROP TABLE usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL unique,
    email varchar(50) NOT NULL unique,
    senha varchar(100) NOT NULL,
    criadoEm timestamp default current_timestamp()
);

CREATE TABLE seguidores(
    usuario_id int not null, 
    seguidor_id int not null, 
    primary key(usuario_id, seguidor_id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE 
);


