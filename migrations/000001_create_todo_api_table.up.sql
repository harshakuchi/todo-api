/*
    migrations are used to manage changed to the postgreSQL database safely and consistently
    it is a database schema version control

    up migration applies changes where as down migration reverts changes
    migrate tool tracks which migrations ran, applies new ones in order, prevents duplicates, supports rollback, keeps all environments synced  
*/
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);