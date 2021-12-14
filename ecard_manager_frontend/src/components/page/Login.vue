<template>
    <div class="login-wrap">
        <div class="title">电子学生证后台管理系统</div>
        <div class="title title-english">SCHOOL BACKGROUND MANAGEMENT SYSTEM</div>
        <div class="bg">
            <div class="ms-login">
                <el-form :model="ruleForm"
                         :rules="rules"
                         ref="ruleForm"
                         label-width="70px"
                         class="demo-ruleForm ms-content"
                         @submit.native.prevent>
                    <el-form-item label="账号："
                                  prop="username">
                        <el-input v-model="ruleForm.username"
                                  autocomplete="on"></el-input>
                    </el-form-item>

                    <el-form-item label="密码："
                                  prop="password">
                        <el-input type="password"
                                  v-model="ruleForm.password"
                                  autocomplete="on"></el-input>
                    </el-form-item>
                    <div class="login-btn">
                        <el-button type="primary"
                                   native-type="submit"
                                   :disabled='disabledbtn'
                                   @click="submitForm('ruleForm',loginUrl,ruleForm)">登录</el-button>
                    </div>
                </el-form>
            </div>
        </div>
        <div class="wzg-foot">
            <!-- <div class="wzg-foot-content wzg-foot-bottom">
              <a href="https://www.attacloud.cn"
                 target="_blank"
                 rel="noopener noreferrer"><span>官网</span></a>
            </div> -->
<!--            <div class="wzg-foot-content">浙ICP备18012214号-1</div>-->
<!--            <div class="wzg-foot-content">Copyright © 2020 杭州电子科技大学</div>-->
        </div>
    </div>
</template>

<script>
    export default {
        name: 'login',
        data: function () {
            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */
                propListUrl: '',
                loginUrl: '',
                disabledbtn: false,/* 可用与不可用 */
                ruleForm: {
                    username: '',
                    password: ''
                },
                rules: {
                    username: [
                        { required: true, message: '请输入用户名', trigger: 'blur' }
                    ],
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' }
                    ]
                }
            }
        },
        created () {
            this.loginUrl = this.Request.getLoginUrl();
            this.initData();
        },
        methods: {
            /* 初始化数据 */
            initData () {
                this.ruleForm.username = localStorage.getItem('username');
            },

            /* 登录提交数据 */
            submitForm (formName, url, formData) {
                this.disabledbtn = true;
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        this.Common.commonAxios(url, 'post', formData, this.loginInfo);
                    } else {
                        this.disabledbtn = false;
                        return false;
                    }
                });
            },

            /* 登录 */
            loginInfo (data) {
                this.disabledbtn = false;
                if (data.status === 200) {
                    let tempData = data.data;
                    if (tempData.code === 200) {
                        this.$message.success("登录成功");
                        localStorage.setItem('username', this.ruleForm.username);
                        localStorage.setItem('token', tempData.data.token);
                        this.$router.push('/dashboard');
                    } else {
                        this.$message.error(tempData.message + '，请检查类型，账号，密码是否一致');
                    }
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },
        }
    }
</script>

<style scoped>
    .title {
        position: absolute;
        left: 50%;
        top: 18%;
        font-size: 35px;
        text-align: center;
        color: #fff;
        transform: translateX(-50%);
    }
    .title-english {
        top: 25%;
        color: #777d8e;
        font-size: 14px;
    }
    .login-wrap {
        position: relative;
        width: 100%;
        height: 100%;
        background: url("../../assets/img/login/dark.jpeg");
        background-size: cover;
    }
    .ms-title {
        width: 100%;
        line-height: 50px;
        text-align: center;
        font-size: 20px;
        color: #fff;
        border-bottom: 1px solid #ddd;
    }
    .bg {
        position: absolute;
        left: 50%;
        top: 31%;
        width: 360px;
        margin: 10px;
        background-color: #bfc1c2;
        border-radius: 5px;
        transform: translateX(-50%);
    }
    .ms-login {
        margin: 10px;
        background: rgba(255, 255, 255, 0.9);
    }
    .ms-content {
        padding: 30px 30px 30px 25px;
    }
    .login-btn {
        text-align: center;
    }
    .login-btn button {
        width: 30%;
        height: 36px;
        margin-bottom: 10px;
        background-color: #579292;
    }
    .login-tips {
        font-size: 12px;
        /* line-height: 30px; */
        /* color: #fff; */
    }

    .lf {
        float: left;
    }
    .box {
        min-width: 350px;
        margin-left: 50px;
        width: 30%;
    }
    .rf {
        float: right;
    }
    .clearfix:after {
        content: ".";
        display: block;
        height: 0;
        visibility: hidden;
        clear: both;
    }
    .clearfix {
        *zoom: 1;
    }
    .wzg-foot {
        position: absolute;
        left: 50%;
        bottom: 0;
        margin: 48px 0 24px 0;
        text-align: center;
        transform: translateX(-50%);
    }
    .wzg-foot-bottom {
        margin-bottom: 5px;
    }
    .wzg-foot .wzg-foot-content {
        text-align: center;
        color: rgba(255, 255, 255, 0.9);
        font-size: 14px;
    }
    .wzg-foot .wzg-foot-content span {
        color: rgba(255, 255, 255, 0.9);
    }
</style>
