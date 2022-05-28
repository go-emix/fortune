<script setup>
import {ref} from 'vue'
import {useI18n} from 'vue-i18n'
import {useRoute, useRouter} from "vue-router"

let rou = useRoute();
let active = ref(rou.name)

const router = useRouter();

const {t} = useI18n()

function select(k) {
    router.push({name: k})
}

</script>

<template>
    <el-menu
        :default-active="active"
        @select="select">
        <el-menu-item index="dashboard">
            <span slot="title">{{ t("dashboard") }}</span>
        </el-menu-item>
        <el-sub-menu index="">
            <el-menu-item index="admin">
                {{ t("admin") }}
            </el-menu-item>
        </el-sub-menu>
    </el-menu>

</template>

<style scoped>

</style>