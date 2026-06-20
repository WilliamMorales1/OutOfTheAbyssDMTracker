import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { Encounter } from '../types.js'

export async function encountersPanel(): Promise<Node> {
  const data = (await api.encounters()) as Encounter[]

  const columns: Column<Encounter>[] = [
    { header: 'Name', render: (e) => h('strong', {}, [e.name]), sortValue: (e) => e.name },
    { header: 'Chapter', render: (e) => String(e.chapter), sortValue: (e) => e.chapter },
    { header: 'Location', render: (e) => e.location, sortValue: (e) => e.location },
    { header: 'Difficulty', render: (e) => e.difficultyStars, sortValue: (e) => e.difficulty },
    {
      header: 'Monsters',
      render: (e) =>
        h(
          'span',
          {},
          e.monsters.map((m) =>
            h('span', { className: 'badge bg-secondary me-1 mb-1' }, [
              m.quantity !== '1' ? `${m.quantity}× ` : '',
              m.name,
              h('small', { className: 'ms-1 text-warning' }, [`CR${m.cr}`]),
            ])
          )
        ),
    },
    { header: 'Level Up', render: (e) => (e.levelup ? h('span', { className: 'badge bg-success' }, ['✓']) : '') },
    { header: 'Notes', render: (e) => e.notes, sortValue: (e) => e.notes },
  ]

  return dataTable(columns, data, 'No encounters found.')
}
