<script setup>
import {useI18n} from 'vue-i18n'
import {langs} from '../pkg/i18n'
import {exit, getState, saveState} from "../pkg/session"
import {ref} from "vue";

const {t, locale} = useI18n()

function syncI18n(va) {
    saveState({i18n: va})
}

const admin = ref(getState().admin)

</script>

<template>
    <span>{{ admin.nickname ? admin.nickname : admin.username }}</span>
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