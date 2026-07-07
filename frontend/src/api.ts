import type { Action, ConditionRow, DemonLord, ExhaustionLevel, MonsterDetail, MonsterRow, MonsterStat, SkillArea, SpellDetail, SpellRow } from './types.js'

async function getJSON<T>(url: string): Promise<T> {
  const res = await fetch(url)
  if (!res.ok) throw new Error(await res.text())
  return res.json() as Promise<T>
}

export const api = {
  sessions: () => getJSON('/api/sessions'),
  demonLords: () => getJSON<DemonLord[]>('/api/demon-lords'),
  actions: () => getJSON<Action[]>('/api/actions'),
  skillAreas: () => getJSON<SkillArea[]>('/api/skill-areas'),
  conditions: () => getJSON<ConditionRow[]>('/api/conditions'),
  exhaustionLevels: () => getJSON<ExhaustionLevel[]>('/api/exhaustion-levels'),
  monsters: () => getJSON<MonsterRow[]>('/api/monsters'),
  monster: (id: number) => getJSON<MonsterDetail>(`/api/monsters/${id}`),
  monsterStats: () => getJSON<MonsterStat[]>('/api/monster-stats'),
  spells: () => getJSON<SpellRow[]>('/api/spells'),
  spell: (id: number) => getJSON<SpellDetail>(`/api/spells/${id}`),
  maps: () => getJSON('/api/maps'),
  search: (q: string) => getJSON(`/api/search?q=${encodeURIComponent(q)}`),
  chat: (q: string) => getJSON(`/api/chat?q=${encodeURIComponent(q)}`),
  notes: () => getJSON<string[]>('/api/notes'),
  note: (name: string) => getJSON<{ name: string; content: string }>(`/api/notes/${encodeURIComponent(name)}`),
  saveNote: async (name: string, content: string) => {
    const res = await fetch(`/api/notes/${encodeURIComponent(name)}`, { method: 'PUT', body: content })
    if (!res.ok) throw new Error(await res.text())
  },
}
