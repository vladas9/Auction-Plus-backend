CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE user(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  address TEXT,
  phone_number VARCHAR(20),
  user_type VARCHAR(20) CHECK (user_type IN ("admin", "client")),
  registred_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE item(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  category VARCHAR(100),
  condition VARCHAR(50),
  image TEXT[]
);

CREATE TABLE auction(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  seller_id UUID REFERENCES user(id),
  item_id UUID REFERENCES item(id),
  starting_bid DECIMAL(10,2) NOT NULL,
  current_bid DECIMAL(10,2),
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  extra_time_enabled BOOLEAN DEFAULT TRUE,
  extra_time_duration INTERVAL,
  extra_time_threshold INTERVAL,
  status VARCHAR(50)
);

CREATE TABLE bid(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  auction_id UUID REFERENCES auction(id),
  bidder_id  UUID REFERENCES user(id),
  amount DECIMAL(10,2) NOT NULL,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE transaction(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  auction_id UUID REFERENCES auction(id),
  buyer_id UUID REFERENCES user(id),
  seller_id UUID REFERENCES user(id),
  transaction_amount DECIMAL(10, 2) NOT NULL,
  transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shipping(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  transaction_id UUID REFERENCES transaction(id),
  shipping_address TEXT NOT NULL,
  tracking_number VARCHAR(255),
  status VARCHAR(50),
  estimated_delivary DATE,
);

CREATE TABLE notification(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES user(id),
  message TEXT NOT NULL,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_readed BOOLEAN NOT NULL,
);
