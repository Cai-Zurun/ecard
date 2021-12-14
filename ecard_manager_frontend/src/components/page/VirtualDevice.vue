<template>
    <div class="table">
        <div class="container" >
            <el-row :gutter="20" class="mgb20" style="margin-top: 10px">
                <el-col :span="8">
                    <el-button type="primary"
                               icon="el-icon-plus"
                               @click="addVirtualDevice">
                        添加虚拟设备
                    </el-button>
                </el-col>
            </el-row>
            <el-divider></el-divider>
            <el-row :gutter="30" class="mgb20">
                <el-col :span="8" v-for="(virtualDevice, index) in virtualDevices"
                        :key="index">
                    <el-card shadow="always"
                             :body-style="{padding: '0'}">
                        <div class="grid-content grid-con-2">
                            <i class="el-icon-mobile-phone grid-con-icon"></i>
                            <div class="grid-cont-right">
                                <div style="margin-bottom: 5px">imei: {{virtualDevice.imei}}</div>
                                <div style="margin-bottom: 15px">学生姓名: {{virtualDevice.student_name}}</div>
                                <el-button type="danger"
                                           icon="el-icon-bell"
                                           @click="alarmRequest(virtualDevice.imei)">
                                    告警
                                </el-button>
                            </div>
                        </div>
                    </el-card>
                </el-col>
            </el-row>
        </div>

        <!-- 重新登录提示框 -->
        <el-dialog title="提示"
                   :visible.sync="loginVisible"
                   width="300px"
                   :close-on-click-modal='false'
                   center>
            <div class="del-dialog-cnt">登录时间过长，请重新登录</div>
            <span slot="footer"
                  class="dialog-footer">
                <el-button @click="loginVisible = false">取 消</el-button>
                <el-button type="primary"
                           @click="Common.signOut('/login')">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 新增编辑查看弹出框 -->
        <el-dialog :title="dialogTitle"
                   :visible.sync="editVisible"
                   width="40%"
                   :close-on-click-modal="false"
                   @closed="closedFormDialog">
            <el-form ref="imeiForm"
                     :model="studentForm"
                     :rules="rules"
                     v-loading="loadingDialog"
                     label-width="120px">
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="学号："
                                      prop="studentNo">
                            <el-input v-model="studentForm.studentNo"
                                      placeholder="请输入学号"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="姓名："
                                      prop="name">
                            <el-input v-model="studentForm.name"
                                      placeholder="请输入姓名"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="年龄："
                                      prop="age">
                            <el-input v-model="studentForm.age"
                                      placeholder="请输入年龄"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-form-item label="性别："
                                  prop="sex">
                        <el-radio-group v-model="studentForm.sex">
                            <el-radio :label="1">男</el-radio>
                            <el-radio :label="0">女</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="手机："
                                      prop="phone">
                            <el-input v-model="studentForm.phone"
                                      placeholder="请输入手机"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="学校："
                                      prop="schoolId">
                            <el-select v-model="studentForm.schoolId"
                                       placeholder="请选择学校"
                                       class="">
                                <el-option v-for="item in selectSchool"
                                           :label="item.SchoolName"
                                           :value="item.ID"
                                           :key="item.ID"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="年级："
                                      prop="grade">
                            <el-input v-model="studentForm.grade"
                                      placeholder="请输入年级"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="班级："
                                      prop="class">
                            <el-input v-model="studentForm.class"
                                      placeholder="请输入班级"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="IMEI号："
                                      prop="imei">
                            <el-input v-model="studentForm.imei"
                                      placeholder="请输入IMEI号"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <span slot="footer"
                  class="dialog-footer">
        <el-button @click="cancelSubmit('imeiForm')">取 消</el-button>

        <el-button type="primary"
                   :disabled='disabledbtn'
                   v-if="!isDisabledbtn"
                   @click="submitForm('studentForm','post', Request.getVirtualDeviceUrl(), studentForm)">确 定</el-button>
      </span>
        </el-dialog>

        <el-dialog :title="告警"
                   :visible.sync="alarmVisible"
                   width="40%"
                   :close-on-click-modal="false"
                   @closed="closedFormDialog">
            <el-form ref="alarmForm"
                     :model="alarmForm"
                     :rules="rules"
                     v-loading="loadingDialog"
                     label-width="120px">
                <el-row :gutter="20">
                    <el-col :span="20">
                        <el-form-item label="告警类型："
                                      prop="alarmType">
                            <el-select v-model="alarmForm.alarmType"
                                       placeholder="请选择告警类型"
                                       class="">
                                <el-option v-for="item in selectAlarmType"
                                           :label="item.alarmContent"
                                           :value="item.id"
                                           :key="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <span slot="footer"
                  class="dialog-footer">
        <el-button @click="cancelSubmit('imeiForm')">取 消</el-button>

        <el-button type="primary"
                   :disabled='disabledbtn'
                   v-if="!isDisabledbtn"
                   @click="submitForm('alarmForm','post', Request.getAlarmUrl(), alarmForm)">确 定</el-button>
      </span>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "VirtualDevice",
        data () {
            return {
                Request: this.$api.api.prototype,/* 用于获取请求地址 */
                Common: this.$common.common.prototype,/* 用于获取公共方法 */

                loginVisible: false,
                virtualDevices:[],
                selectSchool: [],

                dialogTitle: '',
                editVisible: false,
                alarmVisible: false,

                alarmType: '',

                studentForm: {
                    studentNo: '',
                    name: '',
                    sex: '',
                    age: '',
                    phone: '',
                    class: '',
                    grade: '',
                    schoolId: '',
                    imei: '',
                },

                selectAlarmType: [
                    {id: 1, alarmContent: "SOS告警"},
                    {id: 2, alarmContent: "低电量告警"},
                    {id: 3, alarmContent: "出围栏告警"},
                    {id: 4, alarmContent: "进围栏告警"},
                ],
                alarmForm: {
                    alarmType: '',
                    imei: '',
                }
            }
        },

        created() {
            this.getVirtualDevicesRequest();
            this.getSchoolInfoDataRequest();
        },

        methods: {
            // 获取虚拟设备
            getVirtualDevicesRequest () {
                this.Common.getAxios(this.Request.getVirtualDeviceUrl(), '', this.getVirtualDevices);
            },
            getVirtualDevices (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.virtualDevices = tempData.data;
                    console.log("virtualDevices:",this.virtualDevices)
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            /* 请求：获取学校信息 */
            getSchoolInfoDataRequest () {
                this.Common.getAxios(this.Request.getSchoolInfoUrl(), '', this.getSchoolData);
            },
            /* 获取学校数据  */
            getSchoolData (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.loading = false;
                    this.selectSchool = tempData.data;
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.$message.error('服务器繁忙，请稍后再试');
                    this.loading = false;
                }
            },

            addVirtualDevice() {
                this.dialogTitle = '新增';
                this.editVisible = true;
            },

            /* 保存或编辑 */
            submitForm (formName, type, url, formData) {
                /* 防止多次提交 */
                this.disabledbtn = true;
                this.loadingDialog = true;
                this.Common.commonAxios(url, type, formData, this.submitFormInfo);
                this.disabledbtn = false;
                this.loadingDialog = false;
            },

            /* 保存或编辑信息 */
            submitFormInfo (data) {
                if (data.status === 200) {
                    let tempData = data.data;
                    this.disabledbtn = false;
                    this.loadingDialog = false;
                    if (tempData.code === 200) {
                        this.editVisible = false;
                        this.alarmVisible = false;
                        this.$message.success("操作成功");
                        /* 重新刷新数据 */
                        this.getVirtualDevicesRequest();
                    } else {
                        this.$message.error(tempData.message);
                    }
                } else if (data.status === 401) {
                    this.loginVisible = true;
                } else {
                    this.disabledbtn = false;
                    this.loadingDialog = false;
                    this.$message.error('服务器繁忙，请稍后再试');
                }
            },

            alarmRequest(imei) {
                this.alarmVisible = true;
                this.alarmForm.imei = imei;
            },
        },

    }
</script>

<style scoped>
    .grid-content {
        display: flex;
        align-items: center;
        height: 150px;
    }
    .grid-cont-right {
        flex: 1;
        text-align: center;
        font-size: 15px;
        color: #282727;
    }
    .grid-con-icon {
        font-size: 65px;
        width: 100px;
        height: 150px;
        text-align: center;
        line-height: 150px;
        color: #fff;
    }
    .grid-con-2 .grid-con-icon {
        background: rgb(100, 213, 114);
    }
</style>