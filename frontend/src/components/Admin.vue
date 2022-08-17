<script setup>
import {useI18n} from 'vue-i18n'
import {getState} from "../pkg/session"
import {useRoute} from "vue-router"
import {ref} from "vue"
import ax from "../pkg/axios"
import {Nerr} from "../pkg/notify"
import {featureShow} from "../pkg/utils"

const {t} = useI18n()

let rou = useRoute()

let feas = getState().features

let show = ref(featureShow(feas, rou))

let list = ref([])

const loading = ref(true)

async function setList() {
    let da = await ax({
        url: "system/adminList"
    })
    if (!da) {
        loading.value = false
        return
    }
    list.value = da
    loading.value = false
}

async function del(row) {
    await ax({
        url: "system/admin?id=" + row.id,
        method: "delete"
    })
    await setList()
}

const adminDialog = ref(false)

const adminEditDialog = ref(false)

const adminForm = ref({})

function openAdminDialog() {
    adminDialog.value = true
    adminForm.value = {}
    adminForm.value.enabled = true
}

async function edit(row) {
    adminEditDialog.value = true
    let da = await ax({
        url: "system/admin?id=" + row.id,
    })
    if (!da) {
        return
    }
    adminForm.value = da
}

async function newAdmin() {
    let value = adminForm.value;
    if (!value.username || value.username === "") {
        Nerr(t("username") + " " + t("not_empty"))
        return
    }
    let pat = /\w+/i
    let username = value.username;
    if (!pat.test(username)) {
        Nerr(t("must_be_alphanumeric"))
        return
    }
    if (!value.password || value.password === "") {
        Nerr(t("password") + " " + t("not_empty"))
        return
    }
    if (!value.rids || value.rids.length === 0) {
        Nerr(t("lower_role") + " " + t("not_empty"))
        return
    }
    await ax({
        url: "system/admin",
        method: "post",
        data: {
            username: username,
            password: value.password,
            enabled: value.enabled,
            rids: value.rids,
            nickname: value.nickname
        }
    })
    adminDialog.value = false
    await setList()
}

async function editAdmin() {
    let value = adminForm.value;
    if (!value.rids || value.rids.length === 0) {
        Nerr(t("lower_role") + " " + t("not_empty"))
        return
    }
    let pass = ""
    if (value.password && value.password !== "") {
        pass = value.password
    }
    await ax({
        url: "system/admin",
        method: "put",
        data: {
            id: value.id,
            username: value.username,
            password: pass,
            enabled: value.enabled,
            rids: value.rids,
            nickname: value.nickname
        }
    })
    adminEditDialog.value = false
    await setList()
}

const roles = ref([])

async function roleList() {
    let da = await ax({
        url: "system/roleList"
    })
    if (!da) {
        return
    }
    roles.value = da
}

let hcs = {
    background: '#8f8fbe',
    color: 'white',
}

const headCellStyle = ref(hcs)
const cellStyle = ref({background: '#eef7f5'})

setList()

roleList()


</script>

<template>
    <div>
        <el-button v-if="show.add" type="success" @click="openAdminDialog"
                   size="large" class="new">{{ t("new") }}
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
                prop="username"
                :label="t('username')"
                width="180">
            </el-table-column>
            <el-table-column
                prop="nickname"
                :label="t('nickname')"
                width="180">
            </el-table-column>
            <el-table-column
                fixed="right"
                :label="t('operate')"
                width="180"
                align="center">
                <template #default="scope">
                    <div class="operate">
                        <el-link type="primary" v-if="show.edit"
                                 @click="edit(scope.row)" :underline="false"
                                 style="margin-right: 30px">
                            {{ t('edit') }}
                        </el-link>
                        <el-link type="danger" v-if="show.delete"
                                 @click="del(scope.row)" :underline="false">
                            {{ t('delete') }}
                        </el-link>
                    </div>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog v-model="adminDialog" custom-class="adminDialog">
            <el-form :model="adminForm">
                <el-form-item :label="t('username')" label-width="70px">
                    <el-input v-model="adminForm.username" :placeholder="t('must_be_alphanumeric')"
                              class="input"/>
                </el-form-item>
                <el-form-item :label="t('password')" label-width="70px">
                    <el-input v-model="adminForm.password" type="password"
                              class="input" autocomplete="new-password"/>
                </el-form-item>
                <el-form-item :label="t('nickname')" label-width="70px">
                    <el-input v-model="adminForm.nickname" class="input"/>
                </el-form-item>
                <el-form-item :label="t('enabled')" label-width="70px">
                    <el-switch v-model="adminForm.enabled" class="input"/>
                </el-form-item>
                <el-form-item :label="t('lower_role')" label-width="70px">
                    <el-select
                        v-model="adminForm.rids"
                        multiple
                        :placeholder="t('lower_role')"
                        class="input">
                        <el-option
                            v-for="item in roles"
                            :key="item.id"
                            :label="t(item.name)"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>
            </el-form>
            <el-button type="primary" @click="newAdmin" class="save">{{ t("save") }}
            </el-button>
        </el-dialog>

        <el-dialog v-model="adminEditDialog" custom-class="adminEditDialog">
            <el-form :model="adminForm">
                <el-form-item :label="t('username')" label-width="70px">
                    <el-input v-model="adminForm.username" disabled class="input"/>
                </el-form-item>
                <el-form-item :label="t('password')" label-width="70px">
                    <el-input v-model="adminForm.password" type="password" autocomplete="new-password"
                              class="input"/>
                </el-form-item>
                <el-form-item :label="t('nickname')" label-width="70px">
                    <el-input v-model="adminForm.nickname" class="input"/>
                </el-form-item>
                <el-form-item :label="t('enabled')" label-width="70px">
                    <el-switch v-model="adminForm.enabled" class="input"/>
                </el-form-item>
                <el-form-item :label="t('lower_role')" label-width="70px">
                    <el-select
                        v-model="adminForm.rids"
                        multiple
                        :placeholder="t('lower_role')"
                        class="input">
                        <el-option
                            v-for="item in roles"
                            :key="item.id"
                            :label="t(item.name)"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>
            </el-form>
            <el-button type="primary" @click="editAdmin" class="save">{{ t("save") }}</el-button>
        </el-dialog>
    </div>
</template>

<style scoped>
.new {
    margin-left: 10%;
    margin-top: 20px;
}

.table {
    margin-top: 30px;
    margin-left: 10%;
    width: 720px;
    background: transparent;
    border-radius: 8px;
}

.operate {
    text-align: center;
}

:deep(.adminDialog) {
    background: #eef7f5;
    border-radius: 8px;
    width: 500px;
    height: 430px;
}


:deep(.adminEditDialog) {
    background: #eef7f5;
    border-radius: 8px;
    width: 500px;
    height: 430px;
}

.save {
    margin-left: 40px;
    margin-top: 20px;
}

.input {
    width: 270px;
}

</style>