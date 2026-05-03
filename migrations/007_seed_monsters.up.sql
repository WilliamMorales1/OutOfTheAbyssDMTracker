--
-- PostgreSQL database dump
--

\restrict cZs5kI8waB1UcnEIlNbjpFYT8ctjEsFMsytk0y6smDwX0Th404Bkn6wMJc0y2cj

-- Dumped from database version 18.3
-- Dumped by pg_dump version 18.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: monsters; Type: TABLE DATA; Schema: public; Owner: wsm52
--

INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (4, 'Behir', 'monstrosity', '11', 168, '16d12', 17, 'natural', 'walk 50 ft., climb 40 ft.', 23, 16, 18, 7, 14, 12, 'Saves: Stealth+7', NULL, '', 'lightning', '', 'darkvision 90 ft., passive Perception 16', 'Draconic', '', 'Multiattack: The behir makes two attacks: one with its bite and one to constrict.

Bite: Melee Weapon Attack: +10 to hit, reach 10 ft., one target. Hit: 22 (3d10 + 6) piercing damage.

Constrict: Melee Weapon Attack: +10 to hit, reach 5 ft., one Large or smaller creature. Hit: 17 (2d10 + 6) bludgeoning damage plus 17 (2d10 + 6) slashing damage. The target is grappled (escape DC 16) if the behir isn''t already constricting a creature, and the target is restrained until this grapple ends.

Lightning Breath: The behir exhales a line of lightning that is 20 ft. long and 5 ft. wide. Each creature in that line must make a DC 16 Dexterity saving throw, taking 66 (12d10) lightning damage on a failed save, or half as much damage on a successful one.

Swallow: The behir makes one bite attack against a Medium or smaller target it is grappling. If the attack hits, the target is also swallowed, and the grapple ends. While swallowed, the target is blinded and restrained, it has total cover against attacks and other effects outside the behir, and it takes 21 (6d6) acid damage at the start of each of the behir''s turns. A behir can have only one creature swallowed at a time.
If the behir takes 30 damage or more on a single turn from the swallowed creature, the behir must succeed on a DC 14 Constitution saving throw at the end of that turn or regurgitate the creature, which falls prone in a space within 10 ft. of the behir. If the behir dies, a swallowed creature is no longer restrained by it and can escape from the corpse by using 15 ft. of movement, exiting prone.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (6, 'Bugbear', 'humanoid', '1', 27, '5d8', 16, 'armor', 'walk 30 ft.', 15, 14, 13, 8, 11, 9, 'Saves: Survival+2', NULL, '', '', '', 'darkvision 60 ft., passive Perception 10', 'Common, Goblin', 'Brute: A melee weapon deals one extra die of its damage when the bugbear hits with it (included in the attack).

Surprise Attack: If the bugbear surprises a creature and hits it with an attack during the first round of combat, the target takes an extra 7 (2d6) damage from the attack.', 'Morningstar: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 11 (2d8 + 2) piercing damage.

Javelin: Melee or Ranged Weapon Attack: +4 to hit, reach 5 ft. or range 30/120 ft., one target. Hit: 9 (2d6 + 2) piercing damage in melee or 5 (1d6 + 2) piercing damage at range.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (3, 'Animated Armor', 'construct', '1', 33, '6d8', 18, 'natural', 'walk 25 ft.', 14, 11, 13, 1, 3, 1, '', NULL, '', 'poison, psychic', 'Blinded, Charmed, Deafened, Exhaustion, Frightened, Paralyzed, Petrified, Poisoned', 'blindsight 60 ft. (blind beyond this radius), passive Perception 6', '', 'Antimagic Susceptibility: The armor is incapacitated while in the area of an antimagic field. If targeted by dispel magic, the armor must succeed on a Constitution saving throw against the caster''s spell save DC or fall unconscious for 1 minute.

False Appearance: While the armor remains motionless, it is indistinguishable from a normal suit of armor.', 'Multiattack: The armor makes two melee attacks.

Slam: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) bludgeoning damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (5, 'Black Pudding', 'ooze', '4', 85, '10d10', 7, 'dex', 'walk 20 ft., climb 20 ft.', 16, 5, 16, 1, 6, 1, '', NULL, '', 'acid, cold, lightning, slashing', 'Blinded, Charmed, Exhaustion, Frightened, Prone', 'blindsight 60 ft. (blind beyond this radius), passive Perception 8', '', 'Amorphous: The pudding can move through a space as narrow as 1 inch wide without squeezing.

Corrosive Form: A creature that touches the pudding or hits it with a melee attack while within 5 feet of it takes 4 (1d8) acid damage. Any nonmagical weapon made of metal or wood that hits the pudding corrodes. After dealing damage, the weapon takes a permanent and cumulative -1 penalty to damage rolls. If its penalty drops to -5, the weapon is destroyed. Nonmagical ammunition made of metal or wood that hits the pudding is destroyed after dealing damage. The pudding can eat through 2-inch-thick, nonmagical wood or metal in 1 round.

Spider Climb: The pudding can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.', 'Pseudopod: Melee Weapon Attack: +5 to hit, reach 5 ft., one target. Hit: 6 (1d6 + 3) bludgeoning damage plus 18 (4d8) acid damage. In addition, nonmagical armor worn by the target is partly dissolved and takes a permanent and cumulative -1 penalty to the AC it offers. The armor is destroyed if the penalty reduces its AC to 10.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (11, 'Dretch', 'fiend', '1/4', 18, '4d6', 11, 'natural', 'walk 20 ft.', 11, 11, 12, 5, 8, 3, '', NULL, 'cold, fire, lightning', 'poison', 'Poisoned', 'darkvision 60 ft., passive Perception 9', 'Abyssal, telepathy 60 ft. (works only with creatures that understand Abyssal)', '', 'Multiattack: The dretch makes two attacks: one with its bite and one with its claws.

Bite: Melee Weapon Attack: +2 to hit, reach 5 ft., one target. Hit: 3 (1d6) piercing damage.

Claws: Melee Weapon Attack: +2 to hit, reach 5 ft., one target. Hit: 5 (2d4) slashing damage.

Fetid Cloud: A 10-foot radius of disgusting green gas extends out from the dretch. The gas spreads around corners, and its area is lightly obscured. It lasts for 1 minute or until a strong wind disperses it. Any creature that starts its turn in that area must succeed on a DC 11 Constitution saving throw or be poisoned until the start of its next turn. While poisoned in this way, the target can take either an action or a bonus action on its turn, not both, and can''t take reactions.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (117, 'Grell', 'aberration', '3', 55, '10d8+10', 12, '', '10 ft., fly 30 ft. (hover)', 15, 14, 13, 12, 11, 9, 'Skills: Perception +2, Stealth +4', NULL, '', 'lightning', 'blinded, prone', 'blindsight 60 ft. (blind beyond), passive Perception 12', 'Grell', '', 'Multiattack: One tentacle attack and one beak attack.

Tentacles: +4 to hit, reach 10 ft. Hit: 7 (1d8+2) piercing. The target is grappled (escape DC 15) and must make a DC 11 Constitution save or be paralyzed until the grapple ends.

Beak: +4 to hit, reach 5 ft. Hit: 7 (2d4+2) piercing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (8, 'Darkmantle', 'monstrosity', '1/2', 22, '5d6', 11, 'dex', 'walk 10 ft., fly 30 ft.', 16, 12, 13, 2, 10, 5, 'Saves: Stealth+3', NULL, '', '', '', 'blindsight 60 ft., passive Perception 10', '', 'Echolocation: The darkmantle can''t use its blindsight while deafened.

False Appearance: While the darkmantle remains motionless, it is indistinguishable from a cave formation such as a stalactite or stalagmite.', 'Crush: Melee Weapon Attack: +5 to hit, reach 5 ft., one creature. Hit: 6 (1d6 + 3) bludgeoning damage, and the darkmantle attaches to the target. If the target is Medium or smaller and the darkmantle has advantage on the attack roll, it attaches by engulfing the target''s head, and the target is also blinded and unable to breathe while the darkmantle is attached in this way.
While attached to the target, the darkmantle can attack no other creature except the target but has advantage on its attack rolls. The darkmantle''s speed also becomes 0, it can''t benefit from any bonus to its speed, and it moves with the target.
A creature can detach the darkmantle by making a successful DC 13 Strength check as an action. On its turn, the darkmantle can detach itself from the target by using 5 feet of movement.

Darkness Aura: A 15-foot radius of magical darkness extends out from the darkmantle, moves with it, and spreads around corners. The darkness lasts as long as the darkmantle maintains concentration, up to 10 minutes (as if concentrating on a spell). Darkvision can''t penetrate this darkness, and no natural light can illuminate it. If any of the darkness overlaps with an area of light created by a spell of 2nd level or lower, the spell creating the light is dispelled.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (118, 'Grick Alpha', 'monstrosity', '7', 75, '10d10+20', 18, 'natural armor', '30 ft., climb 30 ft.', 18, 13, 15, 4, 14, 9, 'Skills: Perception +5', NULL, 'bludgeoning, piercing, slashing from nonmagical attacks', '', '', 'darkvision 60 ft., passive Perception 15', '—', 'Stone Camouflage: Advantage on Dexterity (Stealth) checks in rocky terrain.', 'Multiattack: One tentacle attack and one beak attack, or two rock attacks.

Tentacles: +7 to hit, reach 10 ft. Hit: 22 (4d8+4) slashing. Target is grappled (escape DC 15).

Beak: +7 to hit, reach 5 ft. Hit: 13 (2d8+4) piercing.

Rock: +7 to hit, range 30/120 ft. Hit: 13 (2d8+4) bludgeoning.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (17, 'Gargoyle', 'elemental', '2', 52, '7d8', 15, 'natural', 'walk 30 ft., fly 60 ft.', 15, 11, 16, 6, 11, 7, '', NULL, 'bludgeoning, piercing, and slashing from nonmagical weapons that aren''t adamantine', 'poison', 'Exhaustion, Petrified, Poisoned', 'darkvision 60 ft., passive Perception 10', 'Terran', 'False Appearance: While the gargoyle remains motion less, it is indistinguishable from an inanimate statue.', 'Multiattack: The gargoyle makes two attacks: one with its bite and one with its claws.

Bite: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) piercing damage.

Claws: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) slashing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (119, 'Hook Horror', 'monstrosity', '3', 75, '10d10+20', 15, 'natural armor', '30 ft., climb 30 ft.', 18, 10, 15, 6, 12, 7, 'Skills: Perception +3', NULL, '', '', '', 'blindsight 10 ft., darkvision 60 ft., passive Perception 13', 'Hook Horror', 'Echolocation: The hook horror can''t use its blindsight while deafened.

Keen Hearing: Advantage on Wisdom (Perception) checks that rely on hearing.', 'Multiattack: Two hook attacks.

Hook: +6 to hit, reach 10 ft. Hit: 11 (2d6+4) piercing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (13, 'Drow', 'humanoid', '1/4', 13, '3d8', 15, 'armor', 'walk 30 ft.', 10, 14, 10, 11, 11, 12, 'Saves: Stealth+4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 12', 'Elvish, Undercommon', 'Fey Ancestry: The drow has advantage on saving throws against being charmed, and magic can''t put the drow to sleep.

Innate Spellcasting: The drow''s spellcasting ability is Charisma (spell save DC 11). It can innately cast the following spells, requiring no material components:
At will: dancing lights
1/day each: darkness, faerie fire

Sunlight Sensitivity: While in sunlight, the drow has disadvantage on attack rolls, as well as on Wisdom (Perception) checks that rely on sight.', 'Shortsword: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) piercing damage.

Hand Crossbow: Ranged Weapon Attack: +4 to hit, range 30/120 ft., one target. Hit: 5 (1d6 + 2) piercing damage, and the target must succeed on a DC 13 Constitution saving throw or be poisoned for 1 hour. If the saving throw fails by 5 or more, the target is also unconscious while poisoned in this way. The target wakes up if it takes damage or if another creature takes an action to shake it awake.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (15, 'Ettin', 'giant', '4', 85, '10d10', 12, 'natural', 'walk 40 ft.', 21, 8, 17, 6, 10, 8, 'Saves: Perception+4', NULL, '', '', '', 'darkvision 60 ft., passive Perception 14', 'Giant, Orc', 'Two Heads: The ettin has advantage on Wisdom (Perception) checks and on saving throws against being blinded, charmed, deafened, frightened, stunned, and knocked unconscious.

Wakeful: When one of the ettin''s heads is asleep, its other head is awake.', 'Multiattack: The ettin makes two attacks: one with its battleaxe and one with its morningstar.

Battleaxe: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 14 (2d8 + 5) slashing damage.

Morningstar: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 14 (2d8 + 5) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (16, 'Fire Elemental', 'elemental', '5', 102, '12d10', 13, 'dex', 'walk 50 ft.', 10, 17, 16, 6, 10, 7, '', NULL, 'bludgeoning, piercing, and slashing from nonmagical weapons', 'fire, poison', 'Exhaustion, Grappled, Paralyzed, Petrified, Poisoned, Prone, Restrained, Unconscious', 'darkvision 60 ft., passive Perception 10', 'Ignan', 'Fire Form: The elemental can move through a space as narrow as 1 inch wide without squeezing. A creature that touches the elemental or hits it with a melee attack while within 5 ft. of it takes 5 (1d10) fire damage. In addition, the elemental can enter a hostile creature''s space and stop there. The first time it enters a creature''s space on a turn, that creature takes 5 (1d10) fire damage and catches fire; until someone takes an action to douse the fire, the creature takes 5 (1d10) fire damage at the start of each of its turns.

Illumination: The elemental sheds bright light in a 30-foot radius and dim light in an additional 30 ft..

Water Susceptibility: For every 5 ft. the elemental moves in water, or for every gallon of water splashed on it, it takes 1 cold damage.', 'Multiattack: The elemental makes two touch attacks.

Touch: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 10 (2d6 + 3) fire damage. If the target is a creature or a flammable object, it ignites. Until a creature takes an action to douse the fire, the target takes 5 (1d10) fire damage at the start of each of its turns.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (20, 'Giant Fire Beetle', 'beast', '0', 4, '1d6', 13, 'natural', 'walk 30 ft.', 8, 10, 12, 1, 7, 3, '', NULL, '', '', '', 'blindsight 30 ft., passive Perception 8', '', 'Illumination: The beetle sheds bright light in a 10-foot radius and dim light for an additional 10 ft..', 'Bite: Melee Weapon Attack: +1 to hit, reach 5 ft., one target. Hit: 2 (1d6 - 1) slashing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (24, 'Goblin', 'humanoid', '1/4', 7, '2d6', 15, 'armor', 'walk 30 ft.', 8, 14, 10, 10, 8, 8, 'Saves: Stealth+6', NULL, '', '', '', 'darkvision 60 ft., passive Perception 9', 'Common, Goblin', 'Nimble Escape: The goblin can take the Disengage or Hide action as a bonus action on each of its turns.', 'Scimitar: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) slashing damage.

Shortbow: Ranged Weapon Attack: +4 to hit, range 80/320 ft., one target. Hit: 5 (1d6 + 2) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (21, 'Giant Spider', 'beast', '1', 26, '4d10', 14, 'natural', 'walk 30 ft., climb 30 ft.', 14, 16, 12, 2, 11, 4, 'Saves: Stealth+7', NULL, '', '', '', 'darkvision 60 ft., blindsight 10 ft., passive Perception 10', '', 'Spider Climb: The spider can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.

Web Sense: While in contact with a web, the spider knows the exact location of any other creature in contact with the same web.

Web Walker: The spider ignores movement restrictions caused by webbing.', 'Bite: Melee Weapon Attack: +5 to hit, reach 5 ft., one creature. Hit: 7 (1d8 + 3) piercing damage, and the target must make a DC 11 Constitution saving throw, taking 9 (2d8) poison damage on a failed save, or half as much damage on a successful one. If the poison damage reduces the target to 0 hit points, the target is stable but poisoned for 1 hour, even after regaining hit points, and is paralyzed while poisoned in this way.

Web: Ranged Weapon Attack: +5 to hit, range 30/60 ft., one creature. Hit: The target is restrained by webbing. As an action, the restrained target can make a DC 12 Strength check, bursting the webbing on a success. The webbing can also be attacked and destroyed (AC 10; hp 5; vulnerability to fire damage; immunity to bludgeoning, poison, and psychic damage).', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (22, 'Giant Wolf Spider', 'beast', '1/4', 11, '2d8', 13, 'dex', 'walk 40 ft., climb 40 ft.', 12, 16, 13, 3, 12, 4, 'Saves: Stealth+7', NULL, '', '', '', 'darkvision 60 ft., blindsight 10 ft., passive Perception 13', '', 'Spider Climb: The spider can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.

Web Sense: While in contact with a web, the spider knows the exact location of any other creature in contact with the same web.

Web Walker: The spider ignores movement restrictions caused by webbing.', 'Bite: Melee Weapon Attack: +3 to hit, reach 5 ft., one creature. Hit: 4 (1d6 + 1) piercing damage, and the target must make a DC 11 Constitution saving throw, taking 7 (2d6) poison damage on a failed save, or half as much damage on a successful one. If the poison damage reduces the target to 0 hit points, the target is stable but poisoned for 1 hour, even after regaining hit points, and is paralyzed while poisoned in this way.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (23, 'Gnoll', 'humanoid', '1/2', 22, '5d8', 15, 'armor', 'walk 30 ft.', 14, 12, 11, 6, 10, 7, '', NULL, '', '', '', 'darkvision 60 ft., passive Perception 10', 'Gnoll', 'Rampage: When the gnoll reduces a creature to 0 hit points with a melee attack on its turn, the gnoll can take a bonus action to move up to half its speed and make a bite attack.', 'Bite: Melee Weapon Attack: +4 to hit, reach 5 ft., one creature. Hit: 4 (1d4 + 2) piercing damage.

Spear: Melee or Ranged Weapon Attack: +4 to hit, reach 5 ft. or range 20/60 ft., one target. Hit: 5 (1d6 + 2) piercing damage, or 6 (1d8 + 2) piercing damage if used with two hands to make a melee attack.

Longbow: Ranged Weapon Attack: +3 to hit, range 150/600 ft., one target. Hit: 5 (1d8 + 1) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (27, 'Grick', 'monstrosity', '2', 27, '6d8', 14, 'natural', 'walk 30 ft., climb 30 ft.', 14, 14, 11, 3, 14, 5, '', NULL, 'bludgeoning, piercing, and slashing from nonmagical weapons', '', '', 'darkvision 60 ft., passive Perception 12', '', 'Stone Camouflage: The grick has advantage on Dexterity (Stealth) checks made to hide in rocky terrain.', 'Multiattack: The grick makes one attack with its tentacles. If that attack hits, the grick can make one beak attack against the same target.

Tentacles: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 9 (2d6 + 2) slashing damage.

Beak: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (31, 'Merrow', 'monstrosity', '2', 45, '6d10', 13, 'natural', 'walk 10 ft., swim 40 ft.', 18, 10, 15, 8, 10, 9, '', NULL, '', '', '', 'darkvision 60 ft., passive Perception 10', 'Abyssal, Aquan', 'Amphibious: The merrow can breathe air and water.', 'Multiattack: The merrow makes two attacks: one with its bite and one with its claws or harpoon.

Bite: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 8 (1d8 + 4) piercing damage.

Claws: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 9 (2d4 + 4) slashing damage.

Harpoon: Melee or Ranged Weapon Attack: +6 to hit, reach 5 ft. or range 20/60 ft., one target. Hit: 11 (2d6 + 4) piercing damage. If the target is a Huge or smaller creature, it must succeed on a Strength contest against the merrow or be pulled up to 20 feet toward the merrow.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (28, 'Grimlock', 'humanoid', '1/4', 11, '2d8', 11, 'dex', 'walk 30 ft.', 16, 12, 12, 9, 8, 6, 'Saves: Stealth+3', NULL, '', '', 'Blinded', 'blindsight 30 ft. or 10 ft. while deafened (blind beyond this radius), passive Perception 13', 'Undercommon', 'Blind Senses: The grimlock can''t use its blindsight while deafened and unable to smell.

Keen Hearing and Smell: The grimlock has advantage on Wisdom (Perception) checks that rely on hearing or smell.

Stone Camouflage: The grimlock has advantage on Dexterity (Stealth) checks made to hide in rocky terrain.', 'Spiked Bone Club: Melee Weapon Attack: +5 to hit, reach 5 ft., one target. Hit: 5 (1d4 + 3) bludgeoning damage plus 2 (1d4) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (29, 'Hezrou', 'fiend', '8', 136, '13d10', 16, 'natural', 'walk 30 ft.', 19, 17, 20, 5, 12, 13, 'Saves: STR+7, CON+8, WIS+4', NULL, 'cold, fire, lightning, bludgeoning, piercing, and slashing from nonmagical weapons', 'poison', 'Poisoned', 'darkvision 120 ft., passive Perception 11', 'Abyssal, telepathy 120 ft.', 'Magic Resistance: The hezrou has advantage on saving throws against spells and other magical effects.

Stench: Any creature that starts its turn within 10 feet of the hezrou must succeed on a DC 14 Constitution saving throw or be poisoned until the start of its next turn. On a successful saving throw, the creature is immune to the hezrou''s stench for 24 hours.', 'Multiattack: The hezrou makes three attacks: one with its bite and two with its claws.

Bite: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 15 (2d10 + 4) piercing damage.

Claws: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 11 (2d6 + 4) slashing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (30, 'Magmin', 'elemental', '1/2', 9, '2d6', 14, 'natural', 'walk 30 ft.', 7, 15, 12, 8, 11, 10, '', NULL, 'bludgeoning, piercing, and slashing from nonmagical weapons', 'fire', '', 'darkvision 60 ft., passive Perception 10', 'Ignan', 'Death Burst: When the magmin dies, it explodes in a burst of fire and magma. Each creature within 10 ft. of it must make a DC 11 Dexterity saving throw, taking 7 (2d6) fire damage on a failed save, or half as much damage on a successful one. Flammable objects that aren''t being worn or carried in that area are ignited.

Ignited Illumination: As a bonus action, the magmin can set itself ablaze or extinguish its flames. While ablaze, the magmin sheds bright light in a 10-foot radius and dim light for an additional 10 ft.', 'Touch: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 7 (2d6) fire damage. If the target is a creature or a flammable object, it ignites. Until a target takes an action to douse the fire, the target takes 3 (1d6) fire damage at the end of each of its turns.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (37, 'Purple Worm', 'monstrosity', '15', 247, '15d20', 18, 'natural', 'walk 50 ft., burrow 30 ft.', 28, 7, 22, 1, 8, 4, 'Saves: CON+11, WIS+4', NULL, '', '', '', 'blindsight 30 ft., tremorsense 60 ft., passive Perception 9', '', 'Tunneler: The worm can burrow through solid rock at half its burrow speed and leaves a 10-foot-diameter tunnel in its wake.', 'Multiattack: The worm makes two attacks: one with its bite and one with its stinger.

Bite: Melee Weapon Attack: +9 to hit, reach 10 ft., one target. Hit: 22 (3d8 + 9) piercing damage. If the target is a Large or smaller creature, it must succeed on a DC 19 Dexterity saving throw or be swallowed by the worm. A swallowed creature is blinded and restrained, it has total cover against attacks and other effects outside the worm, and it takes 21 (6d6) acid damage at the start of each of the worm''s turns.
If the worm takes 30 damage or more on a single turn from a creature inside it, the worm must succeed on a DC 21 Constitution saving throw at the end of that turn or regurgitate all swallowed creatures, which fall prone in a space within 10 feet of the worm. If the worm dies, a swallowed creature is no longer restrained by it and can escape from the corpse by using 20 feet of movement, exiting prone.

Tail Stinger: Melee Weapon Attack: +9 to hit, reach 10 ft., one creature. Hit: 19 (3d6 + 9) piercing damage, and the target must make a DC 19 Constitution saving throw, taking 42 (12d6) poison damage on a failed save, or half as much damage on a successful one.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (45, 'Specter', 'undead', '1', 22, '5d8', 12, 'dex', 'walk 0 ft., fly 50 ft.', 1, 14, 11, 10, 10, 11, '', NULL, 'acid, cold, fire, lightning, thunder, bludgeoning, piercing, and slashing from nonmagical weapons', 'necrotic, poison', 'Charmed, Exhaustion, Grappled, Paralyzed, Petrified, Poisoned, Prone, Restrained, Unconscious', 'darkvision 60 ft., passive Perception 10', 'understands all languages it knew in life but can''t speak', 'Incorporeal Movement: The specter can move through other creatures and objects as if they were difficult terrain. It takes 5 (1d10) force damage if it ends its turn inside an object.

Sunlight Sensitivity: While in sunlight, the specter has disadvantage on attack rolls, as well as on Wisdom (Perception) checks that rely on sight.', 'Life Drain: Melee Spell Attack: +4 to hit, reach 5 ft., one creature. Hit: 10 (3d6) necrotic damage. The target must succeed on a DC 10 Constitution saving throw or its hit point maximum is reduced by an amount equal to the damage taken. This reduction lasts until the creature finishes a long rest. The target dies if this effect reduces its hit point maximum to 0.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (34, 'Ochre Jelly', 'ooze', '2', 45, '6d10', 8, 'dex', 'walk 10 ft., climb 10 ft.', 15, 6, 14, 2, 6, 1, '', NULL, 'acid', 'lightning, slashing', 'Blinded, Charmed, Blinded, Exhaustion, Frightened, Prone', 'blindsight 60 ft. (blind beyond this radius), passive Perception 8', '', 'Amorphous: The jelly can move through a space as narrow as 1 inch wide without squeezing.

Spider Climb: The jelly can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.', 'Pseudopod: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 9 (2d6 + 2) bludgeoning damage plus 3 (1d6) acid damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (35, 'Orc', 'humanoid', '1/2', 15, '2d8', 13, 'armor', 'walk 30 ft.', 16, 12, 16, 7, 11, 10, 'Saves: Intimidation+2', NULL, '', '', '', 'darkvision 60 ft., passive Perception 10', 'Common, Orc', 'Aggressive: As a bonus action, the orc can move up to its speed toward a hostile creature that it can see.', 'Greataxe: Melee Weapon Attack: +5 to hit, reach 5 ft., one target. Hit: 9 (1d12 + 3) slashing damage.

Javelin: Melee or Ranged Weapon Attack: +5 to hit, reach 5 ft. or range 30/120 ft., one target. Hit: 6 (1d6 + 3) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (40, 'Rust Monster', 'monstrosity', '1/2', 27, '5d8', 14, 'natural', 'walk 40 ft.', 13, 12, 13, 2, 13, 6, '', NULL, '', '', '', 'darkvision 60 ft., passive Perception 11', '', 'Iron Scent: The rust monster can pinpoint, by scent, the location of ferrous metal within 30 feet of it.

Rust Metal: Any nonmagical weapon made of metal that hits the rust monster corrodes. After dealing damage, the weapon takes a permanent and cumulative -1 penalty to damage rolls. If its penalty drops to -5, the weapon is destroyed. Nonmagical ammunition made of metal that hits the rust monster is destroyed after dealing damage.', 'Bite: Melee Weapon Attack: +3 to hit, reach 5 ft., one target. Hit: 5 (1d8 + 1) piercing damage.

Antennae: The rust monster corrodes a nonmagical ferrous metal object it can see within 5 feet of it. If the object isn''t being worn or carried, the touch destroys a 1-foot cube of it. If the object is being worn or carried by a creature, the creature can make a DC 11 Dexterity saving throw to avoid the rust monster''s touch.
If the object touched is either metal armor or a metal shield being worn or carried, its takes a permanent and cumulative -1 penalty to the AC it offers. Armor reduced to an AC of 10 or a shield that drops to a +0 bonus is destroyed. If the object touched is a held metal weapon, it rusts as described in the Rust Metal trait.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (44, 'Skeleton', 'undead', '1/4', 13, '2d8', 13, 'armor', 'walk 30 ft.', 10, 14, 15, 6, 8, 5, '', NULL, '', 'poison', 'Poisoned, Exhaustion', 'darkvision 60 ft., passive Perception 9', 'understands all languages it spoke in life but can''t speak', '', 'Shortsword: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d6 + 2) piercing damage.

Shortbow: Ranged Weapon Attack: +4 to hit, range 80/320 ft., one target. Hit: 5 (1d6 + 2) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (39, 'Roper', 'monstrosity', '5', 93, '11d10', 20, 'natural', 'walk 10 ft., climb 10 ft.', 18, 8, 17, 7, 16, 6, 'Saves: Stealth+5', NULL, '', '', '', 'darkvision 60 ft., passive Perception 16', '', 'False Appearance: While the roper remains motionless, it is indistinguishable from a normal cave formation, such as a stalagmite.

Grasping Tendrils: The roper can have up to six tendrils at a time. Each tendril can be attacked (AC 20; 10 hit points; immunity to poison and psychic damage). Destroying a tendril deals no damage to the roper, which can extrude a replacement tendril on its next turn. A tendril can also be broken if a creature takes an action and succeeds on a DC 15 Strength check against it.

Spider Climb: The roper can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.', 'Multiattack: The roper makes four attacks with its tendrils, uses Reel, and makes one attack with its bite.

Bite: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 22 (4d8 + 4) piercing damage.

Tendril: Melee Weapon Attack: +7 to hit, reach 50 ft., one creature. Hit: The target is grappled (escape DC 15). Until the grapple ends, the target is restrained and has disadvantage on Strength checks and Strength saving throws, and the roper can''t use the same tendril on another target.

Reel: The roper pulls each creature grappled by it up to 25 ft. straight toward it.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (42, 'Shadow', 'undead', '1/2', 16, '3d8', 12, 'dex', 'walk 40 ft.', 6, 14, 13, 6, 10, 8, 'Saves: Stealth+4', NULL, 'acid, cold, fire, lightning, thunder, bludgeoning, piercing, and slashing from nonmagical weapons', 'necrotic, poison', 'Exhaustion, Frightened, Grappled, Paralyzed, Petrified, Poisoned, Prone, Restrained', 'darkvision 60 ft., passive Perception 10', '', 'Amorphous: The shadow can move through a space as narrow as 1 inch wide without squeezing.

Shadow Stealth: While in dim light or darkness, the shadow can take the Hide action as a bonus action. Its stealth bonus is also improved to +6.

Sunlight Weakness: While in sunlight, the shadow has disadvantage on attack rolls, ability checks, and saving throws.', 'Strength Drain: Melee Weapon Attack: +4 to hit, reach 5 ft., one creature. Hit: 9 (2d6 + 2) necrotic damage, and the target''s Strength score is reduced by 1d4. The target dies if this reduces its Strength to 0. Otherwise, the reduction lasts until the target finishes a short or long rest.
If a non-evil humanoid dies from this attack, a new shadow rises from the corpse 1d4 hours later.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (43, 'Shrieker', 'plant', '0', 13, '3d8', 5, 'dex', 'walk 0 ft.', 1, 1, 10, 1, 3, 1, '', NULL, '', '', 'Blinded, Blinded, Frightened', 'blindsight 30 ft. (blind beyond this radius), passive Perception 6', '', 'False Appearance: While the shrieker remains motionless, it is indistinguishable from an ordinary fungus.', '', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (50, 'Violet Fungus', 'plant', '1/4', 18, '4d8', 5, 'dex', 'walk 5 ft.', 3, 1, 10, 1, 3, 1, '', NULL, '', '', 'Blinded, Blinded, Frightened', 'blindsight 30 ft. (blind beyond this radius), passive Perception 6', '', 'False Appearance: While the violet fungus remains motionless, it is indistinguishable from an ordinary fungus.', 'Multiattack: The fungus makes 1d4 Rotting Touch attacks.

Rotting Touch: Melee Weapon Attack: +2 to hit, reach 10 ft., one creature. Hit: 4 (1d8) necrotic damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (52, 'Wraith', 'undead', '5', 67, '9d8', 13, 'dex', 'walk 0 ft., fly 60 ft.', 6, 16, 16, 12, 14, 15, '', NULL, 'acid, cold, fire, lightning, thunder, bludgeoning, piercing, and slashing from nonmagical weapons that aren''t silvered', 'necrotic, poison', 'Charmed, Exhaustion, Grappled, Paralyzed, Petrified, Poisoned, Prone, Restrained', 'darkvision 60 ft., passive Perception 12', 'the languages it knew in life', 'Incorporeal Movement: The wraith can move through other creatures and objects as if they were difficult terrain. It takes 5 (1d10) force damage if it ends its turn inside an object.

Sunlight Sensitivity: While in sunlight, the wraith has disadvantage on attack rolls, as well as on Wisdom (Perception) checks that rely on sight.', 'Life Drain: Melee Weapon Attack: +6 to hit, reach 5 ft., one creature. Hit: 21 (4d8 + 3) necrotic damage. The target must succeed on a DC 14 Constitution saving throw or its hit point maximum is reduced by an amount equal to the damage taken. This reduction lasts until the target finishes a long rest. The target dies if this effect reduces its hit point maximum to 0.

Create Specter: The wraith targets a humanoid within 10 feet of it that has been dead for no longer than 1 minute and died violently. The target''s spirit rises as a specter in the space of its corpse or in the nearest unoccupied space. The specter is under the wraith''s control. The wraith can have no more than seven specters under its control at one time.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (53, 'Xorn', 'elemental', '5', 73, '7d8', 19, 'natural', 'walk 20 ft., burrow 20 ft.', 17, 10, 22, 11, 10, 11, 'Saves: Stealth+3', NULL, 'piercing and slashing from nonmagical weapons that aren''t adamantine', '', '', 'darkvision 60 ft., tremorsense 60 ft., passive Perception 16', 'Terran', 'Earth Glide: The xorn can burrow through nonmagical, unworked earth and stone. While doing so, the xorn doesn''t disturb the material it moves through.

Stone Camouflage: The xorn has advantage on Dexterity (Stealth) checks made to hide in rocky terrain.

Treasure Sense: The xorn can pinpoint, by scent, the location of precious metals and stones, such as coins and gems, within 60 ft. of it.', 'Multiattack: The xorn makes three claw attacks and one bite attack.

Bite: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 13 (3d6 + 3) piercing damage.

Claw: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 6 (1d6 + 3) slashing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (47, 'Steam Mephit', 'elemental', '1/4', 21, '6d6', 10, 'dex', 'walk 30 ft., fly 30 ft.', 5, 11, 10, 11, 10, 12, '', NULL, '', 'fire, poison', 'Poisoned', 'darkvision 60 ft., passive Perception 10', 'Aquan, Ignan', 'Death Burst: When the mephit dies, it explodes in a cloud of steam. Each creature within 5 ft. of the mephit must succeed on a DC 10 Dexterity saving throw or take 4 (1d8) fire damage.

Innate Spellcasting: The mephit can innately cast blur, requiring no material components. Its innate spellcasting ability is Charisma.', 'Claws: Melee Weapon Attack: +2 to hit, reach 5 ft., one creature. Hit: 2 (1d4) slashing damage plus 2 (1d4) fire damage.

Steam Breath: The mephit exhales a 15-foot cone of scalding steam. Each creature in that area must succeed on a DC 10 Dexterity saving throw, taking 4 (1d8) fire damage on a failed save, or half as much damage on a successful one.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (48, 'Stirge', 'beast', '1/8', 2, '1d4', 14, 'natural', 'walk 10 ft., fly 40 ft.', 4, 16, 11, 2, 8, 6, '', NULL, '', '', '', 'darkvision 60 ft., passive Perception 9', '', '', 'Blood Drain: Melee Weapon Attack: +5 to hit, reach 5 ft., one creature. Hit: 5 (1d4 + 3) piercing damage, and the stirge attaches to the target. While attached, the stirge doesn''t attack. Instead, at the start of each of the stirge''s turns, the target loses 5 (1d4 + 3) hit points due to blood loss.
The stirge can detach itself by spending 5 feet of its movement. It does so after it drains 10 hit points of blood from the target or the target dies. A creature, including the target, can use its action to detach the stirge.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (49, 'Troll', 'giant', '5', 84, '8d10', 15, 'natural', 'walk 30 ft.', 18, 13, 20, 7, 9, 7, 'Saves: Perception+2', NULL, '', '', '', 'darkvision 60 ft., passive Perception 12', 'Giant', 'Keen Smell: The troll has advantage on Wisdom (Perception) checks that rely on smell.

Regeneration: The troll regains 10 hit points at the start of its turn. If the troll takes acid or fire damage, this trait doesn''t function at the start of the troll''s next turn. The troll dies only if it starts its turn with 0 hit points and doesn''t regenerate.', 'Multiattack: The troll makes three attacks: one with its bite and two with its claws.

Bite: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 7 (1d6 + 4) piercing damage.

Claw: Melee Weapon Attack: +7 to hit, reach 5 ft., one target. Hit: 11 (2d6 + 4) slashing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (7, 'Cloaker', 'aberration', '8', 78, '12d10', 14, 'natural', 'walk 10 ft., fly 40 ft.', 17, 15, 12, 13, 12, 14, 'Saves: Stealth+5', NULL, '', '', '', 'darkvision 60 ft., passive Perception 11', 'Deep Speech, Undercommon', 'Damage Transfer: While attached to a creature, the cloaker takes only half the damage dealt to it (rounded down). and that creature takes the other half.

False Appearance: While the cloaker remains motionless without its underside exposed, it is indistinguishable from a dark leather cloak.

Light Sensitivity: While in bright light, the cloaker has disadvantage on attack rolls and Wisdom (Perception) checks that rely on sight.', 'Multiattack: The cloaker makes two attacks: one with its bite and one with its tail.

Bite: Melee Weapon Attack: +6 to hit, reach 5 ft., one creature. Hit: 10 (2d6 + 3) piercing damage, and if the target is Large or smaller, the cloaker attaches to it. If the cloaker has advantage against the target, the cloaker attaches to the target''s head, and the target is blinded and unable to breathe while the cloaker is attached. While attached, the cloaker can make this attack only against the target and has advantage on the attack roll. The cloaker can detach itself by spending 5 feet of its movement. A creature, including the target, can take its action to detach the cloaker by succeeding on a DC 16 Strength check.

Tail: Melee Weapon Attack: +6 to hit, reach 10 ft., one creature. Hit: 7 (1d8 + 3) slashing damage.

Moan: Each creature within 60 feet of the cloaker that can hear its moan and that isn''t an aberration must succeed on a DC 13 Wisdom saving throw or become frightened until the end of the cloaker''s next turn. If a creature''s saving throw is successful, the creature is immune to the cloaker''s moan for the next 24 hours.

Phantasms: The cloaker magically creates three illusory duplicates of itself if it isn''t in bright light. The duplicates move with it and mimic its actions, shifting position so as to make it impossible to track which cloaker is the real one. If the cloaker is ever in an area of bright light, the duplicates disappear.
Whenever any creature targets the cloaker with an attack or a harmful spell while a duplicate remains, that creature rolls randomly to determine whether it targets the cloaker or one of the duplicates. A creature is unaffected by this magical effect if it can''t see or if it relies on senses other than sight.
A duplicate has the cloaker''s AC and uses its saving throws. If an attack hits a duplicate, or if a duplicate fails a saving throw against an effect that deals damage, the duplicate disappears.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (9, 'Deep Gnome (Svirfneblin)', 'humanoid', '1/2', 16, '3d6', 15, 'armor', 'walk 20 ft.', 15, 14, 14, 12, 10, 9, 'Saves: Stealth+4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 12', 'Gnomish, Terran, Undercommon', 'Stone Camouflage: The gnome has advantage on Dexterity (Stealth) checks made to hide in rocky terrain.

Gnome Cunning: The gnome has advantage on Intelligence, Wisdom, and Charisma saving throws against magic.

Innate Spellcasting: The gnome''s innate spellcasting ability is Intelligence (spell save DC 11). It can innately cast the following spells, requiring no material components:
At will: nondetection (self only)
1/day each: blindness/deafness, blur, disguise self', 'War Pick: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 6 (1d8 + 2) piercing damage.

Poisoned Dart: Ranged Weapon Attack: +4 to hit, range 30/120 ft., one creature. Hit: 4 (1d4 + 2) piercing damage, and the target must succeed on a DC 12 Constitution saving throw or be poisoned for 1 minute. The target can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (12, 'Drider', 'monstrosity', '6', 123, '13d10', 19, 'natural', 'walk 30 ft., climb 30 ft.', 16, 16, 18, 13, 14, 12, 'Saves: Stealth+9', NULL, '', '', '', 'darkvision 120 ft., passive Perception 15', 'Elvish, Undercommon', 'Fey Ancestry: The drider has advantage on saving throws against being charmed, and magic can''t put the drider to sleep.

Innate Spellcasting: The drider''s innate spellcasting ability is Wisdom (spell save DC 13). The drider can innately cast the following spells, requiring no material components:
At will: dancing lights
1/day each: darkness, faerie fire

Spider Climb: The drider can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.

Sunlight Sensitivity: While in sunlight, the drider has disadvantage on attack rolls, as well as on Wisdom (Perception) checks that rely on sight.

Web Walker: The drider ignores movement restrictions caused by webbing.', 'Multiattack: The drider makes three attacks, either with its longsword or its longbow. It can replace one of those attacks with a bite attack.

Bite: Melee Weapon Attack: +6 to hit, reach 5 ft., one creature. Hit: 2 (1d4) piercing damage plus 9 (2d8) poison damage.

Longsword: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 7 (1d8 + 3) slashing damage, or 8 (1d10 + 3) slashing damage if used with two hands.

Longbow: Ranged Weapon Attack: +6 to hit, range 150/600 ft., one target. Hit: 7 (1d8 + 3) piercing damage plus 4 (1d8) poison damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (18, 'Gelatinous Cube', 'ooze', '2', 84, '8d10', 6, 'dex', 'walk 15 ft.', 14, 3, 20, 1, 6, 1, '', NULL, '', '', 'Blinded, Charmed, Deafened, Exhaustion, Frightened, Prone', 'blindsight 60 ft. (blind beyond this radius), passive Perception 8', '', 'Ooze Cube: The cube takes up its entire space. Other creatures can enter the space, but a creature that does so is subjected to the cube''s Engulf and has disadvantage on the saving throw.
Creatures inside the cube can be seen but have total cover.
A creature within 5 feet of the cube can take an action to pull a creature or object out of the cube. Doing so requires a successful DC 12 Strength check, and the creature making the attempt takes 10 (3d6) acid damage.
The cube can hold only one Large creature or up to four Medium or smaller creatures inside it at a time.

Transparent: Even when the cube is in plain sight, it takes a successful DC 15 Wisdom (Perception) check to spot a cube that has neither moved nor attacked. A creature that tries to enter the cube''s space while unaware of the cube is surprised by the cube.', 'Pseudopod: Melee Weapon Attack: +4 to hit, reach 5 ft., one creature. Hit: 10 (3d6) acid damage.

Engulf: The cube moves up to its speed. While doing so, it can enter Large or smaller creatures'' spaces. Whenever the cube enters a creature''s space, the creature must make a DC 12 Dexterity saving throw.
On a successful save, the creature can choose to be pushed 5 feet back or to the side of the cube. A creature that chooses not to be pushed suffers the consequences of a failed saving throw.
On a failed save, the cube enters the creature''s space, and the creature takes 10 (3d6) acid damage and is engulfed. The engulfed creature can''t breathe, is restrained, and takes 21 (6d6) acid damage at the start of each of the cube''s turns. When the cube moves, the engulfed creature moves with it.
An engulfed creature can try to escape by taking an action to make a DC 12 Strength check. On a success, the creature escapes and enters a space of its choice within 5 feet of the cube.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (25, 'Gray Ooze', 'ooze', '1/2', 22, '3d8', 8, 'dex', 'walk 10 ft., climb 10 ft.', 12, 6, 16, 1, 6, 2, 'Saves: Stealth+2', NULL, 'acid, cold, fire', '', 'Blinded, Charmed, Deafened, Exhaustion, Frightened, Prone', 'blindsight 60 ft. (blind beyond this radius), passive Perception 8', '', 'Amorphous: The ooze can move through a space as narrow as 1 inch wide without squeezing.

Corrode Metal: Any nonmagical weapon made of metal that hits the ooze corrodes. After dealing damage, the weapon takes a permanent and cumulative -1 penalty to damage rolls. If its penalty drops to -5, the weapon is destroyed. Nonmagical ammunition made of metal that hits the ooze is destroyed after dealing damage.
The ooze can eat through 2-inch-thick, nonmagical metal in 1 round.

False Appearance: While the ooze remains motionless, it is indistinguishable from an oily pool or wet rock.', 'Pseudopod: Melee Weapon Attack: +3 to hit, reach 5 ft., one target. Hit: 4 (1d6 + 1) bludgeoning damage plus 7 (2d6) acid damage, and if the target is wearing nonmagical metal armor, its armor is partly corroded and takes a permanent and cumulative -1 penalty to the AC it offers. The armor is destroyed if the penalty reduces its AC to 10.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (32, 'Minotaur', 'monstrosity', '3', 76, '9d10', 14, 'natural', 'walk 40 ft.', 18, 11, 16, 6, 16, 9, 'Saves: Perception+7', NULL, '', '', '', 'darkvision 60 ft., passive Perception 17', 'Abyssal', 'Charge: If the minotaur moves at least 10 ft. straight toward a target and then hits it with a gore attack on the same turn, the target takes an extra 9 (2d8) piercing damage. If the target is a creature, it must succeed on a DC 14 Strength saving throw or be pushed up to 10 ft. away and knocked prone.

Labyrinthine Recall: The minotaur can perfectly recall any path it has traveled.

Reckless: At the start of its turn, the minotaur can gain advantage on all melee weapon attack rolls it makes during that turn, but attack rolls against it have advantage until the start of its next turn.', 'Greataxe: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 17 (2d12 + 4) slashing damage.

Gore: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 13 (2d8 + 4) piercing damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (33, 'Nalfeshnee', 'fiend', '13', 184, '16d10', 18, 'natural', 'walk 20 ft., fly 30 ft.', 21, 10, 22, 19, 12, 15, 'Saves: CON+11, INT+9, WIS+6, CHA+7', NULL, 'cold, fire, lightning, bludgeoning, piercing, and slashing from nonmagical weapons', 'poison', 'Poisoned', 'truesight 120 ft., passive Perception 11', 'Abyssal, telepathy 120 ft.', 'Magic Resistance: The nalfeshnee has advantage on saving throws against spells and other magical effects.', 'Multiattack: The nalfeshnee uses Horror Nimbus if it can. It then makes three attacks: one with its bite and two with its claws.

Bite: Melee Weapon Attack: +10 to hit, reach 5 ft., one target. Hit: 32 (5d10 + 5) piercing damage.

Claw: Melee Weapon Attack: +10 to hit, reach 10 ft., one target. Hit: 15 (3d6 + 5) slashing damage.

Horror Nimbus: The nalfeshnee magically emits scintillating, multicolored light. Each creature within 15 feet of the nalfeshnee that can see the light must succeed on a DC 15 Wisdom saving throw or be frightened for 1 minute. A creature can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success. If a creature''s saving throw is successful or the effect ends for it, the creature is immune to the nalfeshnee''s Horror Nimbus for the next 24 hours.

Teleport: The nalfeshnee magically teleports, along with any equipment it is wearing or carrying, up to 120 feet to an unoccupied space it can see.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (116, 'Gas Spore', 'plant', '1/2', 1, '1', 5, '', '0 ft., fly 10 ft. (hover)', 5, 1, 10, 1, 1, 1, '', NULL, '', 'poison', 'blinded, deafened, frightened, paralyzed, poisoned, prone', 'blindsight 30 ft. (blind beyond this radius), passive Perception 6', '—', 'Death Burst: When the gas spore dies, it explodes. Each creature within 20 ft. must succeed on a DC 15 Constitution save or take 26 (4d12) poison damage and become infected with the Flesh Rot disease. On a success, the creature takes half damage and isn''t infected.

Eerie Resemblance: The gas spore resembles a beholder. A creature that sees it and succeeds on a DC 10 Intelligence (Nature) check knows it is a gas spore.

Spore Infector (Disease): Until cured, the infected creature can''t regain hit points. After dying from this disease, the creature''s body produces a gas spore.', 'Touch: +0 to hit, reach 5 ft. The target must succeed on a DC 10 Constitution saving throw or be infected with the Spore Infector disease.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (38, 'Quasit', 'fiend', '1', 7, '3d4', 13, 'dex', 'walk 40 ft.', 5, 17, 10, 7, 10, 10, 'Saves: Stealth+5', NULL, 'cold, fire, lightning, bludgeoning, piercing, and slashing from nonmagical weapons', 'poison', 'Poisoned', 'darkvision 120 ft., passive Perception 10', 'Abyssal, Common', 'Shapechanger: The quasit can use its action to polymorph into a beast form that resembles a bat (speed 10 ft. fly 40 ft.), a centipede (40 ft., climb 40 ft.), or a toad (40 ft., swim 40 ft.), or back into its true form . Its statistics are the same in each form, except for the speed changes noted. Any equipment it is wearing or carrying isn''t transformed . It reverts to its true form if it dies.

Magic Resistance: The quasit has advantage on saving throws against spells and other magical effects.', 'Claw (Bite in Beast Form): Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 5 (1d4 + 3) piercing damage, and the target must succeed on a DC 10 Constitution saving throw or take 5 (2d4) poison damage and become poisoned for 1 minute. The target can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success.

Scare: One creature of the quasit''s choice within 20 ft. of it must succeed on a DC 10 Wisdom saving throw or be frightened for 1 minute. The target can repeat the saving throw at the end of each of its turns, with disadvantage if the quasit is within line of sight, ending the effect on itself on a success.

Invisibility: The quasit magically turns invisible until it attacks or uses Scare, or until its concentration ends (as if concentrating on a spell). Any equipment the quasit wears or carries is invisible with it.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (41, 'Salamander', 'elemental', '5', 90, '12d10', 15, 'natural', 'walk 30 ft.', 18, 14, 15, 11, 10, 12, '', NULL, 'bludgeoning, piercing, and slashing from nonmagical weapons', 'fire', '', 'darkvision 60 ft., passive Perception 10', 'Ignan', 'Heated Body: A creature that touches the salamander or hits it with a melee attack while within 5 ft. of it takes 7 (2d6) fire damage.

Heated Weapons: Any metal melee weapon the salamander wields deals an extra 3 (1d6) fire damage on a hit (included in the attack).', 'Multiattack: The salamander makes two attacks: one with its spear and one with its tail.

Spear: Melee or Ranged Weapon Attack: +7 to hit, reach 5 ft. or range 20/60 ft., one target. Hit: 11 (2d6 + 4) piercing damage, or 13 (2d8 + 4) piercing damage if used with two hands to make a melee attack, plus 3 (1d6) fire damage.

Tail: Melee Weapon Attack: +7 to hit, reach 10 ft., one target. Hit: 11 (2d6 + 4) bludgeoning damage plus 7 (2d6) fire damage, and the target is grappled (escape DC 14). Until this grapple ends, the target is restrained, the salamander can automatically hit the target with its tail, and the salamander can''t make tail attacks against other targets.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (46, 'Spirit Naga', 'monstrosity', '8', 75, '10d10', 15, 'natural', 'walk 40 ft.', 18, 17, 14, 16, 15, 16, 'Saves: DEX+6, CON+5, WIS+5, CHA+6', NULL, '', 'poison', 'Charmed, Poisoned', 'darkvision 60 ft., passive Perception 12', 'Abyssal, Common', 'Rejuvenation: If it dies, the naga returns to life in 1d6 days and regains all its hit points. Only a wish spell can prevent this trait from functioning.

Spellcasting: The naga is a 10th-level spellcaster. Its spellcasting ability is Intelligence (spell save DC 14, +6 to hit with spell attacks), and it needs only verbal components to cast its spells. It has the following wizard spells prepared:

- Cantrips (at will): mage hand, minor illusion, ray of frost
- 1st level (4 slots): charm person, detect magic, sleep
- 2nd level (3 slots): detect thoughts, hold person
- 3rd level (3 slots): lightning bolt, water breathing
- 4th level (3 slots): blight, dimension door
- 5th level (2 slots): dominate person', 'Bite: Melee Weapon Attack: +7 to hit, reach 10 ft., one creature. Hit: 7 (1d6 + 4) piercing damage, and the target must make a DC 13 Constitution saving throw, taking 31 (7d8) poison damage on a failed save, or half as much damage on a successful one.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (51, 'Vrock', 'fiend', '6', 104, '11d10', 15, 'natural', 'walk 40 ft., fly 60 ft.', 17, 15, 18, 8, 13, 8, 'Saves: DEX+5, WIS+4, CHA+2', NULL, 'cold, fire, lightning, bludgeoning, piercing, and slashing from nonmagical weapons', 'poison', 'Poisoned', 'darkvision 120 ft., passive Perception 11', 'Abyssal, telepathy 120 ft.', 'Magic Resistance: The vrock has advantage on saving throws against spells and other magical effects.', 'Multiattack: The vrock makes two attacks: one with its beak and one with its talons.

Beak: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 10 (2d6 + 3) piercing damage.

Talons: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 14 (2d10 + 3) slashing damage.

Spores: A 15-foot-radius cloud of toxic spores extends out from the vrock. The spores spread around corners. Each creature in that area must succeed on a DC 14 Constitution saving throw or become poisoned. While poisoned in this way, a target takes 5 (1d10) poison damage at the start of each of its turns. A target can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success. Emptying a vial of holy water on the target also ends the effect on it.

Stunning Screech: The vrock emits a horrific screech. Each creature within 20 feet of it that can hear it and that isn''t a demon must succeed on a DC 14 Constitution saving throw or be stunned until the end of the vrock''s next turn .', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (107, 'Barlgura', 'fiend (demon)', '5', 68, '8d10+24', 15, 'natural armor', '30 ft., climb 30 ft.', 18, 17, 16, 7, 14, 9, 'Saves: STR +7, DEX +6, CON +6', NULL, '', 'poison', 'poisoned', 'blindsight 30 ft., darkvision 120 ft., passive Perception 12', 'Abyssal, telepathy 120 ft.', 'Innate Spellcasting: (Wis, save DC 12) 2/day: disguise self, mirror image; 1/day: invisibility (self only).

Reckless: At the start of its turn the barlgura can gain advantage on all melee attacks until end of turn, but attack rolls against it have advantage until next turn.

Running Leap: With a 10 ft. running start, the barlgura can long jump up to 40 ft.', 'Multiattack: One bite and two fist attacks.

Bite: +7 to hit, reach 5 ft. Hit: 11 (2d6+4) piercing.

Fist: +7 to hit, reach 5 ft. Hit: 9 (2d4+4) bludgeoning.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (109, 'Chasme', 'fiend (demon)', '6', 84, '13d10+13', 15, 'natural armor', '20 ft., fly 60 ft.', 15, 15, 12, 11, 14, 10, 'Saves: DEX +5, WIS +5 | Skills: Perception +5', NULL, 'cold, fire, lightning; bludgeoning, piercing, slashing from nonmagical attacks', 'poison', 'poisoned', 'blindsight 10 ft., darkvision 120 ft., passive Perception 15', 'Abyssal, telepathy 120 ft.', 'Drone: The chasme produces a horrid droning. Any creature within 30 ft. (except demons) must succeed on a DC 12 Constitution saving throw at the start of each of its turns or fall unconscious for 10 minutes. A creature woken by damage or another creature''s action is immune for 1 hour.

Spider Climb: Can climb difficult surfaces including ceilings without an ability check.', 'Multiattack: One proboscis attack and two claw attacks.

Proboscis: +5 to hit, reach 5 ft. Hit: 16 (4d6+2) piercing. Target must succeed on a DC 13 Constitution save or fall unconscious for 10 minutes (woken by damage or action).

Claws: +5 to hit, reach 5 ft. Hit: 7 (2d4+2) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (110, 'Death Tyrant', 'undead', '14', 187, '22d10+66', 19, 'natural armor', '0 ft., fly 20 ft. (hover)', 10, 14, 16, 19, 15, 19, 'Saves: CON +8, INT +9, WIS +7, CHA +9 | Skills: Perception +12', NULL, '', 'poison', 'charmed, exhaustion, paralyzed, petrified, poisoned, prone', 'darkvision 120 ft., passive Perception 22', 'Deep Speech, Undercommon', 'Negative Energy Cone: The death tyrant''s central eye emits a 150-ft. cone of negative energy. At the start of each of its turns, the tyrant decides which way the cone faces and whether the feature is active. Creatures that die in the cone rise as zombies under the tyrant''s control (max 9 zombies at once).', 'Bite: +5 to hit, reach 5 ft. Hit: 14 (4d6) piercing.

Eye Rays: The death tyrant shoots three of the following magical eye rays at random (reroll duplicates), choosing 1–3 targets within 120 ft. DC 17 save. Rays 1–10 same as Beholder.', 'The death tyrant can take 3 legendary actions (Eye Ray only).

Eye Ray: The death tyrant uses one random eye ray.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (111, 'Drow Elite Warrior', 'humanoid (elf)', '5', 71, '11d8+22', 18, 'studded leather, shield', '30 ft.', 13, 18, 14, 11, 13, 12, 'Saves: DEX +6, CON +4, WIS +3 | Skills: Perception +3, Stealth +6', NULL, '', '', '', 'darkvision 120 ft., passive Perception 13', 'Elvish, Undercommon', 'Fey Ancestry: Advantage on saves against being charmed; magic can''t put the drow to sleep.

Innate Spellcasting: (Cha, save DC 12) At will: dancing lights; 1/day each: darkness, faerie fire, levitate (self only).

Sunlight Sensitivity: Disadvantage on attack rolls and Wisdom (Perception) checks in sunlight.', 'Multiattack: Two shortsword attacks.

Shortsword: +6 to hit, reach 5 ft. Hit: 7 (1d6+4) piercing plus 10 (3d6) poison.

Hand Crossbow: +6 to hit, range 30/120 ft. Hit: 7 (1d6+4) piercing plus 10 (3d6) poison.', 'Parry (Reaction): The drow adds 3 to its AC against one melee attack that would hit it, if it can see the attacker and is wielding a melee weapon.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (112, 'Drow Mage', 'humanoid (elf)', '7', 45, '10d8', 12, '15 with mage armor', '30 ft.', 9, 14, 10, 17, 13, 12, 'Saves: INT +6, WIS +4 | Skills: Arcana +6, Deception +4, Perception +4, Stealth +5', NULL, '', '', '', 'darkvision 120 ft., passive Perception 14', 'Elvish, Undercommon', 'Fey Ancestry: Advantage on saves vs. charmed; magic can''t put to sleep.

Innate Spellcasting: (Cha, save DC 12) At will: dancing lights; 1/day each: darkness, faerie fire, levitate (self only).

Spellcasting: (INT, spell save DC 14, +6 to hit) Cantrips: mage hand, minor illusion, poison spray, ray of frost. 1st (4 slots): mage armor, magic missile, shield. 2nd (3): misty step, web. 3rd (3): fly, lightning bolt. 4th (3): Evard''s black tentacles, greater invisibility. 5th (1): cloudkill.

Sunlight Sensitivity: Disadvantage on attacks and Perception checks in sunlight.', 'Staff: +2 to hit, reach 5 ft. Hit: 2 (1d6-1) bludgeoning, or 3 (1d8-1) two-handed.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (113, 'Drow Priestess of Lolth', 'humanoid (elf)', '8', 71, '13d8+13', 16, 'scale mail', '30 ft.', 10, 14, 12, 13, 17, 18, 'Saves: CON +4, WIS +6, CHA +7 | Skills: Insight +6, Perception +6, Religion +4, Stealth +5', NULL, '', '', '', 'darkvision 120 ft., passive Perception 16', 'Elvish, Undercommon', 'Fey Ancestry: Advantage on saves vs. charmed; magic can''t put to sleep.

Innate Spellcasting: (Cha, save DC 15) At will: dancing lights; 1/day each: darkness, faerie fire, levitate (self only).

Spellcasting: (WIS, spell save DC 14, +6 to hit) Cantrips: guidance, poison spray, resistance, spare the dying, thaumaturgy. 1st (4): animal friendship, cure wounds, detect poison/disease, ray of sickness. 2nd (3): lesser restoration, protection from poison, spiritual weapon. 3rd (3): dispel magic, mass healing word, meld into stone. 4th (3): divine smite (as a cleric), freedom of movement, guardian of faith. 5th (2): contagion, insect plague. 6th (1): harm.

Sunlight Sensitivity: Disadvantage on attacks and Perception checks in sunlight.', 'Multiattack: Two scourge attacks.

Scourge: +5 to hit, reach 5 ft. Hit: 5 (1d6+2) piercing plus 17 (5d6) poison.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (114, 'Fire Snake', 'elemental', '1', 22, '5d8', 14, 'natural armor', '30 ft.', 12, 14, 11, 7, 10, 8, '', NULL, '', 'fire, poison', 'exhaustion, grappled, paralyzed, petrified, poisoned, prone, restrained, unconscious', 'darkvision 60 ft., passive Perception 10', 'Ignan', 'Heated Body: A creature that touches the fire snake or hits it with a melee attack takes 3 (1d6) fire damage.', 'Bite: +3 to hit, reach 5 ft. Hit: 4 (1d4+2) piercing plus 3 (1d6) fire damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (121, 'Ixitxachitl', 'aberration (chaotic evil)', '1/4', 18, '4d6+4', 16, 'natural armor', '0 ft., swim 30 ft.', 11, 14, 12, 10, 12, 12, 'Skills: Religion +2', NULL, '', 'poison', 'poisoned', 'darkvision 60 ft., passive Perception 11', 'Abyssal, telepathy 30 ft.', 'Spellcasting (Cleric, Wis, save DC 11, +3 to hit): Cantrips: guidance, sacred flame, thaumaturgy. 1st (2 slots): bless, cure wounds, detect magic.

Water Breathing: Can breathe only underwater.', 'Barbed Tail: +4 to hit, reach 5 ft. Hit: 4 (1d4+2) piercing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (122, 'Kuo-toa', 'humanoid (kuo-toa)', '1/4', 18, '4d8', 13, 'natural armor and shield', '30 ft., swim 30 ft.', 13, 10, 11, 11, 10, 8, 'Skills: Perception +4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 14', 'Undercommon', 'Amphibious: Can breathe air and water.

Otherworldly Perception: Can sense the presence of any creature within 30 ft. that is invisible or on the Ethereal Plane; such creatures can be detected by sound, movement, smell, etc.

Slippery: Advantage on ability checks and saving throws made to escape a grapple.

Sunlight Sensitivity: Disadvantage on attack rolls and Wisdom (Perception) checks in sunlight.', 'Multiattack: Two attacks: one with its bite and one with its spear, or one with its bite and one with its pincer staff.

Bite: +3 to hit, reach 5 ft. Hit: 3 (1d4+1) piercing.

Spear: +3 to hit, reach 5 ft. or range 20/60 ft. Hit: 4 (1d6+1) piercing or 5 (1d8+1) two-handed.

Net: Range 5/15 ft. The target is restrained (DC 10 Str escape) and can''t move further away. Net has AC 10, 5 hp, and immunity to bludgeoning/poison/psychic damage.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (123, 'Kuo-toa Archpriest', 'humanoid (kuo-toa)', '6', 97, '13d8+39', 13, 'natural armor', '30 ft., swim 30 ft.', 16, 14, 16, 13, 16, 14, 'Saves: CON +6, WIS +6, CHA +5 | Skills: Perception +9, Religion +4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 19', 'Undercommon', 'Amphibious: Can breathe air and water.

Otherworldly Perception: Can sense invisible/ethereal creatures within 30 ft.

Slippery: Advantage on checks/saves to escape grapples.

Spellcasting (Cleric, Wis, save DC 14, +6 to hit): Cantrips: guidance, sacred flame, thaumaturgy. 1st (4): detect magic, sanctuary, shield of faith. 2nd (3): hold person, spiritual weapon. 3rd (3): spirit guardians, tongues. 4th (3): control water, divination. 5th (2): mass cure wounds, scrying.

Sunlight Sensitivity: Disadvantage on attack rolls and Perception checks in sunlight.', 'Multiattack: Two scepter attacks.

Scepter: +6 to hit, reach 5 ft. Hit: 6 (1d6+3) bludgeoning plus 9 (2d8) lightning.

Unarmed Strike: +6 to hit, reach 5 ft. Hit: 4 (1d4+3) bludgeoning.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (125, 'Manes', 'fiend (demon)', '1/8', 9, '2d8', 9, '', '20 ft.', 10, 9, 10, 3, 8, 4, '', NULL, 'cold, fire, lightning', 'poison', 'charmed, frightened, poisoned', 'darkvision 60 ft., passive Perception 9', 'understands Abyssal but can''t speak', '', 'Claws: +2 to hit, reach 5 ft. Hit: 5 (2d4) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (126, 'Mind Flayer', 'aberration (lawful evil)', '7', 71, '13d8+13', 15, 'breastplate', '30 ft.', 11, 12, 12, 19, 17, 17, 'Saves: INT +7, WIS +6, CHA +6 | Skills: Arcana +7, Deception +6, Insight +6, Perception +6, Persuasion +6, Stealth +4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 16', 'Deep Speech, Undercommon, telepathy 120 ft.', 'Magic Resistance: Advantage on saving throws against spells and other magical effects.

Innate Spellcasting (INT, save DC 15): At will: detect thoughts, levitate. 1/day each: dominate monster, plane shift (self only).

Spellcasting (INT, save DC 15, +7 to hit): Cantrips: blade ward, dancing lights, mage hand, shocking grasp. 1st (4): detect magic, disguise self, shield, sleep. 2nd (3): blur, invisibility, ray of enfeeblement. 3rd (3): clairvoyance, lightning bolt, sending. 4th (3): confusion, hallucinatory terrain. 5th (2): telekinesis.', 'Tentacles: +7 to hit, reach 5 ft. Hit: 15 (2d10+4) psychic. Target must succeed on DC 15 Intelligence save or be stunned until end of next turn.

Extract Brain: +7 to hit, reach 5 ft. One grappled/incapacitated Medium or smaller humanoid. The target takes 55 (10d10) piercing damage. If reduced to 0, brain is extracted and devoured; target dies.

Mind Blast (Recharge 5–6): 60-ft. cone, DC 15 Intelligence save. Fail: 22 (4d8+4) psychic and stunned for 1 min. Stunned creature repeats save at end of each turn.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (127, 'Mud Mephit', 'elemental', '1/4', 27, '6d8', 11, '', '20 ft., fly 20 ft., swim 20 ft.', 8, 12, 11, 9, 7, 12, 'Skills: Stealth +3', NULL, '', 'poison', 'poisoned', 'darkvision 60 ft., passive Perception 8', 'Aquan, Terran', 'Death Burst: When the mephit dies, it explodes in a burst of sticky mud. Each Medium or smaller creature within 5 ft. must succeed on a DC 11 Strength saving throw or become restrained for 1 minute. A restrained creature can repeat the save at the end of each turn.

False Appearance: While motionless, the mephit is indistinguishable from an ordinary pool of mud.

Innate Spellcasting: (Cha, save DC 11) 1/day: ray of enfeeblement.', 'Fists: +3 to hit, reach 5 ft. Hit: 4 (1d6+1) bludgeoning.

Mud Breath (Recharge 6): The mephit exhales a 15-ft. cone of mud. Creatures must succeed on a DC 11 Dexterity save or be restrained for 1 minute.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (129, 'Myconid Sovereign', 'plant (lawful neutral)', '2', 60, '8d10+16', 13, 'natural armor', '30 ft.', 12, 10, 14, 13, 15, 10, '', NULL, '', '', '', 'darkvision 120 ft., passive Perception 12', '—', 'Distress Spores: When damaged, all myconids within 240 ft. are aware.

Sun Sickness: Disadvantage on checks, attack rolls, and saves in sunlight; dies after 1 hour.

Rapport Spores: 30-ft.-radius cloud. Affected creatures telepathically communicate for 1 hour.', 'Multiattack: Two fist attacks.

Fist: +3 to hit, reach 5 ft. Hit: 8 (2d6+1) bludgeoning plus 7 (2d6) poison.

Animating Spores (3/Day): Targets one Medium or smaller corpse within 5 ft. In 24 hours it rises as a spore servant under the sovereign''s control (as long as the sovereign is alive, max 10 servants).

Pacifying Spores: One creature within 10 ft. makes a DC 12 Constitution save or is stunned for 1 minute.

Rapport Spores: 120-ft. radius burst. Creatures communicate telepathically for 1 hour.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (130, 'Myconid Sprout', 'plant (lawful neutral)', '0', 7, '2d6', 10, '', '10 ft.', 8, 10, 10, 8, 11, 5, '', NULL, '', '', '', 'darkvision 120 ft., passive Perception 10', '—', 'Distress Spores: When damaged, all myconids within 240 ft. are aware.

Sun Sickness: Disadvantage on checks, attack rolls, and saves in sunlight; dies after 1 hour.', 'Fist: +1 to hit, reach 5 ft. Hit: 1 (1d4-1) bludgeoning plus 2 (1d4) poison.

Rapport Spores (1/Day): 10-ft. radius. Affected creatures communicate telepathically with each other and nearby myconids for 1 hour.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (131, 'Nothic', 'aberration (neutral evil)', '2', 45, '6d8+18', 15, 'natural armor', '30 ft.', 14, 16, 16, 13, 10, 8, 'Skills: Arcana +3, Insight +4, Perception +2, Stealth +5', NULL, '', '', '', 'truesight 120 ft., passive Perception 12', 'Undercommon', 'Keen Sight: Advantage on Wisdom (Perception) checks that rely on sight.', 'Multiattack: Two claw attacks.

Claw: +4 to hit, reach 5 ft. Hit: 6 (1d6+3) slashing.

Rotting Gaze: One creature within 30 ft. must make a DC 12 Constitution saving throw. On a fail: 10 (3d6) necrotic damage; on success: half.

Weird Insight: The nothic targets one creature it can see within 30 ft. The target makes a Dexterity (Deception) check contested by the nothic''s Wisdom (Insight). If the nothic wins, it magically learns one fact or secret about the target.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (132, 'Orc Eye of Gruumsh', 'humanoid (orc)', '2', 45, '6d8+18', 16, 'chain mail', '30 ft.', 16, 12, 16, 9, 13, 12, 'Saves: STR +5, CON +5, WIS +3 | Skills: Intimidation +3, Religion +1', NULL, '', '', '', 'darkvision 60 ft., passive Perception 11', 'Common, Orc', 'Aggressive: As a bonus action, the orc can move up to its speed toward a hostile creature that it can see.

Gruumsh''s Fury: The orc deals an extra 4 (1d8) damage when it hits with a weapon attack.

Spellcasting (Wis, save DC 11, +3 to hit): Cantrips: guidance, thaumaturgy. 1st (3): bless, command, cure wounds. 2nd (2): augury, spiritual weapon.', 'Multiattack: Two attacks.

Spear: +5 to hit, reach 5 ft. or range 20/60 ft. Hit: 6 (1d6+3) piercing, or 7 (1d8+3) two-handed.

Hand Axe: +5 to hit, reach 5 ft. or range 20/60 ft. Hit: 6 (1d6+3) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (133, 'Orog', 'humanoid (orc)', '2', 42, '5d8+20', 18, 'plate', '30 ft.', 18, 12, 18, 12, 11, 12, 'Saves: STR +6, CON +6 | Skills: Intimidation +5, Survival +2', NULL, '', '', '', 'darkvision 60 ft., passive Perception 10', 'Common, Orc', 'Aggressive: As a bonus action, the orog can move up to its speed toward a hostile creature it can see.', 'Multiattack: Two attacks.

Greataxe: +6 to hit, reach 5 ft. Hit: 10 (1d12+4) slashing.

Javelin: +6 to hit, reach 5 ft. or range 30/120 ft. Hit: 7 (1d6+4) piercing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (134, 'Piercer', 'monstrosity', '1/2', 22, '3d8+9', 15, 'natural armor', '5 ft., climb 10 ft.', 10, 13, 16, 1, 7, 3, 'Skills: Stealth +3', NULL, '', '', '', 'blindsight 30 ft., darkvision 60 ft., passive Perception 8', '—', 'False Appearance: While motionless on a ceiling, the piercer is indistinguishable from a normal stalactite.

Spider Climb: Can climb difficult surfaces without an ability check.', 'Drop: One creature directly below the piercer. The target must make a DC 11 Dexterity saving throw. On a fail: takes 3 (1d6) piercing damage per 10 ft. fallen (max 21/6d6) and falls prone. On success: half damage, no prone.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (135, 'Quaggoth', 'humanoid (quaggoth)', '2', 45, '6d8+18', 13, 'natural armor', '30 ft., climb 30 ft.', 17, 12, 16, 6, 12, 7, 'Skills: Athletics +5', NULL, '', 'poison', 'frightened, poisoned', 'darkvision 120 ft., passive Perception 11', 'Undercommon', 'Wounded Fury: While at 10 or fewer hit points, the quaggoth has advantage on attack rolls and deals 7 (2d6) extra damage on each hit.', 'Multiattack: Two claw attacks.

Claws: +5 to hit, reach 5 ft. Hit: 6 (1d6+3) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (136, 'Quaggoth Spore Servant', 'plant (unaligned)', '1', 45, '6d8+18', 13, 'natural armor', '30 ft., climb 30 ft.', 17, 12, 16, 2, 6, 1, '', NULL, '', 'poison', 'blinded, deafened, exhaustion, frightened, paralyzed, poisoned', 'blindsight 30 ft. (blind beyond), passive Perception 8', '—', 'Spore Servant: The servant obeys the myconid sovereign that created it. It has no memory of its former life.', 'Multiattack: Two claw attacks.

Claws: +5 to hit, reach 5 ft. Hit: 6 (1d6+3) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (137, 'Shadow Demon', 'fiend (demon)', '4', 66, '12d8+12', 13, '', '30 ft., fly 30 ft.', 1, 17, 12, 14, 13, 14, 'Saves: DEX +5, CHA +4 | Skills: Stealth +7', NULL, 'acid, fire, necrotic, thunder; bludgeoning, piercing, slashing from nonmagical attacks', 'cold, lightning, poison', 'exhaustion, grappled, paralyzed, petrified, poisoned, prone, restrained', 'darkvision 120 ft., passive Perception 11', 'Abyssal, telepathy 120 ft.', 'Incorporeal Movement: Can move through other creatures and objects as difficult terrain. Takes 5 (1d10) force damage if ending turn inside an object.

Light Sensitivity: Disadvantage on attack rolls and Perception checks in bright light.

Shadow Stealth: While in dim light or darkness, can Hide as a bonus action.

Vulnerability to Radiant: The shadow demon is vulnerable to radiant damage.', 'Claws: +5 to hit, reach 5 ft. Hit: 10 (2d6+3) cold. If in dim light or darkness: 17 (4d6+3) cold instead. Target must also succeed on a DC 13 Strength saving throw or have its Strength reduced by 1d4. The reduction lasts until the target finishes a short or long rest. If reduced to 0 Strength, the target dies.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (139, 'Troglodyte', 'humanoid (troglodyte)', '1/4', 13, '2d8+4', 11, 'natural armor', '30 ft.', 14, 10, 14, 6, 10, 6, 'Skills: Stealth +2', NULL, '', '', '', 'darkvision 60 ft., passive Perception 10', 'Troglodyte', 'Chameleon Skin: The troglodyte has advantage on Dexterity (Stealth) checks.

Stench: Any creature that starts its turn within 5 ft. must succeed on a DC 12 Constitution save or be poisoned until the start of its next turn. On success, the creature is immune to the Stench for 1 hour.

Sunlight Sensitivity: Disadvantage on attack rolls and Perception checks in sunlight.', 'Multiattack: One bite and two claw attacks.

Bite: +4 to hit, reach 5 ft. Hit: 4 (1d4+2) piercing.

Claw: +4 to hit, reach 5 ft. Hit: 4 (1d4+2) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (141, 'Water Weird', 'elemental (neutral)', '3', 58, '9d10+9', 13, '', '0 ft., swim 60 ft.', 17, 16, 13, 11, 10, 10, '', NULL, 'fire; bludgeoning, piercing, slashing from nonmagical attacks', 'poison', 'exhaustion, grappled, paralyzed, petrified, poisoned, prone, restrained, unconscious', 'blindsight 30 ft., passive Perception 10', 'understands Aquan but doesn''t speak', 'Invisible in Water: The water weird is invisible while fully immersed in water.

Water Bound: Dies if it leaves the water to which it is bound.', 'Constrict: +5 to hit, reach 10 ft. Hit: 13 (3d6+3) bludgeoning. Target is grappled (escape DC 13). Until the grapple ends, the target is restrained, the water weird tries to drown it, and has advantage on attack rolls against it.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (142, 'Baphomet', 'fiend (demon)', '23', 275, '22d12+132', 22, 'natural armor', '40 ft.', 30, 14, 22, 18, 24, 16, 'Saves: STR +17, CON +13, WIS +14 | Skills: Intimidation +10, Perception +14', NULL, 'cold, fire, lightning; bludgeoning, piercing, slashing from nonmagical attacks', 'poison', 'charmed, exhaustion, frightened, poisoned', 'truesight 120 ft., passive Perception 24', 'all, telepathy 120 ft.', 'Charge: If Baphomet moves at least 10 ft. straight toward a target and then hits it with a gore attack on the same turn, the target takes an extra 16 (3d10) piercing damage. If the target is a creature, it must succeed on a DC 25 Strength save or be pushed up to 10 ft. and knocked prone.

Legendary Resistance (3/Day): If Baphomet fails a save, he can choose to succeed instead.

Magic Resistance: Advantage on saving throws against spells and magical effects.

Magic Weapons: Weapon attacks are magical.

Reckless: At start of turn, can gain advantage on all melee attacks until end of turn; attack rolls against Baphomet have advantage until next turn.

Innate Spellcasting (CHA, save DC 18): At will: dispel magic, dominate beast, hunter''s mark, maze, pass without trace. 3/day: dispel evil and good, hold monster. 1/day: teleport.', 'Multiattack: Two attacks: one with Heartcleaver and one gore, or three melee attacks.

Heartcleaver (Magical Greataxe): +17 to hit, reach 10 ft. Hit: 23 (3d8+10) slashing plus 13 (3d8) psychic.

Gore: +17 to hit, reach 10 ft. Hit: 17 (3d6+7) piercing.

Bite: +17 to hit, reach 5 ft. Hit: 22 (3d10+6) piercing.', 'Baphomet can take 3 legendary actions, choosing from the options below. Only one at a time, at end of another creature''s turn.

Crowd Fears: Each creature of Baphomet''s choice that can see him within 60 ft. must succeed on a DC 18 Wisdom save or be frightened of him until the end of its next turn.

Gore (Costs 2 Actions): Baphomet makes one gore attack.

Spell (Costs 3 Actions): Baphomet casts one of his innate spells.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (143, 'Demogorgon', 'fiend (demon)', '26', 406, '28d12+196', 22, 'natural armor', '50 ft., swim 50 ft.', 29, 14, 25, 22, 21, 26, 'Saves: DEX +9, CON +14, WIS +12, CHA +15 | Skills: Insight +12, Perception +19', NULL, 'cold, fire, lightning; bludgeoning, piercing, slashing from nonmagical attacks', 'poison', 'charmed, exhaustion, frightened, poisoned', 'truesight 120 ft., passive Perception 29', 'all, telepathy 120 ft.', 'Legendary Resistance (3/Day): If Demogorgon fails a saving throw, he can choose to succeed instead.

Magic Resistance: Advantage on saves against spells and magical effects.

Magic Weapons: Weapon attacks are magical.

Two Heads: Demogorgon has advantage on saves against being blinded, deafened, stunned, or knocked unconscious.

Innate Spellcasting (CHA, save DC 23): At will: detect magic, detect thoughts, telekinesis. 3/day each: fear, teleport. 1/day each: feeblemind, project image.', 'Multiattack: Two tentacle attacks.

Tentacle: +16 to hit, reach 10 ft. Hit: 28 (3d12+9) bludgeoning. Target must succeed on a DC 23 Constitution save or its hit point maximum is reduced by the damage taken. This reduction lasts until the target finishes a long rest. Target dies if this reduces maximum to 0.

Gaze (Recharge 5–6): Demogorgon fixes his gaze on one creature he can see within 120 ft. Target must make a DC 23 Wisdom save or suffer one of the following (Demogorgon''s choice):
• Beguiling Gaze: Charmed 1 hour.
• Hypnotic Gaze: Stunned until end of next turn.
• Insanity Gaze: Suffers the effect of the confusion spell for 1 minute (Wis save at end of each turn to end).', 'Demogorgon can take 3 legendary actions, choosing from the options below.

Tail: +16 to hit, reach 15 ft. Hit: 20 (2d10+9) bludgeoning.

Gaze (Costs 2 Actions): Demogorgon uses his Gaze.

Cast a Spell (Costs 3 Actions): Demogorgon casts one of his innate spells.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (1, 'Aboleth', 'aberration', '10', 135, '18d10', 17, 'natural', 'walk 10 ft., swim 40 ft.', 21, 9, 15, 18, 15, 18, 'Saves: Perception+10', NULL, '', '', '', 'darkvision 120 ft., passive Perception 20', 'Deep Speech, telepathy 120 ft.', 'Amphibious: The aboleth can breathe air and water.

Mucous Cloud: While underwater, the aboleth is surrounded by transformative mucus. A creature that touches the aboleth or that hits it with a melee attack while within 5 ft. of it must make a DC 14 Constitution saving throw. On a failure, the creature is diseased for 1d4 hours. The diseased creature can breathe only underwater.

Probing Telepathy: If a creature communicates telepathically with the aboleth, the aboleth learns the creature''s greatest desires if the aboleth can see the creature.', 'Multiattack: The aboleth makes three tentacle attacks.

Tentacle: Melee Weapon Attack: +9 to hit, reach 10 ft., one target. Hit: 12 (2d6 + 5) bludgeoning damage. If the target is a creature, it must succeed on a DC 14 Constitution saving throw or become diseased. The disease has no effect for 1 minute and can be removed by any magic that cures disease. After 1 minute, the diseased creature''s skin becomes translucent and slimy, the creature can''t regain hit points unless it is underwater, and the disease can be removed only by heal or another disease-curing spell of 6th level or higher. When the creature is outside a body of water, it takes 6 (1d12) acid damage every 10 minutes unless moisture is applied to the skin before 10 minutes have passed.

Tail: Melee Weapon Attack: +9 to hit, reach 10 ft., one target. Hit: 15 (3d6 + 5) bludgeoning damage.

Enslave: The aboleth targets one creature it can see within 30 ft. of it. The target must succeed on a DC 14 Wisdom saving throw or be magically charmed by the aboleth until the aboleth dies or until it is on a different plane of existence from the target. The charmed target is under the aboleth''s control and can''t take reactions, and the aboleth and the target can communicate telepathically with each other over any distance.
Whenever the charmed target takes damage, the target can repeat the saving throw. On a success, the effect ends. No more than once every 24 hours, the target can also repeat the saving throw when it is at least 1 mile away from the aboleth.', 'Detect: The aboleth makes a Wisdom (Perception) check.

Tail Swipe: The aboleth makes one tail attack.

Psychic Drain (Costs 2 Actions): One creature charmed by the aboleth takes 10 (3d6) psychic damage, and the aboleth regains hit points equal to the damage the creature takes.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (2, 'Adult Red Dragon', 'dragon', '17', 256, '19d12', 19, 'natural', 'walk 40 ft., fly 80 ft., climb 40 ft.', 27, 10, 25, 16, 13, 21, 'Saves: Stealth+6', NULL, '', 'fire', '', 'darkvision 120 ft., blindsight 60 ft., passive Perception 23', 'Common, Draconic', 'Legendary Resistance: If the dragon fails a saving throw, it can choose to succeed instead.', 'Multiattack: The dragon can use its Frightful Presence. It then makes three attacks: one with its bite and two with its claws.

Bite: Melee Weapon Attack: +14 to hit, reach 10 ft., one target. Hit: 19 (2d10 + 8) piercing damage plus 7 (2d6) fire damage.

Claw: Melee Weapon Attack: +14 to hit, reach 5 ft., one target. Hit: 15 (2d6 + 8) slashing damage.

Tail: Melee Weapon Attack: +14 to hit, reach 15 ft., one target. Hit: 17 (2d8 + 8) bludgeoning damage.

Frightful Presence: Each creature of the dragon''s choice that is within 120 ft. of the dragon and aware of it must succeed on a DC 19 Wisdom saving throw or become frightened for 1 minute. A creature can repeat the saving throw at the end of each of its turns, ending the effect on itself on a success. If a creature''s saving throw is successful or the effect ends for it, the creature is immune to the dragon''s Frightful Presence for the next 24 hours.

Fire Breath: The dragon exhales fire in a 60-foot cone. Each creature in that area must make a DC 21 Dexterity saving throw, taking 63 (18d6) fire damage on a failed save, or half as much damage on a successful one.', 'Detect: The dragon makes a Wisdom (Perception) check.

Tail Attack: The dragon makes a tail attack.

Wing Attack (Costs 2 Actions): The dragon beats its wings. Each creature within 10 ft. of the dragon must succeed on a DC 22 Dexterity saving throw or take 15 (2d6 + 8) bludgeoning damage and be knocked prone. The dragon can then fly up to half its flying speed.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (145, 'Yeenoghu', 'fiend (demon)', '24', 333, '23d12+184', 20, 'natural armor', '50 ft.', 29, 10, 26, 17, 22, 20, 'Saves: STR +15, CON +14, WIS +12 | Skills: Perception +12', NULL, 'cold, fire, lightning; bludgeoning, piercing, slashing from nonmagical attacks', 'poison', 'charmed, exhaustion, frightened, poisoned', 'truesight 120 ft., passive Perception 22', 'all, telepathy 120 ft.', 'Legendary Resistance (3/Day): If Yeenoghu fails a save, he can choose to succeed instead.

Magic Resistance: Advantage on saves against spells and magical effects.

Magic Weapons: Weapon attacks are magical.

Rampage: When Yeenoghu reduces a creature to 0 hit points with a melee attack, he can take a bonus action to move up to half his speed and make a bite attack.

Innate Spellcasting (CHA, save DC 19): At will: detect magic. 3/day each: dispel magic, fear, invisibility. 1/day each: feeblemind.', 'Multiattack: Three flail attacks.

Flail: +15 to hit, reach 10 ft. The flail has three heads, each dealing a different effect on a hit.
• Head 1 (Confusion): Hit: 17 (2d8+8) bludgeoning. Target must make DC 17 Wis save or be confused for 1 min.
• Head 2 (Paralysis): Hit: 17 (2d8+8) bludgeoning. Target must make DC 17 Con save or be paralyzed for 1 min.
• Head 3 (Death): Hit: 17 (2d8+8) bludgeoning. Target must make DC 17 Con save or drop to 0 hp.

Bite: +15 to hit, reach 10 ft. Hit: 20 (2d10+9) piercing.', 'Yeenoghu can take 3 legendary actions.

Creature of Ruin: Each gnoll within 60 ft. can make one weapon attack as a reaction.

Flail: Yeenoghu makes one flail attack (random head).

Cast a Spell (Costs 3 Actions): Yeenoghu casts one of his innate spells.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (146, 'Zuggtmoy', 'fiend (demon)', '23', 304, '32d10+128', 18, 'natural armor', '30 ft.', 22, 13, 18, 20, 19, 24, 'Saves: DEX +8, CON +11, WIS +11 | Skills: Perception +11', NULL, 'cold, fire, lightning; bludgeoning, piercing, slashing from nonmagical attacks', 'poison', 'charmed, exhaustion, frightened, poisoned', 'truesight 120 ft., passive Perception 21', 'all, telepathy 120 ft.', 'Legendary Resistance (3/Day): If Zuggtmoy fails a save, she can choose to succeed instead.

Magic Resistance: Advantage on saves against spells and magical effects.

Magic Weapons: Weapon attacks are magical.

Innate Spellcasting (CHA, save DC 19): At will: detect magic, locate animals or plants, ray of sickness. 3/day each: dispel magic, ensnaring strike, plant growth. 1/day each: etherealness, teleport.

Spore Infusion: Zuggtmoy radiates a cloud of spores. When a creature starts its turn within 10 ft. it must make a DC 19 Constitution save or become poisoned. While poisoned, it can''t take reactions and must use its action to make one melee attack against a random creature within reach. If it can''t, it does nothing on its turn.', 'Multiattack: Three pseudopod attacks.

Pseudopod: +13 to hit, reach 10 ft. Hit: 15 (2d8+6) bludgeoning plus 9 (2d8) poison.

Animate Corpse (3/Day): Zuggtmoy targets a humanoid corpse within 30 ft. Within 24 hours it rises as a spore servant under her control.

Paralyzing Spores (Recharge 5–6): A 10-ft.-radius cloud of spores erupts from Zuggtmoy. Each creature in area makes a DC 19 Constitution save. Fail: 12 (2d6+5) poison and paralyzed for 1 minute. Succeed: half, not paralyzed.', 'Zuggtmoy can take 3 legendary actions.

Pseudopod: Zuggtmoy makes one pseudopod attack.

Spores (Costs 2 Actions): Zuggtmoy uses Paralyzing Spores if available, otherwise Spore Infusion activates immediately.

Beguiling Spores (Costs 3 Actions): One creature Zuggtmoy can see within 30 ft. must make a DC 19 Wisdom save or be charmed for 1 hour. While charmed, the creature obeys Zuggtmoy''s verbal commands.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (10, 'Doppelganger', 'monstrosity', '3', 52, '8d8', 14, 'dex', 'walk 30 ft.', 11, 18, 14, 11, 12, 14, 'Saves: Insight+3', NULL, '', '', 'Charmed', 'darkvision 60 ft., passive Perception 11', 'Common', 'Shapechanger: The doppelganger can use its action to polymorph into a Small or Medium humanoid it has seen, or back into its true form. Its statistics, other than its size, are the same in each form. Any equipment it is wearing or carrying isn''t transformed. It reverts to its true form if it dies.

Ambusher: In the first round of combat, the doppelganger has advantage on attack rolls against any creature it has surprised.

Surprise Attack: If the doppelganger surprises a creature and hits it with an attack during the first round of combat, the target takes an extra 10 (3d6) damage from the attack.', 'Multiattack: The doppelganger makes two melee attacks.

Slam: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 7 (1d6 + 4) bludgeoning damage.

Read Thoughts: The doppelganger magically reads the surface thoughts of one creature within 60 ft. of it. The effect can penetrate barriers, but 3 ft. of wood or dirt, 2 ft. of stone, 2 inches of metal, or a thin sheet of lead blocks it. While the target is in range, the doppelganger can continue reading its thoughts, as long as the doppelganger''s concentration isn''t broken (as if concentrating on a spell). While reading the target''s mind, the doppelganger has advantage on Wisdom (Insight) and Charisma (Deception, Intimidation, and Persuasion) checks against the target.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (14, 'Duergar', 'humanoid', '1', 26, '4d8', 16, 'armor', 'walk 25 ft.', 14, 11, 14, 11, 10, 9, '', NULL, 'poison', '', '', 'darkvision 120 ft., passive Perception 10', 'Dwarvish, Undercommon', 'Duergar Resilience: The duergar has advantage on saving throws against poison, spells, and illusions, as well as to resist being charmed or paralyzed.

Sunlight Sensitivity: While in sunlight, the duergar has disadvantage on attack rolls, as well as on Wisdom (Perception) checks that rely on sight.', 'Enlarge: For 1 minute, the duergar magically increases in size, along with anything it is wearing or carrying. While enlarged, the duergar is Large, doubles its damage dice on Strength-based weapon attacks (included in the attacks), and makes Strength checks and Strength saving throws with advantage. If the duergar lacks the room to become Large, it attains the maximum size possible in the space available.

War Pick: Melee Weapon Attack: +4 to hit, reach 5 ft., one target. Hit: 6 (1d8 + 2) piercing damage, or 11 (2d8 + 2) piercing damage while enlarged.

Javelin: Melee or Ranged Weapon Attack: +4 to hit, reach 5 ft. or range 30/120 ft., one target. Hit: 5 (1d6 + 2) piercing damage, or 9 (2d6 + 2) piercing damage while enlarged.

Invisibility: The duergar magically turns invisible until it attacks, casts a spell, or uses its Enlarge, or until its concentration is broken, up to 1 hour (as if concentrating on a spell). Any equipment the duergar wears or carries is invisible with it.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (19, 'Ghost', 'undead', '4', 45, '10d8', 11, 'dex', 'walk 0 ft., fly 40 ft.', 7, 13, 10, 10, 12, 17, '', NULL, 'acid, fire, lightning, thunder, bludgeoning, piercing, and slashing from nonmagical weapons', 'cold, necrotic, poison', 'Charmed, Exhaustion, Frightened, Grappled, Paralyzed, Petrified, Poisoned, Prone, Restrained', 'darkvision 60 ft., passive Perception 11', 'any languages it knew in life', 'Ethereal Sight: The ghost can see 60 ft. into the Ethereal Plane when it is on the Material Plane, and vice versa.

Incorporeal Movement: The ghost can move through other creatures and objects as if they were difficult terrain. It takes 5 (1d10) force damage if it ends its turn inside an object.', 'Withering Touch: Melee Weapon Attack: +5 to hit, reach 5 ft., one target. Hit: 17 (4d6 + 3) necrotic damage.

Etherealness: The ghost enters the Ethereal Plane from the Material Plane, or vice versa. It is visible on the Material Plane while it is in the Border Ethereal, and vice versa, yet it can''t affect or be affected by anything on the other plane.

Horrifying Visage: Each non-undead creature within 60 ft. of the ghost that can see it must succeed on a DC 13 Wisdom saving throw or be frightened for 1 minute. If the save fails by 5 or more, the target also ages 1d4 × 10 years. A frightened target can repeat the saving throw at the end of each of its turns, ending the frightened condition on itself on a success. If a target''s saving throw is successful or the effect ends for it, the target is immune to this ghost''s Horrifying Visage for the next 24 hours. The aging effect can be reversed with a greater restoration spell, but only within 24 hours of it occurring.

Possession: One humanoid that the ghost can see within 5 ft. of it must succeed on a DC 13 Charisma saving throw or be possessed by the ghost; the ghost then disappears, and the target is incapacitated and loses control of its body. The ghost now controls the body but doesn''t deprive the target of awareness. The ghost can''t be targeted by any attack, spell, or other effect, except ones that turn undead, and it retains its alignment, Intelligence, Wisdom, Charisma, and immunity to being charmed and frightened. It otherwise uses the possessed target''s statistics, but doesn''t gain access to the target''s knowledge, class features, or proficiencies.
The possession lasts until the body drops to 0 hit points, the ghost ends it as a bonus action, or the ghost is turned or forced out by an effect like the dispel evil and good spell. When the possession ends, the ghost reappears in an unoccupied space within 5 ft. of the body. The target is immune to this ghost''s Possession for 24 hours after succeeding on the saving throw or after the possession ends.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (26, 'Green Hag', 'fey', '3', 82, '11d8', 17, 'natural', 'walk 30 ft.', 18, 12, 16, 13, 14, 14, 'Saves: Stealth+3', NULL, '', '', '', 'darkvision 60 ft., passive Perception 14', 'Common, Draconic, Sylvan', 'Amphibious: The hag can breathe air and water.

Innate Spellcasting: The hag''s innate spellcasting ability is Charisma (spell save DC 12). She can innately cast the following spells, requiring no material components:

At will: dancing lights, minor illusion, vicious mockery

Mimicry: The hag can mimic animal sounds and humanoid voices. A creature that hears the sounds can tell they are imitations with a successful DC 14 Wisdom (Insight) check.', 'Claws: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 13 (2d8 + 4) slashing damage.

Illusory Appearance: The hag covers herself and anything she is wearing or carrying with a magical illusion that makes her look like another creature of her general size and humanoid shape. The illusion ends if the hag takes a bonus action to end it or if she dies.
The changes wrought by this effect fail to hold up to physical inspection. For example, the hag could appear to have smooth skin, but someone touching her would feel her rough flesh. Otherwise, a creature must take an action to visually inspect the illusion and succeed on a DC 20 Intelligence (Investigation) check to discern that the hag is disguised.

Invisible Passage: The hag magically turns invisible until she attacks or casts a spell, or until her concentration ends (as if concentrating on a spell). While invisible, she leaves no physical evidence of her passage, so she can be tracked only by magic. Any equipment she wears or carries is invisible with her.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (36, 'Otyugh', 'aberration', '5', 114, '12d10', 14, 'natural', 'walk 30 ft.', 16, 11, 19, 6, 13, 6, 'Saves: CON+7', NULL, '', '', '', 'darkvision 120 ft., passive Perception 11', 'Otyugh', 'Limited Telepathy: The otyugh can magically transmit simple messages and images to any creature within 120 ft. of it that can understand a language. This form of telepathy doesn''t allow the receiving creature to telepathically respond.', 'Multiattack: The otyugh makes three attacks: one with its bite and two with its tentacles.

Bite: Melee Weapon Attack: +6 to hit, reach 5 ft., one target. Hit: 12 (2d8 + 3) piercing damage. If the target is a creature, it must succeed on a DC 15 Constitution saving throw against disease or become poisoned until the disease is cured. Every 24 hours that elapse, the target must repeat the saving throw, reducing its hit point maximum by 5 (1d10) on a failure. The disease is cured on a success. The target dies if the disease reduces its hit point maximum to 0. This reduction to the target''s hit point maximum lasts until the disease is cured.

Tentacle: Melee Weapon Attack: +6 to hit, reach 10 ft., one target. Hit: 7 (1d8 + 3) bludgeoning damage plus 4 (1d8) piercing damage. If the target is Medium or smaller, it is grappled (escape DC 13) and restrained until the grapple ends. The otyugh has two tentacles, each of which can grapple one target.

Tentacle Slam: The otyugh slams creatures grappled by it into each other or a solid surface. Each creature must succeed on a DC 14 Constitution saving throw or take 10 (2d6 + 3) bludgeoning damage and be stunned until the end of the otyugh''s next turn. On a successful save, the target takes half the bludgeoning damage and isn''t stunned.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (108, 'Beholder', 'aberration', '13', 180, '19d10+76', 18, 'natural armor', '0 ft., fly 20 ft. (hover)', 10, 14, 18, 17, 15, 17, 'Saves: INT +8, WIS +7, CHA +8 | Skills: Perception +12', NULL, '', '', 'prone', 'darkvision 120 ft., passive Perception 22', 'Deep Speech, Undercommon', 'Antimagic Cone: The beholder''s central eye creates a 150-foot cone of antimagic. At the start of each turn, the beholder decides which way the cone faces and whether it''s active. The area works like the antimagic field spell.', 'Bite: +5 to hit, reach 5 ft. Hit: 14 (4d6) piercing.

Eye Rays: The beholder shoots three of the following magical eye rays at random (reroll duplicates), choosing 1–3 targets within 120 ft. DC 16 save.
1 Charm Ray (Wis, charmed 1 hr)
2 Paralyzing Ray (Con, paralyzed 1 min)
3 Fear Ray (Wis, frightened 1 min)
4 Slowing Ray (Dex, speed halved, no reactions, -2 AC/Dex saves, 1 min)
5 Enervation Ray (Con, 8d8 necrotic)
6 Telekinetic Ray (Str, move up to 30 ft.)
7 Sleep Ray (Wis, unconscious 1 min)
8 Petrification Ray (Dex, restrained then petrified)
9 Disintegration Ray (Dex, 10d8+5 force; kills reduce to dust)
10 Death Ray (Dex, 10d10 necrotic; reduces to 0 if killed)', 'The beholder can take 3 legendary actions (Eye Ray only), choosing from the options below. Only one legendary action option can be used at a time and only at the end of another creature''s turn.

Eye Ray: The beholder uses one random eye ray.', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (115, 'Flumph', 'aberration (lawful good)', '1/8', 7, '2d6', 12, '', '5 ft., fly 30 ft.', 6, 15, 10, 14, 14, 11, 'Skills: Arcana +4, History +4, Perception +4, Religion +4', NULL, '', '', '', 'darkvision 60 ft., passive Perception 14', 'Aquan, Auran, Deep Speech, Undercommon, telepathy 60 ft.', 'Advanced Telepathy: Can perceive the content of any telepathic communication within 60 ft. and cannot be surprised by creatures with telepathy.

Prone Deficiency: If knocked prone, the flumph can''t right itself and is incapacitated until someone rights it.

Telepath Shroud: Immune to magic that allows other creatures to read its thoughts, communicate telepathically without its consent, or sense its emotions.', 'Tendrils: +4 to hit, reach 5 ft. Hit: 4 (1d4+2) piercing plus 2 (1d4) acid. Target is covered in pungent slime for 1 hour unless cleansed; creatures within 5 ft. of target have disadvantage on Charisma checks.

Stench Spray (1/Short Rest): Each creature in a 15-ft. cone must succeed on a DC 10 Dex save or be coated in fetid slime. A coated creature gives off a terrible odor for 1d4 hours.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (120, 'Intellect Devourer', 'aberration', '2', 21, '6d4+6', 12, '', '40 ft.', 6, 14, 13, 12, 11, 10, 'Skills: Perception +2, Stealth +4', NULL, 'bludgeoning, piercing, slashing from nonmagical attacks', '', 'blinded', 'blindsight 60 ft. (blind beyond), passive Perception 12', 'understands Deep Speech but can''t speak; telepathy 60 ft.', 'Detect Sentience: Can sense the presence and location of any creature with INT 3+ within 300 ft. that isn''t protected by a mind blank spell.', 'Multiattack: Two claw attacks.

Claws: +4 to hit, reach 5 ft. Hit: 7 (2d4+2) slashing.

Devour Intellect: Target within 10 ft. makes a DC 12 Intelligence save. On a fail: 11 (2d10) psychic damage, and the target''s Intelligence score is reduced by 1d6. Target dies if this reduces Intelligence to 0; otherwise the reduction lasts until the target finishes a short or long rest. On a success: half damage only.

Body Thief: The intellect devourer initiates a contest of Intelligence against an incapacitated humanoid within 5 ft. On a win, the devourer magically consumes the target''s brain, enters the skull, and takes control of the body. The target is blinded/deafened while the devourer occupies the body. The devourer retains its own INT, WIS, CHA and alignment; all other stats are the host''s.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (124, 'Kuo-toa Whip', 'humanoid (kuo-toa)', '1', 65, '10d8+20', 11, 'natural armor', '30 ft., swim 30 ft.', 14, 10, 14, 12, 14, 11, 'Skills: Perception +6, Religion +4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 16', 'Undercommon', 'Amphibious: Can breathe air and water.

Otherworldly Perception: Can sense invisible/ethereal creatures within 30 ft.

Slippery: Advantage on checks/saves to escape grapples.

Spellcasting (Cleric, Wis, save DC 12, +4 to hit): Cantrips: sacred flame, thaumaturgy, toll the dead. 1st (3): detect magic, sanctuary, thunderwave. 2nd (2): hold person, spiritual weapon.

Sunlight Sensitivity: Disadvantage on attack rolls and Perception checks in sunlight.', 'Multiattack: Two attacks with its pincer staff.

Pincer Staff: +4 to hit, reach 10 ft. Hit: 5 (1d6+2) piercing. Target is grappled (escape DC 14) if size Large or smaller and the whip doesn''t have another creature grappled.

Bite: +4 to hit, reach 5 ft. Hit: 4 (1d4+2) piercing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (128, 'Myconid Adult', 'plant (lawful neutral)', '1/2', 22, '4d8+4', 12, 'natural armor', '20 ft.', 10, 10, 12, 10, 11, 7, '', NULL, '', '', '', 'darkvision 120 ft., passive Perception 10', '—', 'Distress Spores: When damaged, all myconids within 240 ft. are immediately aware of it.

Sun Sickness: While in sunlight, the myconid has disadvantage on ability checks, attack rolls, and saving throws. Dies after 1 hour in sunlight.

Rapport Spores: A 20-ft.-radius cloud of spores allows all affected creatures to communicate telepathically with each other for 1 hour.', 'Fist: +2 to hit, reach 5 ft. Hit: 5 (1d10) bludgeoning plus 5 (2d4) poison.

Pacifying Spores: 10-ft. cone; one creature must make a DC 11 Constitution save or be stunned for 1 minute. Repeat save at end of each turn.

Rapport Spores: 20-ft. radius burst. Creatures that inhale them communicate telepathically with each other and all myconids within 1 hour.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (138, 'Succubus/Incubus', 'fiend (shapechanger)', '4', 66, '12d8+12', 13, 'natural armor', '30 ft., fly 60 ft.', 8, 17, 13, 15, 12, 20, 'Skills: Deception +9, Insight +5, Perception +5, Persuasion +9, Stealth +7', NULL, 'cold, fire, lightning, poison; bludgeoning, piercing, slashing from nonmagical attacks', '', '', 'darkvision 60 ft., passive Perception 15', 'Abyssal, Common, Infernal, telepathy 60 ft.', 'Telepathic Bond: The fiend ignores the range restriction of its telepathy when communicating with a creature it has charmed. The two can communicate while on different planes.

Shapechanger: The fiend can use its action to polymorph into a Small or Medium humanoid, or back into its true form. Its statistics are the same in each form.', 'Claw (True Form Only): +5 to hit, reach 5 ft. Hit: 6 (1d6+3) slashing.

Charm: One humanoid that the fiend can see within 30 ft. must succeed on a DC 15 Wisdom save or be magically charmed for 1 day. Charmed target obeys the fiend''s spoken commands. Each time the target takes damage it repeats the save.

Draining Kiss: The fiend kisses a charmed creature or willing creature. Target loses 32 (5d10+5) hit points (no save), and the fiend regains the same number.

Etherealness: The fiend magically enters the Ethereal Plane.

Dreamwalk: While on the Ethereal Plane, the fiend can haunt the dreams of a creature on the Material Plane within 1 mile. The fiend appears in the creature''s dreams.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (140, 'Umber Hulk', 'monstrosity', '5', 93, '11d10+33', 18, 'natural armor', '30 ft., burrow 20 ft.', 20, 13, 16, 9, 10, 10, 'Skills: Perception +3', NULL, '', '', '', 'darkvision 120 ft., tremorsense 60 ft., passive Perception 13', 'Umber Hulk', 'Confusing Gaze: When a creature starts its turn within 30 ft. of the umber hulk and can see its eyes, the umber hulk can magically force that creature to make a DC 15 Charisma saving throw, unless the umber hulk is incapacitated. On a failed save, the creature can''t take reactions until start of its next turn and rolls a d8 to determine what it does: 1–4 do nothing; 5–6 move in random direction; 7–8 make one melee attack against random creature within range.

Tunneler: The umber hulk can burrow through solid rock at half its burrow speed and leaves a 5-ft. wide, 8-ft. tall tunnel in its wake.', 'Multiattack: Two claws and one mandibles attack.

Claw: +8 to hit, reach 5 ft. Hit: 9 (1d8+5) slashing.

Mandibles: +8 to hit, reach 5 ft. Hit: 14 (2d8+5) slashing.', '', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (144, 'Juiblex', 'fiend (demon)', '23', 350, '28d12+168', 18, 'natural armor', '30 ft., climb 30 ft., swim 30 ft.', 24, 9, 23, 20, 20, 16, 'Saves: DEX +6, CON +13, WIS +12 | Skills: Perception +12', NULL, 'cold, fire, lightning; bludgeoning, piercing, slashing from nonmagical attacks', 'acid, poison', 'blinded, charmed, deafened, exhaustion, frightened, grappled, paralyzed, petrified, poisoned, prone, restrained, stunned, unconscious', 'truesight 120 ft., passive Perception 22', 'all, telepathy 120 ft.', 'Legendary Resistance (3/Day): If Juiblex fails a saving throw, it can choose to succeed instead.

Magic Resistance: Advantage on saves against spells and magical effects.

Magic Weapons: Weapon attacks are magical.

Reactive: Juiblex can take one reaction on every turn in combat.

Spew Slime (Recharge 5–6): Juiblex exhales a 60-ft. cone of acid slime. Each creature in area must make a DC 21 Dexterity save. Fail: 45 (10d8) acid damage and is covered in slime (halved speed, disadvantage on Dex checks). Success: half damage only.', 'Multiattack: Three pseudopod attacks.

Pseudopod: +14 to hit, reach 10 ft. Hit: 21 (4d6+7) bludgeoning plus 14 (4d6) acid. Target must make a DC 21 Con save or contract Slime Fever (diseased). At end of each long rest, target makes another save; on fail it loses 14 max HP permanently. Cured by lesser restoration.

Ebon Tendrils (Recharge 5–6): Tentacles erupt from Juiblex. Each creature within 60 ft. must make a DC 21 Dexterity save. Fail: 35 (10d6) necrotic, restrained, and pulled within 5 ft. Held creature takes 17 (5d6) necrotic at start of Juiblex''s turn.', 'Juiblex can take 3 legendary actions.

Pseudopod: Juiblex makes one pseudopod attack.

Slime Wave (Costs 2 Actions): Juiblex causes slime to erupt from the ground at a point within 60 ft. Each creature within 10 ft. must make a DC 21 Dex save or be covered in slime.

Acid Rain (Costs 3 Actions): Acidic slime rains on creatures within 60 ft. (DC 21 Dex save, 21/4d10 acid on fail).', NULL);
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (241, 'Shoor Vandree', 'humanoid (elf)', '5', 71, '11d8+22', 18, 'adamantine studded leather, shield', '30 ft.', 13, 18, 14, 11, 13, 12, 'Saves: DEX +6, CON +4, WIS +3 | Skills: Perception +3, Stealth +6', NULL, '', '', '', 'darkvision 120 ft., passive Perception 13', 'Elvish, Undercommon', 'Fey Ancestry: Advantage on saves vs. charmed; magic can''t put to sleep.

Innate Spellcasting (Cha, DC 12): At will: dancing lights; 1/day: darkness, faerie fire, levitate (self only).

Sunlight Sensitivity: Disadvantage on attacks and Perception in sunlight.', 'Multiattack: Two shortsword attacks.

Shortsword: +6 to hit, reach 5 ft. Hit: 7 (1d6+4) piercing plus 10 (3d6) poison.

Hand Crossbow: +6 to hit, range 30/120 ft. Hit: 7 (1d6+4) piercing plus 10 (3d6) poison.', 'Parry (Reaction): +3 to AC against one melee attack that would hit, if wielding a melee weapon.', 'Ilvara''s favored warrior. Replaced Jorlan after Jorlan''s injury. Wears Jorlan''s confiscated adamantine armor. Overconfident — his arrogance is his weakness.');
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (242, 'Jorlan Duskryn', 'humanoid (elf)', '5', 50, '8d8+16', 17, 'studded leather, shield', '30 ft.', 11, 16, 14, 11, 13, 12, 'Saves: DEX +5, WIS +3 | Skills: Perception +3, Stealth +5', NULL, '', '', '', 'darkvision 120 ft., passive Perception 13', 'Elvish, Undercommon', 'Fey Ancestry: Advantage on saves vs. charmed; magic can''t put to sleep.

Crippled Hands: Jorlan''s hands were badly damaged by a carrion crawler. He has disadvantage on all attack rolls and Strength (Athletics) checks.

Innate Spellcasting (Cha, DC 12): At will: dancing lights; 1/day: darkness, faerie fire.

Sunlight Sensitivity: Disadvantage on attacks and Perception in sunlight.', 'Multiattack: Two shortsword attacks (both at disadvantage due to Crippled Hands).

Shortsword: +5 to hit (disadvantage), reach 5 ft. Hit: 6 (1d6+3) piercing plus 7 (2d6) poison.

Hand Crossbow: +5 to hit (disadvantage), range 30/120 ft. Hit: 6 (1d6+3) piercing plus 7 (2d6) poison.', '', 'Former favored warrior of Ilvara. Bitterly resentful of Shoor. Can be turned against Ilvara if players cultivate his hatred. May provide the master key or a weapon if approached correctly. CR 5 (effectively lower due to injury penalty).');
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (246, 'Chuul', 'aberration', '4', 93, '11d10+33', 16, 'natural armor', '30 ft., swim 30 ft.', 19, 10, 16, 5, 11, 5, 'Skills: Perception +4', NULL, '', 'poison', 'poisoned', 'darkvision 60 ft., tremorsense 30 ft., passive Perception 14', 'understands Deep Speech but can''t speak', 'Amphibious: Can breathe air and water.

Sense Magic: The chuul senses magic within 120 ft. at will. This trait otherwise works like the detect magic spell but isn''t itself magical.', 'Multiattack: Two pincer attacks. If both hit the same target it is grappled (escape DC 14).

Pincer: +6 to hit, reach 10 ft. Hit: 11 (2d6+4) bludgeoning. Target is grappled (escape DC 14) if it is a Large or smaller creature and the chuul doesn''t have two other creatures grappled.

Tentacles: One creature grappled by the chuul must succeed on a DC 13 Constitution save or be poisoned for 1 minute. While poisoned this way, the target is also paralyzed. The creature may repeat the save at end of each of its turns.', '', '');
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (240, 'Ilvara Mizzrym', 'humanoid (elf)', '8', 71, '13d8+13', 16, 'scale mail', '30 ft.', 10, 14, 12, 13, 17, 18, 'Saves: CON +4, WIS +6, CHA +7 | Skills: Insight +6, Perception +6, Religion +4, Stealth +5', NULL, '', '', '', 'darkvision 120 ft., passive Perception 16', 'Elvish, Undercommon', 'Fey Ancestry: Advantage on saves vs. charmed; magic can''t put to sleep.

Innate Spellcasting (Cha, DC 15): At will: dancing lights; 1/day: darkness, faerie fire, levitate (self only).

Spellcasting (Wis, DC 14, +6 to hit): Cantrips: guidance, poison spray, resistance, spare the dying, thaumaturgy. 1st (4): cure wounds, detect magic, ray of sickness, shield of faith. 2nd (3): hold person, spiritual weapon. 3rd (3): clairvoyance, dispel magic. 4th (3): freedom of movement, guardian of faith. 5th (2): contagion, insect plague. 6th (1): harm.

Sunlight Sensitivity: Disadvantage on attacks and Perception checks in sunlight.

Spider Familiar: Ilvara is accompanied by a spider familiar that acts as a scout.', 'Multiattack: Two scourge attacks.

Scourge: +5 to hit, reach 5 ft. Hit: 5 (1d6+2) piercing plus 17 (5d6) poison.

Suggestion (1/Day): Spell effect, DC 14 Wisdom save.', '', 'PRIMARY ANTAGONIST of Part 1. Pursues the party relentlessly after escape. Confrontation at Blingdenstone is the natural endpoint. Uses Suggestion on strongest PC. CR 8.');
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (244, 'Ploopploopeen', 'humanoid (kuo-toa)', '5', 65, '10d8+20', 13, 'natural armor', '30 ft., swim 30 ft.', 14, 14, 14, 12, 16, 14, 'Saves: CON +4, WIS +5, CHA +4 | Skills: Perception +7, Religion +3', NULL, '', '', '', 'darkvision 120 ft., passive Perception 17', 'Undercommon', 'Amphibious: Can breathe air and water.

Otherworldly Perception: Senses invisible/ethereal creatures within 30 ft.

Slippery: Advantage on checks and saves to escape grapples.

Spellcasting (Wis, DC 13, +5 to hit): Cantrips: guidance, sacred flame, thaumaturgy. 1st (4): detect magic, sanctuary, shield of faith. 2nd (3): hold person, spiritual weapon. 3rd (3): spirit guardians, tongues. 4th (2): control water, divination.

Sunlight Sensitivity: Disadvantage on attacks and Perception checks in sunlight.', 'Multiattack: Two scepter attacks.

Scepter: +4 to hit, reach 5 ft. Hit: 5 (1d6+2) bludgeoning plus 7 (2d6) lightning.

Unarmed Strike: +4 to hit, reach 5 ft. Hit: 3 (1d4+2) bludgeoning.', '', 'Archpriest of the Sea Mother (Blibdoolpoolp) at Sloobludop. Plans to use the party as a sacrifice. Is himself being outmaneuvered by his daughter Bloppblippodd.');
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (243, 'The Pudding King', 'humanoid (deep gnome)', '4', 32, '5d6+15', 13, 'natural armor', '25 ft.', 8, 14, 16, 16, 7, 12, 'Saves: CON +5, INT +5 | Skills: Arcana +5, Stealth +4', NULL, '', 'poison', 'poisoned', 'darkvision 120 ft., passive Perception 8', 'Gnomish, Undercommon', 'Gnome Cunning: Advantage on INT, WIS, CHA saves against magic.

Innate Spellcasting (INT, DC 13): At will: nondetection (self only). 1/day each: blindness/deafness, blur, disguise self.

Juiblex''s Blessing: The Pudding King can communicate telepathically with any ooze within 120 ft. Oozes are friendly to him and follow his mental commands.

Speaks in Couplets: The Pudding King always speaks in rhyming couplets; this has no mechanical effect but makes him immediately recognizable.', 'Acidic Touch: +4 to hit, reach 5 ft. Hit: 10 (3d6) acid damage, and the target must succeed on a DC 13 Constitution save or be coated in slime (speed halved, disadvantage on Dex checks, 1 min).

Spawn Ooze (Recharge 5–6): The Pudding King calls 1d2 gray oozes from nearby surfaces. They appear in unoccupied spaces within 10 ft.', '', 'Mad deep gnome consumed by Juiblex. Refers to the two elder black puddings as his children. Defeating him ends the ooze infestation of Blingdenstone. Speaks only in rhyming couplets.');
INSERT INTO public.monsters (id, name, type, cr, hp, hp_formula, ac, ac_desc, speed, str, dex, con, int_score, wis, cha, saving_throws, skills, damage_resistances, damage_immunities, condition_immunities, senses, languages, traits, actions, legendary_actions, notes) OVERRIDING SYSTEM VALUE VALUES (245, 'Bloppblippodd', 'humanoid (kuo-toa)', '6', 97, '13d8+39', 13, 'natural armor', '30 ft., swim 30 ft.', 16, 14, 16, 13, 16, 14, 'Saves: CON +6, WIS +6, CHA +5 | Skills: Perception +9, Religion +4', NULL, '', '', '', 'darkvision 120 ft., passive Perception 19', 'Undercommon', 'Amphibious: Can breathe air and water.

Otherworldly Perception: Senses invisible/ethereal creatures within 30 ft.

Slippery: Advantage on checks and saves to escape grapples.

Spellcasting (Wis, DC 14, +6 to hit): Cantrips: guidance, sacred flame, thaumaturgy. 1st (4): detect magic, sanctuary, shield of faith. 2nd (3): hold person, spiritual weapon. 3rd (3): spirit guardians, tongues. 4th (3): control water, divination. 5th (2): mass cure wounds, scrying.

Sunlight Sensitivity: Disadvantage on attacks and Perception checks in sunlight.

Deep Father''s Blessing: Bloppblippodd has advantage on Charisma (Persuasion) checks made to convert kuo-toa to the cult of the Deep Father.', 'Multiattack: Two scepter attacks.

Scepter: +6 to hit, reach 5 ft. Hit: 6 (1d6+3) bludgeoning plus 9 (2d8) lightning.

Unarmed Strike: +6 to hit, reach 5 ft. Hit: 4 (1d4+3) bludgeoning.', '', 'Ploopploopeen''s daughter and rival. Leads the cult of the Deep Father (actually Demogorgon). Her prayers are actively pulling Demogorgon toward Sloobludop. CR 6.');


--
-- Name: monsters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: wsm52
--

SELECT pg_catalog.setval('public.monsters_id_seq', 267, true);


--
-- PostgreSQL database dump complete
--

\unrestrict cZs5kI8waB1UcnEIlNbjpFYT8ctjEsFMsytk0y6smDwX0Th404Bkn6wMJc0y2cj

