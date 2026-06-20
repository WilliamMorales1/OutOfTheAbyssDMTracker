import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { Npc } from '../types.js'

export async function npcsPanel(): Promise<Node> {
  const data = (await api.npcs()) as Npc[]

  const columns: Column<Npc>[] = [
    { header: 'Name', render: (n) => h('strong', {}, [n.name]), sortValue: (n) => n.name },
    { header: 'Madness', render: (n) => n.madnessStars, sortValue: (n) => n.madness },
    {
      header: 'Disposition',
      render: (n) => h('span', { className: `badge disp-${n.disposition}` }, [n.disposition]),
      sortValue: (n) => n.disposition,
    },
    { header: 'Location', render: (n) => n.location, sortValue: (n) => n.location },
    { header: 'Description', render: (n) => n.description, sortValue: (n) => n.description },
    { header: 'Notes', render: (n) => n.notes, sortValue: (n) => n.notes },
  ]

  return dataTable(columns, data, 'No NPCs found.')
}
