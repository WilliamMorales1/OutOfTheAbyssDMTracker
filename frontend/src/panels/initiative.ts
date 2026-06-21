import { api } from '../api.js'
import { h } from '../dom.js'
import type { MonsterRow } from '../types.js'

interface Combatant {
  id: number
  name: string
  init: number
  hp: number
  maxHp: number
  ac: number
}

let combatants: Combatant[] = []
let nextId = 1
let round = 1
let turn = 0

let monstersLoaded = false
const monsterMap = new Map<string, MonsterRow>()

async function loadMonsters() {
  if (monstersLoaded) return
  const list = (await api.monsters()) as MonsterRow[]
  for (const m of list) monsterMap.set(m.name, m)
  monstersLoaded = true
}

function dexMod(dex: number): number {
  return Math.floor((dex - 10) / 2)
}

export async function initiativePanel(): Promise<Node> {
  await loadMonsters()

  const root = h('div', {}, [])
  let selectedMonster: MonsterRow | null = null

  function sorted(): Combatant[] {
    return [...combatants].sort((a, b) => b.init - a.init)
  }

  function render() {
    root.innerHTML = ''
    document.querySelectorAll('.initiative-suggestions').forEach((el) => el.remove())

    const order = sorted()

    const roundDisplay = h('button', { type: 'button', className: 'btn btn-warning text-dark disabled fw-bold' }, [`Round ${round}`])

    const prevBtn = h('button', {
      type: 'button',
      className: 'btn btn-outline-warning',
      onclick: () => {
        if (order.length === 0) return
        turn--
        if (turn < 0) {
          turn = Math.max(order.length - 1, 0)
          round = Math.max(round - 1, 1)
        }
        render()
      },
    }, ['Prev Turn'])

    const nextBtn = h('button', {
      type: 'button',
      className: 'btn btn-warning',
      onclick: () => {
        if (order.length === 0) return
        turn++
        if (turn >= order.length) {
          turn = 0
          round++
        }
        render()
      },
    }, ['Next Turn'])

    const resetBtn = h('button', {
      type: 'button',
      className: 'btn btn-outline-danger',
      onclick: () => {
        round = 1
        turn = 0
        render()
      },
    }, ['Reset Rounds'])

    const controls = h('div', { className: 'd-flex align-items-stretch gap-2 mb-3' }, [roundDisplay, prevBtn, nextBtn, resetBtn])

    const rows = order.map((c, idx) => {
      const isActive = idx === turn

      const hpDisplay = h('span', { className: c.hp <= 0 ? 'text-danger fw-bold' : '' }, [`${c.hp} / ${c.maxHp}`])

      const dmgInput = h('input', {
        type: 'number',
        className: 'form-control bg-dark text-light border-secondary',
        style: 'width:70px',
        placeholder: 'Amount',
      }) as HTMLInputElement

      const applyDmg = h('button', {
        type: 'button',
        className: 'btn btn-sm btn-outline-danger',
        onclick: () => {
          c.hp = Math.max(0, c.hp - (Number(dmgInput.value) || 0))
          dmgInput.value = ''
          render()
        },
      }, ['Dmg'])

      const applyHeal = h('button', {
        type: 'button',
        className: 'btn btn-sm btn-outline-success',
        onclick: () => {
          c.hp = Math.min(c.maxHp, c.hp + (Number(dmgInput.value) || 0))
          dmgInput.value = ''
          render()
        },
      }, ['Heal'])

      const removeBtn = h('button', {
        type: 'button',
        className: 'btn btn-sm btn-outline-secondary',
        onclick: () => {
          combatants = combatants.filter((x) => x.id !== c.id)
          if (turn >= combatants.length) turn = 0
          render()
        },
      }, ['✕'])

      return h('tr', { className: isActive ? 'table-active' : '' }, [
        h('td', {}, [String(c.init)]),
        h('td', {}, [h('strong', {}, [c.name])]),
        h('td', {}, [String(c.ac)]),
        h('td', {}, [hpDisplay]),
        h('td', {}, [h('div', { className: 'd-flex gap-2 align-items-center' }, [dmgInput, applyDmg, applyHeal])]),
        h('td', {}, [removeBtn]),
      ])
    })

    const suggestions = h('div', {
      className: 'initiative-suggestions list-group shadow rounded-2 overflow-hidden border border-secondary',
      style: 'position:fixed;z-index:2000;max-height:240px;overflow-y:auto;display:none',
    }) as HTMLDivElement
    document.body.append(suggestions)

    function hideSuggestions() {
      suggestions.style.display = 'none'
      suggestions.innerHTML = ''
    }

    function selectMonster(m: MonsterRow) {
      selectedMonster = m
      nameInput.value = m.name
      acInput.value = String(m.ac)
      hpInput.value = String(m.hp)
      hideSuggestions()
    }

    function positionSuggestions() {
      const rect = nameInput.getBoundingClientRect()
      suggestions.style.left = `${rect.left}px`
      suggestions.style.width = `${rect.width}px`
      suggestions.style.bottom = `${window.innerHeight - rect.top + 4}px`
    }

    function showSuggestions(query: string) {
      if (query.length < 2) {
        hideSuggestions()
        return
      }
      const q = query.toLowerCase()
      const matches = [...monsterMap.values()].filter((m) => m.name.toLowerCase().includes(q)).slice(0, 8)
      if (matches.length === 0) {
        hideSuggestions()
        return
      }
      suggestions.innerHTML = ''
      for (const m of matches) {
        suggestions.append(
          h('button', {
            type: 'button',
            className: 'list-group-item list-group-item-action bg-dark text-light border-secondary d-flex justify-content-between align-items-center gap-2',
            onmousedown: (e: Event) => {
              e.preventDefault()
              selectMonster(m)
            },
          }, [
            h('span', {}, [m.name]),
            h('span', { className: 'text-secondary small text-nowrap' }, [`AC ${m.ac} · HP ${m.hp}`]),
          ])
        )
      }
      positionSuggestions()
      suggestions.style.display = 'block'
    }

    const nameInput = h('input', {
      type: 'text',
      className: 'form-control bg-dark text-light border-secondary',
      placeholder: 'e.g. Goblin',
      autocomplete: 'off',
      oninput: (e: Event) => {
        const value = (e.target as HTMLInputElement).value
        selectedMonster = monsterMap.get(value) ?? null
        showSuggestions(value)
      },
      onblur: () => hideSuggestions(),
    }) as HTMLInputElement
    const initInput = h('input', { type: 'number', className: 'form-control bg-dark text-light border-secondary', style: 'width:90px', placeholder: '0' }) as HTMLInputElement
    const acInput = h('input', { type: 'number', className: 'form-control bg-dark text-light border-secondary', style: 'width:90px', placeholder: '0' }) as HTMLInputElement
    const hpInput = h('input', { type: 'number', className: 'form-control bg-dark text-light border-secondary', style: 'width:90px', placeholder: '0' }) as HTMLInputElement

    const rollInitBtn = h('button', {
      type: 'button',
      className: 'btn btn-outline-warning d-flex align-items-center justify-content-center px-2',
      title: selectedMonster ? `1d20 + ${dexMod(selectedMonster.dex)} (DEX)` : '1d20 (select a monster for its DEX mod)',
      onclick: () => {
        const mod = selectedMonster ? dexMod(selectedMonster.dex) : 0
        initInput.value = String(Math.floor(Math.random() * 20) + 1 + mod)
      },
    }, [
      h('span', {
        className: 'd20-icon',
        style: '-webkit-mask-image:url(/images/d20.png);mask-image:url(/images/d20.png)',
      }, []),
    ])

    function addCombatant() {
      const name = nameInput.value.trim()
      if (!name) return
      const hp = Number(hpInput.value) || 0
      combatants.push({
        id: nextId++,
        name,
        init: Number(initInput.value) || 0,
        hp,
        maxHp: hp,
        ac: Number(acInput.value) || 0,
      })
      selectedMonster = null
      render()
    }

    const addBtn = h('button', { type: 'button', className: 'btn btn-warning', onclick: addCombatant }, ['Add to Tracker'])

    const addRow = h('tr', {}, [
      h('td', {}, [h('div', { className: 'd-flex gap-2 align-items-stretch' }, [initInput, rollInitBtn])]),
      h('td', {}, [nameInput]),
      h('td', {}, [acInput]),
      h('td', {}, [hpInput]),
      h('td', {}, [addBtn]),
      h('td', {}, []),
    ])

    const table = h('div', { className: 'table-responsive' }, [
      h('table', { className: 'table table-dark table-hover table-bordered w-100' }, [
        h('thead', {}, [
          h('tr', {}, [
            h('th', {}, ['Initiative']),
            h('th', {}, ['Name']),
            h('th', {}, ['AC']),
            h('th', {}, ['HP']),
            h('th', {}, ['Adjust']),
            h('th', {}, ['']),
          ]),
        ]),
        h('tbody', {}, [...rows, addRow]),
      ]),
    ])

    root.append(controls, table)
  }

  render()
  return root
}
