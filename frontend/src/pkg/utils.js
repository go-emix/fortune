export function appendChildMenu(m, ms) {
    let ln = ms.length
    m.children = []
    for (let i = 0; i < ln; i++) {
        let im = ms[i]
        if (im.parent === m.id) {
            m.children.push(im)
            appendChildMenu(im, ms)
        }
    }
}