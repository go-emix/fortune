<script setup>
import {ref, watch} from "vue"
import {useI18n} from 'vue-i18n'
import ax from "../pkg/axios"
import {Nerr, Nwarn} from "../pkg/notify"
import {featureShow} from "../pkg/utils";
import {useRoute} from "vue-router";
import {getState} from "../pkg/session";

const list = ref([])

const {t, locale} = useI18n()

let rou = useRoute()

let feas = getState().features

let show = ref(featureShow(feas, rou))

const loading = ref(true)

async function setList() {
    let da = await ax({
        url: "system/roleList"
    })
    if (!da) {
        loading.value = false
        return
    }
    list.value = da
    loading.value = false
}

function format(row, col, cell) {
    return t(cell)
}

async function del(row) {
    let msg = t('delete') + " " + t('lower_role') + " ?"
    if (!await Nwarn(msg)) {
        return
    }
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
    let name = roleForm.value.name;
    if (!name) {
        Nerr(t("name") + " " + t("not_empty"))
        return
    }
    let pat = /\w+/i
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
        for (let j = 0; j < as.length; j++) {
            if (ra.id === as[j].id) {
                apiMultiple.value.toggleRowSelection(ra)
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

let hcs = {
    background: '#8f8fbe',
    color: 'white',
}

const headCellStyle = ref(hcs)
const cellStyle = ref({background: '#eef7f5'})


setList()
featureList()
apiList()

</script>

<template>
    <div>
        <el-button type="success" v-if="show.add" @click="openRoleDialog"
                   class="new" size="large">{{ t("new") }}
        </el-button>

        <el-table
            v-if="show.list"
            :data="list"
            row-key="id"
            :header-cell-style="headCellStyle"
            :cell-style="cellStyle"
            class="table"
            stripe
            max-height="350"
            v-loading="loading">
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
                width="250"
                align="center">
                <template #default="scope">
                    <div class="operate">
                        <el-link v-if="show.feature&&scope.row.name!=='root'" type="primary"
                                 @click="feature(scope.row)" :underline="false"
                                 style="margin-right: 30px">
                            {{ t('feature') }}
                        </el-link>
                        <el-link v-if="show.feature&&scope.row.name!=='root'" type="danger"
                                 @click="del(scope.row)" :underline="false">
                            {{ t('delete') }}
                        </el-link>
                        <span v-if="scope.row.name==='root'" style="color: #c8c9cb">
                            {{ t('root_not_edit') }}</span>
                    </div>
                </template>
            </el-table-column>
        </el-table>


        <el-dialog v-model="featureDialog" @open="featureDialogOpen"
                   :title="t(featureDialogTitle)"
                   custom-class="featureDialog">
            <el-tabs active-name="feature" class="feature-tab">
                <el-tab-pane :label="t('feature')" name="feature"
                             class="el-tab">
                    <el-scrollbar max-height="220">
                        <el-tree
                            ref="feaTree"
                            :data="allFs"
                            show-checkbox
                            node-key="id"
                            :props="{label:'name'}"
                            class="feaTree">
                        </el-tree>
                    </el-scrollbar>
                </el-tab-pane>
                <el-tab-pane :label="t('api')" name="api"
                             class="el-tab">
                    <el-table
                        ref="apiMultiple"
                        :data="apiListRef"
                        row-key="id"
                        style="width: 100%"
                        border
                        max-height="270"
                        :cell-style="cellStyle"
                        :header-cell-style="{background: '#b4becb',color:'white'}">
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

        <el-dialog v-model="roleDialog" custom-class="roleDialog">
            <el-form :model="roleForm">
                <el-form-item :label="t('name')" label-width="70px">
                    <el-input v-model="roleForm.name" class="input"
                              :placeholder="t('must_be_alphanumeric')"/>
                </el-form-item>
            </el-form>
            <el-button type="primary" @click="newRole"
                       style="margin-left: 40px;margin-top: 20px">{{ t("save") }}
            </el-button>
        </el-dialog>
    </div>
</template>

<style scoped>

.table {
    margin-top: 30px;
    margin-left: 10%;
    width: 610px;
    background: transparent;
    border-radius: 8px;
}

.new {
    margin-left: 10%;
    margin-top: 20px;
}

.operate {
    text-align: center;
}

:deep(.roleDialog) {
    background: #eef7f5;
    border-radius: 8px;
    width: 400px;
    height: 200px;
}

.input {
    width: 230px;
}

:deep(.featureDialog) {
    background: #eef7f5;
    border-radius: 8px;
    width: 500px;
    height: 480px;
}

.feaTree {
    background: #eef7f5;
    height: 220px;
}

.el-tab {
    height: 300px;
}

.feature-tab {
    margin-top: -20px;
}

</style>