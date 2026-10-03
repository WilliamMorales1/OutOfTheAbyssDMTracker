import { h } from '../dom.js';
function handle(nodes) {
    let stopped = false;
    return {
        stop: () => {
            if (stopped)
                return;
            stopped = true;
            for (const node of nodes) {
                try {
                    node.stop();
                }
                catch {
                    // One-shot nodes can finish before a loop is stopped.
                }
                node.disconnect();
            }
        },
    };
}
function noiseBuffer(context, seconds) {
    const buffer = context.createBuffer(1, Math.ceil(context.sampleRate * seconds), context.sampleRate);
    const data = buffer.getChannelData(0);
    for (let i = 0; i < data.length; i++)
        data[i] = Math.random() * 2 - 1;
    return buffer;
}
function noise(context, destination, seconds, volume, filterType, frequency, loop = false) {
    const source = context.createBufferSource();
    source.buffer = noiseBuffer(context, seconds);
    source.loop = loop;
    const filter = context.createBiquadFilter();
    filter.type = filterType;
    filter.frequency.value = frequency;
    const gain = context.createGain();
    const now = context.currentTime;
    gain.gain.setValueAtTime(volume, now);
    if (!loop)
        gain.gain.exponentialRampToValueAtTime(0.001, now + seconds);
    source.connect(filter).connect(gain).connect(destination);
    source.start();
    if (!loop)
        source.stop(now + seconds);
    return handle([source]);
}
function oscillator(context, destination, startFrequency, endFrequency, seconds, volume, type, loop = false) {
    const oscillator = context.createOscillator();
    const gain = context.createGain();
    const now = context.currentTime;
    oscillator.type = type;
    oscillator.frequency.setValueAtTime(startFrequency, now);
    if (!loop)
        oscillator.frequency.linearRampToValueAtTime(endFrequency, now + seconds);
    gain.gain.setValueAtTime(volume, now);
    if (!loop)
        gain.gain.exponentialRampToValueAtTime(0.001, now + seconds);
    oscillator.connect(gain).connect(destination);
    oscillator.start();
    if (!loop)
        oscillator.stop(now + seconds);
    return handle([oscillator]);
}
function bell(context, destination) {
    const nodes = [];
    for (const [frequency, volume] of [
        [660, 0.35],
        [990, 0.2],
        [1320, 0.1],
    ]) {
        const oscillator = context.createOscillator();
        const gain = context.createGain();
        const now = context.currentTime;
        oscillator.type = 'sine';
        oscillator.frequency.value = frequency;
        gain.gain.setValueAtTime(volume, now);
        gain.gain.exponentialRampToValueAtTime(0.001, now + 1.8);
        oscillator.connect(gain).connect(destination);
        oscillator.start();
        oscillator.stop(now + 1.8);
        nodes.push(oscillator);
    }
    return handle(nodes);
}
function sword(context, destination) {
    const hit = noise(context, destination, 0.35, 0.45, 'bandpass', 1800);
    const ring = oscillator(context, destination, 650, 120, 0.45, 0.16, 'triangle');
    return {
        stop: () => {
            hit.stop();
            ring.stop();
        },
    };
}
function door(context, destination) {
    const creak = oscillator(context, destination, 150, 48, 0.8, 0.28, 'sawtooth');
    const thud = noise(context, destination, 0.18, 0.35, 'lowpass', 220);
    return {
        stop: () => {
            creak.stop();
            thud.stop();
        },
    };
}
function thunder(context, destination) {
    const rumble = noise(context, destination, 1.7, 0.55, 'lowpass', 320);
    const boom = oscillator(context, destination, 75, 28, 1.5, 0.35, 'sine');
    return {
        stop: () => {
            rumble.stop();
            boom.stop();
        },
    };
}
function dice(context, destination) {
    const nodes = [];
    const now = context.currentTime;
    for (let i = 0; i < 8; i++) {
        const source = context.createBufferSource();
        source.buffer = noiseBuffer(context, 0.035);
        const gain = context.createGain();
        const start = now + i * 0.07;
        gain.gain.setValueAtTime(0.28, start);
        gain.gain.exponentialRampToValueAtTime(0.001, start + 0.035);
        source.connect(gain).connect(destination);
        source.start(start);
        source.stop(start + 0.035);
        nodes.push(source);
    }
    return handle(nodes);
}
function goblin(context, destination) {
    const first = oscillator(context, destination, 280, 520, 0.16, 0.2, 'sawtooth');
    const second = oscillator(context, destination, 520, 190, 0.3, 0.2, 'sawtooth');
    return {
        stop: () => {
            first.stop();
            second.stop();
        },
    };
}
function horror(context, destination) {
    const drone = oscillator(context, destination, 90, 32, 1.4, 0.28, 'triangle');
    const scrape = noise(context, destination, 1.1, 0.2, 'bandpass', 700);
    return {
        stop: () => {
            drone.stop();
            scrape.stop();
        },
    };
}
function fire(context, destination) {
    return noise(context, destination, 0.9, 0.35, 'lowpass', 1100);
}
function rain(context, destination) {
    return noise(context, destination, 2, 0.14, 'highpass', 900, true);
}
function campfire(context, destination) {
    const crackle = noise(context, destination, 2, 0.16, 'lowpass', 1400, true);
    const hum = oscillator(context, destination, 72, 72, 2, 0.05, 'sine', true);
    return {
        stop: () => {
            crackle.stop();
            hum.stop();
        },
    };
}
function underdark(context, destination) {
    const low = oscillator(context, destination, 48, 52, 2, 0.13, 'sine', true);
    const high = oscillator(context, destination, 96, 92, 2, 0.04, 'triangle', true);
    return {
        stop: () => {
            low.stop();
            high.stop();
        },
    };
}
const sounds = [
    { key: '1', name: 'Sword Clash', icon: '⚔', description: 'Steel rings in the dark', color: '#dca83d', duration: 0.5, play: sword },
    { key: '2', name: 'Door Creak', icon: '🚪', description: 'Something opened', color: '#a8753b', duration: 0.85, play: door },
    { key: '3', name: 'Thunder', icon: '⚡', description: 'The deep rumbles', color: '#818cf8', duration: 1.7, play: thunder },
    { key: '4', name: 'Dice Roll', icon: '🎲', description: 'Let fate decide', color: '#f0d080', duration: 0.65, play: dice },
    { key: '5', name: 'Goblin Shriek', icon: '👹', description: 'A nasty little surprise', color: '#a3c94a', duration: 0.35, play: goblin },
    { key: '6', name: 'Bell Toll', icon: '🔔', description: 'A warning from below', color: '#f5c76a', duration: 1.85, play: bell },
    { key: '7', name: 'Fire Burst', icon: '🔥', description: 'Flames catch fast', color: '#ef7544', duration: 0.9, play: fire },
    { key: '8', name: 'Dread', icon: '☠', description: 'Something watches', color: '#c084fc', duration: 1.45, play: horror },
    { key: '9', name: 'Rain', icon: '☔', description: 'Looping ambience', color: '#60a5fa', loop: true, duration: 2, play: rain },
    { key: '0', name: 'Campfire', icon: '♨', description: 'Looping ambience', color: '#fb923c', loop: true, duration: 2, play: campfire },
    { key: '-', name: 'Underdark Hum', icon: '◉', description: 'Looping ambience', color: '#2dd4bf', loop: true, duration: 2, play: underdark },
];
class SoundEngine {
    context = null;
    master = null;
    active = new Map();
    getOutput() {
        if (this.context && this.master)
            return { context: this.context, master: this.master };
        const audioWindow = window;
        const AudioContextConstructor = window.AudioContext ?? audioWindow.webkitAudioContext;
        if (!AudioContextConstructor)
            throw new Error('Web Audio is not supported by this browser');
        this.context = new AudioContextConstructor();
        this.master = this.context.createGain();
        this.master.gain.value = 0.65;
        this.master.connect(this.context.destination);
        return { context: this.context, master: this.master };
    }
    setVolume(value) {
        const output = this.getOutput();
        output.master.gain.value = value;
    }
    toggle(sound) {
        const existing = this.active.get(sound.name);
        if (existing) {
            existing.stop();
            this.active.delete(sound.name);
            return false;
        }
        const output = this.getOutput();
        void output.context.resume();
        const newHandle = sound.play(output.context, output.master);
        if (sound.loop)
            this.active.set(sound.name, newHandle);
        return true;
    }
    stopAll() {
        for (const active of this.active.values())
            active.stop();
        this.active.clear();
    }
}
export async function soundboardPanel() {
    const engine = new SoundEngine();
    const active = new Set();
    const root = h('div', {}, []);
    const status = h('span', { className: 'text-gray-500 text-xs' }, ['Ready']);
    const volume = h('input', {
        type: 'range',
        min: '0',
        max: '1',
        step: '0.05',
        value: '0.65',
        className: 'sound-volume w-[120px] accent-yellow-400',
        'aria-label': 'Master volume',
        oninput: (event) => engine.setVolume(Number(event.target.value)),
    });
    function updatePad(pad, sound) {
        const isActive = active.has(sound.name);
        pad.classList.toggle('sound-pad-active', isActive);
        pad.setAttribute('aria-pressed', String(isActive));
        const state = pad.querySelector('.sound-pad-state');
        if (state)
            state.textContent = isActive ? 'Playing' : sound.loop ? 'Loop' : 'Ready';
    }
    const pads = sounds.map((sound) => {
        const pad = h('button', {
            type: 'button',
            className: 'sound-pad group',
            'aria-label': `${sound.name}: ${sound.description}`,
            'aria-pressed': 'false',
            style: { borderColor: sound.color },
            onclick: () => {
                try {
                    const isPlaying = engine.toggle(sound);
                    if (isPlaying)
                        active.add(sound.name);
                    else
                        active.delete(sound.name);
                    updatePad(pad, sound);
                    status.textContent = isPlaying ? `${sound.name} ${sound.loop ? 'looping' : 'played'}` : `${sound.name} stopped`;
                    if (isPlaying && !sound.loop) {
                        window.setTimeout(() => {
                            active.delete(sound.name);
                            updatePad(pad, sound);
                        }, sound.duration * 1000);
                    }
                }
                catch (error) {
                    status.textContent = String(error);
                }
            },
        }, [
            h('span', { className: 'sound-pad-key' }, [sound.key]),
            h('span', { className: 'sound-pad-icon', style: { color: sound.color } }, [sound.icon]),
            h('strong', { className: 'sound-pad-name' }, [sound.name]),
            h('span', { className: 'sound-pad-description' }, [sound.description]),
            h('span', { className: 'sound-pad-state' }, [sound.loop ? 'Loop' : 'Ready']),
        ]);
        return pad;
    });
    const stopButton = h('button', {
        type: 'button',
        className: 'btn btn-outline-danger',
        onclick: () => {
            engine.stopAll();
            active.clear();
            sounds.forEach((sound, index) => updatePad(pads[index], sound));
            status.textContent = 'All loops stopped';
        },
    }, ['Stop All']);
    root.append(h('div', { className: 'flex flex-wrap items-center justify-between gap-3 mb-4' }, [
        h('div', {}, [
            h('h2', { className: 'text-2xl font-bold text-yellow-400 mb-1' }, ['Soundboard']),
            h('p', { className: 'text-gray-400 text-sm' }, ['Quick sounds and ambience for the table.']),
        ]),
        h('div', { className: 'flex items-center gap-3' }, [
            h('label', { className: 'flex items-center gap-2 text-gray-400 text-sm' }, ['Volume', volume]),
            stopButton,
        ]),
    ]), h('div', { className: 'soundboard-grid' }, pads), h('div', { className: 'flex flex-wrap justify-between gap-2 mt-4 items-center' }, [
        h('span', { className: 'text-gray-500 text-xs' }, ['Keyboard: 1–9, 0, -, or click a pad. Loop pads toggle on/off.']),
        status,
    ]));
    window.addEventListener('keydown', (event) => {
        if (!root.isConnected)
            return;
        if (event.target instanceof HTMLInputElement || event.target instanceof HTMLTextAreaElement)
            return;
        const index = sounds.findIndex((sound) => sound.key === event.key);
        if (index < 0)
            return;
        event.preventDefault();
        pads[index].click();
    });
    return root;
}
