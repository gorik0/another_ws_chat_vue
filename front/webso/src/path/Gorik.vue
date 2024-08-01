<script setup>

import Header from '../components/Header.vue';
import Input from '../components/Input.vue';
import Output from '../components/Output.vue';
import { ref } from 'vue';


const sentSuccess = ref(0);
const wsState = ref(2);
const infoMessage = ref("");
const ws = new WebSocket("ws://localhost:9000");
ws.onopen = () => {
    infoMessage.value = ("connected")
    wsState.value = 1
    console.log("all right");
};
const onSubmitMessage = (payload) => {
    console.log("sending...");

    ws.send(payload)

};




ws.onmessage = (event) => {
    console.log(event);
    const data = JSON.parse(event.data);
    switch (data.type) {
        case "ACK":
            if (data.status === "OK") {

                sentSuccess.value = 1

                setTimeout(() => {

                    sentSuccess.value = 0
                }, 2000);

            } else {

                sentSuccess.value = 2
            }
            console.log("ACK");
            break;

        default:
            console.log("NOT ack");
            break;
    }

    infoMessage.value = (event.data)
};
ws.onclose = () => {
    infoMessage.value = ("disconnected")
    wsState.value = 3

};
ws.onerror = () => {
    infoMessage.value = ("error")
};

</script>


<template>
    <div
        class="flex flex-col items-center gap-6 py-6 [&>*]:rounded-xl  [&>*]:border-fuchsia-200 [&>*]:border-8  h-screen w-screen bg-amber-100">
        <div class="flex-1 w-1/3 bg-slate-400 ">

            <Header :wsState=wsState></Header>

        </div>
        <div class="flex-1 w-1/3 bg-pink-400">

            <Output></Output>
        </div>
        <div class="flex-1 w-1/3 bg-lime-200">
            <Input :sentSuccess=sentSuccess @onSubmitMessage="onSubmitMessage"></Input>

        </div>

    </div>


</template>