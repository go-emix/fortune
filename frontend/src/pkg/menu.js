import ax from "./axios"

let da = []

async function init() {
    da = await ax({
        url: "system/menus",
    })
}

await init()

console.log(da)

export const menu = da
