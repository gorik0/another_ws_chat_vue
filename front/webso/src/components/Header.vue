<script setup>
defineProps({
    wsState: Number
})

</script>


<template>



    <div class="  flex flex-col w-full h-full justify-center items-center">
        <div class="flex-auto flex flex-col items-center justify-center">
            <p v-if="wsState === 1 || wsState === 2" class=" text-violet-950">connected</p>
            <p v-else class="animate-pulsik  text-pink-700 ">disconnected</p>
        </div>
        <h1 class="bg-amber-300 text-xl font-bold font-mono text-center">Helo there , let make CONNECT</h1>
        <button class=" bg-red-500 rounded-xl w-6 leading-4 p-1 m-1 text-l font-bold text-slate-800   ">M
            A
            K
            E</Button>

    </div>
</template>