<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router";

import userService from "@/service/user"
import Overlay from "@/component/Overlay.vue";

const router = useRouter()

const loading = ref(false)

const form = reactive({
    username: "",
    password: "",
})

async function create() {
    loading.value = true
    await userService.create({
        username: form.username,
        password: form.password,
    })
    loading.value = false
    router.push({ name: "user-list" })
}
</script>

<template>
    <div class="relative min-h-screen">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Create Store</h1>
            <Overlay :show="loading">
                <form @submit.prevent="create" autocomplete="off">
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Name</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                    </label>
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Password</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.password" />
                    </label>

                    <button class="btn btn-block btn-primary mt-8" type="submit">
                        Create
                    </button>
                </form>
            </Overlay>
        </div>
    </div>
</template>
