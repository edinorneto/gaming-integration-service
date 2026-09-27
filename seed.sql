-- Initial providers
-- Some providers are fictional and exist only for demonstration.
-- This project is independent and does not represent official integrations.

INSERT INTO providers (nome, active)
VALUES
    ('Jungle Originals', true),
    ('Aurora Gaming', true),
    ('Emerald Interactive', true),
    ('NovaPlay', true),
    ('Orbit Gaming', true);


-- Initial games
-- Jungle Originals titles are based on games publicly presented by
-- Jungle Gaming. Other titles are fictional sample data.

INSERT INTO games (nome, provider_id)
VALUES
    -- Jungle Originals
    ('Captain''s Treasure', 1),
    ('Fox the Course Seller', 1),
    ('Chimp Mines', 1),
    ('Goblin''s Gold', 1),

    -- Aurora Gaming
    ('Emerald Rush', 2),
    ('Golden Jungle', 2),
    ('Mystic Fortune', 2),
    ('Treasure Spins', 2),

    -- Emerald Interactive
    ('Neon Reels', 3),
    ('Diamond Vault', 3),
    ('Wild Horizon', 3),
    ('Lucky Temple', 3),

    -- NovaPlay
    ('Golden Quest', 4),
    ('Treasure Temple', 4),
    ('Moonlit Fortune', 4),
    ('Wild Expedition', 4),

    -- Orbit Gaming
    ('Cyber Fortune', 5),
    ('Jungle Nights', 5),
    ('Crystal Reels', 5),
    ('Lucky Orbit', 5);