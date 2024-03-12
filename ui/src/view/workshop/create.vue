<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router";
import { MagnifyingGlassIcon, PencilSquareIcon } from "@heroicons/vue/24/solid";

import userService from "@/service/user"
import Overlay from "@/component/Overlay.vue";
import OrderItemsModal from "@/component/workshop/OrderItemsModal.vue"

const router = useRouter()

const loading = ref(false)

const form = reactive({
    username: "",
    password: "",
    items: []
})

const itemsModal = reactive({ show: false })

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
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <label class="form-control w-full ">
                            <div class="label">
                                <span class="label-text">Name</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <label class="form-control w-full ">
                            <div class="label">
                                <span class="label-text">Address</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.password" />
                        </label>
                        <label class="form-control w-full ">
                            <div class="label">
                                <span class="label-text">DNI</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">RUC</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">Brand</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">Model</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">Color</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <div class="form-control w-full">
                            <div class="label">
                                <label class="label-text" for="i1">Plate</label>
                                <span class="label-text-alt">
                                    <button class="btn btn-sm btn-square" @click.prevent="void 0">
                                        <MagnifyingGlassIcon class="w-4 h-4" />
                                    </button>
                                </span>
                            </div>
                            <input class="input input-bordered w-full" type="text" id="i1" required
                                v-model="form.username" />
                        </div>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">Mileage</span>
                            </div>
                            <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                        </label>
                        <label class="form-control w-full">
                            <div class="label">
                                <span class="label-text">Observations</span>
                            </div>
                            <textarea class="textarea textarea-bordered w-full" v-model="form.username" />
                        </label>
                        <div class="form-control w-full">
                            <div class="label">
                                <label class="label-text" for="i2">Items</label>
                            </div>
                            <div class="join">
                                <input class="input input-bordered join-item w-full" type="number" id="i2" readonly
                                    :value="99" />
                                <button class="btn btn-square join-item" role="button"
                                    @click.prevent="itemsModal.show = !itemsModal.show">
                                    <PencilSquareIcon class="w-6 h-6" />
                                </button>
                            </div>
                        </div>
                    </div>

                    <button class="btn btn-block btn-primary mt-8" type="submit">
                        Create
                    </button>
                </form>
            </Overlay>
        </div>
        <OrderItemsModal :show="itemsModal.show" v-model="form.items" @closing="itemsModal.show = false" />
    </div>
</template>
