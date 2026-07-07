import { api } from '../api.js'
import { h } from '../dom.js'
import { dataTable, type Column } from '../dataTable.js'
import type { ConditionRow, DemonLord, ExhaustionLevel, SkillArea } from '../types.js'

async function actionsView(): Promise<Node> {
  const actions = await api.actions()
  const skillAreas = await api.skillAreas()

  const columns: Column<SkillArea>[] = [
    { header: 'Skill', render: (s) => s.skill, sortValue: (s) => s.skill },
    { header: 'Areas', render: (s) => s.areas, sortValue: (s) => s.areas },
  ]

  const children: Node[] = []
  for (const a of actions) {
    children.push(
      h('div', { className: 'mb-3' }, [
        h('div', { className: 'flex items-baseline gap-3 mb-2' }, [
          h('h2', { className: 'text-base font-bold text-yellow-400 uppercase tracking-wide' }, [a.name]),
          h('span', { className: 'text-xs font-semibold text-gray-400 uppercase tracking-widest' }, [a.tag]),
        ]),
        h('p', { className: 'text-gray-200 leading-relaxed mb-3' }, [a.description]),
      ])
    )
  }
  children.push(dataTable(columns, skillAreas as SkillArea[], 'No skill areas found.'))

  return h('div', {}, children)
}

async function conditionsView(): Promise<Node> {
  const conditions = (await api.conditions()) as ConditionRow[]
  const exhaustionLevels = (await api.exhaustionLevels()) as ExhaustionLevel[]

  const columns: Column<ConditionRow>[] = [
    { header: 'Condition', render: (c) => c.name, sortValue: (c) => c.name },
    {
      header: 'Effects',
      render: (c) => {
        if (c.effects) {
          return h('ul', { className: 'list-disc list-outside ml-5 space-y-1' },
            c.effects.split('\n').map((b) => h('li', {}, [b]))
          )
        }
        return c.description
      },
      sortValue: (c) => c.name,
    },
  ]

  const exhaustionColumns: Column<ExhaustionLevel>[] = [
    { header: 'Level', render: (e) => e.level, sortValue: (e) => e.level },
    { header: 'Effect', render: (e) => e.effect, sortValue: (e) => e.level },
  ]

  const exhaustion = conditions.find((c) => c.name === 'Exhaustion')

  return h('div', {}, [
    dataTable(columns, conditions, 'No conditions found.'),
    ...(exhaustion
      ? [
          h('h2', { className: 'text-base font-bold text-yellow-400 uppercase tracking-wide mt-4 mb-2' }, ['Exhaustion Levels']),
          dataTable(exhaustionColumns, exhaustionLevels, 'No exhaustion levels found.'),
          ...exhaustion.descriptionAfter.split('\n\n').map((para) =>
            h('p', { className: 'text-gray-200 leading-relaxed mt-3' }, [para])
          ),
        ]
      : []),
  ])
}

async function demonLordsView(): Promise<Node> {
  const demonLords = (await api.demonLords()) as DemonLord[]

  const columns: Column<DemonLord>[] = [
    { header: 'Name', render: (d) => d.name, sortValue: (d) => d.name },
    { header: 'Dominions', render: (d) => d.dominions, sortValue: (d) => d.dominions },
    { header: 'Epithets', render: (d) => d.epithets, sortValue: (d) => d.epithets },
    { header: 'Layer', render: (d) => d.layer, sortValue: (d) => d.layer },
    { header: 'Description', render: (d) => d.description, sortValue: (d) => d.description },
    { header: 'Servants', render: (d) => d.servants, sortValue: (d) => d.servants },
    { header: 'Component', render: (d) => d.component, sortValue: (d) => d.component },
    { header: 'Component Location', render: (d) => d.componentLocation, sortValue: (d) => d.componentLocation },
  ]

  return dataTable(columns, demonLords, 'No demon lords found.')
}

const REFS = [
  { id: 'actions', title: 'Actions', load: actionsView },
  { id: 'conditions', title: 'Conditions', load: conditionsView },
  { id: 'demon-lords', title: 'Demon Lords', load: demonLordsView },
]

export async function refsPanel(): Promise<Node> {
  const content = h('div', {}, [])

  async function showRef(id: string) {
    const ref = REFS.find((r) => r.id === id)
    content.innerHTML = ''
    if (!ref) return
    content.append(await ref.load())
  }

  const select = h(
    'select',
    { className: 'form-select w-auto' },
    [
      h('option', { value: '' }, ['Select a reference...']),
      ...REFS.map((r) => h('option', { value: r.id }, [r.title])),
    ]
  ) as HTMLSelectElement

  select.addEventListener('change', () => showRef(select.value))

  select.value = REFS[0].id
  await showRef(REFS[0].id)

  return h('div', {}, [
    h('div', { className: 'mb-4' }, [select]),
    content,
  ])
}
