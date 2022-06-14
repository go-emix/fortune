<script setup>
import {useI18n} from 'vue-i18n'
import Left from './Left.vue'
import Top from "./Top.vue"
import {getState} from "../pkg/session"
import {useRoute} from "vue-router"
import {ref, shallowRef} from "vue"
import AddAdmin from './AddAdmin.vue'

const {t} = useI18n()

let com = shallowRef(AddAdmin)

let rou = useRoute()

let feas = getState().features

let show = ref({})

for (let i = 0; i < feas.length; i++) {
    let f = feas[i];
    if (f.menu === rou.meta.id) {
        show.value[f.name] = true
    }
}


function del() {
    com.value = {}
}


</script>

<template>
    <Top></Top>
    <Left></Left>
    <el-button v-if="show.add">{{ t("add") }}</el-button>
    <el-button v-if="show.delete" @click="del">{{ t("delete") }}</el-button>
    <component :is="com"></component>
</template>

<style scoped>

</style>