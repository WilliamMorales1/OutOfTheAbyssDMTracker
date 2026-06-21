import { api } from '../api.js'
import { h, onSubmitAsync } from '../dom.js'
import type { SearchResult } from '../types.js'

function escapeRegExp(s: string): string {
  return s.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
}

// Splits text into nodes, wrapping case-insensitive matches of any query
// term in a <mark> element. Builds DOM nodes directly (no innerHTML) so
// result content can never be interpreted as markup.
function highlight(text: string, query: string): (Node | string)[] {
  const terms = query
    .split(/\s+/)
    .map((t) => t.trim())
    .filter(Boolean)
    .map(escapeRegExp)
  if (terms.length === 0) return [text]

  const pattern = `(${terms.join('|')})`
  const parts = text.split(new RegExp(pattern, 'gi'))
  const isMatch = new RegExp(`^${pattern}$`, 'i')
  return parts.map((part) => (isMatch.test(part) ? h('mark', { className: 'bg-warning text-dark' }, [part]) : part))
}

export function searchPanel(): Node {
  const status = h('div', {}, [])
  const results = h('div', { className: 'mt-3' }, [h('p', { className: 'text-secondary' }, ['Enter a query above.'])])

  const input = h('input', {
    name: 'q',
    type: 'text',
    className: 'form-control bg-dark text-light border-secondary',
    placeholder: "Search lore semantically... (e.g. 'drow priestess tactics')",
    required: true,
  }) as HTMLInputElement

  const submitBtn = h('button', { className: 'btn btn-warning', type: 'submit' }, ['Search']) as HTMLButtonElement

  const form = h('form', {}, [
    h('div', { className: 'input-group mb-3' }, [input, submitBtn]),
    status,
  ]) as HTMLFormElement

  onSubmitAsync(form, submitBtn, status, 'Searching...', async () => {
    const q = input.value.trim()
    if (!q) return
    const res = (await api.search(q)) as SearchResult[]
    results.innerHTML = ''
    if (res.length === 0) {
      results.append(h('p', { className: 'text-secondary' }, ['No results found.']))
    } else {
      for (const r of res) {
        results.append(
          h('div', { className: 'card bg-dark border-secondary mb-3' }, [
            h('div', { className: 'card-header d-flex justify-content-between align-items-center' }, [
              h('strong', { className: 'text-warning' }, [r.chapterTitle]),
              h('span', { className: 'badge bg-secondary' }, [`${(r.score * 100).toFixed(0)}% match`]),
            ]),
            h(
              'div',
              { className: 'card-body text-light', style: 'white-space:pre-wrap;font-size:.875rem' },
              highlight(r.content, q)
            ),
          ])
        )
      }
    }
  })

  return h('div', {}, [form, results])
}
