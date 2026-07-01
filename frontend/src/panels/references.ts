import { api } from '../api.js'
import { h } from '../dom.js'
import type { Reference, RefEntry } from '../types.js'

function renderTable(table: { headers: string[]; rows: string[][] }): HTMLElement {
  return h('div', { className: 'overflow-x-auto mb-3' }, [
    h('table', { className: 'w-full text-sm border-collapse' }, [
      h('thead', {}, [
        h('tr', { className: 'border-b border-yellow-900/60' },
          table.headers.map((hdr) =>
            h('th', { className: 'px-4 py-2 text-left text-yellow-500 uppercase text-xs tracking-wider' }, [hdr])
          )
        ),
      ]),
      h('tbody', {}, table.rows.map((row, i) =>
        h('tr', { className: i % 2 === 0 ? 'bg-gray-800/40' : '' },
          row.map((cell, ci) =>
            h('td', { className: `px-4 py-2 align-top ${ci === 0 ? 'font-semibold text-yellow-300 whitespace-nowrap' : 'text-gray-200'}` }, [cell])
          )
        )
      )),
    ]),
  ])
}

function renderEntry(entry: RefEntry): HTMLElement {
  const children: HTMLElement[] = []

  if (entry.description) {
    for (const para of entry.description.split('\n\n')) {
      children.push(h('p', { className: 'text-gray-200 leading-relaxed mb-2' }, [para]))
    }
  }

  if (entry.bullets?.length) {
    children.push(
      h('ul', { className: 'list-disc list-outside ml-5 space-y-1 mb-2' },
        entry.bullets.map((b) => h('li', { className: 'text-gray-200 text-sm' }, [b]))
      )
    )
  }

  if (entry.table) {
    children.push(renderTable(entry.table))
  }

  if (entry.descriptionAfter) {
    for (const para of entry.descriptionAfter.split('\n\n')) {
      children.push(h('p', { className: 'text-gray-200 leading-relaxed mb-2' }, [para]))
    }
  }

  return h('div', { className: 'mb-5' }, [
    h('div', { className: 'flex items-baseline gap-3 mb-2' }, [
      h('h2', { className: 'text-base font-bold text-yellow-400 uppercase tracking-wide' }, [entry.name]),
      ...(entry.tag
        ? [h('span', { className: 'text-xs font-semibold text-gray-400 uppercase tracking-widest' }, [entry.tag])]
        : []),
    ]),
    ...children,
  ])
}

export async function refsPanel(): Promise<Node> {
  const refs = (await api.refs()) as Reference[]

  const content = h('div', { className: 'max-w-3xl' }, [])

  function showRef(id: string) {
    const ref = refs.find((r) => r.id === id)
    content.innerHTML = ''
    if (!ref) return
    for (const entry of ref.content.entries) {
      content.append(renderEntry(entry))
    }
  }

  const select = h(
    'select',
    { className: 'form-select w-auto' },
    [
      h('option', { value: '' }, ['Select a reference...']),
      ...refs.map((r) => h('option', { value: r.id }, [r.title])),
    ]
  ) as HTMLSelectElement

  select.addEventListener('change', () => showRef(select.value))

  if (refs.length) {
    select.value = refs[0].id
    showRef(refs[0].id)
  }

  return h('div', {}, [
    h('div', { className: 'mb-4' }, [select]),
    content,
  ])
}
