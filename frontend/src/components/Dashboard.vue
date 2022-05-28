<script setup>
import {useI18n} from 'vue-i18n'
import {langs} from '../pkg/i18n'
import {clearState, saveState} from "../pkg/session"
import {useRouter} from "vue-router"
import Left from './Left.vue'

const {t, locale} = useI18n()

const router = useRouter();


function syncI18n(va) {
    saveState({i18n: va})
}

function exit() {
    clearState()
    router.push({name: "login"})
}

</script>

<template>
    <Left></Left>
    <div>{{ t("dashboard") }}</div>
    <el-select v-model="locale" @change="syncI18n">
        <el-option
            v-for="item in langs"
            :key="item.i18n"
            :label="item.lang"
            :value="item.i18n">
        </el-option>
    </el-select>
    <el-button @click="exit">{{ t("exit") }}</el-button>
</template>

<style scoped>

</style>