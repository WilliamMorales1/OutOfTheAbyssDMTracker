import { api } from '../api.js'
import { h, svg } from '../dom.js'
import type { Reference } from '../types.js'

export async function refsPanel(): Promise<Node> {
  const refs = (await api.refs()) as Reference[]

  const container = h('div', {}, [])

  const select = h(
    'select',
    { className: 'form-select form-select-sm max-w-[300px] w-auto' },
    [
      h('option', { value: '' }, ['Select a reference...']),
      ...refs.map((ref) => h('option', { value: ref.id }, [ref.id])),
    ]
  ) as HTMLSelectElement

  const svgEls: SVGSVGElement[] = []
  const refDivs: HTMLElement[] = []

  function showRef(id: string) {
    refDivs.forEach((div) => {
      div.classList.toggle('hidden', div.dataset.refId !== id)
    })
    if (id) requestAnimationFrame(updateScales)
  }

  select.addEventListener('change', () => showRef(select.value))

  for (const ref of refs) {
    const card = h('div', { className: 'dc hidden' })

    const svgEl = svg('svg', {
      class: 'smap',
      viewBox: ref.vb,
      style: 'display:block;width:100%',
    }) as unknown as SVGSVGElement
    svgEl.setAttribute('xmlns', 'http://www.w3.org/2000/svg')

    const image = svg('image', { href: ref.img, width: ref.w, height: ref.h })
    svgEl.append(image)

    svgEls.push(svgEl)

    const refDiv = h('div', {
      className: 'sm relative bg-void w-full mb-4 hidden',
    }, [
      svgEl as unknown as Node,
      card,
    ]) as HTMLElement
    refDiv.dataset.refId = ref.id
    refDivs.push(refDiv)

    container.append(refDiv)
  }

  function updateScales() {
    svgEls.forEach((svgEl) => {
      const ctm = svgEl.getScreenCTM()
      if (!ctm) return
      const scale = Math.sqrt(ctm.a * ctm.a + ctm.b * ctm.b)
      if (!scale) return
      const inv = 1 / scale
      svgEl.querySelectorAll<SVGGElement>('.loc').forEach((g) => {
        const mx = g.dataset.mx
        const my = g.dataset.my
        g.setAttribute('transform', `translate(${mx},${my}) scale(${inv})`)
      })
    })
  }

  window.addEventListener('resize', updateScales)

  if (refs.length) {
    select.value = refs[0].id
    showRef(refs[0].id)
  }

  return container
}
