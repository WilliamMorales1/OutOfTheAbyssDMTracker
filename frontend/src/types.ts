export interface Session {
  sessionNum: number
  title: string
  chapters: string
  levelStart: number
  levelEnd: number
  summary: string
  keyEncounters: string
  keyNpcs: string
  checkpoint: string
}

export interface Npc {
  name: string
  madness: number
  madnessStars: string
  disposition: string
  location: string
  description: string
  notes: string
}

export interface MonsterRow {
  id: number
  name: string
  type: string
  cr: string
  imageUrl: string
}

export interface MonsterStat {
  id: number
  name: string
  ac: number
  hp: number
  dex: number
}

export interface StatBlockEntry {
  name: string
  text: string
}

export interface MonsterDetail {
  id: number
  name: string
  type: string
  size: string
  alignment: string
  cr: string
  source: string
  hp: number
  hpFormula: string
  ac: number
  acDesc: string
  speed: string
  str: number
  dex: number
  con: number
  int: number
  wis: number
  cha: number
  savingThrows: string
  skills: string
  damageResistances: string
  damageImmunities: string
  vulnerabilities: string
  conditionImmunities: string
  senses: string
  passivePerception: number
  languages: string
  environment: string
  imageUrl: string
  tokenUrl: string
  traits: StatBlockEntry[] | null
  actions: StatBlockEntry[] | null
  reactions: StatBlockEntry[] | null
  legendaryActions: StatBlockEntry[] | null
  spellcasting: StatBlockEntry[] | null
  notes: string
}

export interface SpellRow {
  id: number
  name: string
  level: number
  school: string
  classes: string
}

export interface SpellDetail {
  id: number
  name: string
  level: number
  school: string
  ritual: boolean
  castingTime: string
  range: string
  components: string
  duration: string
  concentration: boolean
  classes: string
  description: string
  higherLevel: string
  source: string
}

export interface Marker {
  i: number
  x: number
  y: number
  title: string
  body: string
}

export interface GameMap {
  id: string
  img: string
  vb: string
  markers: Marker[]
}

export interface RefEntry {
  name: string
  tag?: string
  description?: string
  bullets?: string[]
  table?: { headers: string[]; rows: string[][] }
  descriptionAfter?: string
}

export interface Reference {
  id: string
  title: string
  content: { entries: RefEntry[] }
}

export interface SearchResult {
  chapterTitle: string
  content: string
  score: number
}

export interface ChatExchange {
  question: string
  answer: string
}
