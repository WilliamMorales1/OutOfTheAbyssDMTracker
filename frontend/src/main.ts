import { h, mount } from './dom.js'
import { sessionsPanel } from './panels/sessions.js'
import { npcsPanel } from './panels/npcs.js'
import { monstersPanel } from './panels/monsters.js'
import { mapsPanel } from './panels/maps.js'
import { chatPanel } from './panels/chat.js'
import { searchPanel } from './panels/search.js'
import { notesPanel } from './panels/notes.js'
import { initiativePanel } from './panels/initiative.js'
import { refsPanel } from './panels/references.js'

const tabs = [
  { name: 'Sessions', path: 'sessions', load: sessionsPanel },
  { name: 'Notes', path: 'notes', load: notesPanel },
  { name: 'NPCs', path: 'npcs', load: npcsPanel },
  { name: 'Monsters', path: 'monsters', load: monstersPanel },
  { name: 'Maps', path: 'maps', load: mapsPanel },
  { name: 'References', path: 'refs', load: refsPanel },
  { name: 'Initiative', path: 'initiative', load: async () => initiativePanel() },
  { name: 'Ask Agent', path: 'chat', load: async () => chatPanel() },
  { name: 'Lore Search', path: 'search', load: async () => searchPanel() },
] as const

const root = document.getElementById('root')!

const header = h('header', { className: 'bg-black py-3 mb-4 border-b border-gray-600' }, [
  h('div', { className: 'container' }, [
    h('h1', { className: 'text-5xl font-bold text-yellow-400 mb-0' }, ['Out of the Abyss']),
  ]),
])

const navList = h('ul', { className: 'nav-tabs mb-3' }, [])
const panel = h('div', { className: 'bg-gray-600/10 rounded p-3' }, [])

const panelCache = new Map<string, Node>()

async function activate(path: (typeof tabs)[number]['path']) {
  const currentChild = panel.firstChild as any
  if (typeof currentChild?.__saveIfDirty === 'function') await currentChild.__saveIfDirty()

  navList.querySelectorAll('button').forEach((btn) => {
    const isActive = btn.dataset.path === path
    btn.className = `nav-link ${isActive ? 'active' : 'text-gray-100'}`
  })

  const cached = panelCache.get(path)
  if (cached) {
    mount(panel, cached)
    return
  }

  panel.innerHTML = '<p>Loading...</p>'
  const tab = tabs.find((t) => t.path === path)!
  try {
    const node = await tab.load()
    panelCache.set(path, node)
    mount(panel, node)
  } catch (err) {
    mount(panel, h('p', { className: 'text-red-500' }, [String(err)]))
  }
}

for (const t of tabs) {
  const btn = h('button', { 'data-path': t.path, onclick: () => activate(t.path) }, [t.name])
  navList.append(h('li', {}, [btn]))
}

const container = h('div', { className: 'container' }, [navList, panel])

root.append(header, container)
activate('sessions')
