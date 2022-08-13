<script setup>
import {ref, watch} from "vue"
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import ax from "../pkg/axios"
import Top from "./Top.vue"
import {Nerr} from "../pkg/notify"

const list = ref([])

const {t, locale} = useI18n()

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

async function del(row) {
    await ax({
        url: "system/role?id=" + row.id,
        method: "delete"
    })
    await setList()
}

const roleForm = ref({})

const roleDialog = ref(false)

function openRoleDialog() {
    roleDialog.value = true
    roleForm.value = {}
}

async function newRole() {
    let pat = /\w+/i
    let name = roleForm.value.name;
    if (!pat.test(name)) {
        Nerr(t("must_be_alphanumeric"))
        return
    }
    await ax({
        url: "system/role",
        method: "post",
        data: {
            name: name
        }
    })
    roleDialog.value = false
    await setList()
}

const allFs = ref([])

const featureDialog = ref(false)

const feaTree = ref(null)

const roleRow = ref(null)

const featureDialogTitle = ref("role")

async function feature(row) {
    roleRow.value = row
    featureDialogTitle.value = row.name
    featureDialog.value = true
}

async function featureDialogOpen() {
    let feas = await ax({
        url: "system/featureListByRole?role=" + roleRow.value.id,
    })
    if (!feas) {
        return
    }
    let fidsVal = []
    for (let i = 0; i < feas.length; i++) {
        fidsVal.push(feas[i].id)
    }
    feaTree.value.setCheckedKeys(fidsVal)
    let as = await ax({
        url: "system/apiListByRole?role=" + roleRow.value.id,
    })
    if (!as) {
        return
    }
    apiMultiple.value.clearSelection()
    let aln = apiListRef.value.length
    for (let i = 0; i < aln; i++) {
        let ra = apiListRef.value[i]
        for (let i = 0; i < as.length; i++) {
            if (ra.id === as[i].id) {
                apiMultiple.value.toggleRowSelection(apiListRef.value[i])
                break
            }
        }
    }
}

let feaList = undefined

const apiListRef = ref([])

const apiMultiple = ref(null)

async function featureList() {
    feaList = await ax({
        url: "system/featureList"
    })
    localFeatures()
}

async function apiList() {
    let da = await ax({
        url: "system/apiList"
    })
    if (!da) {
        return
    }
    apiListRef.value = da
}

function localFeatures() {
    if (!feaList) {
        return
    }
    let mp = new Map()
    for (let i = 0; i < feaList.length; i++) {
        // 对象拷贝
        let fe = Object.assign({}, feaList[i])
        fe.name = t(fe.name)
        let fs = mp.get(fe.menu.name)
        if (fs) {
            fs.push(fe)
        } else {
            mp.set(fe.menu.name, [fe])
        }
    }
    let allfsVal = []
    for (let item of mp) {
        allfsVal.push({name: t(item[0]), children: item[1]})
    }
    allFs.value = allfsVal
}

function saveRole() {
    let keys = feaTree.value.getCheckedKeys()
    ax({
        url: "system/features",
        method: "put",
        data: {
            rid: roleRow.value.id,
            fids: keys
        }
    })
    let rows = apiMultiple.value.getSelectionRows()
    let aids = []
    for (let i = 0; i < rows.length; i++) {
        aids.push(rows[i].id)
    }
    ax({
        url: "system/apis",
        method: "put",
        data: {
            rid: roleRow.value.id,
            aids: aids
        }
    })
    featureDialog.value = false
}

watch(locale, function () {
    localFeatures()
})

setList()
featureList()
apiList()

</script>

<template>
    <Top></Top>
    <Left></Left>
    <el-button type="success" @click="openRoleDialog">{{ t("new") }}</el-button>
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
                <el-link v-if="scope.row.name!=='root'" type="primary"
                         @click="feature(scope.row)" :underline="false">
                    {{ t('feature') }}
                </el-link>
                <el-link v-if="scope.row.name!=='root'" type="danger"
                         @click="del(scope.row)" :underline="false">
                    {{ t('delete') }}
                </el-link>
                <span v-else style="color: #c8c9cb">{{ t('root_not_edit') }}</span>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog v-model="featureDialog" @open="featureDialogOpen"
               :title="t(featureDialogTitle)">
        <el-tabs active-name="feature">
            <el-tab-pane :label="t('feature')" name="feature">
                <el-tree
                    ref="feaTree"
                    :data="allFs"
                    show-checkbox
                    node-key="id"
                    :props="{label:'name'}">
                </el-tree>
            </el-tab-pane>
            <el-tab-pane :label="t('api')" name="api">
                <el-table
                    ref="apiMultiple"
                    :data="apiListRef"
                    row-key="id"
                    style="width: 100%"
                    border>
                    <el-table-column type="selection" width="55"/>
                    <el-table-column
                        prop="id"
                        label="ID">
                    </el-table-column>
                    <el-table-column
                        prop="name"
                        :label="t('name')"
                        :formatter="format">
                    </el-table-column>
                    <el-table-column
                        prop="path"
                        :label="t('path')">
                    </el-table-column>
                    <el-table-column
                        prop="method"
                        :label="t('method')">
                    </el-table-column>
                </el-table>
            </el-tab-pane>
        </el-tabs>

        <el-button type="primary" @click="saveRole">{{ t("save") }}</el-button>
    </el-dialog>

    <el-dialog v-model="roleDialog">
        <el-form :model="roleForm">
            <el-form-item :label="t('name')">
                <el-input v-model="roleForm.name" :placeholder="t('must_be_alphanumeric')"/>
            </el-form-item>
        </el-form>
        <el-button type="primary" @click="newRole">{{ t("save") }}</el-button>
    </el-dialog>

</template>

<style scoped>

</style>