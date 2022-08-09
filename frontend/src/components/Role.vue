<script setup>
import {ref} from "vue"
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import ax from "../pkg/axios"
import Top from "./Top.vue"

const list = ref([])

const {t} = useI18n()

async function setList() {
    let da = await ax({
        url: "system/roleList"
    })
    if (!da) {
        return
    }
    list.value = da
}

function format(row, col, cell) {
    return t(cell)
}

function edit(row) {
    console.log(row)
}

async function feature(row) {
    let da = await ax({
        url: "system/roleList"
    })
    if (!da) {
        return
    }
    console.log(row)
}


setList()

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
            :label="t('name')"
            width="180"
            :formatter="format">
        </el-table-column>
        <el-table-column
            fixed="right"
            :label="t('operate')"
            width="180">
            <template #default="scope">
                <el-button type="text" @click="edit(scope.row)">{{ t('edit') }}</el-button>
                <el-button type="text" @click="feature(scope.row)">{{ t('feature') }}</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<style scoped>

</style>