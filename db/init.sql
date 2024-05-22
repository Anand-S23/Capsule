CREATE TABLE IF NOT EXISTS users (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email      VARCHAR(256) NOT NULL UNIQUE,
    phone      VARCHAR(20) NOT NULL,
    name       VARCHAR(75) NOT NULL,
    password   VARCHAR(128) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS connections (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id   UUID REFERENCES users(id) NOT NULL,
    first_name VARCHAR(35) NOT NULL,
    last_name  VARCHAR(35) NOT NULL,
    email      VARCHAR(256),
    phone      VARCHAR(20),
    linkedin   VARCHAR(128),
    company    VARCHAR(128),
    dob        DATE,
    notes      VARCHAR(1048576) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS meetings (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id    UUID REFERENCES users(id) NOT NULL,
    time        DATETIME NOT NULL,
    location    VARCHAR(256),
    description VARCHAR(256),
    notes       VARCHAR(1048576) DEFAULT '',
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS participants (
    id            SERIAL PRIMARY KEY NOT NULL,
    meeting_id    UUID REFERENCES meetings(id) NOT NULL,
    connection_id UUID REFERENCES connections(id) NOT NULL,
    owner_id      UUID REFERENCES users(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS reminders (
    id            SERIAL PRIMARY KEY NOT NULL,
    connection_id UUID REFERENCES connections(id) NOT NULL,
    owner_id      UUID REFERENCES users(id) NOT NULL,
    time          DATETIME NOT NULL,
    description   VARCHAR(256),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

