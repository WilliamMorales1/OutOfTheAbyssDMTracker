import { api } from '../api.js'
import { h, mount } from '../dom.js'
import type { MonsterDetail, MonsterRow, StatBlockEntry } from '../types.js'
import { CR_OPTIONS, crToNumber, scaleMonster } from '../monster-scaling.js'

// baseType strips a parenthetical suffix so "aberration (beholder)" filters
// together with plain "aberration".
function baseType(type: string): string {
  return type.replace(/\s*\(.*\)\s*$/, '')
}

function statRow(label: string, value: string): Node | null {
  if (!value) return null
  return h('div', {}, [h('strong', {}, [`${label}. `]), value])
}

function entryBlock(title: string, entries: StatBlockEntry[] | null): Node | null {
  if (!entries || entries.length === 0) return null
  return h('div', { className: 'mb-3' }, [
    h('h5', { className: 'text-yellow-400 border-b border-gray-600 pb-1 font-semibold' }, [title]),
    ...entries.map((e) =>
      h('p', { className: 'mb-2' }, [h('strong', { className: 'italic' }, [`${e.name}. `]), e.text])
    ),
  ])
}

function abilityScore(label: string, score: number): Node {
  const mod = Math.floor((score - 10) / 2)
  const modText = mod >= 0 ? `+${mod}` : `${mod}`
  return h('div', { className: 'text-center min-w-[56px]' }, [
    h('div', { className: 'text-sm text-gray-400' }, [label]),
    h('div', { className: 'font-bold' }, [String(score)]),
    h('div', { className: 'text-sm text-gray-400' }, [`(${modText})`]),
  ])
}

function monsterStatBlock(m: MonsterDetail, onScale: (e: Event) => void, baseCr: string): Node {
  const headerLine = [m.size, m.type, m.alignment].filter(Boolean).join(', ')

  return h('div', { className: 'card' }, [
    h('div', { className: 'card-body' }, [
      h('div', { className: 'flex flex-wrap gap-4' }, [
        h('div', { className: 'flex-1 min-w-[280px]' }, [
          h('div', { className: 'flex items-center gap-2 mb-0' }, [
            m.tokenUrl
              ? h('div', { className: 'group relative z-10 w-12 h-12 shrink-0' }, [
                  h('img', {
                    src: m.tokenUrl,
                    alt: '',
                    className:
                      'absolute inset-0 w-full h-full rounded-full border border-gray-600 object-cover cursor-zoom-in group-hover:shadow-xl group-hover:border-gray-400 group-hover:scale-[5] transition-transform duration-150 origin-center',
                    onerror: (e: Event) => {
                      ;(e.target as HTMLImageElement).style.display = 'none'
                    },
                  }),
                ])
              : null,
            h('h3', { className: 'text-yellow-400 mb-0 text-xl font-bold' }, [m.name]),
            h('select', {
              className: 'form-select inline-block w-auto ml-2',
              onchange: onScale,
            }, CR_OPTIONS.map((cr) => h('option', { value: cr, selected: cr === m.cr }, [`CR ${cr}${cr === baseCr ? ' (actual)' : ''}`]))),
          ]),
          h('div', { className: 'italic text-gray-400 mb-2' }, [headerLine || '—']),

          h('div', {}, [h('strong', {}, ['Armor Class. ']), `${m.ac}${m.acDesc ? ` (${m.acDesc})` : ''}`]),
          h('div', {}, [h('strong', {}, ['Hit Points. ']), `${m.hp}${m.hpFormula ? ` (${m.hpFormula})` : ''}`]),
          h('div', { className: 'mb-2' }, [h('strong', {}, ['Speed. ']), m.speed || '—']),

          h('div', { className: 'flex gap-3 my-3 border-t border-b border-gray-600 py-2' }, [
            abilityScore('STR', m.str),
            abilityScore('DEX', m.dex),
            abilityScore('CON', m.con),
            abilityScore('INT', m.int),
            abilityScore('WIS', m.wis),
            abilityScore('CHA', m.cha),
          ]),

          statRow('Saving Throws', m.savingThrows),
          statRow('Skills', m.skills),
          statRow('Damage Vulnerabilities', m.vulnerabilities),
          statRow('Damage Resistances', m.damageResistances),
          statRow('Damage Immunities', m.damageImmunities),
          statRow('Condition Immunities', m.conditionImmunities),
          h('div', {}, [
            h('strong', {}, ['Senses. ']),
            [m.senses, `passive Perception ${m.passivePerception}`].filter(Boolean).join(', '),
          ]),
          statRow('Languages', m.languages || '—'),
          h('div', { className: 'mb-2' }, [h('strong', {}, ['Challenge. ']), `${m.cr || '—'}`]),
          statRow('Environment', m.environment),
          h('div', { className: 'text-sm text-gray-400' }, [m.source ? `Source: ${m.source}` : '']),
        ]),

        m.imageUrl
          ? h('img', {
              src: m.imageUrl,
              alt: m.name,
              className: 'rounded border border-gray-600 max-w-[320px] max-h-[420px] object-contain',
            })
          : null,
      ]),

      h('hr', { className: 'border-gray-600 my-3' }, []),

      entryBlock('Traits', m.traits),
      entryBlock('Actions', m.actions),
      entryBlock('Reactions', m.reactions),
      entryBlock('Legendary Actions', m.legendaryActions),
      entryBlock('Spellcasting', m.spellcasting),

      m.notes ? h('div', { className: 'mt-3 text-gray-400' }, [h('strong', {}, ['Notes. ']), m.notes]) : null,
    ]),
  ])
}

