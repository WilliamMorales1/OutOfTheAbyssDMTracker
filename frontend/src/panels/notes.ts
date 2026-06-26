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

  const saveBtn = h('button', { type: 'submit', className: 'btn btn-outline-warning' }, ['Save'])
  saveBtn.setAttribute('disabled', '')

  const downloadBtn = h(
    'button',
    {
      type: 'button',
      className: 'btn btn-outline-warning',
      onclick: () => {
        if (!select.value) return
        const blob = new Blob([editor.getValue()], { type: 'text/markdown' })
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

  // --- LSP client ---
  let ws: WebSocket | null = null
  let lspReady = false
  let openUri = ''
  let docVersion = 0
  let debounce: ReturnType<typeof setTimeout> | null = null
  let tokenReqId = 100
  let pendingTokenUri = ''
  let tokenTypes: string[] = []
  let decorations = editor.createDecorationsCollection([])
  let nextReqId = 200
  const pendingRequests = new Map<number, (result: any) => void>()

  function sendLSP(msg: object) {
    if (ws?.readyState === WebSocket.OPEN) ws.send(JSON.stringify(msg))
  }

  function requestTokens() {
    if (!lspReady || !openUri) return
    tokenReqId++
    pendingTokenUri = openUri
    sendLSP({ jsonrpc: '2.0', id: tokenReqId, method: 'textDocument/semanticTokens/full', params: { textDocument: { uri: openUri } } })
  }

  function applyTokens(data: number[]) {
    const model = editor.getModel()
    if (!model) return
    const decors: any[] = []
    let line = 0, startChar = 0
    for (let i = 0; i + 4 < data.length; i += 5) {
      const deltaLine = data[i], deltaChar = data[i + 1], length = data[i + 2], typeIdx = data[i + 3]
      if (deltaLine > 0) { line += deltaLine; startChar = deltaChar } else { startChar += deltaChar }
      const tokenType = tokenTypes[typeIdx]
      if (!tokenType) continue
      decors.push({
        range: new monaco.Range(line + 1, startChar + 1, line + 1, startChar + length + 1),
        options: { inlineClassName: `ns-token-${tokenType}` },
      })
    }
    decorations.set(decors)
  }

  function didOpen(name: string, text: string) {
    if (!lspReady) return
    openUri = `file:///notes/${name}`
    docVersion = 1
    sendLSP({
      jsonrpc: '2.0', method: 'textDocument/didOpen',
      params: { textDocument: { uri: openUri, languageId: 'markdown', version: 1, text } },
    })
    requestTokens()
  }

  function didClose() {
    if (!lspReady || !openUri) return
    sendLSP({ jsonrpc: '2.0', method: 'textDocument/didClose', params: { textDocument: { uri: openUri } } })
    decorations.set([])
    openUri = ''
  }

  function connectLSP() {
    const proto = location.protocol === 'https:' ? 'wss:' : 'ws:'
    ws = new WebSocket(`${proto}//${location.host}/api/notes/lsp`)
    ws.onopen = () => {
      sendLSP({
        jsonrpc: '2.0', id: 1, method: 'initialize',
        params: { processId: null, capabilities: { textDocument: { semanticTokens: { requests: { full: true }, tokenTypes: [], tokenModifiers: [], formats: ['relative'] }, hover: { contentFormat: ['markdown', 'plaintext'] }, publishDiagnostics: {} } }, rootUri: null },
      })
    }
    ws.onmessage = (ev) => {
      let msg: any
      try { msg = JSON.parse(ev.data as string) } catch { return }
      if (msg.id === 1 && msg.result) {
        tokenTypes = msg.result.capabilities?.semanticTokensProvider?.legend?.tokenTypes ?? []
        sendLSP({ jsonrpc: '2.0', method: 'initialized', params: {} })
        lspReady = true
        if (select.value) didOpen(select.value, editor.getValue())
      }
      if (msg.id === tokenReqId && pendingTokenUri === openUri && msg.result?.data) {
        applyTokens(msg.result.data)
      }
      if (msg.id != null && pendingRequests.has(msg.id)) {
        const resolve = pendingRequests.get(msg.id)!
        pendingRequests.delete(msg.id)
        resolve(msg.result ?? null)
      }
    }
    ws.onclose = () => { lspReady = false; ws = null; openUri = '' }
  }

  editor.onDidChangeModelContent(() => {
    if (debounce !== null) clearTimeout(debounce)
    debounce = setTimeout(() => {
      if (!lspReady || !openUri) return
      docVersion++
      sendLSP({
        jsonrpc: '2.0', method: 'textDocument/didChange',
        params: { textDocument: { uri: openUri, version: docVersion }, contentChanges: [{ text: editor.getValue() }] },
      })
      requestTokens()
    }, 300)
  })

  connectLSP()

  monaco.languages.registerHoverProvider('markdown', {
    provideHover: (_model: any, position: any) => {
      if (!lspReady || !openUri) return null
      const id = nextReqId++
      return new Promise((resolve) => {
        pendingRequests.set(id, (result) => {
          if (!result?.contents) { resolve(null); return }
          const contents = Array.isArray(result.contents)
            ? result.contents
            : [result.contents]
          resolve({
            contents: contents.map((c: any) => ({
              value: typeof c === 'string' ? c : c.value ?? '',
              isTrusted: false,
            })),
          })
        })
        sendLSP({
          jsonrpc: '2.0', id,
          method: 'textDocument/hover',
          params: {
            textDocument: { uri: openUri },
            position: { line: position.lineNumber - 1, character: position.column - 1 },
          },
        })
        setTimeout(() => {
          if (pendingRequests.has(id)) { pendingRequests.delete(id); resolve(null) }
        }, 3000)
      })
    },
  })

  async function loadNote(name: string) {
    didClose()
    if (!name) {
      editor.setValue('')
      editor.updateOptions({ readOnly: true })
      saveBtn.setAttribute('disabled', '')
      downloadBtn.setAttribute('disabled', '')
      return
    }
    const note = await api.note(name)
    editor.setValue(note.content)
    editor.updateOptions({ readOnly: false })
    saveBtn.removeAttribute('disabled')
    downloadBtn.removeAttribute('disabled')
    didOpen(name, note.content)
  }

  select.addEventListener('change', () => loadNote(select.value))

  onSubmitAsync(saveForm, saveBtn, saveStatus, 'Saving...', async () => {
    if (!select.value) throw new Error('No note selected')
    await api.saveNote(select.value, editor.getValue())
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

  return h('div', {}, [
    h('div', { className: 'd-flex gap-3 align-items-center mb-3' }, [select, newForm]),
    editorDiv,
    saveForm,
  ])
}
