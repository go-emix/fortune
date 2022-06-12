<script setup>
import {ref} from 'vue'
import ax from '../pkg/axios'
import {Nerr, Nsucc} from '../pkg/notify'
import {useI18n} from 'vue-i18n'
import {langs} from '../pkg/i18n'
import {saveState} from "../pkg/session"
import {useRouter} from 'vue-router'
import menu from '../pkg/menu'

const {t, locale} = useI18n()

const router = useRouter();

let form = ref({})

function syncI18n(va) {
    saveState({i18n: va})
}

async function login() {
    if (!form.value.username) {
        Nerr(t('username') + ' ' + t('not_empty'))
        return
    }
    if (!form.value.password) {
        Nerr(t('password') + ' ' + t('not_empty'))
        return
    }
    let da = await ax({
        url: "system/login",
        method: "post",
        data: form.value
    })
    if (!da) {
        return
    }
    saveState({token: da.token})
    let msg = da.nickname
    if (locale.value === "en") {
        msg = da.username
    }
    await menu(router)
    await router.push({name: "dashboard"})
    Nsucc(t("welcome") + " " + msg)
}

async function tq() {
    let da = await ax({
        url: "system/tq",
    })
    if (!da) {
        return
    }
    Nsucc(da)
}

</script>

<template>
    <el-form :model="form" label-width="80px">
        <el-form-item :label="t('username')">
            <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item :label="t('password')">
            <el-input v-model="form.password" type="password"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="login">{{ t("login") }}</el-button>
        </el-form-item>
    </el-form>

    <el-select v-model="locale" @change="syncI18n">
        <el-option
            v-for="item in langs"
            :key="item.i18n"
            :label="item.lang"
            :value="item.i18n">
        </el-option>
    </el-select>
    <el-button @click="tq">{{ t("tq") }}</el-button>

</template>

<style scoped>

</style>
