<template>
    <div class="wrapper">
        <v-head></v-head>
        <v-sidebar></v-sidebar>
        <div class="content-box"
             :class="{'content-collapse':collapse}">
            <div class="content"
                 id="target">
                <transition name="move"
                            mode="out-in">
                    <router-view></router-view>
                </transition>
            </div>
        </div>
    </div>
</template>

<script>
    import vHead from './Header.vue';
    import vSidebar from './Sidebar.vue';
    import bus from './bus';
    export default {
        data () {
            return {
                Request: this.$api.api.prototype /* 用于获取请求地址 */,
                Common: this.$common.common.prototype /* 用于获取公共方法 */,

                tagsList: [],
                collapse: false,

                alarmMsg: '',
            }
        },
        components: {
            vHead, vSidebar
        },
        created () {
            bus.$on('collapse', msg => {
                this.collapse = msg;
            })
        },

        mounted() {
            let token = "token";
            let event = "alarm";
            let ws = new WebSocket(this.Request.getWsAddr());

            ws.onopen = function () {
                ws.send("{\"token\": \"" + token + "\", \"event\": \"" + event + "\"}")
            };
            ws.onmessage = (e) => {
                this.$notify({
                    title: '告警',
                    message: e.data,
                    type: 'warning'
                });
                // this.getAlarmsCountRequest();
            };
        },

        methods: {
        }
    }
</script>
<style scoped>
    .content {
        overflow-y: auto;
    }
</style>