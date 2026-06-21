import { h } from '../dom.js'

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

export function initiativePanel(): Node {
  const root = h('div', {}, [])

  function sorted(): Combatant[] {
    return [...combatants].sort((a, b) => b.init - a.init)
  }

  function render() {
    root.innerHTML = ''

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

    const nameInput = h('input', { type: 'text', className: 'form-control bg-dark text-light border-secondary', placeholder: 'e.g. Goblin' }) as HTMLInputElement
    const initInput = h('input', { type: 'number', className: 'form-control bg-dark text-light border-secondary', style: 'width:90px', placeholder: '0' }) as HTMLInputElement
    const acInput = h('input', { type: 'number', className: 'form-control bg-dark text-light border-secondary', style: 'width:90px', placeholder: '0' }) as HTMLInputElement
    const hpInput = h('input', { type: 'number', className: 'form-control bg-dark text-light border-secondary', style: 'width:90px', placeholder: '0' }) as HTMLInputElement
    
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
      render()
    }

    const addBtn = h('button', { type: 'button', className: 'btn btn-warning', onclick: addCombatant }, ['Add to Tracker'])

    const addRow = h('tr', {}, [
      h('td', {}, [initInput]),
      h('td', {}, [
        h('div', { className: 'd-flex gap-2 align-items-center' }, [
          nameInput,
          h('div', { className: 'form-check d-flex align-items-center text-nowrap' }),
        ]),
      ]),
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