// Wraps the stat block with a CR scaler (mirroring 5etools' scale-creature
// control): picking a target CR recomputes HP/AC/to-hit/DC/damage via
// monster-scaling.ts and re-renders the block without refetching.
function monsterDetailView(base: MonsterDetail): Node {
  const baseCr = base.cr || '0'
  const body = h('div', {}, []) as HTMLDivElement

  function render(m: MonsterDetail) {
    mount(body, monsterStatBlock(m, onScale, baseCr))
  }

  function onScale(e: Event) {
    const value = (e.target as HTMLSelectElement).value
    render(value === baseCr ? base : scaleMonster(base, crToNumber(value)))
  }

  render(base)
  return body
}

export async function monstersPanel(): Promise<Node> {
  const list = (await api.monsters()) as MonsterRow[]

  if (list.length === 0) {
    return h('div', { className: 'text-gray-400' }, ['No monsters found. Run: go run ./cmd/ingest-5etools'])
  }

  const detail = h('div', { className: 'mt-3' }, []) as HTMLDivElement

  const types = Array.from(new Set(list.map((m) => baseType(m.type)).filter(Boolean))).sort((a, b) =>
    a.localeCompare(b)
  )
  let typeFilter = ''

  let selectSeq = 0

  async function selectMonster(m: MonsterRow) {
    const seq = ++selectSeq
    nameInput.value = m.name
    hideSuggestions()
    detail.innerHTML = ''
    detail.append(h('div', { className: 'text-gray-400' }, ['Loading...']))
    const full = await api.monster(m.id)
    if (seq !== selectSeq) return
    mount(detail, monsterDetailView(full))
  }

  const suggestions = h('div', {
    className: 'monster-suggestions list-group shadow border border-gray-600 fixed z-[2000] max-h-[280px] overflow-y-auto hidden',
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
        (m) => (typeFilter === '' || baseType(m.type) === typeFilter) && (q === '' || m.name.toLowerCase().includes(q))
      )
      .slice()
      .sort((a, b) => a.name.localeCompare(b.name))
    if (matches.length === 0) {
      hideSuggestions()
      return
    }
    suggestions.innerHTML = ''
    for (const m of matches) {
      suggestions.append(
        h('button', {
          type: 'button',
          className: 'list-group-item list-group-item-action flex justify-between items-center gap-2',
          onmousedown: (e: Event) => {
            e.preventDefault()
            selectMonster(m)
          },
        }, [
          h('span', {}, [m.name]),
          h('span', { className: 'text-gray-400 text-sm whitespace-nowrap' }, [
            [m.type, m.cr ? `CR ${m.cr}` : null].filter(Boolean).join(' · '),
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
    placeholder: 'Type a monster name...',
    autocomplete: 'off',
    oninput: (e: Event) => showSuggestions((e.target as HTMLInputElement).value),
    onfocus: (e: Event) => showSuggestions((e.target as HTMLInputElement).value),
    onblur: () => hideSuggestions(),
  }) as HTMLInputElement

  const typeSelect = h('select', {
    className: 'form-select w-auto',
    onchange: (e: Event) => {
      typeFilter = (e.target as HTMLSelectElement).value
      if (document.activeElement === nameInput) showSuggestions(nameInput.value)
    },
  }, [
    h('option', { value: '' }, ['All Types']),
    ...types.map((t) => h('option', { value: t }, [t])),
  ]) as HTMLSelectElement

  return h('div', {}, [
    h('div', { className: 'flex gap-2 mb-2' }, [nameInput, typeSelect]),
    detail,
  ])
}
