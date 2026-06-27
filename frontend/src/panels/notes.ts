import { api } from '../api.js'
import { h, onSubmitAsync } from '../dom.js'

function loadMonaco(): Promise<any> {
  return new Promise((resolve) => {
    const w = window as any
    if (w.monaco) { resolve(w.monaco); return }
    const script = document.createElement('script')
    script.src = 'https://cdn.jsdelivr.net/npm/monaco-editor@0.52.2/min/vs/loader.js'
    script.onload = () => {
      w.require.config({ paths: { vs: 'https://cdn.jsdelivr.net/npm/monaco-editor@0.52.2/min/vs' } })
      w.require(['vs/editor/editor.main'], () => resolve(w.monaco))
    }
    document.head.appendChild(script)
  })
}

export async function notesPanel(): Promise<Node> {
  const [names, monaco] = await Promise.all([api.notes(), loadMonaco()])

  const select = h(
    'select',
    { className: 'form-select bg-dark text-light border-secondary' },
    [
      h('option', { value: '' }, ['Select a note...']),
      ...names.map((n) => h('option', { value: n }, [n])),
    ]
  ) as HTMLSelectElement

  const newNameInput = h('input', {
    type: 'text',
    className: 'form-control bg-dark text-light border-secondary',
    placeholder: 'new-note.md',
  }) as HTMLInputElement
  const newBtn = h('button', { type: 'submit', className: 'btn btn-outline-warning' }, ['Create'])
  const newStatus = h('span', { className: 'ms-2' }, [])
  const newForm = h('form', { className: 'd-flex gap-2 align-items-center' }, [newNameInput, newBtn, newStatus]) as HTMLFormElement

  const editorDiv = h('div', {
    className: 'border border-secondary rounded',
    style: 'height:500px;',
  }, []) as HTMLElement

  const editor = monaco.editor.create(editorDiv, {
    language: 'markdown',
    theme: 'vs-dark',
    automaticLayout: true,
    minimap: { enabled: false },
    wordWrap: 'on',
    readOnly: true,
  })

  let isDirty = false
  let currentNoteName = ''

  editor.onDidChangeModelContent(() => { isDirty = true })

  async function saveIfDirty() {
    if (!isDirty || !currentNoteName) return
    await api.saveNote(currentNoteName, editor.getValue())
    isDirty = false
  }

  async function loadNote(name: string) {
    await saveIfDirty()
    currentNoteName = name
    isDirty = false
    if (!name) {
      editor.setValue('')
      editor.updateOptions({ readOnly: true })
      return
    }
    const note = await api.note(name)
    editor.setValue(note.content)
    editor.updateOptions({ readOnly: false })
  }

  select.addEventListener('change', () => loadNote(select.value))

  onSubmitAsync(newForm, newBtn, newStatus, 'Creating...', async () => {
    let name = newNameInput.value.trim()
    if (!name) throw new Error('Name required')
    if (!name.endsWith('.md')) name += '.md'
    if (!/^[A-Za-z0-9_-]+\.md$/.test(name)) throw new Error('Use letters, numbers, _ or - only')
    await api.saveNote(name, '')
    select.append(h('option', { value: name }, [name]))
    select.value = name
    newNameInput.value = ''
    await loadNote(name)
  })

  const beforeUnloadHandler = () => {
    if (isDirty && currentNoteName) {
      fetch(`/api/notes/${encodeURIComponent(currentNoteName)}`, { method: 'PUT', body: editor.getValue(), keepalive: true })
    }
  }
  window.addEventListener('beforeunload', beforeUnloadHandler)

  const root = h('div', {}, [
    h('div', { className: 'd-flex gap-3 align-items-center mb-3' }, [select, newForm]),
    editorDiv,
  ]) as any
  root.__saveIfDirty = saveIfDirty
  return root as Node
}
