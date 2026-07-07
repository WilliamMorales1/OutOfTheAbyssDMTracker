CREATE TABLE IF NOT EXISTS DemonLords (
    id                 INTEGER PRIMARY KEY AUTOINCREMENT,
    name               TEXT NOT NULL,
    dominions          TEXT NOT NULL,
    epithets           TEXT NOT NULL,
    layer              TEXT NOT NULL,
    description        TEXT NOT NULL,
    servants           TEXT NOT NULL,
    component          TEXT NOT NULL,
    component_location TEXT NOT NULL
);

INSERT INTO DemonLords (name, dominions, epithets, layer, description, servants, component, component_location) VALUES
('Demogorgon', 'Madness', 'Prince of Demons, the Sibilant Beast', '88: Gaping Maw', 'jungle continent, ocean, brine flats, twin fanged-skull towers, underwater labyrinth Abysm, drives visitors mad', 'Ixitxachitl', 'The central eye of a beholder', 'Vast Oblivium'),
('Fraz-Urb''luu', 'Lies', 'Prince of Deception, Lord of Illusions', '176: Hollow''s Heart', 'continent of white powder under starless sky, artificially sunlit', 'Rakshasa', 'The lier''s gem', 'Mantol-Derith'),
('Graz''zt', 'Lust', 'Prince of Pleasure, the Dark Prince', '45-47: Azzagrat', 'three linked layers, steppe, luminous dark sky, blue-sun realm with reversed heat/cold', 'Cambion', 'The unhatched egg of a purple worm', 'Wormwrithings'),
('Orcus', 'Death', 'Prince of Undeath, the Blood Lord', '113: Thanatos', 'perpetual moonlit night, thin air, decaying land, dead rise as undead within an hour', 'Undead', 'Six feathers from six different angels', 'Labyrinth'),
('Yeenoghu', 'Hunt', 'Beast of Butchery, the Gnoll Lord', '422: Death Dells', 'barren hunting grounds of hills and ravines, scarce civilization', 'Gnoll', 'The butchered remains of a gnoll', 'Labyrinth'),
('Baphomet', 'Savagery', 'Prince of Beasts, the Horned King', '600: Endless Maze', 'infinite labyrinth home to minotaurs, ogres, goristros', 'Minotaur', 'The heart of a goristro', 'Labyrinth'),
('Juiblex', 'Gluttony', 'Faceless Lord, the Oozing Hunger', '222: Shedaklah (contested)', 'fetid swamp of oozes and molds under filth-raining green clouds', 'Oozes', 'Ooze of Hunger', 'Blingdenstone'),
('Zuggtmoy', 'Parasitism', 'Queen of Fungi, Lady of Rot and Decay', '222: Shedaklah (contested)', 'giant fungal palace, shelf-fungi bridges, acidic puffballs, poisonous vapors', 'Myconid', 'Mushrooms of a Yestobrod', 'Araumycos');
