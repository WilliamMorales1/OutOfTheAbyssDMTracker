//go:build ignore

package main

import (
	"context"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://wsm52:H&pg@localhost/oota")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	seedLocations(ctx, conn)
	seedNPCs(ctx, conn)
	seedEncounters(ctx, conn)
	seedEvents(ctx, conn)

	log.Println("seed complete")
}

// normalizeLocation strips sub-location suffixes so they match Locations.name FK.
// e.g. "Velkynvelve - prisoner" → "Velkynvelve"
// Returns empty string for generic/travel locations that have no Locations row.
func normalizeLocation(loc string) string {
	knownLocations := map[string]bool{
		"Velkynvelve":            true,
		"The Silken Paths":       true,
		"Sloobludop":             true,
		"Gracklstugh":            true,
		"Neverlight Grove":       true,
		"Blingdenstone":          true,
		"The Darklake":           true,
		"Gauntlgrym":             true,
		"Mantol-Derith":          true,
		"Araj - Vizeran's Tower": true,
		"Menzoberranzan":         true,
		"The Wormwrithings":      true,
		"The Labyrinth":          true,
		"Araumycos":              true,
	}
	if knownLocations[loc] {
		return loc
	}
	// Try stripping " - ..." suffix
	if idx := strings.Index(loc, " - "); idx != -1 {
		base := loc[:idx]
		if knownLocations[base] {
			return base
		}
	}
	// No matching location row — return empty to use NULL
	return ""
}

func difficultyInt(s string) int {
	switch s {
	case "trivial":
		return 1
	case "medium":
		return 2
	case "hard":
		return 3
	case "deadly":
		return 4
	default:
		return 0
	}
}

// nullableLocation returns nil when loc is empty so pgx inserts NULL (satisfying FK).
func nullableLocation(loc string) interface{} {
	n := normalizeLocation(loc)
	if n == "" {
		return nil
	}
	return n
}

func seedLocations(ctx context.Context, conn *pgx.Conn) {
	type row struct {
		name, typ, status string
		danger            int
		desc, secrets     string
	}
	rows := []row{
		{"Velkynvelve", "dungeon", "visited", 3,
			"A drow outpost built into a cluster of stalactites and natural columns above a subterranean stream. Spider silk bridges connect the stalactite \"buildings.\" Eight drow guards, two elite warriors, a priestess, and roughly 12 quaggoth thralls. The party begins here as prisoners.",
			"Equipment room holds all confiscated party gear. Jorlan has a master key and can be bribed/turned. The spider cult shrine has a hidden cache of potions. Chasm below is 300 ft deep - escape route if players are creative."},
		{"The Silken Paths", "wilderness", "unknown", 2,
			"A vast network of enormous spider webs strung across a massive cavern, connecting stalactites and rock spires. Some webs are strong enough to walk on; others are deceptively fragile. Multiple species of giant spider inhabit the region.",
			"Web Sack: contains a lost drow supply pack with rope, pitons, and a potion of healing. A giant spider queen guards a cache of cocooned travelers - one may still be alive."},
		{"Sloobludop", "city", "unknown", 4,
			"A kuo-toa settlement on the shores of the Darklake. Crude buildings of scavenged stone and bone. The town is split between two fanatical cults. The population (about 500 kuo-toa) has been driven into a collective fever by proximity to Demogorgon's influence.",
			"Demogorgon will physically arrive here at the chapter's climax, rising from the Darklake. The party's only goal when this happens is survival. Ploopploopeen plans to sacrifice the party. Bloppblippodd's \"Deep Father\" is real - and is coming."},
		{"Gracklstugh", "city", "unknown", 3,
			"The City of Blades. A massive duergar city built around volcanic activity and lit by the red glow of lava flows. Population ~10,000 duergar plus slaves. Center of duergar industry, weapons manufacture, and trade. Divided into Darklake District, Blade Bazaar, Laduguer's Furrow, and Themberchaud's Lair.",
			"Demogorgon's influence is causing waves of madness among the populace - civilians attacking each other, stone giants going berserk. The Whorlstone Tunnels beneath the city hold cultists worshipping Demogorgon. Themberchaud suspects he is being used and lied to. The Gray Ghosts thieves guild stole a dragon egg. A secret duergar clan (Clan Cairngorm) is infected with derro tunneling madness."},
		{"Neverlight Grove", "wilderness", "unknown", 3,
			"A massive cavern lit by bioluminescent fungi of breathtaking beauty. Home to thousands of myconids led by two sovereigns. Normally peaceful and welcoming - but now deeply corrupted by Zuggtmoy's presence. Strange flowers bloom in impossible colors. Spores hang thick in the air.",
			"Zuggtmoy is hiding in the Fungal Court at the center of the grove, performing a \"wedding\" ritual with the myconid population. Sovereign Basidia is not yet corrupted and is afraid. The Circle of Masters contains evidence of the corruption. Attending the wedding ceremony risks permanent spore infection (Zuggtmoy's domain power)."},
		{"Blingdenstone", "city", "unknown", 3,
			"The deep gnome city, partially rebuilt after its destruction by drow. Currently under siege from within by a colony of oozes and the spirit of the Pudding King (a mad gnome consumed by Juiblex). The city is divided into functional districts and the haunted Goldwhisker Quarter, held by wererats.",
			"The Pudding King is in the deepest cavern with his \"children\" - two massive oozes named Glabbagool and the Royal Ooze. Topsy and Turvy's wererat kin live in Goldwhisker Quarter. The gnomes are desperate - they will grant the party full access and rewards for clearing the ooze infestation."},
		{"The Darklake", "wilderness", "unknown", 3,
			"A vast underground lake system - hundreds of miles of interconnected waterways, tunnels, and islands. The primary travel route between Sloobludop, Gracklstugh, and the western Underdark. Inhabited by aquatic monsters: plesiosaurs, giant octopuses, kuo-toa, and worse.",
			"Demogorgon's physical form has been partially manifesting beneath the Darklake's deepest waters since his summoning. Islands in the Darklake include: Isle of Rezz (kuo-toa trading post), Flowing Stones (home to a water weird), and the Sunken Tower (submerged wizard's tower with a waterbreathing amulet)."},
		{"Gauntlgrym", "city", "unknown", 2,
			"The ancient dwarven city reclaimed by Bruenor Battlehammer and his clan. A monumental forge-city built around the Hosttower's power and a primordial fire titan imprisoned beneath. Heavily defended, politically significant.",
			"Bruenor will receive the party if they have Eldeth's shield or credible news of demon lords. The city is a staging ground for a potential Underdark military expedition. Lords of Waterdeep can be contacted here via sending stone."},
		{"Mantol-Derith", "city", "unknown", 3,
			"A secret trading post at the intersection of several Underdark factions - drow, duergar, svirfneblin, and illithid all conduct neutral trade here under a strict code of conduct. Demogorgon's madness has broken the truce and the settlement is tearing itself apart.",
			"A Harper agent (Khalessa Draga) is here and desperate for rescue. The trading post's four factions have turned on each other - each blaming the others for a theft. Demogorgon's corrupting influence caused it all. Access to Vizeran's tower requires navigating the ruins of this truce."},
		{"Araj - Vizeran's Tower", "dungeon", "unknown", 2,
			"A remote, cunningly hidden tower of black stone in a desolate stretch of the Underdark. Vizeran DeVir has lived here in secret for centuries, amassing vast arcane knowledge in exile. The tower is trapped, warded, and absolutely isolated.",
			"Vizeran's real plan: the Dark Heart ritual, if completed, will unleash a catastrophic wave of power centered on Menzoberranzan - not merely banishing demons but potentially destroying the city and killing thousands. The party must decide if helping him is worth it."},
		{"Menzoberranzan", "city", "unknown", 5,
			"The City of Spiders. The greatest drow city in the Underdark. Population ~20,000. Ruled by the Ruling Council (eight matron mothers of noble houses). Lit by faerie fire in blues and greens. Cradle of drow civilization and the source of the crisis.",
			"Gromph Baenre's ritual caused the demon summoning. House Baenre is suppressing this information. The city is barely holding together as demon lords' influence drives citizens to violence. The party can negotiate with Matron Mother Triel for support in exchange for keeping the city's secrets."},
		{"The Wormwrithings", "wilderness", "unknown", 4,
			"A vast network of enormous tunnels bored by purple worms. The passages are massive - wide enough for wagons but prone to sudden purple worm emergence. Also inhabited by cloakers, ropers, and gibbering mouthers.",
			"One tunnel system leads directly to the lair of Yeenoghu, who has established a hunting ground here and is driving gnoll warbands through the region in a frenzy of slaughter."},
		{"The Labyrinth", "dungeon", "unknown", 5,
			"Baphomet's domain - a vast, impossible maze carved into the Underdark over centuries of his influence. The tunnels shift, dead ends appear and vanish, and the entire region is saturated with his minotaur-hunting madness. The Maze Engine (a ancient modron artifact) is located here.",
			"The Maze Engine can alter reality locally when activated. A series of wild magic effects occurs when it runs. Baphomet hunts here personally. The engine can be used to banish a demon lord if triggered correctly during the final ritual."},
		{"Araumycos", "wilderness", "unknown", 3,
			"The largest living organism in the Forgotten Realms - a colossal fungal colony that occupies much of the northern Underdark. Zuggtmoy has been communing with Araumycos, attempting to use it as a vessel for her power and a weapon of world-ending scope.",
			"If Zuggtmoy fully merges with Araumycos, the resulting entity would be effectively impossible to stop. The merger must be disrupted as part of the ritual components Vizeran requires (her spores). Time pressure: the longer Part 2 takes, the closer the merger gets."},
	}

	for _, r := range rows {
		_, err := conn.Exec(ctx,
			`INSERT INTO Locations (name, type, status, danger, description, secrets)
			 VALUES ($1,$2,$3,$4,$5,$6)
			 ON CONFLICT (name) DO NOTHING`,
			r.name, r.typ, r.status, r.danger, r.desc, r.secrets)
		if err != nil {
			log.Fatalf("location %q: %v", r.name, err)
		}
	}
	log.Println("locations seeded")
}

