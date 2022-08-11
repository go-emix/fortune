<script setup>
import {ref} from "vue"
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import ax from "../pkg/axios"
import Top from "./Top.vue"
import {ElTree} from 'element-plus'

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

const allfs = ref([])

const fids = ref([])

const featureDialog = ref(false)

const treeRef = ref(ElTree)

async function feature(row) {
    let feas = await ax({
        url: "system/featureListByRole?role=" + row.id,
    })
    if (!feas) {
        return
    }
    let fidsVal = []
    for (let i = 0; i < feas.length; i++) {
        fids.value.push(feas[i].id)
    }
    console.log(treeRef.value)
    treeRef.value.setCheckedKeys(fidsVal)
    featureDialog.value = true
}

async function featureList() {
    let feaList = await ax({
        url: "system/featureList"
    })
    if (!feaList) {
        return
    }
    let mp = new Map()
    for (let i = 0; i < feaList.length; i++) {
        let fe = feaList[i]
        let fs = mp.get(fe.menu.name)
        if (fs) {
            fs.push(fe)
        } else {
            mp.set(fe.menu.name, [fe])
        }
    }
    let allfsVal = []
    for (let item of mp) {
        allfsVal.push({name: item[0], children: item[1]})
    }
    allfs.value = allfsVal
}

setList()
featureList()

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

    <el-dialog v-model="featureDialog">
        <el-tree
            ref="treeRef"
            :data="allfs"
            show-checkbox
            node-key="id"
            :props="{label:'name'}">
        </el-tree>
    </el-dialog>

</template>

<style scoped>

</style>