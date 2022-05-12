create TABLE nfe (
  access_key varchar(54) not null primary key,
  url varchar(256) not null,
  issue_date TIMESTAMP WITH TIME ZONE not null,
  store_name varchar(256) not null,
  store_cnpj varchar(18) not null,
  total numeric(10,2) not null,
  discount numeric(10,2) not null,
  created_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)