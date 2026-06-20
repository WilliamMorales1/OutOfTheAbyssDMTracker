import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { Session } from '../types.js'

export async function sessionsPanel(): Promise<Node> {
  const data = (await api.sessions()) as Session[]

  const columns: Column<Session>[] = [
    { header: '#', render: (s) => h('strong', {}, [String(s.sessionNum)]), sortValue: (s) => s.sessionNum },
    { header: 'Title', render: (s) => h('strong', {}, [s.title]), sortValue: (s) => s.title },
    { header: 'Chapters', render: (s) => h('span', { className: 'badge bg-secondary' }, [s.chapters]), sortValue: (s) => s.chapters },
    {
      header: 'Levels',
      render: (s) =>
        h('span', { className: 'badge bg-info text-dark' }, [
          String(s.levelStart) + (s.levelStart !== s.levelEnd ? `→${s.levelEnd}` : ''),
        ]),
      sortValue: (s) => s.levelStart,
    },
    { header: 'Summary', render: (s) => s.summary, sortValue: (s) => s.summary },
    { header: 'Key Encounters', render: (s) => s.keyEncounters, sortValue: (s) => s.keyEncounters },
    { header: 'Key NPCs', render: (s) => s.keyNpcs, sortValue: (s) => s.keyNpcs },
    { header: 'Checkpoint', render: (s) => s.checkpoint, sortValue: (s) => s.checkpoint },
  ]

  return dataTable(columns, data, 'No sessions found. Run: go run ./database/seed_sessions.go')
}
