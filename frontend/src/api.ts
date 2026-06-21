async function getJSON<T>(url: string): Promise<T> {
  const res = await fetch(url)
  if (!res.ok) throw new Error(await res.text())
  return res.json() as Promise<T>
}

export const api = {
  sessions: () => getJSON('/api/sessions'),
  npcs: () => getJSON('/api/npcs'),
  encounters: () => getJSON('/api/encounters'),
  monsters: () => getJSON('/api/monsters'),
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
