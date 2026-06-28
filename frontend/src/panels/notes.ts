import { api } from '../api.js'
import { h, onSubmitAsync } from '../dom.js'
import { marked } from 'marked'

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
    { className: 'form-select bg-dark text-light border-secondary', style: 'width: auto;' },
    [
      h('option', { value: '' }, ['Select a note...']),
      ...names.map((n) => h('option', { value: n }, [n])),
    ]
  ) as HTMLSelectElement

  const newNameInput = h('input', {
    type: 'text',
    className: 'form-control bg-dark text-light border-secondary',
    placeholder: 'new-note.md',
    style: 'width: 160px;',
  }) as HTMLInputElement
  const newBtn = h('button', { type: 'submit', className: 'btn btn-outline-warning' }, ['Create'])
  const newStatus = h('span', {}, [])
  const previewBtn = h('button', {
    type: 'button',
    className: 'btn btn-outline-warning',
  }, ['Preview']) as HTMLButtonElement
  const newForm = h('form', { className: 'd-flex gap-2 align-items-center' }, [newNameInput, newBtn, previewBtn, newStatus]) as HTMLFormElement

  const editorDiv = h('div', {
    className: 'border border-secondary rounded',
    style: 'height:500px;',
  }, []) as HTMLElement

  const previewDiv = h('div', {
    className: 'border border-secondary rounded p-3 overflow-auto markdown-preview',
    style: 'height:500px; display:none;',
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
  let isPreviewing = false

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
    if (isPreviewing) togglePreview()
    if (!name) {
      editor.setValue('')
      editor.updateOptions({ readOnly: true })
      return
    }
    const note = await api.note(name)
    editor.setValue(note.content)
    editor.updateOptions({ readOnly: false })
  }

  function togglePreview() {
    isPreviewing = !isPreviewing
    if (isPreviewing) {
      previewDiv.innerHTML = marked(editor.getValue()) as string
      editorDiv.style.display = 'none'
      previewDiv.style.display = ''
      previewBtn.textContent = 'Edit'
    } else {
      editorDiv.style.display = ''
      previewDiv.style.display = 'none'
      previewBtn.textContent = 'Preview'
    }
  }

  previewBtn.addEventListener('click', togglePreview)

  select.addEventListener('change', () => loadNote(select.value))

  onSubmitAsync(newForm, newBtn, newStatus, 'Creating...', async () => {
    let name = newNameInput.value.trim()
    if (!name) { alert('Name required'); return }
    if (!name.endsWith('.md')) name += '.md'
    if (!/^[A-Za-z0-9_-]+\.md$/.test(name)) { alert('Use letters, numbers, _ or - only'); return }
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
    h('div', { className: 'd-flex gap-2 align-items-center mb-3' }, [select, newForm]),
    editorDiv,
    previewDiv,
  ]) as any
  root.__saveIfDirty = saveIfDirty
  return root as Node
}
