<script setup lang="ts">
import { } from "vue"

import Toast from "@/component/Toast.vue"
import Confirmation from "@/component/Confirmation.vue"
</script>

<template>
    <div>
        <router-view />
        <Toast />
        <Confirmation />
    </div>
</template>
