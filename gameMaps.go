package main

import (
	"strings"
)

type Marker struct {
	I     int
	X, Y  int
	Title string
	Body  string
}

type GameMap struct {
	ID      string
	Img     string
	VB      string
	Markers []Marker
}

func (gm GameMap) Width() string {
	p := strings.Fields(gm.VB)
	if len(p) == 4 {
		return p[2]
	}
	return "100%"
}

func (gm GameMap) Height() string {
	p := strings.Fields(gm.VB)
	if len(p) == 4 {
		return p[3]
	}
	return "100%"
}

var gameMaps = []GameMap{
	{
		ID:  "underdark",
		Img: "./images/underdark.webp",
		VB:  "0 0 2404 3000",
	},
	{
		ID:  "sloobludop",
		Img: "./images/sloobludop.webp",
		VB:  "0 0 680 554",
		Markers: []Marker{
			{
				I: 1, X: 663, Y: 286,
				Title: "North Gate",
				Body:  "Heavy netting reinforced with bone hooks encloses the village to the north. A central gate allows safe passage through. Crawling through the netting requires DC 15 Dexterity (Acrobatics) — failure deals 1d8 piercing damage and leaves the creature restrained (DC 12 Strength to break free). Four kuo-toa whips guard the gate and challenge all who approach. Escorted parties pass unchallenged. Unescorted captives are distributed by faction: d20 roll of 1–7 sends them to area 3, 8–18 to area 4, 19–20 causes the guards to brawl immediately over the right to claim them.",
			},
			{
				I: 2, X: 23, Y: 271,
				Title: "South Gate",
				Body:  "The southern boundary of Sloobludop mirrors the north — the same bone-hooked netting wall, the same central gate, and four kuo-toa whips standing watch. The same DC 15 Acrobatics check applies to any who attempt to crawl through. Unescorted captives face the same faction-split dice. A 19–20 result means guards of opposing cults come to immediate blows — potentially exploitable chaos for a cunning party.",
			},
			{
				I: 3, X: 471, Y: 146,
				Title: "Docks",
				Body:  "Six kuo-toa keelboats moor at the water's edge of the Darklake. Ferry service is available but no vessel departs without auguries — a half-hour bone-casting ritual. Results: 1–8 favorable, 9–18 unfavorable, 19–20 unclear (archpriest of the Deep Father must be consulted). Five patrols of kuo-toa monitors guard the platforms at all hours. Stealing a boat requires DC 16 Stealth — the kuo-toa can sense invisible creatures and will pursue any thieves to reclaim them as offerings.",
			},
			{
				I: 4, X: 391, Y: 328,
				Title: "Shrine of the Sea Mother",
				Body:  "A nine-foot idol stands at the center: a crudely carved wooden body capped with the severed head and claws of a giant albino crayfish, lashed on with gut cord and reeking of rot. Four kuo-toa monitors guard it at all times alongside 2d4 circling worshippers. Glooglugogg — kuo-toa whip and loyal son of Ploopploopeen — tends the shrine and is deeply hostile to outsiders. Ploopploopeen's adjacent quarters hold: 500 cp, 2,000 sp, 150 gp, 27 pp, matched pearls worth 1,000 gp, two potions of healing, one potion of water breathing, and a scroll of light.",
			},
			{
				I: 5, X: 324, Y: 240,
				Title: "Altar of the Deep Father",
				Body:  "The idol: a manta ray hide stretched between two support poles, a dead manta ray pinned at its center, two dead octopuses draped over the top with tentacles arranged in dark display. The broad stone altar beneath runs dark with accumulated blood. Klibdoloogut (kuo-toa whip) stands before it; six kuo-toa tend the offerings. Any non-kuo-toa humanoid encountered here is immediately seized for sacrifice. Archpriest Bloppblippodd — demon-touched, Ploopploopeen's estranged daughter — lives in an adjacent hovel. A duergar arms smuggler named Hemeth lies bound against the far wall, willing to bargain for his freedom.",
			},
		},
	},
	{
		ID:  "gracklstugh",
		Img: "./images/gracklstugh.webp",
		VB:  "0 0 4050 3300",
		Markers: []Marker{
			{
				I: 1, X: 1311, Y: 740,
				Title: "The Shattered Spire",
				Body:  `A broken stalagmite juts out from the Darklake about forty feet from the shore, forming the foundation of a tavern built with fungi stalks in a manner similar to a log cabin. A bridge woven of rothé wool allows patrons to cross the water to visit. Where Werz meets with them. A bar fight breaks out as they leave, neither duergar remembers why they started fighting`,
			},
			{
				I: 2, X: 2200, Y: 750,
				Title: "Darklake Docks",
				Body:  `These busy docks are used primarily by flat-bottomed rafts made of zurkhwood and lacquered puffball floats. Some of these ramshackle barges come with oars or paddle wheels. The rafts look ungainly, but each can carry tons of trade goods.`,
			},
			{
				I: 3, X: 3268, Y: 650,
				Title: "Darklake Brewery",
				Body:  `This huge, ramshackle brewery is built of stone blocks stacked to make walls between the petrified stems of a small forest of gigantic mushrooms. Big copper vats steam within, filling the air with a heavy, yeasty stink. Dozens of copper kegs stand nearby, and burly gray dwarves swarm over the place, mashing fungus, mixing fermenting masses, and filling casks with freshly brewed ale. This complex is the workplace and home of Clan Muzgardt, the duergar clan in charge of brewing Darklake Stout and in control of the brewing and importation of other spirits. Non-duergar aren't welcome inside the brewery.`,
			},
			{
				I: 4, X: 2678, Y: 889,
				Title: "Blade Bazaar",
				Body:  `This marketplace is named after the most abundant goods the duergar offer, but the shops here sell almost everything available in the city, along with stalls set up by visiting merchants. The din of people arguing, mostly in Dwarvish, nearly drowns out the hammering coming from the city's forges, and the crowds here offer a good chance to slip away from pursuers.Characters can unload some of the treasures they might be carrying. Nonmagical weapons, armor, and shields can be purchased in the Blade Bazaar.`,
			},
			{
				I: 5, X: 2177, Y: 1002,
				Title: "Darklake District",
				Body:  `The Darklake District gives an illusion of openness. The streets are relatively wide to allow for merchant carts and wagons to pass, and the buildings aren't as crowded around stalagmites as in the southern districts. Openness doesn't mean welcoming, however. The duergar who ply their trades here are wary of all the foreigners confined by law to this part of the city. A wave of heat slams against you as an acrid smog rises to choke the air out of your lungs. The Darklake spreads out beyond a jumble of buildings and streets, reflecting the lights of countless fires burning across the city within hollowed-out columns and stalagmites. Though the streets are crowded, you move easily within the surging throng of buyers, merchants. You aren't the only outsiders here, as you spy drow, svirfneblin, derro, orcs, and other races in the crowds. The shouting of people blends with the sound of distant hammering to create a constant, distracting din. Behind the forbidding walls separating the Darklake District from the rest of the city stand the docks, markets, and shops where Gracklstugh's commerce and trade are conducted. The many duergar merchants—along with drow, svirfneblin, orcs, and others—pay little attention to the characters unless they are looking to do business.`,
			},
			{
				I: 6, X: 3100, Y: 964,
				Title: "Overlake Hold",
				Body:  `Dunglorrin Torune, which translates as Overlake Hold, is a fortress and temple dedicated to Laduguer carved into the heart of a massive stalagmite on the shore of the Darklake. It is also the home of the Deepking and the center of government. Dunglorrin Torune bristles with forge chimneys from which smoke billows and ledges from which catapults can hurl stones at waterborne invaders.`,
			},
			{
				I: 7, X: 2671, Y: 1080,
				Title: "Ghorlborn's Lair",
				Body:  `This inn is the only establishment in Gracklstugh that accepts non-duergar guests. "Ghohlbrorn" means "bulette" in Dwarvish, and the inn is built inside a small cavern complex beneath the Blade Bazaar at the northern end of the Darklake District. Its halls are cold and damp. A central chamber serves as a dining room, branching out into different small, twisting halls along which the rooms are excavated. It's dark, cramped, and uncomfortable, but safe and defensible.`,
			},
			{
				I: 8, X: 638, Y: 1185,
				Title: "Halls of Sacred Scrolls",
				Body:  "The Halls of Sacred Spells comprise a temple of Diirinka carved into a stalagmite in Northfurrow District. Here, the derro Council of Savants meets and plots, living in luxurious quarters and hiding such opulence from their fellow derro. All areas of the Halls of Sacred Spells except the central worship chamber are forbidden to derro who aren't savants. Duergar don't enter this place, whose main doors are false and carved into the rock. The savants enter and leave using spells such as dimension door and passwall, while lesser derro access the worship chamber through secret tunnels from the West Cleft.",
			},
			{
				I: 9, X: 694, Y: 1555,
				Title: "West Cleft District",
				Body:  `The West Cleft District was the original abode for the city's derro slaves and remains a dark and dangerous ghetto. A residential and light-industry district on the western side of Laduguer's Furrow. Home to lower-ranking duergar craftspeople and their households. The district is quieter than the Darklake District but not without violence and madness.`,
			},
			{
				I: 10, X: 2000, Y: 1850,
				Title: "Laduguer's Furrow",
				Body:  `Long ago, an earthquake split the cavern that houses Gracklstugh, leaving a rift nearly two hundred feet deep and five hundred feet wide. Laduguer's Furrow has a packed-gravel floor and extends roughly a quarter mile beyond the natural walls of the city in both directions. Each end of the rift has a steeply sloping floor, carved with a set of stairs and a wide ramp for both pedestrians and wagons. Vents along the walls release potent gases that sappers of Clan Xardelvar tap for industrial applications, including the crafting of the magical flame lances used by xarrorn warriors. The chasm is Gracklstugh's main residential zone, with homes built on the top part of its north and south sides. Outsiders are normally forbidden from this area.`,
			},
			{
				I: 11, X: 3272, Y: 2078,
				Title: "East Cleft District",
				Body:  `The industrial quarter east of Laduguer's Furrow, home to foundries, smelters, and the workshops of weapon-crafting clans. The air here is thick with smoke and the ring of hammers never ceases. Clan Steelshadow maintains several private forges in this district. The heat is oppressive even by Underdark standards. The East Cleft District was more recently settled after the derro earned their freedom, though it is only slightly less rough than West Cleft.`,
			},
			{
				I: 12, X: 937, Y: 2234,
				Title: "Cairngorm Cavern",
				Body:  `A long tunnel opens in Southfurrow District, extending several hundred feet and into the home of the stone giants of Clan Cairngorm. The tribe is named after the ancient oath of fealty their ancestors swore to the bearers of the Cairngorm Crown, the traditional regalia of Deepkingdom monarchs. The giants lead simple, uncomplicated lives, and their dwellings reflect this. The stone giants value their privacy, and duergar are normally not allowed inside Cairngorm Cavern. An exception is made for the Deepking, who holds meetings here with the giants' leader, Stonespeaker Hgraam, when necessary.`,
			},
			{
				I: 13, X: 2749, Y: 2384,
				Title: "Hold of the Deep King",
				Body:  `The Hold of the Deepking stands south of Laduguer's Furrow and north of Themberchaud's lair. The Hold of the Deepking is a dark and foreboding edifice lodged between two great columns that rise up into thick clouds of smoke that conceal the cavern ceiling. Giant basalt braziers filled with molten lava bathe the palace facade in a hellish glow, and the thick stone walls bristle with iron turrets and battlements. There appears to be no one guarding the palace. This is an illusion. All the palace guards are invisible, and characters who observe the palace for some time can hear duergar guards in heavy armor marching to and fro. 200 invisible guards protect and 50 stand at attention in perfect rows before the palace, ready to cut down anyone who approaches the palace gates unescorted. Another 50 watch from the turrets and battlements. 100 more in the palace are the Deepking's honor guard.`,
			},
			{
				I: 14, X: 2600, Y: 2967,
				Title: "Themberchaud's Cavern",
				Body:  `At the far southeast corner of Gracklstugh's cavern, the entrance to Themberchaud's lair is guarded by the Keepers of the Flame. Not that anyone would be foolish enough to trespass into the Wyrmsmith's home, but ever since the Gray Ghosts stole a red dragon egg meant to hatch Themberchaud's successor, the Keepers aren't taking any chances. For some time now, the Keepers have been actively seeking capable mercenaries in Gracklstugh and taking any opportunity to press them into service. If an agent of the order intervened in the characters' arrest, the leader of the Keepers of the Flame—Gartokkar Xundorn—is notified by magical messaging. He waits for the characters as they are brought to the dragon's cavern—but Themberchaud is watching too.`,
			},
		},
	},
	{
		ID:  "whorlstoneTunnels",
		Img: "./images/whorlstoneTunnels.webp",
		VB:  "0 0 3300 4750",
		Markers: []Marker{
			{
				I: 1, X: 523, Y: 1993,
				Title: "Entrance",
				Body:  `A long stalactite-lined cavern suffused with faerzress that mutes sound and distorts the mind. Characters waiting here for Droki must make DC 14 Wisdom saves — failure causes disadvantage on checks and saves until a short rest; a natural 20 grants the Lucky feat temporarily. Faerzress prevents combat noise from carrying beyond this area. Buppido, if present, uses any distraction to slip away to area 1b.`,
			},
			{
				I: 101, X: 642, Y: 1811,
				Title: "Pool Bypass",
				Body:  `A small side cave with a narrow crack in the far wall — the entrance to a tight natural tunnel bypassing the diseased pool. The crack is surrounded by patches of pygmywort and bigwig mushrooms (1d10 + 10 of each). Droki uses this route, shrinking himself with a pygmywort before squeezing through. Characters must similarly reduce their size to follow him.`,
			},
			{
				I: 102, X: 894, Y: 2161,
				Title: "Buppido's Lair",
				Body:  `A grisly shrine built by the derro Buppido, who believes himself a god. The floor is carpeted in humanoid remains arranged in a spiral. Buppido attacks on sight and raises six two-skulled skeletons on his first turn. After combat, the ghost of svirfneblin Pelek emerges — killed by Buppido — and asks the party to carry his remains to Blingdenstone for burial. He notes that Droki passes through these tunnels regularly.`,
			},
			{
				I: 2, X: 1056, Y: 1960,
				Title: "Diseased Pool",
				Body:  `A large, warm pool fed by a river from the Darklake. The water is disease-ridden — any creature starting its turn in the pool makes a DC 13 Constitution save or contracts cackle fever. Gnomes are immune. Characters can bypass swimming by climbing the slippery cavern walls with a DC 13 Athletics check. Failure drops them into the water.`,
			},
			{
				I: 3, X: 1424, Y: 2000,
				Title: "Parade of Fools",
				Body:  `A group of myconids — three adults, five sprouts, and two quaggoth spore servants — dance in silent, faerzress-induced ecstasy. Their leader Voosbur offers to share Zuggtmoy's gift, granting a corrupted tree stride ability but imposing mounting Wisdom saves to avoid euphoric paralysis. One sprout, Rumpadump, silently warns against accepting. Sarith, if present, eventually breaks away and joins the myconids permanently.`,
			},
			{
				I: 4, X: 1786, Y: 1951,
				Title: "Fungi Thicket",
				Body:  `A dense fungi forest blocking a tunnel junction, humming with harmless air whistling through perforated mushrooms. Difficult terrain for Small and Medium creatures; Tiny creatures have half cover. Droki shrinks to sneak through undetected. The first character to reach the junction triggers two swarms of centipedes, with spider and centipede swarms joining on subsequent rounds. Rich foraging available, including bigwigs and pygmyworts.`,
			},
			{
				I: 5, X: 2249, Y: 1759,
				Title: "Raucous Mesa",
				Body:  `A faerzress-saturated mesa that traps and echoes sounds from all of Gracklstugh above. Characters at the top can attempt to focus on specific sounds with a DC 12 Wisdom save — failure deals 2d6 psychic damage; failure by 5+ causes madness. Exceeding the DC by 5+ allows one specific question about Gracklstugh, including Droki's location. A narrow crack in the west wall shortcut leads to area 7.`,
			},
			{
				I: 6, X: 2687, Y: 1199,
				Title: "Dire Den",
				Body:  `A narrow-tunnel network serving as the lair of the Spider King — a two-headed, demonic giant spider with 44 HP, advantage on Perception checks, and two attendant giant spiders. The Spider King senses intruders automatically and positions itself to bottleneck them at the chamber entrance while its attendants attack from the ceiling. All three fight to the death. Combat noise here alerts the duergar in area 7.`,
			},
			{
				I: 7, X: 2383, Y: 1427,
				Title: "Gray Ghost Garden",
				Body:  `A locked zurkhwood chamber used by three Gray Ghost duergar to cultivate fungi for alchemical use. A central fungi pit is irrigated by a copper sprinkler tank that doubles as a stinking cloud weapon (save DC 12). Characters entering through the east crack can attempt a surprise attack. The duergar alchemist from area 8 joins combat in the second round. Crates along the north wall hold fungi worth 25 gp each.`,
			},
			{
				I: 8, X: 2304, Y: 1175,
				Title: "Gray Alchemist",
				Body:  `A clean, two-level Gray Ghost safe house and laboratory. Duergar alchemist Lorthio Bukbukken works here, armed with acid vials and alchemist's fire. The lower level connects via trapdoor and a 60-foot ladder to the Darklake Docks above. Treasure includes potions of healing, fire breath, and psychic resistance. A hidden letter in the desk implicates Stone Guard warrior Gorglak in a shady weapons deal and a planned murder.`,
			},
			{
				I: 9, X: 2124, Y: 888,
				Title: "Fountain of Evil",
				Body:  `A large cavern split by a pool with a churning central vortex, filled by Darklake runoff. A demonically corrupted water weird guards the western elevated path, attacking from 30 feet below by erupting the pool like a geyser. Droki takes the eastern lower path. Characters shadowing him must pass a group Stealth check against the weird's Perception, or be attacked — blowing their cover and sending Droki fleeing.`,
			},
			{
				I: 10, X: 1500, Y: 894,
				Title: "Cultist Pens",
				Body:  `A locked iron-gated chamber housing two derro cultists and three caged cave bears (polar bear stats). The bears feign rest but are alert — Stealth checks required to pass. On their first turn, the cultists free the bears; the third breaks out independently. A spiral path channels dormant enchantment magic used to tame the bears. A narrow crack in the south wall leads toward area 12, impassable for Medium and larger creatures.`,
			},
			{
				I: 11, X: 1275, Y: 1065,
				Title: "Quasit Playground",
				Body:  `Four quasits use this narrow tunnel as a shortcut and hiding spot, currently wrestling each other. They attack intruders on sight. When two are dropped, the survivors turn invisible and flee — making enough noise to be tracked by sound. If any escape, the cultists in area 12 cannot be surprised. A captured quasit reveals Narrak's name and the cult's plan to curse Gracklstugh's stone giants with demonic madness.`,
			},
			{
				I: 12, X: 1019, Y: 1315,
				Title: "Cultist Hideout",
				Body:  `Derro savant Narrak leads five cultists in a Demogorgon ritual atop a 5-foot natural platform, slowly growing a second head onto a stone giant effigy. A caged death dog and ettin bodyguard Grula-Munga reinforce the room. A shrieker near the western tunnel acts as an alarm; a pit trap beyond it is lined with green slime. Narrak's chest holds Keoghtom's ointment and 45 gp. Books detail two forbidden head-grafting rituals.`,
			},
			{
				I: 13, X: 1506, Y: 465,
				Title: "Dumping Pit",
				Body:  `A 15-foot-deep pit reeking of death, watched by a single derro from a scrap-metal-lined ledge. Seven zombies — three duergar and four grimlocks — shamble below, animated by faerzress. A crawling claw wearing an obsidian ring lurks in the offal mounds; the ring is a single-use magic item granting stoneskin for one hour. The hand belongs to Pelek the svirfneblin. Burying it in Blingdenstone lays his ghost to rest.`,
			},
			{
				I: 14, X: 2557, Y: 295,
				Title: "Obelisk",
				Body:  `A massive chamber containing a 50-foot black metal obelisk of alien origin, guarded by mad derro savant Pliinki and a spectator beholder. Expending any spell slot while touching the obelisk teleports all mesa occupants to the northwest gate of Gracklstugh — a temporary effect caused by faerzress disruption. A northeast mesa holds a stolen red dragon egg. Any lump of matching black metal is absorbed into the obelisk, repairing a crack on its surface.`,
			},
			{
				I: 141, X: 2219, Y: 94,
				Title: "Fungi-Covered Doors",
				Body:  `Double zurkhwood doors hidden behind a fungi patch containing 2d6 bigwigs and 2d6 pygmyworts. Characters approaching from the west must succeed on a DC 15 Perception check to notice them. The doors are barred from the east side; a DC 20 Strength check breaks them down, but doing so alerts Pliinki and the spectator in area 14.`,
			},
		},
	},
}
