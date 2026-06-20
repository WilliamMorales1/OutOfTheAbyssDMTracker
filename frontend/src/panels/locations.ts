import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { Location } from '../types.js'

export async function locationsPanel(): Promise<Node> {
  const data = (await api.locations()) as Location[]

  const columns: Column<Location>[] = [
    { header: 'Name', render: (l) => h('strong', {}, [l.name]), sortValue: (l) => l.name },
    { header: 'Type', render: (l) => l.type, sortValue: (l) => l.type },
    { header: 'Danger', render: (l) => l.dangerStars, sortValue: (l) => l.danger },
    { header: 'Description', render: (l) => l.description, sortValue: (l) => l.description },
    { header: 'Secrets', render: (l) => l.secrets, sortValue: (l) => l.secrets },
  ]

  return dataTable(columns, data, 'No locations found.')
}
