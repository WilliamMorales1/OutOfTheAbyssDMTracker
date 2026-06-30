type Attrs = Record<string, string | number | boolean | ((e: Event) => void) | Partial<CSSStyleDeclaration> | undefined>


export function h<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  attrs: Attrs = {},
  children: (Node | string | null | undefined)[] = []
): HTMLElementTagNameMap[K] {
  const el = document.createElement(tag)
  for (const [k, v] of Object.entries(attrs)) {
    if (v === undefined) continue
    if (k.startsWith('on') && typeof v === 'function') {
      el.addEventListener(k.slice(2).toLowerCase(), v as (e: Event) => void)
    } else if (k === 'className') {
      el.setAttribute('class', String(v))
    } else if (k === 'style' && typeof v === 'object') {
      Object.assign(el.style, v)
    } else if (typeof v === 'boolean') {
      if (v) el.setAttribute(k, '')
    } else {
      el.setAttribute(k, String(v))
    }
  }
  for (const c of children) {
    if (c === null || c === undefined) continue
    el.append(typeof c === 'string' ? document.createTextNode(c) : c)
  }
  return el
}

export function svg(tag: string, attrs: Record<string, string | number> = {}, children: (Node | string)[] = []): SVGElement {
  const el = document.createElementNS('http://www.w3.org/2000/svg', tag)
  for (const [k, v] of Object.entries(attrs)) {
    el.setAttribute(k, String(v))
  }
  for (const c of children) {
    el.append(typeof c === 'string' ? document.createTextNode(c) : c)
  }
  return el
}

export function clear(el: Element) {
  el.innerHTML = ''
}

export function mount(parent: Element, child: Node) {
  clear(parent)
  parent.append(child)
}

// Wires a form's submit handler to run an async action with a spinner in
// `status` while busy, the submit button disabled, and errors shown on failure.
export function onSubmitAsync(
  form: HTMLFormElement,
  submitBtn: HTMLButtonElement,
  status: HTMLElement,
  busyText: string,
  run: () => Promise<void>
) {
  form.addEventListener('submit', async (e) => {
    e.preventDefault()
    submitBtn.setAttribute('disabled', '')
    status.innerHTML = ''
    status.append(
      h('span', { className: 'spinner-border spinner-border-sm text-warning me-2' }),
      h('span', { className: 'text-secondary small' }, [busyText])
    )
    try {
      await run()
      status.innerHTML = ''
    } catch (err) {
      status.innerHTML = ''
      status.append(h('span', { className: 'text-danger small' }, [String(err)]))
    } finally {
      submitBtn.removeAttribute('disabled')
    }
  })
}
