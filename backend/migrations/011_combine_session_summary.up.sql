ALTER TABLE Sessions DROP COLUMN key_encounters;
ALTER TABLE Sessions DROP COLUMN key_npcs;

UPDATE Sessions SET
  chapters = 'Ch 1',
  summary = 'Velkynvelve, drow slave outpost. Escape the slave pen, flee into the Underdark before Ilvara breaks their will. Jorlan, her resentful ex-consort, quietly leaves the pen gate unlocked; rival demon bands clash through the outpost during a guard change, opening a gap for the prisoners to slip past the drow into the dark.'
WHERE session_num = 1;

UPDATE Sessions SET
  chapters = 'Ch 2',
  summary = 'Silken Paths, web-choked chasm past Velkynvelve. Cross the shifting webs toward the Darklake, evade Ilvara''s drow pursuit. Route passes the Oozing Temple, a cultist shrine sunk in the muck, ruled by the glabbagool and its followers.'
WHERE session_num = 2;

UPDATE Sessions SET
  chapters = 'Ch 3',
  summary = 'Sloobludop, kuo-toa settlement on the Darklake shore. Secure a boat, cross toward Gracklstugh. Succession feud between archpriest Ploopploopeen and daughter Bloppblippodd turns into a botched sacrifice ritual; the dark energy rips Demogorgon, Prince of Demons, out of the lake to rampage the settlement - escape by land or water before he notices.'
WHERE session_num = 3;

UPDATE Sessions SET
  chapters = 'Ch 3-4',
  summary = 'Darklake crossing, gates of Gracklstugh, the duergar City of Blades. Cross the Darklake by boat, talk or sneak past duergar border patrols into the Darklake District, the only part of the city open to outsiders.'
WHERE session_num = 4;

UPDATE Sessions SET
  chapters = 'Ch 4',
  summary = 'Gracklstugh''s Blade Bazaar, Overlake Hold, derro-infested Whorlstone Tunnels. Investigate demonic corruption for Stone Guard commander Errde Blackskull; side errands for Werz Saltbaron (gem delivery to Kazook Pickshine in Blingdenstone) and Gartokkar Xundorn of the Keepers of the Flame (recover Themberchaud''s stolen egg from derro courier Droki). Trail ends in the Whorlstone Tunnels: duergar alchemist and Droki both captured.'
WHERE session_num = 5;

UPDATE Sessions SET
  title = 'Neverlight Grove: Two Sovereigns',
  chapters = 'Ch 5',
  summary = 'Neverlight Grove, bioluminescent myconid enclave. Reach the grove, contact its two sovereigns Phylo and Basidia, learn what''s corrupting their people. Zuggtmoy''s spores already grip the circle sewing her wedding gown; confronting her pulls the party into a dreamlike back-and-forth fight against fungal servants before Basidia offers safe passage to Blingdenstone. Sarith Kzekarit, infected by spores in the fight, survives but stays under her influence.'
WHERE session_num = 6;

UPDATE Sessions SET
  chapters = 'Ch 6',
  summary = 'Blingdenstone, svirfneblin ruins built over an old war. Navigate the deep gnomes'' maze gate, deal with clan Goldwhisker''s wererats (secretly sheltering Topsy and Turvy''s kin), confront the Pudding King - a Juiblex-corrupted deep gnome running the city''s ooze infestation from its depths. Killing him yields the first ritual relic and a stone of controlling earth elementals; drow pursuers holding Jimjar captive surface here too.'
WHERE session_num = 7;

UPDATE Sessions SET
  title = 'Escape to the Surface & Gauntlgrym',
  chapters = 'Ch 7-8',
  summary = 'Surface world, then Gauntlgrym, King Bruenor Battlehammer''s dwarven hold. Break free of the Underdark at last, answer Bruenor''s summons, recount the demon lords'' rising, negotiate aid from five factions - Harpers, Order of the Gauntlet, Emerald Enclave, Lords'' Alliance, Zhentarim - to fund the return trip below. Eldeth Feldrun, killed by the party after her feud with Ront turned to attempted murder, honored at Gauntlgrym in her family''s name.'
