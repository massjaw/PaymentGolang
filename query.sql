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

CREATE TABLE payment_log (
		id serial primary key,
		sender_id int not null,
		receipt_id int not null,
		amount double precision,
		date_time date not null default current_date,
		CONSTRAINT sender_log 
			FOREIGN KEY(sender_id)
			REFERENCES users(id),
		CONSTRAINT receipt_log 
				FOREIGN KEY(receipt_id)
				REFERENCES users(id)
);

CREATE TABLE payment_log_detail (
		id serial primary key,
		payment_id int not null,
		user_id int not null,
		amount double precision,
		date_time date not null default current_date,
		CONSTRAINT trx_log 
			FOREIGN KEY(user_id)
			REFERENCES users(id),
		CONSTRAINT payment_detail
				FOREIGN KEY(payment_id)
				REFERENCES payment_log(id)
);


