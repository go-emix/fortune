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

export function featureShow(features, rou) {
    let show = {}
    for (let i = 0; i < features.length; i++) {
        let f = features[i];
        if (f.menu_id === rou.meta.id) {
            show[f.name] = true
        }
    }
    return show
}