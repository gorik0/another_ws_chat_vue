<script setup>

import { ref } from 'vue';


defineProps({
    sentSuccess: Number
})
const emited = defineEmits(['onSubmitMessage'])


const messageInput = ref("")


const handleSubmitMessage = () => {

    console.log("sending...");

    messageInput.value = "(msg has been sent)"
    setTimeout(() => {
        messageInput.value = ""
    }, 1000);
};

</script>


<template>

    <div class="relative w-full h-full ">



        <div v-if="sentSuccess === 1" class="flex z-0 absolute w-full items-center justify-between h-full">
            <h1 class="mt-20 text-center text-slate-400 text-6xl w-full">
                {{ sentSuccess === 1 ? 'G O T T A' : 'E R R O R' }}</h1>
        </div>
        <form @submit.prevent="() => {
            emited('onSubmitMessage', messageInput)
            handleSubmitMessage()


        }" class=" z-10 absolute  flex flex-col w-full h-full justify-center items-center">
            <h1 class="bg-amber-300 text-xl font-bold font-mono text-center">Helo there , let make CONNECT</h1>
            <textarea v-model="messageInput"
                class=" bg-amber-200 outline-none rounded-b-xl p-5 line-clamp-4 flex-auto flex flex-col items-center justify-center">

            </textarea>
            <button type="submit"
                class=" bg-red-500 rounded-xl w-6 leading-4 p-1 m-1 text-l font-bold text-slate-800   ">M
                A
                K
                E</Button>

        </form>
    </div>


</template>