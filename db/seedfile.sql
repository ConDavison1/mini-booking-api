
-- Create Users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL -- Note: Plaintext for demo purposes only
);

-- Insert example user
INSERT INTO users (id, email, password)
VALUES ('33333333-3333-3333-3333-333333333333', 'admin@example.com', 'password123');

-- Create Program table
CREATE TABLE IF NOT EXISTS programs (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    capacity INT NOT NULL,
    registered INT DEFAULT 0,
    visibility TEXT CHECK (visibility IN ('PUBLIC', 'PRIVATE', 'TEAM_ONLY')) NOT NULL
);

-- Create Booking table
CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY,
    program_id UUID REFERENCES programs(id),
    user_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert example programs
INSERT INTO programs (id, name, description, start_date, end_date, capacity, registered, visibility)
VALUES 
  ('11111111-1111-1111-1111-111111111111', 'Elite Camp', 'Advanced training program.', '2025-07-01', '2025-07-10', 30, 10, 'PUBLIC'),
  ('22222222-2222-2222-2222-222222222222', 'Beginners Bootcamp', 'Introductory sessions.', '2025-08-01', '2025-08-05', 20, 5, 'PRIVATE');
