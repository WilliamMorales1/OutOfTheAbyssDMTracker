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
				I: 1, X: 680, Y: 680,
				Title: "The Shattered Spire",
				Body:  "A crumbling stalagmite tower on the Darklake shore, structurally compromised and avoided by most duergar. Used occasionally as a lookout post despite the danger of collapse. The upper levels are accessible but each floor requires a DC 12 Dexterity saving throw to avoid dislodging loose stone — failure rains debris on creatures below (1d6 bludgeoning).",
			},
			{
				I: 2, X: 1700, Y: 580,
				Title: "Darklake Docks",
				Body:  "The main harbor of Gracklstugh, where trade vessels and patrol boats moor along stone quays. Kuo-toa workers unload cargo alongside duergar dock hands. Clan Muzgardt controls the docking fees and ferry contracts. The docks are busy at all hours — a natural point of entry for travelers arriving by Darklake and a hub for black-market dealings conducted in plain sight.",
			},
			{
				I: 3, X: 3200, Y: 480,
				Title: "Darklake Brewery",
				Body:  "Clan Muzgardt's most profitable enterprise, producing the potent mushroom ale and fungi-based spirits that fuel Gracklstugh's workforce. The brewery operates day and night, its vats tended by indentured laborers. The clan's brewmaster, Ulvara Muzgardt, is also the de facto liaison between Clan Muzgardt and the Stone Guards — she trades information as freely as she trades kegs.",
			},
			{
				I: 4, X: 2800, Y: 950,
				Title: "Blade Bazaar",
				Body:  "The largest open market in Gracklstugh, specializing in weapons, armor, and war materiel forged in the city's renowned foundries. Dozens of stalls line the cavern floor. The bazaar is also where travelers can find information, hire guides, and discreetly acquire goods that never appear on any official manifest. Prices are non-negotiable with strangers — DC 15 Persuasion or Deception to haggle.",
			},
			{
				I: 5, X: 2000, Y: 1080,
				Title: "Darklake District",
				Body:  "The bustling commercial heart of Gracklstugh, sprawling between Laduguer's Furrow and the Darklake shore. A mix of duergar merchants, off-duty warriors, and non-duergar residents crowd narrow streets between squat stone buildings. The district is the most permissive part of the city — outsiders are tolerated here more than anywhere else, though Stone Guards still patrol in force.",
			},
			{
				I: 6, X: 3400, Y: 1080,
				Title: "Overlake Hold",
				Body:  "The fortress barracks of the Stone Guards, Gracklstugh's elite military police. Carved directly into the cavern wall overlooking the Darklake, the hold houses several hundred warriors at any given time. The captain of the Stone Guards, Errde Blackskull, runs a tight operation from here. Outsiders brought before the Stone Guards are held in cells beneath the hold pending interrogation.",
			},
			{
				I: 7, X: 2500, Y: 1200,
				Title: "Ghost Rocks",
				Body:  "A cluster of pale stalagmites in the Darklake District said to be the petrified remains of duergar who displeased Laduguer. The stones are avoided at night. Locals claim to hear whispering from within them — most dismiss this as superstition, but a successful DC 14 Perception check while standing among the rocks reveals faint Dwarvish murmuring that speaks fragmented warnings about the Deep King.",
			},
			{
				I: 8, X: 480, Y: 1380,
				Title: "Halls of Sacred Scrolls",
				Body:  "The archive and temple complex of the Keepers of the Flame, duergar diviners and loremasters who serve as Gracklstugh's religious scholars. Thousands of scrolls and stone tablets line the shelves, covering history, theology, and arcane theory. The high keeper, Gartokkar Xundorn, is deeply suspicious of outsiders and will not share the collection freely — but the right price or a compelling story can buy access to specific records.",
			},
			{
				I: 9, X: 300, Y: 1650,
				Title: "West Cleft District",
				Body:  "A residential and light-industry district on the western side of Laduguer's Furrow. Home to lower-ranking duergar craftspeople and their households. The district is quieter than the Darklake District but not without tension — inter-clan disputes here are handled with blades rather than words, and the Stone Guards take their time responding to calls from this side of the furrow.",
			},
			{
				I: 10, X: 2000, Y: 1800,
				Title: "Laduguer's Furrow",
				Body:  "A vast natural rift that splits Gracklstugh from north to south, named for the duergar god Laduguer. The chasm is hundreds of feet deep; the sound of distant hammering rises from forges below. Stone bridges cross it at several points, each guarded by pairs of Stone Guards. The Furrow is a hard boundary — which side of it you live on signals your status and clan allegiance to anyone who knows the city.",
			},
			{
				I: 11, X: 3600, Y: 1900,
				Title: "East Cleft District",
				Body:  "The industrial quarter east of Laduguer's Furrow, home to foundries, smelters, and the workshops of weapon-crafting clans. The air here is thick with smoke and the ring of hammers never ceases. Clan Steelshadow maintains several private forges in this district. The heat is oppressive even by Underdark standards — Constitution saves may be called for during extended activity.",
			},
			{
				I: 12, X: 500, Y: 2500,
				Title: "Cairngorm Cavern",
				Body:  "A vast side cavern at the city's southwestern edge where Themberchaud the Wyrmsmith, Gracklstugh's red dragon, is supposed to maintain his lair — but Themberchaud actually dwells in his own cavern to the south. Cairngorm is instead home to Cairngorm the svirfneblin trader, a deep gnome who has somehow maintained neutrality in Gracklstugh's clan politics for decades. His shop is a known haven for surface-worlders passing through.",
			},
			{
				I: 13, X: 2700, Y: 2400,
				Title: "Hold of the Deep King",
				Body:  "The palace complex of Horgar Steelshadow V, the Deep King of Gracklstugh. The hold is a fortress within the city — heavily guarded, its inner chambers inaccessible to outsiders without a direct royal summons. The Deep King is ill, possibly poisoned, and the Stone Guard and the Keepers of the Flame are both maneuvering around his weakness. Audiences can be arranged but require sponsorship from a recognized clan or officer.",
			},
			{
				I: 14, X: 2000, Y: 2900,
				Title: "Themberchaud's Cavern",
				Body:  "The private lair of Themberchaud, the adult red dragon who serves as Gracklstugh's Wyrmsmith — responsible for keeping the city's forge fires burning. Themberchaud is intelligent, vain, and increasingly aware that the duergar have kept him dependent and isolated. He will speak to outsiders who reach his cavern and may offer rewards for information about the surface world or tasks that serve his growing ambitions. Approaching without invitation requires bypassing two checkpoints of duergar guards.",
			},
		},
	},
	{
		ID:  "whorlstoneTunnels",
		Img: "./images/whorlstoneTunnels.webp",
		VB:  "0 0 3300 4750",
		Markers: []Marker{
			{
				I: 1, X: 260, Y: 1700,
				Title: "Area 1: Entry Passage",
				Body:  "The main entrance to the Whorlstone Tunnels from Gracklstugh — a narrow crack in the cavern wall, easily missed. The passage smells of rot and old blood. Recent tracks in the fungal growth suggest regular use. The tunnel narrows to 5 feet in places; Medium creatures must succeed on a DC 10 Acrobatics check to move at full speed.",
			},
			{
				I: 2, X: 820, Y: 1580,
				Title: "Area 2: Fungi Cavern",
				Body:  "A humid cave choked with bioluminescent fungi — pale violet light illuminates the floor. Dozens of edible and toxic varieties grow here in clusters. A DC 13 Nature check identifies which are safe. Harvesting the mushrooms takes 10 minutes per dose. A derro patrol (2d4 derro) passes through every 2 hours.",
			},
			{
				I: 3, X: 1150, Y: 1640,
				Title: "Area 3: Dark Pool",
				Body:  "A still black pool fills most of this chamber. The water is 20 feet deep at center and cold enough to require DC 10 Constitution saves per minute of immersion. Something stirs in the depths — a pair of giant blind albino crayfish (use giant crab stats) that surface if disturbed. The pool connects to a submerged passage leading south (Athletics DC 14 to swim through).",
			},
			{
				I: 4, X: 1430, Y: 1600,
				Title: "Area 4: Tunnel Crossroads",
				Body:  "Four passages meet here. Derro graffiti covers the walls — mostly obscene, but mixed in are coded directional markings used by the Ironslag cult. A DC 14 Investigation check and understanding Dwarvish deciphers the markings, revealing the cult uses this as a waypoint and the symbols indicate where prisoners are taken.",
			},
			{
				I: 5, X: 1700, Y: 1430,
				Title: "Area 5: Darkcap Forest",
				Body:  "A dense field of darkcap mushrooms, their broad caps reaching shoulder height. Movement through the patch requires a DC 12 Dexterity (Stealth) check or the caps release clouds of spores — creatures in a 10-foot radius must succeed on a DC 13 Constitution save or be poisoned for 1 hour. The derro cultivate these deliberately as a perimeter alarm.",
			},
			{
				I: 6, X: 1620, Y: 1170,
				Title: "Area 6: Lair of the Spider King",
				Body:  "See Spider King map. Home of Yestabrod, a myconid mutated beyond recognition by demonic influence — now a bloated, web-trailing horror that calls itself the Spider King. Thick webbing fills the upper reaches. Yestabrod commands a mixed force of myconids and derro thralls. Killing Yestabrod causes its body to explode in a burst of spores (DC 15 Con save or poisoned 1 hour and hallucinating).",
			},
			{
				I: 7, X: 2080, Y: 1000,
				Title: "Area 7: Gray Ghost Garden",
				Body:  "See Gray Ghost Garden map. A grove of gray, desaturated fungi tended by myconids who have withdrawn from the myconid sovereign's influence. The garden has an eerie stillness — sounds are muffled and colors appear washed out within its bounds. The myconids here are peaceful and will meld with visitors, sharing fragmented visions of Gracklstugh's past. The Gray Ghost mushrooms they cultivate are alchemical components worth 10 gp each.",
			},
			{
				I: 8, X: 1870, Y: 790,
				Title: "Area 8: Gray Alchemist's Workshop",
				Body:  "See Gray Ghost Garden map. The laboratory of the Gray Alchemist — a duergar who rejected clan society and has lived in the tunnels for decades. The workshop is cluttered with alembics, fungal extracts, and experimental notes. The Alchemist is eccentric but not hostile; he will trade information about the tunnels and the cult for rare ingredients or surface-world news. His notes contain partial formulas for a potion that suppresses demonic corruption.",
			},
			{
				I: 9, X: 2380, Y: 680,
				Title: "Area 9: Crystal Cavern",
				Body:  "Pale blue crystals line every surface of this cave, refracting light into geometric patterns. The crystals hum at a frequency felt in the teeth. Magical effects are enhanced here — spell save DCs increase by 2 and damage dice for spells are rolled with advantage. The crystals are fragile; heavy impacts shatter them and cause the harmonic hum to cease, ending the magical enhancement.",
			},
			{
				I: 10, X: 700, Y: 1100,
				Title: "Area 10: Cultist Pens",
				Body:  "See Cultist Pens map. Derro cultists of the Elder Elemental Eye hold prisoners here — a mix of Underdark wanderers, surface-world slaves, and duergar dissidents. The pens are crude iron cages bolted to the cavern wall. 1d6+2 derro guards are present at all times, led by a derro savant. Prisoners can be freed; they know the layout of the adjacent cult areas and will guide rescuers to the shrine if freed.",
			},
			{
				I: 11, X: 1010, Y: 1040,
				Title: "Area 11: Fungal Warren",
				Body:  "A low-ceilinged cave threaded with mycelium networks so dense the floor is spongy underfoot. Moving silently requires DC 14 Stealth — footsteps pop and squelch audibly. Several gas spore creatures drift at ceiling height. A derro alchemist has been experimenting on captured myconids here; three myconid sprouts in crude wire cages can be found and freed.",
			},
			{
				I: 12, X: 1900, Y: 1760,
				Title: "Area 12: Cultist Hideout",
				Body:  "See Cultist Hideout map. A shrine and meeting chamber for the derro cult of the Elder Elemental Eye. A crude stone altar occupies the center, stained with offerings. The walls bear elemental sigils carved in a frenzied hand. A pitfall trap guards the entrance (DC 14 Perception to notice; 20-foot drop, 2d6 bludgeoning). A shrieker fungus serves as a living alarm near the inner door.",
			},
			{
				I: 13, X: 1390, Y: 790,
				Title: "Area 13: Webbed Passage",
				Body:  "Thick spider webbing fills this corridor from floor to ceiling. Movement costs double and requires a DC 12 Strength check to avoid becoming restrained (DC 14 Strength to break free). The webs are dry and extremely flammable — a single torch ignites a 10-foot section per round. Two phase spiders lair in the webs and attack from the Ethereal Plane.",
			},
			{
				I: 14, X: 830, Y: 640,
				Title: "Area 14: The Obelisk",
				Body:  "See Obelisk Chamber map. A four-foot black stone obelisk of pre-derro origin dominates the center of a roughly circular cavern. Concentric ridges of raised stone surround it. The obelisk radiates faint conjuration magic (DC 14 Arcana to identify). Touching it while concentrating (DC 13 Arcana) reveals fragmented visions of Gracklstugh before the duergar arrived. The derro cultists believe it is a conduit to the Elder Elemental Eye and conduct rituals here on new moons.",
			},
			{
				I: 15, X: 2240, Y: 490,
				Title: "Area 14a: Hidden Chamber",
				Body:  "A concealed side chamber accessed through a stone door disguised as natural rock (DC 16 Investigation to find). Inside: a cache of cult supplies — 3 potions of healing, 200 gp in mixed coin, a derro-made dagger +1, and a coded ledger in Dwarvish detailing prisoner transfers and cult contact names in Gracklstugh.",
			},
			{
				I: 16, X: 700, Y: 1840,
				Title: "Area 16: Derro Warren",
				Body:  "A cluster of crude stone hovels where off-duty cultists sleep and eat. Twelve derro are here at any given time, split between sleeping and gaming with bone dice. Alarm is raised immediately if combat sounds from any adjacent area. The warren has one exit back toward Gracklstugh and one leading deeper into the tunnels. Ransacking the hovels yields 2d6 × 10 sp in scattered coin.",
			},
		},
	},
	{
		ID:  "obeliskChamber",
		Img: "./images/obeliskChamber.webp",
		VB:  "0 0 1000 945",
		Markers: []Marker{
			{
				I: 1, X: 370, Y: 530,
				Title: "The Obelisk",
				Body:  "A four-foot black stone obelisk of unknown pre-derro origin stands at the center of concentric raised ridges carved into the stone floor. The obelisk radiates faint conjuration magic (DC 14 Arcana). Touching it while concentrating (DC 13 Arcana) triggers fragmented visions — Gracklstugh as a natural cavern before duergar settlement, shadowy figures performing rituals, and a brief overwhelming sensation of being observed from somewhere vast and cold. The derro cultists treat this as their holiest site and perform lunar rituals here.",
			},
		},
	},
	{
		ID:  "cultistHideout",
		Img: "./images/cultistHideout.webp",
		VB:  "0 0 1000 618",
		Markers: []Marker{
			{
				I: 1, X: 155, Y: 310,
				Title: "Pitfall Trap",
				Body:  "A 5-foot section of flooring conceals a 20-foot pit. DC 14 Perception (passive or active) to notice the slightly discolored stone. Creatures that fail fall for 2d6 bludgeoning damage and land prone. The pit walls are smooth; climbing out requires DC 14 Athletics. The derro reset this trap after use — a 10-minute task requiring tools.",
			},
			{
				I: 2, X: 460, Y: 280,
				Title: "Shrieker",
				Body:  "A shrieker fungus planted near the inner door serves as a living alarm. It begins screaming when any creature without darkvision moves within 30 feet with a light source, or when any creature moves within 10 feet in darkness. The screaming lasts for 1d4 rounds and alerts all creatures in areas 12 and 16. It can be destroyed (AC 5, 3 hp) or pacified with a DC 12 Nature check and 1 minute of careful handling.",
			},
			{
				I: 3, X: 810, Y: 310,
				Title: "Cult Shrine",
				Body:  "A crude stone altar occupies the center of the inner chamber, its surface dark with dried offerings. Elemental Eye sigils cover every wall, carved in a frenzied style suggesting multiple hands over many years. A locked iron chest behind the altar (DC 14 Thieves' Tools) holds: cult robes, a ritual dagger with the Eye sigil, 80 gp, and a partial map of the Whorlstone Tunnels marking areas the cult considers sacred.",
			},
		},
	},
	{
		ID:  "cultistPens",
		Img: "./images/cultistPens.webp",
		VB:  "0 0 1000 724",
		Markers: []Marker{
			{
				I: 1, X: 220, Y: 380,
				Title: "Guard Post",
				Body:  "A crude derro guard station at the pen entrance — a stone table, two stools, and a rack of javelins. 1d4+1 derro are here at all times. They raise alarm immediately if threatened, alerting the rest of the pen guards. A ring of keys on the table opens all the cages. A ledger on the table records prisoner names, origins, and 'intended use' in cramped Dwarvish.",
			},
			{
				I: 2, X: 620, Y: 362,
				Title: "Prisoner Cages",
				Body:  "Iron cages bolted to the cavern wall hold 2d6 prisoners — a mix of Underdark wanderers, surface-world captives, and duergar accused of disloyalty. The prisoners are malnourished but alive. Several know useful information about the tunnels or Gracklstugh. Freeing them creates an obligation — they have nowhere safe to go without escort. One prisoner, a deep gnome named Senni, knows a hidden exit from the tunnels leading to a rarely used Darklake cove.",
			},
		},
	},
	{
		ID:  "greyGhostGarden",
		Img: "./images/greyGhostGarden.webp",
		VB:  "0 0 1000 875",
		Markers: []Marker{
			{
				I: 1, X: 200, Y: 220,
				Title: "Gray Alchemist's Workshop",
				Body:  "A duergar exile known only as the Gray Alchemist has lived in these tunnels for decades, rejected by clan society after his experiments drew suspicion. The workshop is cluttered with fungal specimens, bubbling alembics, and notes in cramped Dwarvish. The Alchemist is eccentric but not hostile. He will trade information about the tunnels, the cult, and Gracklstugh's political situation in exchange for rare surface-world ingredients or news from above. His partial formula for a demonic-corruption suppressant could be invaluable.",
			},
			{
				I: 2, X: 740, Y: 560,
				Title: "Gray Ghost Garden",
				Body:  "A grove of gray, desaturated fungi tended by a small community of myconids who have withdrawn from their sovereign's influence. The garden has an uncanny stillness — sounds are muffled, colors appear washed out, and mundane anxiety seems to fade within its bounds. The myconids will meld with any creature that approaches peacefully, sharing fragmented visual memories of Gracklstugh's past. The Gray Ghost mushrooms they cultivate (10 gp each, 2d6 available) are potent alchemical components used in spore-resistance draughts.",
			},
		},
	},
	{
		ID:  "lairOfTheSpiderKing",
		Img: "./images/lairOfTheSpiderKing.webp",
		VB:  "0 0 944 1000",
		Markers: []Marker{
			{
				I: 1, X: 470, Y: 400,
				Title: "Lair of Yestabrod",
				Body:  "Yestabrod was once a myconid, but demonic influence has warped it into something unrecognizable — a bloated, pale mass trailing thick webs, calling itself the Spider King. The lair is strung with webbing from floor to ceiling; movement costs double. Yestabrod commands 2d6 myconid sprouts and 1d4 derro thralls (treat as charmed). When reduced to 0 hp, its body ruptures: all creatures within 20 feet must succeed on a DC 15 Constitution saving throw or become poisoned for 1 hour and suffer hallucinations (DM's choice of effect).",
			},
		},
	},
}
