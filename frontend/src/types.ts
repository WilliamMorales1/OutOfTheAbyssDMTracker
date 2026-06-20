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
  dmNotes: string
}

export interface Location {
  name: string
  type: string
  danger: number
  dangerStars: string
  description: string
  secrets: string
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
  name: string
  type: string
  cr: string
  hp: number
  hpFormula: string
  ac: number
  acDesc: string
  speed: string
}

export interface EventRow {
  title: string
  category: string
  description: string
}

export interface EncounterMonster {
  name: string
  cr: string
  quantity: string
}

export interface Encounter {
  id: number
  name: string
  chapter: number
  location: string
  difficulty: number
  difficultyStars: string
  levelup: boolean
  notes: string
  monsters: EncounterMonster[]
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
  w: string
  h: string
  markers: Marker[]
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