func seedNPCs(ctx context.Context, conn *pgx.Conn) {
	type row struct {
		name, disposition, location string
		madness                     int
		desc, notes                 string
	}
	rows := []row{
		{"Buppido", "neutral", "Velkynvelve - prisoner", 4,
			"Short, stocky derro with wild eyes and a disarming grin. Talkative, charming, cunning. Secretly believes himself to be the reincarnation of the derro god Diinkarazan and a chosen divine instrument of murder.",
			"SECRET: Plans to murder companions one by one as ritual sacrifices. Wants to reach Gracklstugh where he will gather a small murderous cult. His friendliness is a facade - betray this when players start to notice missing companions or supplies."},
		{"Prince Derendil", "ally", "Velkynvelve - prisoner", 6,
			"A massive quaggoth who speaks Elvish with courtly grace. Claims to be Derendil, elven prince of Nelrindenvane in the High Forest, polymorphed into a quaggoth by a jealous rival. Composed and regal - until stressed.",
			"Delusion is implanted by Fraz-Urb'luu. The kingdom of Nelrindenvane does not exist. Under stress or when his story is questioned he flies into a quaggoth berserk rage. Goes berserk if reduced below 10 HP. Treat with compassion - his tragedy is genuine, even if the backstory is false."},
		{"Eldeth Feldrun", "ally", "Velkynvelve - prisoner", 0,
			"Proud, defiant shield dwarf scout from Gauntlgrym. Strong warrior, hates drow with a personal intensity. Captured while scouting Underdark passages near her city.",
			"Bond: Recover her shield and warhammer from the equipment room. Goal: Return to Gauntlgrym. If she dies, her wish is for her shield to be returned to her kin at Gauntlgrym. Key NPC hookup for Gauntlgrym in Part 2. May sacrifice herself protecting the party - let that land."},
		{"Jimjar", "ally", "Velkynvelve - prisoner", 1,
			"Wiry, cheerful deep gnome whose most defining trait is a compulsive gambling addiction. Bets on absolutely everything - tomorrow's marching order, whether a stalactite will fall, who eats first. Always smiling.",
			"Excellent scout and navigator - knows the routes to Blingdenstone cold. Will offer to guide the party there. His gambling is harmless but constant comic relief. Reliable in a fight despite his demeanor."},
		{"Ront", "neutral", "Velkynvelve - prisoner", 2,
			"Hulking orc warrior, captured after his tribe was scattered by a drow raid. Bully to those weaker than him, but crumbles in the face of firm authority. Deeply ashamed of his capture.",
			"Responds to strength - if a PC wins a contest of strength or calls him out decisively, he will follow and eventually be loyal. Starts hostile to gnomes and halflings. Can evolve into a genuine ally by Part 2 if handled well."},
		{"Sarith Kzekarit", "neutral", "Velkynvelve - prisoner", 3,
			"Brooding, guilt-ridden drow warrior. Imprisoned for murdering a fellow guard during a psychotic episode he cannot explain. Morose, withdrawn, occasionally whispers to himself.",
			"INFECTED by Zuggtmoy's spores - he does not know it. Will become increasingly erratic near Neverlight Grove. Excellent navigator; knows Underdark routes but must be drawn out. His infection is a slow-burn horror subplot. He will eventually die or fully turn."},
		{"Shuushar the Awakened", "ally", "Velkynvelve - prisoner", 0,
			"A rare kuo-toa who has achieved genuine Buddhist-style enlightenment. Serene, pacifist, soft-spoken. Finds beauty in everything. Completely refuses to harm any creature.",
			"Best guide for the Darklake and for reaching Sloobludop - his home settlement. His pacifism will cause tension during fights. His people at Sloobludop have gone dangerously mad. His calm will be shattered by what he finds at home - play this for emotional weight."},
		{"Stool", "ally", "Velkynvelve - prisoner", 0,
			"A small, frightened young myconid, stolen from Neverlight Grove. Communicates only through rapport spores - a cloud of telepathic fungal spores shared among those who inhale them.",
			"CRITICAL UTILITY NPC. Rapport spores allow all affected creatures to communicate telepathically, bypassing language barriers entirely. Desperately wants to return to Neverlight Grove. Do not kill Stool - players will bond fiercely with it. Sovereign Phylo will want it back."},
		{"Topsy", "neutral", "Velkynvelve - prisoner", 2,
			"One of the twin deep gnome siblings. Female. Filed teeth, furtive eyes, whispers to her brother constantly. Seems nervous far beyond their situation warrants.",
			"SECRET: Wererat, infected along with her brother Turvy. Both are fighting their lycanthropy and terrified of discovery. Will transform within the first 20 days of travel - time this for maximum drama. Will eventually be forced to flee or reveal themselves to protect each other."},
		{"Turvy", "neutral", "Velkynvelve - prisoner", 2,
			"Topsy's twin brother. Male. Same filed teeth and nervous energy. The two are inseparable and share whispered conversations the party cannot hear.",
			"Same wererat secret as Topsy. If one is cornered or threatened, both transform. They are not evil - just terrified and ashamed. A sympathetic resolution (helping them control the curse) is possible but difficult."},
		{"Ilvara Mizzrym", "hostile", "Velkynvelve", 1,
			"Commander of Velkynvelve outpost. Tall, imperious, dressed in silken priestess robes with a black spider-silk scourge. Punishes failure and weakness with casual cruelty. Ambitious and frustrated at being stationed at a minor outpost.",
			"PRIMARY ANTAGONIST of Part 1. Pursues the party relentlessly through the Underdark after escape. CR 8. Spells: Clairvoyance, Dispel Magic, Faerie Fire, Levitate, Suggestion, Poison Spray. Uses her spider familiar for scouting. Has a complex rivalry with Jorlan. Confrontation at Blingdenstone is the natural endpoint."},
		{"Jorlan Duskryn", "neutral", "Velkynvelve", 0,
			"Former favored warrior of Ilvara, badly disfigured by a carrion crawler attack that left his hands nearly useless. His status has been destroyed - replaced by Shoor. Bitter, volatile, and potentially useful to the players.",
			"Can be turned against Ilvara - he has motive and knowledge of the outpost. May slip the party a key or a weapon if they approach him correctly and his hatred of Ilvara is cultivated. CR 5. Could become an uneasy NPC ally if played long-term."},
		{"Shoor Vandree", "hostile", "Velkynvelve", 0,
			"Young, vain drow warrior who has replaced Jorlan as Ilvara's favored. Overconfident, cruel, and eager to prove himself. Wears Jorlan's confiscated adamantine armor.",
			"Bully with something to prove. Will take the lead on recapturing escaped prisoners. CR 5. Can be used to drive a wedge between him and Jorlan in the early game. His overconfidence is his greatest weakness."},
		{"Themberchaud", "neutral", "Gracklstugh", 0,
			"The Wyrmsmith of Gracklstugh - an adult red dragon who has served as the city's living forge for decades. Enormous, vain, and increasingly suspicious that the Keepers of the Flame are keeping him isolated intentionally.",
			"CR 17. Can be an asset or a catastrophic threat. Wants information about the outside world and suspects he is being manipulated. Promises rewards for information. If he learns the full truth of his captivity - he may raze Gracklstugh. Handle his questline with care; his paranoia is justified."},
		{"Errde Blackskull", "neutral", "Gracklstugh", 0,
			"Captain of the Stone Guard, Gracklstugh's duergar military police. Pragmatic and dutiful. Assigned to investigate the demonic madness spreading through the city's populace.",
			"Hires the party to investigate the madness in the Blade Bazaar and the Whorlstone Tunnels. Neutral but will turn hostile if the party causes visible chaos. Represents the duergar's last attempt at rational governance before full demonic breakdown."},
		{"Ylsa Henstak", "neutral", "Gracklstugh", 0,
			"A duergar merchant who makes contact with the party in Gracklstugh. Member of a merchant clan worried about the chaos. Provides market access and city information.",
			"Quest hook: her clan has noticed strange stone growth patterns and violent outbursts from citizens. Access to Blade Bazaar vendors. Will pay for information and recovered goods."},
		{"Sovereign Phylo", "hostile", "Neverlight Grove", 8,
			"One of two myconid sovereigns of Neverlight Grove. Deeply corrupted by Zuggtmoy - utterly devoted to the Demon Queen of Fungi. Speaks in rapturous, honeyed terms about the coming \"Great Wedding.\" Eyes weep spores.",
			"CORRUPTED. Do not trust. Will try to draw party deeper into the grove with false hospitality. Wants Stool back for indoctrination. If players attend the \"wedding ceremony\" in the rotunda, they risk spore infection. Zuggtmoy herself may manifest here."},
		{"Sovereign Basidia", "ally", "Neverlight Grove", 1,
			"The second myconid sovereign. Deeply troubled by Phylo's behavior and the changes sweeping the grove. Seeks outsiders who can help - or confirm - her fears.",
			"KEY ALLY. Will communicate the truth about Phylo's corruption if the party reaches her privately. Can guide them to the Circle of Masters and the evidence of Zuggtmoy's presence. Will help the party escape if Phylo turns hostile."},
		{"Ploopploopeen", "neutral", "Sloobludop", 5,
			"Archpriest of the Sea Mother (Blibdoolpoolp) at Sloobludop. Seeks the party's help in a power struggle against his own daughter Bloppblippodd, who leads a rival cult worshiping a \"new god.\"",
			"Plans to use the party as a sacrifice to impress the Sea Mother - but is himself being outmaneuvered by his daughter. The party is caught between two factions of fanatics. Demogorgon's arrival will end both factions' plans violently."},
		{"Bloppblippodd", "hostile", "Sloobludop", 7,
			"Ploopploopeen's daughter. Leader of the new cult worshipping \"the Deep Father\" - which is in fact Demogorgon himself, accidentally summoned into existence by the kuo-toa's collective feverish belief.",
			"Her prayers are actively pulling Demogorgon toward Sloobludop. The climax of Ch. 3: Demogorgon rises from the Darklake and massacres everything. The party's only goal at that point is survival and escape."},
		{"Bruenor Battlehammer", "ally", "Gauntlgrym", 0,
			"The legendary dwarf king, restored to the throne of Gauntlgrym. Gruff, proud, and deeply suspicious of surface-dwellers bringing news of demon lords. Commands respect from every dwarf in the Underdark.",
			"Key Part 2 contact. Will grant resources and dwarven support for the assault on Menzoberranzan if convinced. Eldeth's shield (if returned) makes him immediately sympathetic to the party. His council is one of the primary locations for the lords' briefing on the situation."},
		{"Vizeran DeVir", "neutral", "Araj - Vizeran's Tower", 0,
			"Ancient, exiled drow archmage living alone in a remote Underdark tower. Brilliant, arrogant, thoroughly evil - but the only being with a plan to banish the demon lords. Hates Menzoberranzan with obsessive intensity.",
			"Part 2 primary quest-giver. His \"Dark Heart\" ritual requires gathering: the eye of a demon lord, Demogorgon's brain, Yeenoghu's heart, Zuggtmoy's spores, Juiblex's slime, a dark elf's exile-heart, and the Maze Engine components. His real goal is to destroy Menzoberranzan using the ritual - not just banish demons. The party must decide whether to trust him."},
		{"Gromph Baenre", "hostile", "Menzoberranzan", 2,
			"The Archmage of Menzoberranzan, the most powerful wizard in the Underdark. His reckless use of an artifact (the Demon Weave) is directly responsible for summoning all the demon lords.",
			"Reveals the cause of the demon incursion in the Menzoberranzan chapter. Not a combat encounter - more of a revelation scene. His guilt and denial are central to the final act's moral stakes. Does not believe he needs to answer to anyone for what he did."},
		{"Triel Baenre", "hostile", "Menzoberranzan", 0,
			"Matron Mother of House Baenre and effective ruler of Menzoberranzan. Cold, calculating, utterly ruthless. Views the demon lord incursion as a political problem to be managed, not a catastrophe to be stopped.",
			"Will meet with the party if they reach Menzoberranzan through proper channels. Wants the demon lords gone but is unwilling to appear weak or indebted to outsiders. Key to the political maneuvering in Ch. 15."},
		{"Grazilaxx", "neutral", "Mantol-Derith", 0,
			"A mind flayer who has rejected the illithid collective and travels alone as a merchant of information. Appears calm and businesslike. Trades knowledge for knowledge.",
			"Useful recurring information broker in Part 2. Knows the locations of several demon lords and the routes to Vizeran's tower. Cannot be fully trusted - will sell the party's information to others. CR 7."},
	}

	for _, r := range rows {
		_, err := conn.Exec(ctx,
			`INSERT INTO NPCS (name, disposition, location, madness, description, notes)
			 VALUES ($1,$2,$3,$4,$5,$6)`,
			r.name, r.disposition, nullableLocation(r.location), r.madness, r.desc, r.notes)
		if err != nil {
			log.Fatalf("npc %q: %v", r.name, err)
		}
	}
	log.Println("npcs seeded")
}

