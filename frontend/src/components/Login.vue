<script setup>
import {ref} from 'vue'
import ax from '../pkg/axios'
import {Nerr, Nsucc} from '../pkg/notify'
import {useI18n} from 'vue-i18n'
import {langs} from '../pkg/i18n'
import {exit, saveState} from "../pkg/session"
import {useRouter} from 'vue-router'
import menu from '../pkg/menu'
import feature from '../pkg/feature'

const {t, locale} = useI18n()

const router = useRouter();

let form = ref({})

function syncI18n(va) {
    saveState({i18n: va})
}

const fullscreenLoading = ref(false)

async function login() {
    if (!form.value.username) {
        Nerr(t('username') + ' ' + t('not_empty'))
        return
    }
    if (!form.value.password) {
        Nerr(t('password') + ' ' + t('not_empty'))
        return
    }
    fullscreenLoading.value = true
    let da = await ax({
        url: "system/login",
        method: "post",
        data: form.value
    })
    if (!da) {
        fullscreenLoading.value = false
        return
    }
    saveState({admin: da})
    let msg = da.nickname
    if (locale.value === "en") {
        msg = da.username
    }
    await menu(router)
    await feature()
    try {
        await router.push({name: "dashboard"})
        Nsucc(t("welcome") + " " + msg)
    } catch (e) {
        exit()
    }
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

function clean() {
    form.value = {}
}


</script>

<template>
    <el-container class="container">
        <el-main id="main">
            <el-form :model="form" label-width="80px" class="form">
                <el-form-item :label="t('username')" class="item">
                    <el-input v-model="form.username" class="input"></el-input>
                </el-form-item>
                <el-form-item :label="t('password')" class="item">
                    <el-input v-model="form.password" type="password" class="input"></el-input>
                </el-form-item>
                <el-form-item class="item">
                    <el-button type="primary" @click="login" style="margin-left: 20px"
                               v-loading.fullscreen.lock="fullscreenLoading">{{ t("login") }}
                    </el-button>
                    <el-button @click="clean" style="margin-left: 30px">{{ t("clean") }}</el-button>
                </el-form-item>
            </el-form>

            <div class="div">
                <el-select v-model="locale" @change="syncI18n" class="select">
                    <el-option
                        v-for="item in langs"
                        :key="item.i18n"
                        :label="item.lang"
                        :value="item.i18n">
                    </el-option>
                </el-select>
                <el-button @click="tq">{{ t("tq") }}</el-button>
            </div>
        </el-main>
    </el-container>
</template>

<style scoped>

.container {
    width: 100%;
    background: #babaeb;
    height: 100%;
}

.input {
    width: 250px;
}

.item {
    margin-top: 20px;
}

.div {
    width: 460px;
    margin: -40px auto auto;
    padding-left: 10%;
}

.select {
    width: 80px;
    margin-right: 30px;
    margin-left: 12%;
}

.form {
    width: 460px;
    margin: 10% auto 5%;
    padding-top: 20px;
    padding-left: 50px;
    background: hsla(0, 0%, 100%, 0.2);
    padding-bottom: 20px;
    border-radius: 15px;
}

.form::before {
    content: '';
    position: relative;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    filter: blur(20px);
}

</style>
