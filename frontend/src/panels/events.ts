import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { EventRow } from '../types.js'

export async function eventsPanel(): Promise<Node> {
  const data = (await api.events()) as EventRow[]

  const columns: Column<EventRow>[] = [
    { header: 'Title', render: (e) => h('strong', {}, [e.title]), sortValue: (e) => e.title },
    {
      header: 'Category',
      render: (e) => h('span', { className: `badge cat-${e.category}` }, [e.category]),
      sortValue: (e) => e.category,
    },
    { header: 'Description', render: (e) => e.description, sortValue: (e) => e.description },
  ]

  return dataTable(columns, data, 'No events found.')
}
