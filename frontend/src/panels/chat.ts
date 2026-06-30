import { api } from '../api.js'
import { h, onSubmitAsync } from '../dom.js'
import type { ChatExchange } from '../types.js'

function renderAnswer(text: string): Node {
  const wrap = h('div', {}, [])
  for (const para of text.split('\n\n')) {
    const p = h('p', {}, [])
    const lines = para.split('\n')
    lines.forEach((line, j) => {
      p.append(document.createTextNode(line))
      if (j < lines.length - 1) p.append(h('br'))
    })
    wrap.append(p)
  }
  return wrap
}

export function chatPanel(): Node {
  const history = h('div', { id: 'chat-history' }, [])
  const status = h('div', { className: 'mt-2' }, [])

  const input = h('input', {
    name: 'q',
    type: 'text',
    className: 'form-control',
    placeholder: 'Ask anything about the campaign...',
    required: true,
  }) as HTMLInputElement

  const submitBtn = h('button', { className: 'btn btn-primary', type: 'submit' }, ['Ask']) as HTMLButtonElement

  const form = h('form', { className: 'mt-3' }, [
    h('div', { className: 'input-group' }, [input, submitBtn]),
    status,
  ]) as HTMLFormElement

  onSubmitAsync(form, submitBtn, status, 'Thinking...', async () => {
    const q = input.value.trim()
    if (!q) return
    const res = (await api.chat(q)) as ChatExchange
    history.append(
      h('div', {}, [
        h('div', { className: 'chat-msg user flex justify-end' }, [
          h('div', { className: 'chat-bubble' }, [res.question]),
        ]),
        h('div', { className: 'chat-msg agent flex justify-start' }, [
          h('div', { className: 'chat-bubble' }, [renderAnswer(res.answer)]),
        ]),
      ])
    )
    input.value = ''
  })

  return h('div', {}, [history, form])
}
