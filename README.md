# websocket_im

The front-end code in vue:

```
<template>
    <div>
        <div>
            <input placeholder="senderId" v-model="fromId">
            <input placeholder="receiverId" v-model="toId">
            <br>
            <button @click="connect()">Connect</button>
        </div>

        <div>
            <textarea rows="10" cols="20" v-model="sendMsg"></textarea>

            <textarea rows="10" cols="20" :value="recvMsg" disabled></textarea>
            <br>
            <button @click="send()">Send</button>

        </div>
    </div>
</template>

<script>
export default {
    name: 'VueIndex',

    data() {
        return {
            ws: null,
            fromId: "",
            toId: "",
            sendMsg: "",
            recvMsg: "",
        };
    },

    methods: {
        connect() {
            this.ws = new WebSocket(`ws://localhost:9090/user/ws?fromId=${this.fromId}&toId=${this.toId}`)

            this.ws.onopen = function () {  // 连接建立时触发
                alert("ws连接成功")
                console.log("ws连接建立")
            }

            this.ws.onmessage = this.onmessage   // 收到消息时触发

            this.ws.onclose = function () {     // 连接关闭时触发
                console.log("ws连接关闭")
            }

            this.ws.onerror = function () {     // 出错时触发
                console.log('ws连接失败...')
            }
        },

        onmessage(event) {
            this.recvMsg = String(event.data)
            console.log(event.data)
        },

        send() {
            if (this.sendMsg.length > 0) {
                this.ws.send(this.sendMsg)
            }
        }

    },
};
</script>

<style scoped></style>
```
