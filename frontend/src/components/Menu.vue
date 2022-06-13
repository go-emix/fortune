<script setup>
import {ref} from "vue"
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import ax from "../pkg/axios"
import {appendChildMenu} from "../pkg/utils"
import Top from "./Top.vue"
import {useRouter} from "vue-router";

const list = ref([])

const router=useRouter()

console.log(router.getRoutes())

async function setList() {
    let da = await ax({
        url: "system/menuList"
    })
    if (!da) {
        return
    }
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

const {t} = useI18n()

</script>

<template>
    <Top></Top>
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