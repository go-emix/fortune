<script setup>
import {ref} from "vue"
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import ax from "../pkg/axios"

const list = ref([])

async function setList() {
    let da = await ax({
        url: "system/menuList"
    })
    let ln = da.length
    let ms = []
    for (let i = 0; i < ln; i++) {
        let v = da[i]
        if (v.parent === 0) {
            ms.push(v)
            appendChildMenu(v, da)
        }
    }
    list.value = ms
}

setList()

function appendChildMenu(m, ms) {
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

const {t} = useI18n()

</script>

<template>
    <Left></Left>
    <el-table
        :data="list"
        row-key="id"
        style="width: 100%"
        border>
        <el-table-column
            prop="id"
            label="ID"
            width="180">
        </el-table-column>
        <el-table-column
            prop="name"
            label="name"
            width="180">
        </el-table-column>
        <el-table-column
            prop="parent"
            label="parent"
            width="180">
        </el-table-column>
        <el-table-column
            prop="path"
            label="path"
            width="180">
        </el-table-column>
        <el-table-column
            prop="component"
            label="component"
            width="180">
        </el-table-column>
    </el-table>
</template>

<style scoped>

</style>