func seedEncounters(ctx context.Context, conn *pgx.Conn) {
	type row struct {
		name, location, status, enemies, ms, notes string
		difficulty                                 string
	}
	rows := []row{
		{"[Ch1] Escape from Velkynvelve", "Velkynvelve", "planned",
			"Ilvara Mizzrym (CR 8), Shoor Vandree (CR 5), 6× Drow (CR 1/4), 12× Quaggoth (CR 1), 2× Giant Spider",
			"Party reaches Level 2 upon successful escape.",
			"Party starts without equipment. Find gear in the equipment room; escape via webbed shafts or waterfall pool below (300 ft chasm). Jorlan can be turned - he has a master key. The rope bridge can be cut for a last-ditch escape. Dawn gong triggers Ilvara's full response. Let creativity shine.",
			"hard"},
		{"[Ch2] Encounter Check (d20 every 8 hrs)", "Underdark tunnels", "planned", "— roll below -", "",
			"Roll d20 every 8 hrs of travel. 1–13 = No encounter. 14–15 = Terrain (roll d20 on Terrain table). 16–17 = Creature (roll d20 on Creature table). 18–20 = Both terrain + creature.", "trivial"},
		{"[Ch2] Terrain Encounters (d20)", "Underdark tunnels", "planned", "Environmental", "",
			"1 Boneyard | 2 Cliff and ladder | 3 Crystal clusters | 4 Fungus cavern | 5 Gas leak | 6 Gorge | 7 High ledge | 8 Horrid sounds | 9 Lava swell | 10 Muck pit | 11 Rockfall | 12 Rope bridge | 13 Ruins | 14 Shelter | 15 Sinkhole | 16 Slime or mold | 17 Steam vent | 18 Underground stream | 19 Warning sign | 20 Webs", "trivial"},
		{"[Ch2] Creature Encounters (d20)", "Underdark tunnels", "planned", "See notes for full table", "",
			"1–2 Ambushers (reroll if resting) | 3 Carrion crawler | 4–5 Escaped slaves | 6–7 Fungi | 8–9 Giant fire beetles | 10–11 Giant \"rocktopus\" | 12 Mad creature | 13 Ochre jelly | 14–15 Raiders | 16 Scouts | 17 Society of Brilliance | 18 Spore servants | 19–20 Traders", "medium"},
		{"[Ch2] Ambushers - Cliff/Ledge Sub-table (d20)", "Underdark tunnels", "planned",
			"Chuul, Giant Spiders, Grell, Gricks, Orogs, Piercers, Umber Hulk", "",
			"1–2 1 chuul lurking in a pool | 3 1d6 giant spiders on walls/ceiling | 4–5 1 grell near high ceiling | 6–9 1d4 gricks in crevice | 10–15 1d4 orogs on ledges | 16–17 1d6 piercers (fake stalactites) | 18–20 1 umber hulk bursting from wall", "medium"},
		{"[Ch2] Escaped Slaves Sub-table (d4)", "Underdark tunnels", "planned", "Commoners / Goblins", "",
			"1 1d2 moon elf commoners | 2 1d3 shield dwarf commoners | 3 1d4 human commoners | 4 1d6 goblins", "trivial"},
		{"[Ch2] Fungi Sub-table (d6)", "Underdark tunnels", "planned", "Gas Spores, Shriekers, Violet Fungi", "",
			"1–2 1d4 gas spores | 3–4 1d4 shriekers | 5–6 1d4 violet fungi", "trivial"},
		{"[Ch2] Mad Creature Sub-table (d4)", "Underdark tunnels", "planned",
			"Deep Gnome / Drow / Duergar / Stone Giant (all under demonic madness)", "",
			"1 1 deep gnome | 2 1 drow | 3 1 duergar | 4 1 stone giant - all afflicted with short-term or long-term madness", "medium"},
		{"[Ch2] Raiders Sub-table (d6)", "Underdark tunnels", "planned", "Bandits, Goblins, or Orcs", "",
			"1–2 1d6 human bandits + 1 bandit captain | 3–4 2d4 goblins + 1 goblin boss | 5–6 1d6 orcs + 1 orc Eye of Gruumsh", "medium"},
		{"[Ch2] Scouts Sub-table (d6)", "Underdark tunnels", "planned", "Drow, Myconid Adults, or Dwarf Scouts", "",
			"1–2 1 drow | 3–4 1d4 myconid adults | 5–6 1d6 shield dwarf scouts", "trivial"},
		{"[Ch2] Society of Brilliance Sub-table (d10)", "Underdark tunnels", "planned", "Named NPCs (non-hostile)", "",
			"1–2 Y the derro savant | 3–4 Blurg the orog | 5–6 Grazilaxx the mind flayer | 7–8 Skriss the troglodyte | 9–10 Sloopidoop the kuo-toa archpriest", "trivial"},
		{"[Ch2] Spore Servants Sub-table (d10)", "Underdark tunnels", "planned",
			"Spore Servants (Drow / Duergar / Hook Horror / Quaggoth)", "",
			"1–3 1d4 drow spore servants | 4–6 1d6 duergar spore servants | 7–8 1d4 hook horror spore servants | 9–10 1d8 quaggoth spore servants", "medium"},
		{"[Ch2] Traders Sub-table (d4)", "Underdark tunnels", "planned", "Friendly (unless threatened)", "",
			"1 2d4 deep gnomes | 2 2d4 drow | 3 2d4 duergar | 4 2d4 kuo-toa", "trivial"},
		{"[Ch2] Silken Paths Encounters (d12)", "The Silken Paths", "planned",
			"Darkmantles, Drow+Quaggoth, Giant Spiders, Mimic, Spectator", "",
			"1 Cocooned lightfoot halfling (possibly alive) | 2 1d4 darkmantles | 3 1d4 drow + 1d4 quaggoth slaves | 4–8 2d4 giant spiders | 9 1 mimic | 10 1 spectator | 11–12 Web break (structural failure, DC 10 Acrobatics or fall)", "medium"},
		{"[Ch2] Boneyard Undead Sub-table (d20)", "Underdark tunnels", "planned", "Skeletons, Minotaur Skeletons", "",
			"1–14 No encounter | 15–18 3d4 skeletons | 19–20 1d3 minotaur skeletons", "medium"},
		{"[Ch2] Slime/Mold Sub-table (d6)", "Underdark tunnels", "planned", "Environmental hazard", "",
			"1–3 Patch of green slime | 4–5 Patch of yellow mold | 6 Patch of brown mold", "trivial"},
		{"[Ch2] Demon Encounters (d20) - near corrupted zones", "Underdark tunnels", "planned",
			"Barlgura, Dretches, Shadow Demons", "",
			"1–14 No encounter | 15–16 1 invisible barlgura | 17–18 3d4 dretches | 19–20 1d2 shadow demons", "hard"},
		{"[Ch3] Darklake Encounter Check (d20)", "The Darklake", "planned", "— roll below -", "",
			"Roll d20 every 8 hrs on the Darklake. 1–13 No encounter | 14–15 Terrain (roll d10) | 16–17 Creature (roll d12) | 18–20 Both terrain + creature.", "trivial"},
		{"[Ch3] Darklake Terrain (d10)", "The Darklake", "planned", "Environmental", "",
			"1 Collision | 2 Falls or locks | 3 Island | 4 Low ceiling | 5 Rockfall | 6 Rough current | 7 Run aground | 8 Stone teeth | 9 Tight passage | 10 Whirlpool", "trivial"},
		{"[Ch3] Darklake Creature Encounters (d12)", "The Darklake", "planned",
			"Aquatic Troll, Darkmantles, Duergar, Green Hag, Grell, Ixitxachitl, Kuo-toa, Merrow, Stirges, Quippers, Water Weird", "",
			"1 1 aquatic troll | 2 2d4 darkmantles | 3 1d4+2 duergar in a keelboat | 4 1 green hag | 5 1 grell | 6–7 1d6+2 ixitxachitl | 8 1d4 kuo-toa in a keelboat | 9 1d4 merrow | 10 3d6 stirges | 11 1 swarm of quippers | 12 1 water weird", "medium"},
		{"[Ch4] Gracklstugh City Encounters (d20)", "Gracklstugh", "planned",
			"Duergar guards, Derro rioters, Drow emissary, Orc mercenaries, Themberchaud", "",
			"1–2 Abusive duergar guards | 3–4 Deep gnome merchant* | 5–7 Derro rioters* | 8–9 Drow emissary* | 10–12 Duergar patrol | 13–14 Mad duergar | 15–16 Orc mercenaries* | 17–18 Slave caravan | 19 Steeder handlers | 20 Themberchaud. (*) = potential ally or info source", "medium"},
		{"[Ch4] Whorlstone Tunnels Encounters (d20)", "Gracklstugh", "planned",
			"Carrion Crawler, Demons, Flumph, Gray Ooze, Spore Servants, Grimlock, Swarm of Centipedes, Xorn", "",
			"1–10 No encounter | 11–12 1 carrion crawler | 13 Demon pack | 14 1 flumph | 15 1 gray ooze | 16 1d4 moldy quaggoth spore servants | 17 1d4 two-headed grimlocks | 18 1 swarm of centipedes | 19 1 xorn | 20 Yellow mold", "medium"},
		{"[Ch4] Darklake District Treasure (d4)", "Gracklstugh", "planned", "None - found treasure", "",
			"1 Humanoid skeleton wearing ring of water walking | 2 Zurkhwood chest with 1d6×100 gp + 1d6 gems (50 gp each) | 3 Skeleton in leather armor: rusted shortsword + rotted quiver with 1d20 +1 arrows OR pouch with 1d10 +2 sling stones OR case with 1d4 +3 crossbow bolts | 4 +1 shield (rusted and non-magical on repeat)", "trivial"},
		{"[Ch4] Themberchaud - Dragon Audience", "Gracklstugh", "planned",
			"Themberchaud (CR 17 Adult Red Dragon)",
			"Level 4 upon completing the Gracklstugh arc (negotiating with Dragon/Keepers).",
			"Themberchaud wants outside world information. Paranoid the Keepers are hiding things. Reward info generously. Do NOT let this become combat. If he goes berserk the city-escape scene is spectacular. Use his paranoia as a lever to engineer a distraction for the party's exit.", "deadly"},
		{"[Ch5] Neverlight Grove Encounters (d20)", "Neverlight Grove", "planned",
			"Nothics, Chasme, Vrock", "",
			"1–8 No encounter | 9–16 Fungi patch (roll d20 on Fungi table) | 17–18 1d4 nothics | 19–20 1 chasme on ceiling OR 1 vrock on ledge", "medium"},
		{"[Ch5] Neverlight Fungi/Creature Table (d20)", "Neverlight Grove", "planned",
			"Carrion Crawlers, Spore Servants, Myconid Adults, Giant Fire Beetles, Otyugh, Shriekers, Violet Fungi", "",
			"1 1d6 barrelstalks | 2 2d6 bluecaps | 3 1d3 carrion crawlers | 4 1d4 drow spore servants + 1d4 quaggoth spore servants | 5 Fire lichen near thermal vent | 6 3d6 giant fire beetles | 7 1d4 myconid adults | 8 1d6 nightlights | 9 1 otyugh under offal | 10 Brown mold patch | 11 1d4 awakened zurkhwood | 12 2d4 sheets of ripplebark | 13 1d4 shriekers | 14 2d4 timmasks | 15 1d6 tongues of madness | 16 2d6 torchstalks | 17 2d6 trillimacs | 18 1d4 violet fungi | 19 2d4 waterorbs near spring | 20 1d4 zurkhwoods", "medium"},
		{"[Ch5] Neverlight Grove - Zuggtmoy's Wedding", "Neverlight Grove", "planned",
			"Zuggtmoy (CR 23), corrupted Myconid Sovereigns, Myconid Adults",
			"Level 5 upon escaping the Grove and warning Basidia.",
			"Zuggtmoy is performing her wedding to Araumycos. DC 15 Con save vs. spore infection (long-term madness). Primary goal: escape with evidence, warn Basidia. Basidia can seal tunnels. Do not fight Zuggtmoy directly.", "deadly"},
		{"[Ch6] Blingdenstone Encounters (d20)", "Blingdenstone", "planned",
			"Animated Armor, Cave Badgers, Fiendish Spiders, Ghost, Mephits, Oozes, Wererats, Xorn", "",
			"1–10 No encounter | 11 1d4+1 animated drow statues | 12 1d4+2 cave badgers | 13 Dungeon hazard (roll d6) | 14 Elemental vagabonds | 15 1d4+2 fiendish giant spiders | 16 1 ghost | 17 Mephit gang | 18 Roaming ooze (roll d4) | 19 1d4+1 svirfneblin wererats | 20 1 xorn", "medium"},
		{"[Ch6] Blingdenstone Dungeon Hazards (d6)", "Blingdenstone", "planned", "Environmental", "",
			"1–3 Patch of brown mold | 4–5 Patch of green slime | 6 Patch of yellow mold", "trivial"},
		{"[Ch6] Roaming Ooze Sub-table (d4)", "Blingdenstone", "planned",
			"Black Pudding, Gelatinous Cube, Gray Ooze, Ochre Jelly", "",
			"1 1 black pudding | 2 1 gelatinous cube | 3 1d4+1 gray oozes (one is psychic variant) | 4 1d2 ochre jellies", "medium"},
		{"[Ch6] Pudding King's Domain Oozes (d6)", "Blingdenstone", "planned",
			"Black Puddings, Gelatinous Cubes, Gray Oozes, Ochre Jellies", "",
			"1–2 1 black pudding + 2 gray oozes | 3 1 gelatinous cube + 1 ochre jelly | 4–5 3 gray oozes + 1 ochre jelly | 6 2 black puddings", "hard"},
		{"[Ch6] The Pudding King's Court", "Blingdenstone", "planned",
			"The Pudding King (CR 4), 2× Black Pudding, 1× Elder Ooze",
			"Level 6 upon defeating the Pudding King and cleansing the city.",
			"Mad deep gnome consumed by Juiblex. Speaks in rhyming couplets. His \"children\" (the two black puddings) fight alongside him. Defeating him ends the ooze infestation. Rewards from gnomes: gem cache, safe passage, surface route info.", "hard"},
		{"[Ch8] Gauntlgrym Approach Encounters (d20)", "Gauntlgrym", "planned",
			"Cloaker, Driders, Dwarf Ghost, Veterans, Priest, Gargoyles, Gricks, Rust Monsters", "",
			"1–12 No encounter | 13 1 cloaker | 14 1d2 driders | 15 1 dwarf ghost (friendly unless attacked) | 16 Patrol: 6 shield dwarf veterans | 17 1 shield dwarf priest + 1d4+1 acolytes | 18 1d6+1 gargoyles | 19 1 grick alpha + 1d4+1 gricks | 20 1d4 rust monsters", "medium"},
		{"[Ch8] Gauntlgrym Interior Encounters (d20)", "Gauntlgrym", "planned",
			"Doppelgangers, Fire Elementals, Magmins, Salamander, Fire Snakes, Spirit Naga, Troglodytes, Wraith + Specters", "",
			"1–14 No encounter | 15 1d4+1 doppelgangers (disguised as dwarves) | 16 1d2 fire elementals + 3d6 magmins | 17 1 salamander + 1d4+1 fire snakes | 18 1 spirit naga | 19 3d6 troglodytes | 20 1 wraith leading 1d6+1 specters", "hard"},
		{"[Ch10] Part 2 Travel Event Table (d20)", "Underdark tunnels", "planned", "Various - see notes", "",
			"1–2 Battle aftermath | 3–6 Creature encounter (roll d20 on creature table) | 7–9 Demon encounter | 10–11 Discipline problem (NPC morale) | 12–13 Disease | 14–15 Madness | 16–17 Poisoned NPCs | 18–19 Spoiled supplies | 20 Vanishing NPCs", "medium"},
		{"[Ch10] Part 2 Demon Encounters (d20)", "Underdark tunnels", "planned",
			"Barlguras, Chasmes, Hezrous, Shadow Demons, Vrocks, Juiblex", "",
			"1–4 1d4 barlguras | 5–8 1d4 chasmes | 9–10 1d2 hezrous | 11–14 1d4 shadow demons | 15–18 1d3 vrocks | 19–20 Juiblex (CR 23 - survival only)", "hard"},
		{"[Ch13] Wormwrithings Encounters (d20)", "The Wormwrithings", "planned",
			"Drider, Drow, Dwarf Commoners, Ettins, Flumphs, Gricks, Purple Worm, Troglodytes, Trolls, Umber Hulk", "",
			"1–10 No encounter | 11 1 drider | 12 Drow hunting party | 13 3d6 dwarf commoners | 14 1d6 ettins | 15 3d6 flumphs | 16 Grick nest | 17 1 purple worm | 18 2d6 troglodytes | 19 1d4 trolls | 20 1 umber hulk", "hard"},
		{"[Ch13] Wormwrithings Tunnel Encounters (d20)", "The Wormwrithings", "planned",
			"Giant Spiders, Purple Worm", "",
			"1–15 No encounter | 16–18 1 giant spider (from area 11) | 19–20 1 purple worm", "hard"},
		{"[Ch14] Labyrinth Encounters (d20)", "The Labyrinth", "planned",
			"Behir, Flumphs, Gnolls, Grells, Hezrous, Manes, Minotaurs, Monodrone, Quaggoths, Shriekers", "",
			"1–10 No encounter | 11 1 behir | 12 2d4 flumphs | 13 Gnoll pack | 14 1d4 grells | 15 1d4 hezrous | 16 4d8 manes | 17 2d4 minotaurs | 18 1 monodrone | 19 2d6 quaggoths | 20 1d4 shriekers", "hard"},
		{"[Ch14] Labyrinth Hunting Ground (d20) - Baphomet's territory", "The Labyrinth", "planned",
			"Corpse/bones (hazard), Gnolls, Minotaurs, Baphomet", "",
			"1–10 No encounter | 11–12 Corpse | 13–14 Gnawed bones | 15–17 2d4 gnolls | 18–20 1d6 minotaurs. Baphomet patrols - players hear bellowing, getting closer each hour.", "deadly"},
		{"[Ch14] The Maze Engine Activation", "The Labyrinth", "planned",
			"Baphomet (CR 23), 4× Minotaur",
			"Level 13 upon securing the Maze Engine components.",
			"Activating the Maze Engine triggers a Wild Magic surge (d10 table). Can banish a demon lord if triggered correctly during the final ritual. Baphomet is hunting - escape is a valid strategy. Minotaurs guard the engine's antechamber.", "deadly"},
		{"[Ch15] Menzoberranzan Entry Patrols (d20)", "Menzoberranzan", "planned",
			"Drow Patrol A–D (escalating size)", "",
			"1–10 No encounter | 11–14 Drow patrol A (4 drow) | 15–17 Drow patrol B (6 drow + 1 mage) | 18–19 Drow patrol C (8 drow + 1 elite warrior) | 20 Drow patrol D (10 drow + 1 priestess)", "hard"},
		{"[Ch15] Menzoberranzan Bazaar Encounters (d20)", "Menzoberranzan", "planned",
			"Bugbears, Driders, Drow, Spore Servants, Goblins, Intellect Devourers", "",
			"1–2 2d4 bugbears | 3–4 Clandestine meeting | 5 1d4 driders | 6–10 Drow patrol | 11–12 1d4+1 drow spore servants | 13–14 Escaped slaves | 15–16 1d4+1 goblins | 17–19 (empty - reroll) | 20 1d4 intellect devourers", "medium"},
		{"[Ch15] Menzoberranzan Narbondellyn District (d20)", "Menzoberranzan", "planned",
			"Bugbears, Drow, Giant Wolf Spiders", "",
			"1–2 1d4+2 bugbears | 3–8 Drow adolescents | 9–10 Drow pickpocket | 11–12 3d6 giant wolf spiders | 13–14 Infected drow | 15–16 Mad drow | 17–18 1 shield dwarf berserker | 19–20 Svirfneblin lure", "medium"},
		{"[Ch15] Menzoberranzan Eastmyr Slave District (d20)", "Menzoberranzan", "planned",
			"Drow Patrols, Spore Servants, Giant Wolf Spiders", "",
			"1–5 Drow patrol | 6–8 2d4+2 drow spore servants | 9–10 Escaped slaves | 11–14 1d6+2 giant wolf spiders | 15–20 Slave farmers", "medium"},
		{"[Ch15] Menzoberranzan West Wall District (d20)", "Menzoberranzan", "planned",
			"Bregan D'aerthe Spy, Drow Patrol, Drow Priestess, Spiders", "",
			"1–4 Bregan D'aerthe spy | 5–8 Drow foot patrol | 9–12 Drow priestess of Lolth | 13–16 Spider nest | 17–20 Statue of Lolth", "hard"},
		{"[Ch15] Menzoberranzan Duthcloim District (d20)", "Menzoberranzan", "planned",
			"Bregan D'aerthe, Cult of Y, Drow Patrol", "",
			"1–5 Bregan D'aerthe spy | 6–10 Cult of 'Y' (Demogorgon cultists) | 11–15 Drow foot patrol | 16–20 Scroll from Narbondel's Shadow (encoded message - intelligence opportunity)", "hard"},
		{"[Ch15] Menzoberranzan Tier Breche (d20)", "Menzoberranzan", "planned",
			"Black Pudding, Spore Servants, Elite Drow, Giant Spiders, Gricks, Stirges, Violet Fungi", "",
			"1–10 No encounter | 11 1 black pudding | 12 3d6 drow spore servants | 13 Elite drow foot patrol | 14 Exotic fungi (roll d6) | 15 1d4 giant spiders | 16 1d4 gricks | 17 Hunting party | 18 1 shrieker | 19 3d6 stirges | 20 1d4+1 violet fungi", "hard"},
		{"[Ch15] Menzoberranzan Exotic Fungi (d6)", "Menzoberranzan", "planned", "None", "",
			"1 1d6 nightlights (50% unlit) | 2 2d6 Nilhogg's noses | 3 1d6 patches of ormu | 4 2d6 timmasks | 5 1d6 tongues of madness | 6 3d6 torchstalks", "trivial"},
		{"[Ch15] Menzoberranzan Qu'ellarz'orl (Noble Tier, d20)", "Menzoberranzan", "planned",
			"Beholder, Bregan D'aerthe Mercenaries, Elite Drow, Noble Entourage", "",
			"1–3 Beholder | 4–7 Bregan D'aerthe mercenaries | 8–12 Elite drow patrol | 13–16 Noble entourage | 17–20 Statue of Lolth", "deadly"},
		{"[Ch15] Menzoberranzan Donigarten (Farms, d20)", "Menzoberranzan", "planned",
			"Elite Drow Patrol, Gargoyles, Giant Wolf Spiders, Groundskeepers", "",
			"1–5 Elite drow patrol | 6–8 2d4 gargoyles | 9–10 1d6+2 giant wolf spiders | 11–14 Groundskeepers | 15–20 Slave parade", "medium"},
		{"[Ch15] Menzoberranzan Sorcere (Mage Tower, d20)", "Menzoberranzan", "planned",
			"Drow Acolytes, Drow Mages, Drow Warriors", "",
			"1–6 Drow acolytes | 7–12 Drow mages | 13–20 Drow warriors", "hard"},
		{"[Ch15] Menzoberranzan Arach-Tinilith (Spider Temple, d20)", "Menzoberranzan", "planned",
			"Bandersnatches, Bregan D'aerthe, Drow Foot Patrol", "",
			"1–3 Bandersnatches | 4–7 Bregan D'aerthe spy | 8–14 Drow foot patrol | 15–17 Slave abuse (social encounter) | 18–20 Statue of Lolth", "hard"},
		{"[Ch15] Menzoberranzan Web Tunnels (d20)", "Menzoberranzan", "planned",
			"Drow Mages, Giant Spider, Quasit, Shadow Demon, Slaves, Succubus/Incubus", "",
			"1–6 1d4 drow mages | 7–8 1 giant spider | 9–10 1 invisible quasit | 11–12 1 mad drow mage | 13–14 1 shadow demon | 15–18 1d4 slaves | 19–20 1 succubus or incubus", "hard"},
		{"[Ch16] Araumycos Travel Encounters (d20)", "Araumycos", "planned",
			"Death Tyrant, Demons, Gnolls, Gricks, Myconids, Oozes, Two-headed Trolls", "",
			"1 Death tyrant | 2–6 Demons (roll d12) | 7–8 Gnoll pack | 9–10 Gricks | 11–14 Myconid parade | 15–18 Oozes | 19–20 Two-headed trolls", "hard"},
		{"[Ch16] Fetid Wedding Demon Sub-table (d12)", "Araumycos", "planned",
			"Barlguras, Chasmes, Hezrous, Manes, Nalfeshnee, Vrocks", "",
			"1–2 2d4 barlguras | 3–4 2d4 chasmes | 5–6 1d4 hezrous | 7–8 1d100 manes | 9–10 1 nalfeshnee | 11–12 2d4 vrocks", "deadly"},
		{"[Ch16] Fetid Wedding Fungi Hazards (d20)", "Araumycos", "planned",
			"Gas Spores, Violet Fungi, Edible Fungi, Exotic Fungi", "",
			"1–5 No encounter | 6–10 Fungi (roll d6) | 11–14 Mold pit | 15–17 Myconid parade | 18–20 Oozes", "trivial"},
		{"[Ch16] Fetid Wedding Fungi Detail (d6)", "Araumycos", "planned", "Minor", "",
			"1 1d6 gas spores | 2 1d6 violet fungi | 3–4 3d6 edible fungi (choose variety from ch2) | 5–6 3d6 exotic fungi (choose variety from ch2)", "trivial"},
		{"[Ch17] Demon Sortie Waves (d4)", "Menzoberranzan", "planned",
			"Barlguras, Chasmes, Hezrous, Vrocks", "",
			"Waves hit every 2 rounds during the ritual: 1 = 4 barlguras | 2 = 4 chasmes | 3 = 2 hezrous | 4 = 3 vrocks", "deadly"},
		{"[Ch17] Battle Chaos Events (d6)", "Menzoberranzan", "planned", "Environmental damage only", "",
			"1–2 Explosion: DC 13 Dex save, 3d6 fire, 20 ft radius | 3–4 Flying Debris: DC 13 Dex save, 3d6 bludgeoning + prone, 30 ft radius | 5–6 Close Call: DC 13 Dex save, 3d6 bludgeoning to one character", "deadly"},
		{"[Ch17] Dark Heart Ritual - Final Battle", "Menzoberranzan", "planned",
			"Demogorgon (CR 26), Juiblex / Zuggtmoy / Baphomet / Yeenoghu avatars (CR 12 each)",
			"Campaign finale.",
			"Vizeran's ritual draws all demon lords to one place. Protect the ritual site for 10 rounds while Vizeran casts. Each demon lord attacks one PC per round. On round 10 the ritual fires - all banished. If party discovered Vizeran's true plan: DC 20 Arcana to modify the ritual to banish without destroying Menzoberranzan. Describe each lord's banishment individually.", "deadly"},
		{"[Story] Drow Pursuit Escalation", "Underdark tunnels", "planned",
			"Ilvara Mizzrym + 4 drow + 2 giant spiders", "",
			"Roll 1d20 every 3 days of travel. On 15+ the pursuit catches up. Ilvara grows more reckless each failure - escalating danger but also more intelligence opportunities if a scout is captured.", "medium"},
		{"[Story] Arrival at Sloobludop", "Sloobludop", "planned",
			"Ploopploopeen cultists (6× kuo-toa + 1 archpriest), Bloppblippodd cultists (8× kuo-toa + 1 archpriest)",
			"Level 3 upon reaching Sloobludop or the first major settlement.",
			"Party is intended as sacrifices. Mid-ceremony Bloppblippodd's faction attacks. Before anyone wins - Demogorgon arrives. Phase 3 = PURE SURVIVAL. Demogorgon is unkillable. Use him to destroy the town around them. Goal: reach the boats.", "hard"},
		{"[Story] Demogorgon Rises at Sloobludop", "Sloobludop", "planned",
			"Demogorgon (CR 26) - DO NOT run as normal combat", "",
			"CRITICAL SET PIECE. Gaze attack (DC 23 Wis, madness on fail) fires every round at random targets. Buildings collapse. Goal: escape on boats. Anyone who stays to fight takes 70+ damage per round. Give each PC a personal escape moment.", "deadly"},
		{"[Story] Ilvara's Final Pursuit", "Blingdenstone", "planned",
			"Ilvara Mizzrym (CR 8, wounded), Jorlan Duskryn (CR 5), 4× drow",
			"Level 7 or 8 upon reaching the Surface (Part 1 Conclusion).",
			"Final Part 1 confrontation. Ilvara is furious and reckless. Jorlan may switch sides if his resentment was cultivated. Ilvara uses Suggestion on strongest PC. Defeating her ends drow pursuit permanently.", "hard"},
		{"[Story] Mantol-Derith Faction War", "Mantol-Derith", "planned",
			"Drow faction, Duergar faction, Svirfneblin faction, Mind Flayer faction (all hostile to each other)",
			"Level 9 upon reaching Vizeran's Tower.",
			"Four factions tearing each other apart. Party must find Harper agent Khalessa Draga (poisoned). First PC to draw a weapon triggers everyone. Each faction has a piece of the truth about the demonic corruption.", "hard"},
	}

	for _, r := range rows {
		_, err := conn.Exec(ctx,
			`INSERT INTO Encounters (name, location, difficulty, status, enemies, levelup, notes)
			 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
			r.name, nullableLocation(r.location), difficultyInt(r.difficulty),
			r.status, r.enemies, r.ms != "", r.notes)
		if err != nil {
			log.Fatalf("encounter %q: %v", r.name, err)
		}
	}
	log.Println("encounters seeded")
}

func seedEvents(ctx context.Context, conn *pgx.Conn) {
	type row struct{ title, category, desc string }
	rows := []row{
		{"Gromph's Ritual Shatters the Wards", "story",
			"Gromph Baenre, Archmage of Menzoberranzan, performs a catastrophic ritual using the demon-tainted Demon Weave. Every ward imprisoning demon lords in the Abyss is weakened simultaneously. Demogorgon, Baphomet, Juiblex, Zuggtmoy, Orcus, Yeenoghu, and others pour through into the Underdark. This is the event that caused everything - the party does not learn the full truth until Chapter 15."},
		{"Party Captured at Velkynvelve", "story",
			"The player characters are captured by a drow patrol and imprisoned at the Velkynvelve outpost under Mistress Ilvara Mizzrym's command. Chapter 1 begins here."},
		{"Demogorgon Destroys Sloobludop", "demon",
			"Chapter 3 climax. Demogorgon physically rises from the Darklake at Sloobludop, destroying the settlement and massacring its kuo-toa population in minutes. The party flees on boats. This is the players' first direct encounter with the scale of the demon lord threat."},
		{"Gracklstugh Falls to Madness", "demon",
			"Chapter 4. Demogorgon's influence drives waves of madness through Gracklstugh's duergar population. Stone giants go berserk. Derro cultists perform open rituals. The city's rigid order begins to crack. Themberchaud's patience with the Keepers of the Flame may snap."},
		{"Zuggtmoy's Wedding in Neverlight Grove", "demon",
			"Chapter 5. Zuggtmoy is conducting a ritual \"wedding\" to Araumycos, the continent-spanning fungal organism. If completed, she will possess Araumycos and become an unstoppable force. The myconid population of Neverlight Grove has been almost entirely subverted."},
		{"Party Reaches the Surface", "story",
			"End of Part 1 (Chapter 9). After weeks in the Underdark, the party escapes to the surface world. They must deliver news of the demon lord incursion to the Lords' Alliance or appropriate authorities. This is the chapter break between Part 1 and Part 2."},
		{"Council of Waterdeep - Call to Return", "faction",
			"Chapter 10. Surface authorities - Lords of Waterdeep, Order of the Gauntlet, Harpers, Emerald Enclave, Zhentarim - convene an emergency council. After debriefing the party, they are commissioned to return to the Underdark with resources and allies to address the demon lord threat."},
		{"Mantol-Derith Faction War", "faction",
			"Chapter 11. The neutral trading ground of Mantol-Derith has collapsed into violence as Demogorgon's madness infects each faction. The party must navigate chaos, rescue Harper agent Khalessa Draga, and find the path to Vizeran's tower."},
		{"Vizeran's Bargain Accepted", "story",
			"Chapter 12. The party meets Vizeran DeVir and learns of the Dark Heart ritual. He provides the components list and the knowledge to execute it. His assistance is invaluable - and comes with a hidden cost."},
		{"Vizeran's True Plan Revealed", "story",
			"Late Chapter 12 / Chapter 15. Investigation or magical means reveal that Vizeran's ritual, as written, will not merely banish the demon lords - it will release the accumulated energy of their banishment directly into Menzoberranzan, likely killing tens of thousands. The party must decide whether to help him, modify the ritual, or find another way."},
		{"Dark Heart Ritual - Demon Lords Banished", "story",
			"Chapter 17 (Finale). The ritual is completed at Menzoberranzan. Demon lords are drawn to one place and banished back to the Abyss. The exact outcome depends on whether the party discovered and addressed Vizeran's true plan. Either way, the campaign ends here."},
	}

	for _, r := range rows {
		_, err := conn.Exec(ctx,
			`INSERT INTO Events (title, category, description) VALUES ($1,$2,$3)`,
			r.title, r.category, r.desc)
		if err != nil {
			log.Fatalf("event %q: %v", r.title, err)
		}
	}
	log.Println("events seeded")
}
