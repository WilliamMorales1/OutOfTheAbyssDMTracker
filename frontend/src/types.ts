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
  name: string
  type: string
  cr: string
  hp: number
  hpFormula: string
  ac: number
  acDesc: string
  speed: string
  dex: number
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
