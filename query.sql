CREATE TABLE users (
		id serial primary key,
		user_username varchar(50) unique not null,
		user_email varchar(50) not null,
		user_password varchar(50) not null
);

CREATE TABLE users_wallet (
		id serial primary key,
		user_id int unique not null,
		balance double precision default 0,
		CONSTRAINT user_wallet
			FOREIGN KEY(user_id)
			REFERENCES users(id)
);

CREATE TABLE transaction_log (
		id serial primary key,
		user_id int not null,
		amount double precision,
		date_time date not null default current_date,
		CONSTRAINT trx_log 
			FOREIGN KEY(user_id)
			REFERENCES users(id)
);

