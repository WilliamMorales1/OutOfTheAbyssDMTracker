import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { MonsterRow } from '../types.js'

export async function monstersPanel(): Promise<Node> {
  const data = (await api.monsters()) as MonsterRow[]

  const columns: Column<MonsterRow>[] = [
    { header: 'Name', render: (m) => h('strong', {}, [m.name]), sortValue: (m) => m.name },
    { header: 'Type', render: (m) => m.type, sortValue: (m) => m.type },
    { header: 'CR', render: (m) => m.cr, sortValue: (m) => m.cr },
    {
      header: 'HP',
      render: (m) => h('span', {}, [String(m.hp), m.hpFormula ? h('small', { className: 'text-secondary' }, [` (${m.hpFormula})`]) : null]),
      sortValue: (m) => m.hp,
    },
    {
      header: 'AC',
      render: (m) => h('span', {}, [String(m.ac), m.acDesc ? h('small', { className: 'text-secondary' }, [` (${m.acDesc})`]) : null]),
      sortValue: (m) => m.ac,
    },
    { header: 'Speed', render: (m) => m.speed, sortValue: (m) => m.speed },
  ]

  return dataTable(columns, data, 'No monsters found. Run: go run ./database/seed_monsters.go')
}
