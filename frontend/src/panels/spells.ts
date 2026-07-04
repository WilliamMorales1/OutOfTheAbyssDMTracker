import { api } from '../api.js'
import { h, mount } from '../dom.js'
import type { SpellDetail, SpellRow } from '../types.js'

function statRow(label: string, value: string): Node | null {
  if (!value) return null
  return h('div', {}, [h('strong', {}, [`${label}. `]), value])
}

function levelSchoolLine(s: SpellDetail): string {
  const level = s.level === 0 ? 'Cantrip' : `${s.level}${ordinal(s.level)}-level`
  const school = s.school.toLowerCase()
  const ritual = s.ritual ? ' (ritual)' : ''
  return s.level === 0 ? `${school} cantrip${ritual}` : `${level} ${school}${ritual}`
}

function ordinal(n: number): string {
  if (n === 1) return 'st'
  if (n === 2) return 'nd'
  if (n === 3) return 'rd'
  return 'th'
}

function spellStatBlock(s: SpellDetail): Node {
  return h('div', { className: 'card' }, [
    h('div', { className: 'card-body' }, [
      h('h3', { className: 'text-yellow-400 mb-0 text-xl font-bold' }, [s.name]),
      h('div', { className: 'italic text-gray-400 mb-2' }, [levelSchoolLine(s)]),

      statRow('Casting Time', s.castingTime),
      statRow('Range', s.range),
      statRow('Components', s.components),
      statRow('Duration', s.concentration ? `Concentration, ${s.duration}` : s.duration),
      statRow('Classes', s.classes),

      h('hr', { className: 'border-gray-600 my-3' }, []),

      s.description
        ? h('div', { className: 'whitespace-pre-line' }, [s.description])
        : null,

      s.higherLevel
        ? h('p', { className: 'mt-2' }, [h('strong', { className: 'italic' }, ['At Higher Levels. ']), s.higherLevel])
        : null,

      h('div', { className: 'text-sm text-gray-400 mt-3' }, [s.source ? `Source: ${s.source}` : '']),
    ]),
  ])
}

export async function spellsPanel(): Promise<Node> {
  const list = (await api.spells()) as SpellRow[]

  if (list.length === 0) {
    return h('div', { className: 'text-gray-400' }, ['No spells found. Run: go run ./cmd/ingest-5etools'])
  }

  const detail = h('div', { className: 'mt-3' }, []) as HTMLDivElement

  const schools = Array.from(new Set(list.map((s) => s.school).filter(Boolean))).sort((a, b) => a.localeCompare(b))
  let schoolFilter = ''
  let levelFilter = ''

  let selectSeq = 0

  async function selectSpell(s: SpellRow) {
    const seq = ++selectSeq
    nameInput.value = s.name
    hideSuggestions()
    detail.innerHTML = ''
    detail.append(h('div', { className: 'text-gray-400' }, ['Loading...']))
    const full = await api.spell(s.id)
    if (seq !== selectSeq) return
    mount(detail, spellStatBlock(full))
  }

  const suggestions = h('div', {
    className: 'spell-suggestions list-group shadow border border-gray-600 fixed z-[2000] max-h-[280px] overflow-y-auto hidden',
  }) as HTMLDivElement
  document.body.append(suggestions)

  function hideSuggestions() {
    suggestions.classList.add('hidden')
    suggestions.innerHTML = ''
  }

  function positionSuggestions() {
    const rect = nameInput.getBoundingClientRect()
    suggestions.style.left = `${rect.left}px`
    suggestions.style.width = `${rect.width}px`
    suggestions.style.top = `${rect.bottom + 4}px`
  }

  function showSuggestions(query: string) {
    const q = query.toLowerCase()
    const matches = list
      .filter(
        (s) =>
          (schoolFilter === '' || s.school === schoolFilter) &&
          (levelFilter === '' || String(s.level) === levelFilter) &&
          (q === '' || s.name.toLowerCase().includes(q))
      )
      .slice()
      .sort((a, b) => a.name.localeCompare(b.name))
    if (matches.length === 0) {
      hideSuggestions()
      return
    }
    suggestions.innerHTML = ''
    for (const s of matches) {
      suggestions.append(
        h('button', {
          type: 'button',
          className: 'list-group-item list-group-item-action flex justify-between items-center gap-2',
          onmousedown: (e: Event) => {
            e.preventDefault()
            selectSpell(s)
          },
        }, [
          h('span', {}, [s.name]),
          h('span', { className: 'text-gray-400 text-sm whitespace-nowrap' }, [
            [s.level === 0 ? 'Cantrip' : `Lvl ${s.level}`, s.school].filter(Boolean).join(' · '),
          ]),
        ])
      )
    }
    positionSuggestions()
    suggestions.classList.remove('hidden')
  }

  const nameInput = h('input', {
    type: 'text',
    className: 'form-control',
    placeholder: 'Type a spell name...',
    autocomplete: 'off',
    oninput: (e: Event) => showSuggestions((e.target as HTMLInputElement).value),
    onfocus: (e: Event) => showSuggestions((e.target as HTMLInputElement).value),
    onblur: () => hideSuggestions(),
  }) as HTMLInputElement

  const schoolSelect = h('select', {
    className: 'form-select w-auto',
    onchange: (e: Event) => {
      schoolFilter = (e.target as HTMLSelectElement).value
      if (document.activeElement === nameInput) showSuggestions(nameInput.value)
    },
  }, [
    h('option', { value: '' }, ['All Schools']),
    ...schools.map((sc) => h('option', { value: sc }, [sc])),
  ]) as HTMLSelectElement

  const levelSelect = h('select', {
    className: 'form-select w-auto',
    onchange: (e: Event) => {
      levelFilter = (e.target as HTMLSelectElement).value
      if (document.activeElement === nameInput) showSuggestions(nameInput.value)
    },
  }, [
    h('option', { value: '' }, ['All Levels']),
    ...[0, 1, 2, 3, 4, 5, 6, 7, 8, 9].map((lvl) =>
      h('option', { value: String(lvl) }, [lvl === 0 ? 'Cantrip' : `Level ${lvl}`])
    ),
  ]) as HTMLSelectElement

  return h('div', {}, [
    h('div', { className: 'flex gap-2 mb-2' }, [nameInput, schoolSelect, levelSelect]),
    detail,
  ])
}
