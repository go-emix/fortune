<script setup>
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import Top from "./Top.vue"
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

async function setList() {
    let da = await ax({
        url: "system/adminList"
    })
    if (!da) {
        return
    }
    list.value = da
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

setList()

roleList()


</script>

<template>
    <Top></Top>
    <Left></Left>
    <el-button v-if="show.add" type="success" @click="openAdminDialog">{{ t("new") }}</el-button>
    <el-table
        v-if="show.list"
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
            width="180">
            <template #default="scope">
                <el-link type="primary" v-if="show.edit"
                         @click="edit(scope.row)" :underline="false">
                    {{ t('edit') }}
                </el-link>
                <el-link type="danger" v-if="show.delete"
                         @click="del(scope.row)" :underline="false">
                    {{ t('delete') }}
                </el-link>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog v-model="adminDialog">
        <el-form :model="adminForm">
            <el-form-item :label="t('username')">
                <el-input v-model="adminForm.username" :placeholder="t('must_be_alphanumeric')"/>
            </el-form-item>
            <el-form-item :label="t('password')">
                <el-input v-model="adminForm.password" type="password" autocomplete="new-password"/>
            </el-form-item>
            <el-form-item :label="t('nickname')">
                <el-input v-model="adminForm.nickname"/>
            </el-form-item>
            <el-form-item :label="t('enabled')">
                <el-switch v-model="adminForm.enabled"/>
            </el-form-item>
            <el-form-item>
                <el-select
                    v-model="adminForm.rids"
                    multiple
                    :placeholder="t('lower_role')"
                    style="width: 240px">
                    <el-option
                        v-for="item in roles"
                        :key="item.id"
                        :label="t(item.name)"
                        :value="item.id"
                    />
                </el-select>
            </el-form-item>
        </el-form>
        <el-button type="primary" @click="newAdmin">{{ t("save") }}</el-button>
    </el-dialog>

    <el-dialog v-model="adminEditDialog">
        <el-form :model="adminForm">
            <el-form-item :label="t('username')">
                <el-input v-model="adminForm.username" disabled/>
            </el-form-item>
            <el-form-item :label="t('password')">
                <el-input v-model="adminForm.password" type="password" autocomplete="new-password"/>
            </el-form-item>
            <el-form-item :label="t('nickname')">
                <el-input v-model="adminForm.nickname"/>
            </el-form-item>
            <el-form-item :label="t('enabled')">
                <el-switch v-model="adminForm.enabled"/>
            </el-form-item>
            <el-form-item>
                <el-select
                    v-model="adminForm.rids"
                    multiple
                    :placeholder="t('lower_role')"
                    style="width: 240px">
                    <el-option
                        v-for="item in roles"
                        :key="item.id"
                        :label="t(item.name)"
                        :value="item.id"
                    />
                </el-select>
            </el-form-item>
        </el-form>
        <el-button type="primary" @click="editAdmin">{{ t("save") }}</el-button>
    </el-dialog>

</template>

<style scoped>

</style>