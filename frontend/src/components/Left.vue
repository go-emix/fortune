<script setup>
import {ref} from 'vue'
import {useI18n} from 'vue-i18n'
import {useRoute, useRouter} from "vue-router"
import {getState} from "../pkg/session"

let rou = useRoute();
let active = ref(rou.name)

const router = useRouter();

const {t} = useI18n()

function select(k) {
    router.push({name: k})
}

let show = ref({})

function setShow() {
    let state = getState()
    let ln = state.menus.length
    for (let i = 0; i < ln; i++) {
        show.value[state.menus[i].name] = true
    }
}

setShow()

</script>

<template>
    <el-menu
        :default-active="active"
        @select="select"
        text-color="#fff"
        background-color="#494986"
        active-text-color="#4ac927">
        <el-menu-item index="dashboard" v-if="show.dashboard">
            <span slot="title">{{ t("dashboard") }}</span>
        </el-menu-item>
        <el-sub-menu index="system" v-if="show.admin||show.role">
            <template #title>
                <span>{{ t("system") }}</span>
            </template>
            <el-menu-item index="admin" v-if="show.admin">
                {{ t("admin") }}
            </el-menu-item>
            <el-menu-item index="role" v-if="show.role">
                {{ t("role") }}
            </el-menu-item>
        </el-sub-menu>
    </el-menu>

</template>

<style scoped>

</style>