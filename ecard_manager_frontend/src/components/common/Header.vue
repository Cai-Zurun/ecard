<template>
    <div class="header">
        <!-- 折叠按钮 -->
        <div class="collapse-btn" @click="collapseChage">
            <i :class="menu"></i>
        </div>

        <div class="logo">电子学生证后台管理系统</div>
        <div class="header-right">
            <div class="header-user-con">
                <!-- 用户头像 -->
                <div class="user-avator">
<!--                    <router-link v-if="tagsList.length > 0" :to="tagsList[0].path">-->
                    <img src="../../assets/img/head.png" width="1.5" height="1.5">
<!--                    </router-link>-->
                </div>

                <!-- 用户名下拉菜单 -->
                <el-dropdown class="user-name" trigger="click" @command="handleCommand">
                    <span class="el-dropdown-link">
                      {{ username }} <i class="el-icon-caret-bottom"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item divided command="loginout">
                            退出登录
                        </el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </div>
        </div>
    </div>
</template>
<script>
    import bus from "../common/bus";
    export default {
        data() {
            return {
                Request: this.$api.api.prototype /* 用于获取请求地址 */,
                Common: this.$common.common.prototype /* 用于获取公共方法 */,
                imageUrl: "",
                collapse: true,
                menu: "el-icon-s-fold",
                fullscreen: false,
                message: 2,
                userImg: "../../assets/logo.png",
                tagsList: [],
                username: "总管理员",
                loginVisible: false /* 提示需要登录的对话框 */,

                limitShow: localStorage.getItem("userType")
            };
        },
        created() {
            /* 兄弟组件之间的传值 */
            bus.$on("userImg", arg => {
                this.userImg = arg; // 接收
            });
        },
        computed: {},
        methods: {
            // 用户名下拉菜单选择事件
            handleCommand(command) {
                if (command == "loginout") {
                    localStorage.removeItem("token");
                    this.$router.push("/login");
                }
            },

            // 侧边栏折叠
            collapseChage() {
                this.collapse = !this.collapse;
                if (this.collapse) {
                    this.menu = "el-icon-s-unfold";
                } else {
                    this.menu = "el-icon-s-fold";
                }
                bus.$emit("collapse", this.collapse);
            },
        },
        mounted() {
            if (document.body.clientWidth < 1500) {
                this.collapseChage();
            }
        }
    };
</script>
<style scoped>
    .header {
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 70px;
        /* background-color: #1e3c6e; */
        font-size: 22px;
        color: #fff;
    }

    .collapse-btn {
        float: left;
        padding: 0 21px;
        cursor: pointer;
        line-height: 70px;
    }
    .header .logo {
        float: left;
        /* width: 250px; */
        line-height: 70px;
    }
    .header-right {
        float: right;
        padding-right: 50px;
    }
    .header-user-con {
        display: flex;
        height: 70px;
        align-items: center;
    }
    .aky {
        float: left;
        height: 100%;
        margin-right: 21px;
        padding: 10px;
        line-height: 70px;
        box-sizing: border-box;
    }
    .aky-img {
        height: 47px;
    }
    /* .btn-help {
      margin-right: 5px;
      font-size: 24px;
       cursor: pointer;
    } */
    .btn-fullscreen {
        transform: rotate(45deg);
        margin: 0 5px;
        font-size: 24px;
    }
    .btn-bell,
    .btn-fullscreen {
        position: relative;
        width: 30px;
        height: 30px;
        text-align: center;
        border-radius: 15px;
        cursor: pointer;
    }
    .btn-bell-badge {
        position: absolute;
        right: 0;
        top: -2px;
        width: 8px;
        height: 8px;
        border-radius: 4px;
        background: #f56c6c;
        color: #fff;
    }
    .btn-bell .el-icon-bell {
        color: #fff;
    }
    .user-name {
        margin-left: 10px;
    }
    .user-avator {
        margin-left: 20px;
    }
    .user-avator img {
        display: block;
        width: 40px;
        height: 40px;
        border-radius: 50%;
        object-fit: cover;
    }
    .el-dropdown-link {
        color: #fff;
        cursor: pointer;
    }
    .el-dropdown-menu__item {
        text-align: center;
    }
    .el-dropdown-menu__item a {
        color: #606266;
    }

    .to-console,
    .btn-help {
        margin-right: 15px;
        font-size: 22px;
        cursor: pointer;
    }
    .to-console:hover,
    .btn-help:hover {
        color: #606266;
        transition: 0.3s;
    }
</style>
