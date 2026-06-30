import { api } from '../api.js'
import { h, svg } from '../dom.js'
import type { GameMap, Marker } from '../types.js'

export async function mapsPanel(): Promise<Node> {
  const maps = (await api.maps()) as GameMap[]

  let coordTracking = false

  const container = h('div', {}, [])

  const coordDisplay = h('div', {
    className: 'fixed hidden text-gold font-mono text-sm bg-surface/85 px-[7px] py-[2px] rounded pointer-events-none z-[200]',
  })

  const toggleBtn = h(
    'button',
    {
      className: 'btn btn-outline-warning',
      onclick: () => {
        coordTracking = !coordTracking
        if (!coordTracking) coordDisplay.classList.add('hidden')
      },
    },
    ['Toggle Coordinate Tooltip']
  )

  const select = h(
    'select',
    { className: 'form-select form-select-sm max-w-[300px] w-auto' },
    [
      h('option', { value: '' }, ['Select a map...']),
      ...maps.map((gm) => h('option', { value: gm.id }, [gm.id])),
    ]
  ) as HTMLSelectElement

  container.append(
    h('div', { className: 'flex gap-3 items-center mb-3' }, [select, toggleBtn]),
    coordDisplay
  )

  const svgEls: SVGSVGElement[] = []
  const mapDivs: HTMLElement[] = []

  function showMap(id: string) {
    mapDivs.forEach((div) => {
      div.classList.toggle('hidden', div.dataset.mapId !== id)
    })
    if (id) requestAnimationFrame(updateScales)
  }

  select.addEventListener('change', () => showMap(select.value))

  for (const gm of maps) {
    const card = h('div', { className: 'dc hidden' })

    const svgEl = svg('svg', {
      class: 'smap',
      viewBox: gm.vb,
      style: 'display:block;width:100%',
    }) as unknown as SVGSVGElement
    svgEl.setAttribute('xmlns', 'http://www.w3.org/2000/svg')

    const image = svg('image', { href: gm.img, width: gm.w, height: gm.h })
    svgEl.append(image)

    function showCard(m: Marker) {
      const titleEl = h('div', { className: 'dt' }, [m.title])
      const bodyEl = h('div', { className: 'db' }, [m.body])
      card.innerHTML = ''
      card.append(titleEl, bodyEl)
      card.classList.remove('hidden')
    }

    for (const m of gm.markers) {
      const circle = svg('circle', { class: 'pi', r: 12, fill: '#231d09', stroke: '#8a6428', 'stroke-width': 1.5 })
      const text = svg('text', {
        class: 'pn',
        'text-anchor': 'middle',
        'dominant-baseline': 'middle',
        'font-size': 12,
        'font-weight': 'bold',
        fill: '#c89030',
      }, [String(m.i)])

      const g = svg('g', { class: 'loc', tabindex: 0, 'data-mx': m.x, 'data-my': m.y, style: 'cursor:pointer;outline:none' }, [
        circle,
        text,
      ])
      g.addEventListener('mouseenter', () => showCard(m))
      g.addEventListener('mouseleave', () => {
        card.classList.add('hidden')
      })
      svgEl.append(g)
    }

    svgEls.push(svgEl)

    const mapDiv = h('div', {
      className: 'sm relative bg-void w-full mb-4 hidden',
    }, [
      svgEl as unknown as Node,
      card,
    ]) as HTMLElement
    mapDiv.dataset.mapId = gm.id
    mapDivs.push(mapDiv)

    mapDiv.addEventListener('mousemove', (e) => {
      const me = e as MouseEvent
      if (card.classList.contains('hidden')) return
      const pad = 16
      let x = me.clientX + pad
      let y = me.clientY + pad
      card.style.left = x + 'px'
      card.style.top = y + 'px'
      const cw = card.offsetWidth
      const ch = card.offsetHeight
      const vw = window.innerWidth
      const vh = window.innerHeight
      if (x + cw > vw) x = me.clientX - cw - pad
      if (y + ch > vh) y = me.clientY - ch - pad
      card.style.left = Math.max(0, x) + 'px'
      card.style.top = Math.max(0, y) + 'px'
    })

    container.append(mapDiv)
  }

  container.addEventListener('mousemove', (e) => {
    if (!coordTracking) {
      coordDisplay.classList.add('hidden')
      return
    }
    const me = e as MouseEvent
    const target = me.target as Element
    const svgTarget = target.closest('svg.smap') as SVGSVGElement | null
    if (!svgTarget) {
      coordDisplay.classList.add('hidden')
      return
    }
    const pt = svgTarget.createSVGPoint()
    pt.x = me.clientX
    pt.y = me.clientY
    const ctm = svgTarget.getScreenCTM()
    if (!ctm) return
    const cursor = pt.matrixTransform(ctm.inverse())
    coordDisplay.textContent = `${Math.round(cursor.x)}, ${Math.round(cursor.y)}`
    coordDisplay.classList.remove('hidden')

    const pad = 14
    let x = me.clientX + pad
    let y = me.clientY + pad
    if (x + coordDisplay.offsetWidth > window.innerWidth) x = me.clientX - coordDisplay.offsetWidth - pad
    if (y + coordDisplay.offsetHeight > window.innerHeight) y = me.clientY - coordDisplay.offsetHeight - pad
    coordDisplay.style.left = Math.max(0, x) + 'px'
    coordDisplay.style.top = Math.max(0, y) + 'px'
  })

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

  if (maps.length) {
    select.value = maps[0].id
    showMap(maps[0].id)
  }

  return container
}
