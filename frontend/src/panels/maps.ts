import { api } from '../api.js'
import { h, svg } from '../dom.js'
import type { GameMap, Marker } from '../types.js'

export async function mapsPanel(): Promise<Node> {
  const maps = (await api.maps()) as GameMap[]

  let coordTracking = false

  const container = h('div', {}, [])

  const coordDisplay = h('div', {
    style:
      'position:fixed;display:none;color:#f0d080;font-family:monospace;font-size:0.85rem;background:rgba(13,14,22,0.85);padding:2px 7px;border-radius:4px;pointer-events:none;z-index:200;',
  })

  const toggleBtn = h(
    'button',
    {
      className: 'btn btn-sm btn-outline-warning',
      onclick: () => {
        coordTracking = !coordTracking
        if (!coordTracking) coordDisplay.style.display = 'none'
      },
    },
    ['Toggle Coordinate Tooltip']
  )

  container.append(h('div', { style: 'margin-bottom:10px;' }, [toggleBtn]), coordDisplay)

  const svgEls: SVGSVGElement[] = []

  for (const gm of maps) {
    const card = h('div', {
      className: 'dc',
      style:
        'position:fixed;display:none;background:#0d0e16;border:1px solid #352c10;padding:12px;width:max-content;max-width:min(320px,85vw);box-sizing:border-box;white-space:normal;word-wrap:break-word;pointer-events:none;z-index:100;border-radius:8px;box-shadow:0 4px 15px rgba(0,0,0,0.5);',
    })

    const svgEl = svg('svg', {
      class: 'smap',
      viewBox: gm.vb,
      style: 'display:block;width:100%',
    }) as unknown as SVGSVGElement
    svgEl.setAttribute('xmlns', 'http://www.w3.org/2000/svg')

    const image = svg('image', { href: gm.img, width: gm.w, height: gm.h })
    svgEl.append(image)

    function showCard(m: Marker) {
      const titleEl = h('div', { style: "font-family:'Trajan Pro', serif;font-size:1.25rem;color:#f0d080;margin-bottom:10px;" }, [m.title])
      const bodyEl = h('div', { style: 'line-height:1.5;color:#d6d6d6;' }, [m.body])
      card.innerHTML = ''
      card.append(titleEl, bodyEl)
      card.style.display = 'block'
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

      const g = svg('g', { class: 'loc', tabindex: 0, 'data-mx': m.x, 'data-my': m.y, style: 'cursor:pointer;outline:none;' }, [
        circle,
        text,
      ])
      g.addEventListener('mouseenter', () => showCard(m))
      g.addEventListener('mouseleave', () => {
        card.style.display = 'none'
      })
      svgEl.append(g)
    }

    svgEls.push(svgEl)

    const mapDiv = h('div', { className: 'sm', style: 'position:relative;background:#080910;width:100%;margin-bottom:16px;' }, [
      svgEl as unknown as Node,
      card,
    ])

    mapDiv.addEventListener('mousemove', (e) => {
      const me = e as MouseEvent
      if (card.style.display !== 'block') return
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
      coordDisplay.style.display = 'none'
      return
    }
    const me = e as MouseEvent
    const target = me.target as Element
    const svgTarget = target.closest('svg.smap') as SVGSVGElement | null
    if (!svgTarget) {
      coordDisplay.style.display = 'none'
      return
    }
    const pt = svgTarget.createSVGPoint()
    pt.x = me.clientX
    pt.y = me.clientY
    const ctm = svgTarget.getScreenCTM()
    if (!ctm) return
    const cursor = pt.matrixTransform(ctm.inverse())
    coordDisplay.textContent = `${Math.round(cursor.x)}, ${Math.round(cursor.y)}`
    coordDisplay.style.display = 'block'

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
  requestAnimationFrame(updateScales)

  return container
}
