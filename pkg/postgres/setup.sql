CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE users(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  image UUID,
  address TEXT,
  phone_number VARCHAR(20),
  user_type VARCHAR(20) CHECK (user_type IN ('admin', 'client')),
  registered_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  category VARCHAR(100),
  condition VARCHAR(50),
  images UUID[] 
);

CREATE TABLE auctions(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  seller_id UUID REFERENCES users(id),
  item_id UUID REFERENCES items(id),
  starting_bid DECIMAL(10,2) NOT NULL,
  current_bid DECIMAL(10,2),
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  extra_time_enabled BOOLEAN DEFAULT TRUE,
  extra_time_duration INTERVAL,
  extra_time_threshold INTERVAL,
  status BOOLEAN NOT NULL
);

CREATE TABLE bids(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  auction_id UUID REFERENCES auctions(id),
  bidder_id  UUID REFERENCES users(id),
  amount DECIMAL(10,2) NOT NULL,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  auction_id UUID REFERENCES auctions(id),
  buyer_id UUID REFERENCES users(id),
  seller_id UUID REFERENCES users(id),
  transaction_amount DECIMAL(10, 2) NOT NULL,
  transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shipping(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  transaction_id UUID REFERENCES transactions(id),
  shipping_address TEXT NOT NULL,
  tracking_number VARCHAR(255),
  status VARCHAR(50),
  estimated_delivery DATE
);

CREATE TABLE notifications(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id),
  message TEXT NOT NULL,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_read BOOLEAN NOT NULL
);
