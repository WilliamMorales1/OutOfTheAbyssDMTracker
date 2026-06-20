import { api } from '../api.js'
import { h, onSubmitAsync } from '../dom.js'

export async function notesPanel(): Promise<Node> {
  const names = await api.notes()

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
    placeholder: 'new-note-name.md',
  }) as HTMLInputElement
  const newBtn = h('button', { type: 'submit', className: 'btn btn-outline-warning' }, ['Create'])
  const newStatus = h('span', { className: 'ms-2' }, [])
  const newForm = h('form', { className: 'd-flex gap-2 align-items-center' }, [newNameInput, newBtn, newStatus]) as HTMLFormElement

  const editor = h('textarea', {
    className: 'form-control bg-dark text-light border-secondary font-monospace',
    rows: 20,
  }) as HTMLTextAreaElement
  editor.setAttribute('disabled', '')

  const saveBtn = h('button', { type: 'submit', className: 'btn btn-outline-warning' }, ['Save'])
  saveBtn.setAttribute('disabled', '')

  const downloadBtn = h(
    'button',
    {
      type: 'button',
      className: 'btn btn-outline-warning',
      onclick: () => {
        if (!select.value) return
        const blob = new Blob([editor.value], { type: 'text/markdown' })
        const url = URL.createObjectURL(blob)
        const a = h('a', { href: url, download: select.value }, []) as HTMLAnchorElement
        a.click()
        URL.revokeObjectURL(url)
      },
    },
    ['Download']
  ) as HTMLButtonElement
  downloadBtn.setAttribute('disabled', '')

  const saveStatus = h('span', { className: 'ms-2' }, [])
  const saveForm = h('form', { className: 'd-flex align-items-center gap-2 mt-2' }, [downloadBtn, saveBtn, saveStatus]) as HTMLFormElement

  async function loadNote(name: string) {
    if (!name) {
      editor.value = ''
      editor.setAttribute('disabled', '')
      saveBtn.setAttribute('disabled', '')
      downloadBtn.setAttribute('disabled', '')
      return
    }
    const note = await api.note(name)
    editor.value = note.content
    editor.removeAttribute('disabled')
    saveBtn.removeAttribute('disabled')
    downloadBtn.removeAttribute('disabled')
  }

  select.addEventListener('change', () => loadNote(select.value))

  onSubmitAsync(saveForm, saveBtn, saveStatus, 'Saving...', async () => {
    if (!select.value) throw new Error('No note selected')
    await api.saveNote(select.value, editor.value)
  })

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

  const root = h('div', {}, [
    h('div', { className: 'd-flex gap-3 align-items-center mb-3' }, [select, newForm]),
    editor,
    saveForm,
  ])

  return root
}
