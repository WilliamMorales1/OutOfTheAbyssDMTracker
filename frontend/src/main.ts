import { h, mount } from './dom.js'
import { sessionsPanel } from './panels/sessions.js'
import { locationsPanel } from './panels/locations.js'
import { npcsPanel } from './panels/npcs.js'
import { encountersPanel } from './panels/encounters.js'
import { monstersPanel } from './panels/monsters.js'
import { eventsPanel } from './panels/events.js'
import { mapsPanel } from './panels/maps.js'
import { chatPanel } from './panels/chat.js'
import { searchPanel } from './panels/search.js'
import { notesPanel } from './panels/notes.js'

const tabs = [
  { name: 'Sessions', path: 'sessions', load: sessionsPanel },
  { name: 'Session Notes', path: 'notes', load: notesPanel },
  { name: 'Locations', path: 'locations', load: locationsPanel },
  { name: 'NPCs', path: 'npcs', load: npcsPanel },
  { name: 'Encounters', path: 'encounters', load: encountersPanel },
  { name: 'Monsters', path: 'monsters', load: monstersPanel },
  { name: 'Events', path: 'events', load: eventsPanel },
  { name: 'Maps', path: 'maps', load: mapsPanel },
  { name: 'Ask Agent', path: 'chat', load: async () => chatPanel() },
  { name: 'Lore Search', path: 'search', load: async () => searchPanel() },
] as const

const root = document.getElementById('root')!

const header = h('header', { className: 'bg-black py-3 mb-4 border-bottom border-secondary' }, [
  h('div', { className: 'container' }, [
    h('h1', { className: 'display-5 fw-bold text-warning mb-0' }, ['Out of the Abyss']),
  ]),
])

const navList = h('ul', { className: 'nav nav-tabs border-secondary mb-3' }, [])
const panel = h('div', { className: 'bg-secondary bg-opacity-10 rounded p-3' }, [])

async function activate(path: (typeof tabs)[number]['path']) {
  navList.querySelectorAll('button').forEach((btn) => {
    const isActive = btn.dataset.path === path
    btn.className = `nav-link bg-transparent ${isActive ? 'active text-warning border-warning' : 'text-light'}`
  })
  panel.innerHTML = '<p>Loading...</p>'
  const tab = tabs.find((t) => t.path === path)!
  try {
    const node = await tab.load()
    mount(panel, node)
  } catch (err) {
    mount(panel, h('p', { className: 'text-danger' }, [String(err)]))
  }
}

for (const t of tabs) {
  const btn = h('button', { 'data-path': t.path, onclick: () => activate(t.path) }, [t.name])
  navList.append(h('li', { className: 'nav-item' }, [btn]))
}

const container = h('div', { className: 'container' }, [navList, panel])

root.append(header, container)
activate('sessions')