WHERE session_num = 8;

UPDATE Sessions SET
  title = 'Return to the Underdark: Mantol-Derith & Gravenhollow',
  chapters = 'Ch 9-11',
  summary = 'Mantol-Derith''s four-faction trade enclave, then the stone library of Gravenhollow. Lead a dwarven-backed expedition back into the Underdark, broker passage through Mantol-Derith''s drow/duergar/svirfneblin/Zhentarim standoff, destroy Fraz-Urb''luu''s madness-gem before it ignites the enclave into war. Gravenhollow holds answers on the demon lords'' escape, points onward to Vizeran DeVir''s tower.'
WHERE session_num = 9;

UPDATE Sessions SET
  chapters = 'Ch 12',
  summary = 'Araj, tower of drow archmage Vizeran DeVir. Meet Vizeran, learn his plan to banish the demon lords - a ritual talisman, the dark heart, built from components scattered across the Underdark - agree to gather them for him.'
WHERE session_num = 10;

UPDATE Sessions SET
  title = 'The Wormwrithings: The Worm Nursery',
  chapters = 'Ch 13',
  summary = 'Wormwrithings, purple-worm-riddled tunnels south of the Labyrinth. Fight through troglodyte territory and the worm nursery, claim the intact unhatched purple worm egg for the dark heart ritual - Graz''zt''s component.'
WHERE session_num = 11;

UPDATE Sessions SET
  title = 'The Labyrinth: Maze Engine & Yeenoghu''s Hunt',
  chapters = 'Ch 14',
  summary = 'Baphomet''s Endless Maze, deep in the Labyrinth. Reach and activate the Maze Engine at its heart, survive Yeenoghu''s hunt through the maze and the goristro he kills along the way, claim two more components - goristro heart (Baphomet''s) and six feathers chiseled from petrified fallen angels in the Gallery of Angels (Orcus''s).'
WHERE session_num = 12;

UPDATE Sessions SET
  title = 'The Vast Oblivium: Karazikar''s Maw',
  chapters = 'Ch 13',
  summary = 'Vast Oblivium, beholder Karazikar''s lair off the Wormwrithings tunnels. Infiltrate the eye tyrant''s chasm-spanning domain, steal the central eye needed for the ritual - Demogorgon''s component - from Karazikar and its mad slave-priest Shedrak.'
WHERE session_num = 13;

UPDATE Sessions SET
  chapters = 'Ch 15',
  summary = 'Menzoberranzan, City of Spiders. Infiltrate Sorcere with the Council of Spiders'' help, steal Gromph Baenre''s demon-summoning grimoire, decide where to place the finished dark heart talisman to draw the demon lords into their final battle.'
WHERE session_num = 14;

UPDATE Sessions SET
  title = 'Araumycos: The Fetid Wedding',
  chapters = 'Ch 16',
  summary = 'Araumycos, vast fungal super-organism, site of Zuggtmoy''s wedding. Rapport with Basidia into Araumycos''s sleeping mind, turn it against Zuggtmoy before she claims it, fight through her wedding procession and Juiblex''s ooze spies to gather the last component - Zuggtmoy''s own mushrooms. Sarith Kzekarit, still carrying her infection since Neverlight Grove, dies here.'
WHERE session_num = 15;

UPDATE Sessions SET
  chapters = 'Ch 17',
  summary = 'Menzoberranzan, for Vizeran''s ritual and the final battle against the demon lords. Hold the site through nine hours of ritual while the summoned demon lords tear each other apart, finish off whoever''s left standing - Demogorgon, who crushes Orcus before turning on the party, wounded, without his usual resistance worn down.'
WHERE session_num = 16;
