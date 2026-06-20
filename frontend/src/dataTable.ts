import { h } from './dom.js'

export interface Column<T> {
  header: string
  render: (row: T) => Node | string
  sortValue?: (row: T) => string | number
}

export function dataTable<T>(columns: Column<T>[], rows: T[], emptyMessage: string): Node {
  if (rows.length === 0) {
    return h('p', { className: 'empty p-3' }, [emptyMessage])
  }

  let filter = ''
  let sortCol: number | null = null
  let sortAsc = true

  const tbody = h('tbody', {}, [])
  const theadRow = h('tr', {}, [])
  const wrapper = h('div', { className: 'table-responsive' }, [])

  function getRows(): T[] {
    let result = rows
    if (filter.trim()) {
      const f = filter.toLowerCase()
      result = result.filter((row) =>
        columns.some((c) => String(c.sortValue ? c.sortValue(row) : '').toLowerCase().includes(f))
      )
    }
    if (sortCol !== null) {
      const col = columns[sortCol]
      if (col.sortValue) {
        result = [...result].sort((a, b) => {
          const av = col.sortValue!(a)
          const bv = col.sortValue!(b)
          const cmp = av < bv ? -1 : av > bv ? 1 : 0
          return sortAsc ? cmp : -cmp
        })
      }
    }
    return result
  }

  function renderHead() {
    theadRow.innerHTML = ''
    columns.forEach((c, i) => {
      const th = h(
        'th',
        {
          style: c.sortValue ? 'cursor:pointer;white-space:nowrap' : 'white-space:nowrap',
          onclick: c.sortValue
            ? () => {
                if (sortCol === i) {
                  sortAsc = !sortAsc
                } else {
                  sortCol = i
                  sortAsc = true
                }
                renderHead()
                renderBody()
              }
            : undefined,
        },
        [c.header + (sortCol === i ? (sortAsc ? ' ▲' : ' ▼') : '')]
      )
      theadRow.append(th)
    })
  }

  function renderBody() {
    tbody.innerHTML = ''
    for (const row of getRows()) {
      const tr = h('tr', {}, [])
      for (const c of columns) {
        tr.append(h('td', {}, [c.render(row)]))
      }
      tbody.append(tr)
    }
  }

  const filterInput = h('input', {
    type: 'text',
    className: 'form-control bg-dark text-light border-secondary mb-2',
    placeholder: 'Filter columns...',
    oninput: (e: Event) => {
      filter = (e.target as HTMLInputElement).value
      renderBody()
    },
  })

  renderHead()
  renderBody()

  const table = h('table', { className: 'table table-dark table-hover table-bordered w-100' }, [
    h('thead', {}, [theadRow]),
    tbody,
  ])

  wrapper.append(filterInput, table)
  return wrapper
}
