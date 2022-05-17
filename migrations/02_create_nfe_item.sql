create table nfe_item (
	access_key varchar(54) not null,
	description varchar (256) not null,
	quantity decimal not null,
	unit_measure varchar (50),
	unit_price numeric(10,2) not null,
	total_price numeric(10,2) not null,
	CONSTRAINT fk_nfe
      FOREIGN KEY(access_key) 
	  REFERENCES nfe(access_key)
)

